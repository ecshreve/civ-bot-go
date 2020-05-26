package discord

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/ecshreve/civ-bot-go/bot/constants"
)

func (cs *CivSession) pick(s *discordgo.Session, m *discordgo.MessageCreate) {
	possibles := []*Civ{}
	for _, c := range cs.Civs {
		if c.Banned == false {
			possibles = append(possibles, c)
		}
	}

	picks := make(map[*discordgo.User][]*Civ, 0)
	for _, u := range cs.Players {
		picks[u] = []*Civ{}
		rand.Seed(time.Now().Unix())
		for i := 0; i < 3; i++ {
			n := rand.Int() % len(possibles)
			p := possibles[n]
			if p.Picked != true {
				picks[u] = append(picks[u], p)
				p.Picked = true
			} else {
				i--
			}
		}
	}

	var p []*discordgo.MessageEmbedField
	for k, v := range picks {
		f := &discordgo.MessageEmbedField{
			Name:  k.Username,
			Value: formatCivs(v),
		}
		p = append(p, f)
	}
	cs.PickTime = time.Now()

	pickMessage, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "picks",
		Description: "here's this round of picks, if 50% or more players react with ♻️ in the next 60 seconds then we'll re pick",
		Color:       constants.ColorDARKBLUE,
		Fields:      p,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "pick",
		},
	})

	if err != nil {
		fmt.Println("error sending pick message")
	}

	s.MessageReactionAdd(m.ChannelID, pickMessage.ID, "♻️")
	countdown(s, m, pickMessage, cs.PickTime, 60)
}

func countdown(s *discordgo.Session, m *discordgo.MessageCreate, msg *discordgo.Message, start time.Time, seconds int64) {
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

	Session.handleRePick(s, m)
}

func (cs *CivSession) handleRePick(s *discordgo.Session, m *discordgo.MessageCreate) {
	if cs.RePickVotes*2 >= len(cs.Players) {
		cs.Picks = map[*discordgo.User][]*Civ{}
		cs.RePickVotes = 0
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title: "alright looks like we're picking again",
			Color: constants.ColorORANGE,
		})
		cs.pick(s, m)
	} else {
		cs.reset()
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title: "great, have fun! see y'all next time 👋",
			Color: constants.ColorORANGE,
		})
	}
}
