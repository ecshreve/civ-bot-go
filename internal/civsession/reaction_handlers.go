package civsession

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/ecshreve/civ-bot-go/internal/constants"
	"github.com/ecshreve/civ-bot-go/internal/util"
)

// newReactionHandler handles all new related reactions.
func (cs *CivSession) newReactionHandler(s *discordgo.Session, r *discordgo.MessageReactionAdd, m *discordgo.Message, user *discordgo.User) {
	if r.Emoji.Name == "✋" {
		cs.Players[user.ID] = user
	}
	if r.Emoji.Name == "✅" {
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title:       "ℹ️ okay, now that we have our players",
			Description: "- everyone gets to ban a civ, enter `/civ ban <civ name>` to choose\n- if you change your mind just enter `/civ ban <new civ name>` to update your choice\n\nnote: you can enter a ban by either the civ or leader name",
			Color:       constants.ColorGREEN,
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  "Players",
					Value: util.FormatUsers(cs.Players),
				},
			},
		})
	}
}

func (cfg *CivConfig) setConfigFieldHelper(s *discordgo.Session, m *discordgo.Message, r string) {
	embed := m.Embeds[0]
	if cfg.NumBans < 0 {
		cfg.NumBans = constants.EmojiNumMap[r]
		embed.Description = "updating config...\nselect NumPicks -- the number of Civs each player gets to choose from"
		s.ChannelMessageEditEmbed(m.ChannelID, m.ID, embed)
		s.MessageReactionsRemoveAll(m.ChannelID, m.ID)
		s.MessageReactionAdd(m.ChannelID, m.ID, "0️⃣")
		s.MessageReactionAdd(m.ChannelID, m.ID, "1️⃣")
		s.MessageReactionAdd(m.ChannelID, m.ID, "2️⃣")
		s.MessageReactionAdd(m.ChannelID, m.ID, "3️⃣")
		s.MessageReactionAdd(m.ChannelID, m.ID, "4️⃣")
		s.MessageReactionAdd(m.ChannelID, m.ID, "5️⃣")
		return
	}
	if cfg.NumPicks < 0 {
		cfg.NumPicks = constants.EmojiNumMap[r]
		embed.Description = "updating config...\nselect NumRepicks -- the max number of times allowed to re-pick Civs"
		s.ChannelMessageEditEmbed(m.ChannelID, m.ID, embed)
		s.MessageReactionsRemoveAll(m.ChannelID, m.ID)
		s.MessageReactionAdd(m.ChannelID, m.ID, "0️⃣")
		s.MessageReactionAdd(m.ChannelID, m.ID, "1️⃣")
		s.MessageReactionAdd(m.ChannelID, m.ID, "2️⃣")
		s.MessageReactionAdd(m.ChannelID, m.ID, "3️⃣")
		s.MessageReactionAdd(m.ChannelID, m.ID, "4️⃣")
		s.MessageReactionAdd(m.ChannelID, m.ID, "5️⃣")
		return
	}
	if cfg.NumRepicks < 0 {
		cfg.NumRepicks = constants.EmojiNumMap[r]
		embed.Description = "updating config...\nselect UseFilthyTiers -- true/false make picks based on Filthy's tier list"
		s.ChannelMessageEditEmbed(m.ChannelID, m.ID, embed)
		s.MessageReactionsRemoveAll(m.ChannelID, m.ID)
		s.MessageReactionAdd(m.ChannelID, m.ID, "👍")
		s.MessageReactionAdd(m.ChannelID, m.ID, "👎")
		return
	}
}

func (cs *CivSession) configReactionHandler(s *discordgo.Session, r *discordgo.MessageReactionAdd, m *discordgo.Message, user *discordgo.User) {
	embed := m.Embeds[0]
	if r.Emoji.Name == "✅" {
		embed.Description = "✅ **starting civ picker session with the current config** ✅"
		s.ChannelMessageEditEmbed(m.ChannelID, m.ID, embed)
	}
	if r.Emoji.Name == "🛠" {
		embed.Description = "updating config...\nselect NumBans -- the number of Civs each player gets to ban"
		embed.Fields = nil
		s.ChannelMessageEditEmbed(m.ChannelID, m.ID, embed)
		s.MessageReactionsRemoveAll(m.ChannelID, m.ID)
		s.MessageReactionAdd(m.ChannelID, m.ID, "0️⃣")
		s.MessageReactionAdd(m.ChannelID, m.ID, "1️⃣")
		s.MessageReactionAdd(m.ChannelID, m.ID, "2️⃣")
		s.MessageReactionAdd(m.ChannelID, m.ID, "3️⃣")
		cs.Config = &CivConfig{
			NumBans:        -1,
			NumPicks:       -1,
			NumRepicks:     -1,
			UseFilthyTiers: false,
		}
	}
	if _, ok := constants.EmojiNumMap[r.Emoji.Name]; ok {
		cs.Config.setConfigFieldHelper(s, m, r.Emoji.Name)
	}
	if r.Emoji.Name == "👍" || r.Emoji.Name == "👎" {
		cs.Config.UseFilthyTiers = r.Emoji.Name == "👍"
		embed.Description = "here's the current game config\nselect ✅ to accept config\nselect 🛠 to change config"
		embed.Fields = []*discordgo.MessageEmbedField{
			{
				Name:  "NumBans -- the number of Civs each player gets to ban",
				Value: fmt.Sprintf("%d", cs.Config.NumBans),
			},
			{
				Name:  "NumPicks -- the number of Civs each player gets to choose from",
				Value: fmt.Sprintf("%d", cs.Config.NumPicks),
			},
			{
				Name:  "NumRepicks -- the max number of times allowed to re-pick Civs",
				Value: fmt.Sprintf("%d", cs.Config.NumRepicks),
			},
			{
				Name:  "UseFilthyTiers -- true/false make picks based on Filthy's tier list",
				Value: fmt.Sprintf("%v", cs.Config.UseFilthyTiers),
			},
		}
		s.ChannelMessageEditEmbed(m.ChannelID, m.ID, embed)
		s.MessageReactionsRemoveAll(m.ChannelID, m.ID)
		s.MessageReactionAdd(m.ChannelID, m.ID, "✅")
		s.MessageReactionAdd(m.ChannelID, m.ID, "🛠")
	}
}

func (cs *CivSession) pickReactionHandler(s *discordgo.Session, r *discordgo.MessageReactionAdd, m *discordgo.Message, user *discordgo.User) {
	if r.Emoji.Name == "♻️" {
		cs.RePickVotes++
	}
}
