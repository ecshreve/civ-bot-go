package civsession

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/samsarahq/go/oops"

	"github.com/ecshreve/civ-bot-go/internal/civ"
	"github.com/ecshreve/civ-bot-go/internal/constants"
)

// banCiv does a fuzzy match on the given string, if it finds a match it sets that
// Civ's Banned value to true and updates the CivSession's slice of Bans.
func (cs *CivSession) banCiv(civToBan string, userID string) (*civ.Civ, error) {
	if cs.Config.NumBans == 0 {
		return nil, oops.Errorf("config numBans set to 0")
	}
	if civToBan == "" {
		return nil, oops.Errorf("empty civToBan argument")
	}
	if userID == "" {
		return nil, oops.Errorf("empty userID argument")
	}

	// If we didn't find a match, or the matched Civ is already banned then
	// return nil and an error.
	c := civ.GetCivByString(civToBan, cs.Civs)
	if c == nil {
		return nil, oops.Errorf("unable to get civ by string: %s", civToBan)
	}
	if c.Banned {
		return nil, oops.Errorf("found civ is already banned: %+v", c)
	}

	// If this player had previously banned the max number of Civs as defined by
	// the cs.Config, then unban the oldest one.
	if len(cs.Bans[userID]) == cs.Config.NumBans {
		cs.Bans[userID][0].Banned = false
		cs.Bans[userID] = cs.Bans[userID][1:]
	}

	// Add this Civ to the CivSession's slice of Bans.
	c.Banned = true
	cs.Bans[userID] = append(cs.Bans[userID], c)

	return c, nil
}

// makePick returns a random Civ from the given slice of Civs that has not been
// marked as Picked. If the provided list is empty, or all Civs in the list are
// already Banned or Picked then it returns nil and an error.
func (cs *CivSession) makePick(civs []*civ.Civ) (*civ.Civ, error) {
	if len(civs) == 0 {
		return nil, oops.Errorf("empty civs arg")
	}

	rand.Seed(cs.Clock.Now().Unix())
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
func (cs *CivSession) makePicks(civs []*civ.Civ, numPicks int) ([]*civ.Civ, error) {
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

// makePicksWithTier returns random picks for each Player ensuring that each
// Player gets at minimum one top tier Civ. It directly alters the CivSession
// that's the pointer receiver of the function. If unable to make picks for all
// Players then it returns an error.
func (cs *CivSession) makePicksWithTier() error {
	civsByTier := getCivsByTier(cs.Civs)
	topTierCivs := append(civsByTier[1], civsByTier[2]...)

	picks := make(map[string][]*civ.Civ)
	for _, u := range cs.Players {
		picks[u.ID] = []*civ.Civ{}

		// Pick a top tier Civ for this player.
		topTierPick, err := cs.makePick(topTierCivs)
		if err != nil {
			cs.resetPicks()
			return oops.Wrapf(err, "unable to pick top tier civ for player: %s", u.Username)
		}
		picks[u.ID] = append(picks[u.ID], topTierPick)
	}

	// Pick remaining Civs for each Player.
	if cs.Config.NumPicks > 1 {
		for _, u := range cs.Players {
			lowTierPicks, err := cs.makePicks(cs.Civs, cs.Config.NumPicks-1)
			if err != nil {
				cs.resetPicks()
				return oops.Wrapf(err, "unable to pick remaining civs for player: %s", u.Username)
			}
			picks[u.ID] = append(picks[u.ID], lowTierPicks...)
		}
	}
	cs.Picks = picks
	cs.PickTime = time.Now()
	return nil
}

// makePicksWithoutTier returns random picks for each Player with no guarantees
// related to the Civ tiers.
func (cs *CivSession) makePicksWithoutTier() error {
	picks := make(map[string][]*civ.Civ)
	for _, u := range cs.Players {
		// Pick Civs for this player.
		picksForPlayer, err := cs.makePicks(cs.Civs, cs.Config.NumPicks)
		if err != nil {
			return oops.Wrapf(err, "unable to make picks for player: %s", u.Username)
		}
		picks[u.ID] = picksForPlayer
	}
	cs.Picks = picks
	cs.PickTime = time.Now()

	return nil
}

// pick selects Civs at random and assigns them to Players. It also handles the
// logic surrounding re-picking. It returns an error if we encounter a problem
// making picks at any point during the pick flow.
//
// TODO: add test
func (cs *CivSession) pick(s *discordgo.Session, m *discordgo.MessageCreate) error {
	embedDescription := "here's this round of picks"
	if cs.RePicksRemaining > 0 {
		rePickThreshold := int(math.Ceil(float64(len(cs.Players)) / 2))
		embedDescription = embedDescription + fmt.Sprintf(", if %d or more players react with ♻️ in the next 60 seconds then we'll pick again\n\n%s re-picks remainging", rePickThreshold, constants.NumEmojiMap[cs.RePicksRemaining])
	}

	var err error
	if cs.Config.UseFilthyTiers {
		err = cs.makePicksWithTier()
	} else {
		err = cs.makePicksWithoutTier()
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

	pickMessage, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
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
		s.MessageReactionAdd(m.ChannelID, pickMessage.ID, "♻️")
		cs.countdown(s, m, pickMessage, cs.PickTime, 60)
	} else {
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
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
func (cs *CivSession) countdown(s *discordgo.Session, m *discordgo.MessageCreate, msg *discordgo.Message, start time.Time, seconds int64) {
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
		s.ChannelMessageEditEmbed(channelID, messageID, embed)
		if timeRemaining <= 0 {
			break
		}
	}

	cs.handleRePick(s, m)
}

// handleRePick checks to see if the required number of re-pick votes have been
// reached, if so then pick again, if not then reset the CivSession and display
// a goodbye message.
//
// TODO: add test
func (cs *CivSession) handleRePick(s *discordgo.Session, m *discordgo.MessageCreate) {
	cs.RePicksRemaining--

	if cs.RePickVotes*2 >= len(cs.Players) {
		cs.Picks = map[string][]*civ.Civ{}
		cs.RePickVotes = 0
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title: "alright looks like we're picking again",
			Color: constants.ColorORANGE,
		})
		cs.pick(s, m)
	} else {
		cs.Reset()
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title: "great, have fun! see y'all next time 👋",
			Color: constants.ColorORANGE,
		})
	}
}
