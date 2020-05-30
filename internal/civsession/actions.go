package civsession

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/ecshreve/civ-bot-go/internal/civ"
	"github.com/ecshreve/civ-bot-go/internal/constants"
)

// banCiv does a fuzzy match on the given string, if it finds a match it sets that
// Civ's Banned value to true and updates the CivSession's slice of Bans.
func (cs *CivSession) banCiv(civToBan string, userID string) *civ.Civ {
	if civToBan == "" || userID == "" {
		return nil
	}

	// If we didn't find a match, or the matched Civ is already banned then just
	// return nil.
	c := civ.GetCivByString(civToBan, cs.Civs)
	if c == nil || c.Banned == true {
		return nil
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

	return c
}

// makePick returns a random Civ from the given slice of Civs that has not been
// marked as Picked.
func makePick(civs []*civ.Civ) *civ.Civ {
	rand.Seed(time.Now().Unix())

	var p *civ.Civ
	foundPick := false

	// Keep picking at random until we find a Civ that hasn't been picked yet.
	// Once we find one, mark it as picked.
	for !foundPick {
		n := rand.Int() % len(civs)
		p = civs[n]
		if p.Picked != true {
			foundPick = true
		}
	}
	p.Picked = true

	return p
}

func (cs *CivSession) makePicksWithTier() []*discordgo.MessageEmbedField {
	possibles := []*civ.Civ{}
	for _, c := range cs.Civs {
		if c.Banned == false {
			possibles = append(possibles, c)
		}
	}

	possiblesByTier := getCivsByTier(possibles)
	topTierPossibles := append(possiblesByTier[1], possiblesByTier[2]...)
	picks := make(map[string][]*civ.Civ)
	for _, u := range cs.Players {
		picks[u.ID] = []*civ.Civ{}

		// Pick a top tier Civ for this player.
		picks[u.ID] = append(picks[u.ID], makePick(topTierPossibles))
	}

	// Pick remaining Civs for each Player.
	for _, u := range cs.Players {
		for i := 0; i < cs.Config.NumPicks-1; i++ {
			picks[u.ID] = append(picks[u.ID], makePick(possibles))
		}
	}
	cs.Picks = picks
	cs.PickTime = time.Now()

	// Generate MessageEmbedFields for the Picks.
	var p []*discordgo.MessageEmbedField
	for k, v := range picks {
		f := &discordgo.MessageEmbedField{
			Name:  cs.Players[k].Username,
			Value: civ.FormatCivs(v),
		}
		p = append(p, f)
	}

	return p
}

func (cs *CivSession) makePicks() []*discordgo.MessageEmbedField {
	possibles := []*civ.Civ{}
	for _, c := range cs.Civs {
		if c.Banned == false {
			possibles = append(possibles, c)
		}
	}

	picks := make(map[string][]*civ.Civ)
	for _, u := range cs.Players {
		picks[u.ID] = []*civ.Civ{}
		rand.Seed(time.Now().Unix())
		for i := 0; i < cs.Config.NumPicks; i++ {
			n := rand.Int() % len(possibles)
			p := possibles[n]
			if p.Picked != true {
				picks[u.ID] = append(picks[u.ID], p)
				p.Picked = true
			} else {
				i--
			}
		}
	}
	cs.Picks = picks

	var p []*discordgo.MessageEmbedField
	for k, v := range picks {
		f := &discordgo.MessageEmbedField{
			Name:  cs.Players[k].Username,
			Value: civ.FormatCivs(v),
		}
		p = append(p, f)
	}
	cs.PickTime = time.Now()

	return p
}

// pick selects Civs at random and assigns them to Players. It also handles the
// logic surrounding re-picking.
func (cs *CivSession) pick(s *discordgo.Session, m *discordgo.MessageCreate) {
	embedDescription := "here's this round of picks"
	if cs.RePicksRemaining > 0 {
		rePickThreshold := int(math.Ceil(float64(len(cs.Players)) / 2))
		embedDescription = embedDescription + fmt.Sprintf(", if %d or more players react with ♻️ in the next 60 seconds then we'll pick again\n\n%s re-picks remainging", rePickThreshold, constants.NumEmojiMap[cs.RePicksRemaining])
	}

	var embedFields []*discordgo.MessageEmbedField
	if cs.Config.UseFilthyTiers {
		embedFields = cs.makePicksWithTier()
	} else {
		embedFields = cs.makePicks()
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
}

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
