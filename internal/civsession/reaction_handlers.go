package civsession

import (
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

func addNumberReactionHelper(s *discordgo.Session, m *discordgo.Message, nums []int) {
	for _, n := range nums {
		s.MessageReactionAdd(m.ChannelID, m.ID, constants.NumEmojiMap[n])
	}
}

func (cfg *CivConfig) setConfigFieldHelper(s *discordgo.Session, m *discordgo.Message, r string) {
	embed := m.Embeds[0]
	if cfg.NumBans < 0 {
		cfg.NumBans = constants.EmojiNumMap[r]
		embed.Description = "updating config...\nselect NumPicks -- the number of Civs each player gets to choose from"
		s.ChannelMessageEditEmbed(m.ChannelID, m.ID, embed)
		s.MessageReactionsRemoveAll(m.ChannelID, m.ID)
		addNumberReactionHelper(s, m, []int{1, 2, 3, 4, 5})
		return
	}
	if cfg.NumPicks < 0 {
		cfg.NumPicks = constants.EmojiNumMap[r]
		embed.Description = "updating config...\nselect NumRepicks -- the max number of times allowed to re-pick Civs"
		s.ChannelMessageEditEmbed(m.ChannelID, m.ID, embed)
		s.MessageReactionsRemoveAll(m.ChannelID, m.ID)
		addNumberReactionHelper(s, m, []int{0, 1, 2, 3, 4, 5})
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
		embed.Description = "✅ **starting new civ picker session with the current config** ✅"
		s.ChannelMessageEditEmbed(m.ChannelID, m.ID, embed)
		cs.Reset()
		cs.newCommandHandler(s, &discordgo.MessageCreate{Message: m})
	}
	if r.Emoji.Name == "🛠" {
		embed.Description = "updating config...\nselect NumBans -- the number of Civs each player gets to ban"
		embed.Fields = nil
		s.ChannelMessageEditEmbed(m.ChannelID, m.ID, embed)
		s.MessageReactionsRemoveAll(m.ChannelID, m.ID)
		addNumberReactionHelper(s, m, []int{0, 1, 2, 3})

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
		embed.Description = "here's the current game config\nselect 🛠 to change config\nselect ✅ to accept config"
		embed.Fields = cs.getConfigEmbedFields()
		s.ChannelMessageEditEmbed(m.ChannelID, m.ID, embed)
		s.MessageReactionsRemoveAll(m.ChannelID, m.ID)
		s.MessageReactionAdd(m.ChannelID, m.ID, "🛠")
		s.MessageReactionAdd(m.ChannelID, m.ID, "✅")
	}
}

func (cs *CivSession) pickReactionHandler(s *discordgo.Session, r *discordgo.MessageReactionAdd, m *discordgo.Message, user *discordgo.User) {
	if r.Emoji.Name == "♻️" {
		cs.RePickVotes++
	}
}
