package bot

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/ecshreve/civ-bot-go/internal/civ"
	"github.com/ecshreve/civ-bot-go/internal/constants"
	"github.com/samsarahq/go/oops"
)

// makePick returns a random Civ from the given slice of Civs that has not been
// marked as Picked. If the provided list is empty, or all Civs in the list are
// already Banned or Picked then it returns nil and an error.
func (cs *CivState) makePick(civs []*civ.Civ) (*civ.Civ, error) {
	if len(civs) == 0 {
		return nil, oops.Errorf("empty civs arg")
	}

	rand.Seed(cs.Clk.Now().Unix())
	possibles := civs
	var p *civ.Civ
	foundPick := false

	// Keep picking at random until we find a Civ that hasn't been picked yet.
	// Once we find one, mark it as picked.
	for !foundPick && len(possibles) > 0 {
		n := rand.Int() % len(possibles)
		if civs[n].Picked != true && civs[n].Banned != true {
			p = civs[n]
			p.Picked = true
			foundPick = true

		} else {
			possibles = append(possibles[:n], possibles[n+1:]...)
		}
	}

	if p == nil {
		return nil, oops.Errorf("all civs are already Banned or Picked")
	}

	return p, nil
}

// makePicks returns a slice of numPicks random Civs from the given slice of
// Civs. If the provided slice is empty or if we are unable to get numPicks
// random Civs, it returns nil and an error.
func (cs *CivState) makePicks(civs []*civ.Civ, numPicks int) ([]*civ.Civ, error) {
	if len(civs) < numPicks {
		return nil, oops.Errorf("can't make %d picks from a slice of length %d", numPicks, len(civs))
	}

	var picks []*civ.Civ
	for i := 0; i < numPicks; i++ {
		pick, err := cs.makePick(civs)
		if err != nil {
			return nil, oops.Wrapf(err, "unable to make picks")
		}
		picks = append(picks, pick)
	}

	civ.SortCivs(picks)
	return picks, nil
}

// resetPicks resets the Picked value to false for all Civs and resets the CivState
// Picks field to a nil map.
func (cs *CivState) resetPicks() {
	for _, c := range cs.Civs {
		c.Picked = false
	}

	cs.Picks = make(map[PlayerID][]*civ.Civ)
}

// makePicksWithTier returns random picks for each Player ensuring that each
// Player gets at minimum one top tier Civ. It directly alters the CivSession
// that's the pointer receiver of the function. If unable to make picks for all
// Players then it returns an error.
func (b *Bot) makePicksWithTier() error {
	cs := b.CivState
	civsByTier := civ.GetCivsByTier(cs.Civs)
	topTierCivs := append(civsByTier[1], civsByTier[2]...)

	picks := make(map[PlayerID][]*civ.Civ)
	for _, p := range cs.Players {
		picks[p.PlayerID] = []*civ.Civ{}

		// Pick a top tier Civ for this player.
		topTierPick, err := cs.makePick(topTierCivs)
		if err != nil {
			cs.resetPicks()
			return oops.Wrapf(err, "unable to pick top tier civ for player: %s", p.Username)
		}
		picks[p.PlayerID] = append(picks[p.PlayerID], topTierPick)
	}

	// Pick remaining Civs for each Player.
	if b.CivConfig.Picks > 1 {
		for _, p := range cs.Players {
			lowTierPicks, err := cs.makePicks(cs.Civs, b.CivConfig.Picks-1)
			if err != nil {
				cs.resetPicks()
				return oops.Wrapf(err, "unable to pick remaining civs for player: %s", p.Username)
			}
			picks[p.PlayerID] = append(picks[p.PlayerID], lowTierPicks...)
		}
	}
	cs.Picks = picks
	cs.PickTime = time.Now()
	return nil
}

// makePicksWithoutTier returns random picks for each Player with no guarantees
// related to the Civ tiers.
func (b *Bot) makePicksWithoutTier() error {
	picks := make(map[PlayerID][]*civ.Civ)
	for _, p := range b.CivState.Players {
		// Pick Civs for this Player.
		picksForPlayer, err := b.CivState.makePicks(b.CivState.Civs, b.CivConfig.Picks)
		if err != nil {
			return oops.Wrapf(err, "unable to make picks for player: %s", p.PlayerID)
		}
		picks[p.PlayerID] = picksForPlayer
	}
	b.CivState.Picks = picks
	b.CivState.PickTime = time.Now()

	return nil
}

// Pick selects Civs at random and assigns them to Players. It also handles the
// logic surrounding re-picking. It returns an error if we encounter a problem
// making picks at any point during the pick flow.
//
// TODO: add test
func (b *Bot) Pick(channelID string) error {
	cs := b.CivState
	embedDescription := "here's this round of picks"
	if cs.RePicksRemaining > 0 {
		rePickThreshold := int(math.Ceil(float64(len(cs.Players)) / 2))
		embedDescription = embedDescription + fmt.Sprintf(", if %d or more players react with ♻️ in the next 60 seconds then we'll pick again\n\n%s re-picks remainging", rePickThreshold, constants.NumEmojiMap[cs.RePicksRemaining])
	}

	var err error
	if b.CivConfig.UseTiers {
		err = b.makePicksWithTier()
	} else {
		err = b.makePicksWithoutTier()
	}

	// If we encounter an error making picks then make sure to reset pick values
	// for the CivSession and return the error.
	if err != nil {
		cs.resetPicks()
		return oops.Wrapf(err, "unable to make picks")
	}

	var embedFields []*discordgo.MessageEmbedField
	for k, v := range cs.Picks {
		f := &discordgo.MessageEmbedField{
			Name:  cs.PlayerMap[k].Username,
			Value: civ.FormatCivs(v),
		}
		embedFields = append(embedFields, f)
	}

	pickMessage, err := b.DS.ChannelMessageSendEmbed(channelID, &discordgo.MessageEmbed{
		Title:       "picks",
		Description: embedDescription,
		Color:       constants.ColorDARKBLUE,
		Fields:      embedFields,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "pick",
		},
	})
	if err != nil {
		fmt.Println("error sending pick message")
	}

	if cs.RePicksRemaining > 0 {
		b.DS.MessageReactionAdd(channelID, pickMessage.ID, "♻️")
		b.countdown(pickMessage, cs.PickTime, 60)
	} else {
		b.DS.ChannelMessageSendEmbed(channelID, &discordgo.MessageEmbed{
			Title: "no more re-picks, those are your choices, deal with it",
			Color: constants.ColorORANGE,
		})
	}

	return nil
}

// countown handles editing the existing embed with Picks to display the
// amount of time remaining before the option to vote for a re-pick expires.
//
// TODO: add test
func (b *Bot) countdown(msg *discordgo.Message, start time.Time, seconds int64) {
	end := start.Add(time.Duration(time.Second * time.Duration(seconds)))
	channelID := msg.ChannelID
	messageID := msg.ID

	if len(msg.Embeds) != 1 {
		return
	}
	embed := msg.Embeds[0]

	for range time.Tick(1 * time.Second) {
		timeRemaining := int(end.Sub(time.Now()).Seconds())
		siren := ""
		if timeRemaining <= 10 && timeRemaining > 0 {
			siren = "🚨"
		}
		embed.Title = fmt.Sprintf("picks     %s -- %d seconds remaining -- %s", siren, timeRemaining, siren)
		b.DS.ChannelMessageEditEmbed(channelID, messageID, embed)
		if timeRemaining <= 0 {
			break
		}
	}

	b.handleRePick(channelID)
}

// handleRePick checks to see if the required number of re-pick votes have been
// reached, if so then pick again, if not then reset the CivSession and display
// a goodbye message.
//
// TODO: add test
func (b *Bot) handleRePick(channelID string) {
	cs := b.CivState
	cs.RePicksRemaining--

	if cs.RePickVotes*2 >= len(cs.Players) {
		cs.Picks = make(map[PlayerID][]*civ.Civ)
		cs.RePickVotes = 0
		b.DS.ChannelMessageSendEmbed(channelID, &discordgo.MessageEmbed{
			Title: "alright looks like we're picking again",
			Color: constants.ColorORANGE,
		})
		b.Pick(channelID)
	} else {
		b.DS.ChannelMessageSendEmbed(channelID, &discordgo.MessageEmbed{
			Title: "great, have fun! see y'all next time 👋",
			Color: constants.ColorORANGE,
		})
	}
}
