package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/ecshreve/civ-bot-go/internal/constants"
	"github.com/samsarahq/go/oops"
)

type ReactionInfo struct {
	Command CommandID
	Emojis  []string
}

type Reaction interface {
	Info() *ReactionInfo
	Process(*Bot, *discordgo.MessageReaction) (*discordgo.Message, error)
}

var AllReactions = []Reaction{
	&newReaction{},
}

type newReaction struct{}

func (r *newReaction) Info() *ReactionInfo {
	return &ReactionInfo{
		Command: CommandID("new"),
		Emojis:  []string{"✋", "✅"},
	}
}

func (r *newReaction) Process(b *Bot, mr *discordgo.MessageReaction) (*discordgo.Message, error) {
	if mr.Emoji.Name == "✋" {
		user, err := b.DS.User(mr.UserID)
		if err != nil {
			return nil, oops.Wrapf(err, "unable to get user for userID: %s", mr.UserID)
		}

		player := NewPlayer(user)
		b.CivState.Players = append(b.CivState.Players, player)
		b.CivState.PlayerMap[player.PlayerID] = player
		return nil, nil
	}
	if mr.Emoji.Name == "✅" {
		b.DS.ChannelMessageSendEmbed(mr.ChannelID, &discordgo.MessageEmbed{
			Title:       "ℹ️ okay, now that we have our players",
			Description: "- everyone gets to ban a civ, enter `/civ ban <civ name>` to choose\n- if you change your mind just enter `/civ ban <new civ name>` to update your choice\n\nnote: you can enter a ban by either the civ or leader name",
			Color:       constants.ColorGREEN,
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  "Players",
					Value: FormatPlayers(b.CivState.Players),
				},
			},
		})
	}
	return nil, nil
}

func getCommandIDToReactionMap(reactions []Reaction) map[CommandID]Reaction {
	reactionMap := make(map[CommandID]Reaction)
	for _, r := range reactions {
		reactionMap[r.Info().Command] = r
	}

	return reactionMap
}
