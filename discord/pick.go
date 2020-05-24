package discord

import (
	"fmt"
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

	var p []*discordgo.MessageEmbedField
	for k, v := range picks {
		f := &discordgo.MessageEmbedField{
			Name:  k.Username,
			Value: formatCivs(v),
		}
		p = append(p, f)
	}
	cs.PickTime = time.Now()
	time.AfterFunc(60*time.Second, func() { cs.handleRePick(s, m) })

	pickMessage, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "picks",
		Description: "here's this round of picks, if 50% or more players react with ♻️ in the next 60 seconds then we'll re pick",
		Color:       cDARKBLUE,
		Fields:      p,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "pick",
		},
	})

	if err != nil {
		fmt.Println("error sending pick message")
	}

	s.MessageReactionAdd(m.ChannelID, pickMessage.ID, "♻️")
}

func (cs *CivSession) pickReactionHandler(s *discordgo.Session, r *discordgo.MessageReactionAdd, m *discordgo.Message, user *discordgo.User) {
	if r.Emoji.Name == "♻️" {
		cs.RePickVotes++
	}
}

func (cs *CivSession) handleRePick(s *discordgo.Session, m *discordgo.MessageCreate) {
	if cs.RePickVotes*2 >= len(cs.Players) {
		cs.Picks = map[*discordgo.User][]*Civ{}
		cs.RePickVotes = 0
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title: "alright looks like we're picking again",
			Color: cORANGE,
		})
		cs.pick(s, m)
	} else {
		cs.reset()
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title: "great, have fun! see y'all next time 👋",
			Color: cORANGE,
		})
	}
}
