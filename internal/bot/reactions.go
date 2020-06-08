package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/ecshreve/civ-bot-go/internal/constants"
	"github.com/samsarahq/go/oops"
)

// ReactionInfo defines the properties of a Reaction.
type ReactionInfo struct {
	Command CommandID
	Emojis  []string
}

// Reaction is a reaction on a Message that the Bot sent.
type Reaction interface {
	Info() *ReactionInfo
	Process(*Bot, *discordgo.MessageReaction) (*discordgo.Message, error)
}

// AllReactions is a slice contianing all the valid Reactions for the Bot.
var AllReactions = []Reaction{
	&newReaction{},
	&pickReaction{},
}

// Reaction interface implementation for reactions on the Message that results from
// processing the newCommand.
type newReaction struct{}

func (r *newReaction) Info() *ReactionInfo {
	return &ReactionInfo{
		Command: CommandID("new"),
		Emojis:  []string{"✋", "✅"},
	}
}

func (r *newReaction) Process(b *Bot, mr *discordgo.MessageReaction) (*discordgo.Message, error) {
	var embed *discordgo.MessageEmbed

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
		embed = &discordgo.MessageEmbed{
			Title:       "ℹ️ okay, now that we have our players",
			Description: "- everyone gets to ban a civ, enter `/civ ban <civ name>` to choose\n- if you change your mind just enter `/civ ban <new civ name>` to update your choice\n\nnote: you can enter a ban by either the civ or leader name",
			Color:       constants.ColorGREEN,
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  "Players",
					Value: FormatPlayers(b.CivState.Players),
				},
			},
		}
	}

	if embed == nil {
		return nil, oops.Errorf("invalid reaction")
	}

	newReactionMessage, err := b.DS.ChannelMessageSendEmbed(mr.ChannelID, embed)
	if err != nil {
		return nil, oops.Errorf("error sending embed %+v", embed)
	}

	return newReactionMessage, nil
}

type pickReaction struct{}

func (r *pickReaction) Info() *ReactionInfo {
	return &ReactionInfo{
		Command: CommandID("pick"),
		Emojis:  []string{"♻️"},
	}
}

func (r *pickReaction) Process(b *Bot, mr *discordgo.MessageReaction) (*discordgo.Message, error) {
	if mr.Emoji.Name == "♻️" {
		b.CivState.RePickVotes++
	}

	if b.CivState.RePickVotes*2 > len(b.CivState.Players) && b.CivState.RePicksRemaining > 0 {
		b.CivState.DoRepick = true
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
