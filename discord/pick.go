package discord

import (
	"math/rand"
	"time"

	"github.com/bwmarrin/discordgo"
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

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title: "picks",
		Color: cDARKBLUE,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "picks",
				Value: formatPicks(picks),
			},
		},
	})
}
