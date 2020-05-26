package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/ecshreve/civ-bot-go/bot/civsession"
	"github.com/ecshreve/civ-bot-go/bot/constants"
	"github.com/ecshreve/civ-bot-go/bot/util"
)

func banCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) == 1 {
		s.ChannelMessageSend(m.ChannelID, util.ErrorMessage("ban missing", "🤔  "+util.FormatUser(m.Author)+" you have to actually ban someone"))
		return
	}

	cs := civsession.CS
	c := civsession.BanCiv(args[1], m.Author.ID)
	if c == nil {
		s.ChannelMessageSend(m.ChannelID, util.ErrorMessage("invalid ban", "🤔  "+util.FormatUser(m.Author)+" can you pick a valid civ to ban?"))
		return
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title: "🍌 current bans",
		Color: constants.ColorRED,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "bans",
				Value: cs.FormatBans(),
			},
		},
	})

	// If all players have entered a Ban then pick Civs for all players.
	if len(cs.Bans) == len(cs.Players) {
		pick(s, m)
	}
}

// helpCommandHandler sends a help message. If a specific command is provided it
// provides that specific help message, otherwise it provides the default summary
// help message for all commands.
func helpCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	var title string
	var description string
	var fields []*discordgo.MessageEmbedField

	topic := ""
	if len(args) > 1 {
		topic = args[1]
	}

	// TODO: this is dumb, should make it better.
	switch topic {
	case "new":
		title = "🆕  new"
		description = "starts a new civ-bot session in the current channel \n whoever wants to play reacts with  ✋\n someone adds a  ✅ react when ready to continue"
	case "oops":
		title = "🤷‍♀️  oops"
		description = "abandon current session and start over"
	case "info":
		title = "ℹ️  info"
		description = "output the current state of the session"
	case "ban":
		title = "🍌  ban"
		description = "ban a civ so it can't be part of a player's picks"
	case "list":
		title = "☁︎  list"
		description = "lists all civs and leaders"
	default:
		title = "ℹ️  topics - civ-bot"
		description = "pick a topic to get help"
		fields = []*discordgo.MessageEmbedField{
			{
				Name:  "🆕  new",
				Value: "`/civ help new`: instructions on starting a new civ-bot session",
			},
			{
				Name:  "🤷‍♀️  oops",
				Value: "`/civ help oops`: abandon current session and start over",
			},
			{
				Name:  "ℹ️  info",
				Value: "`/civ help info`: output the current state of the session",
			},
			{
				Name:  "🍌  ban",
				Value: "`/civ help ban`: ban a civ so it can't be part of a player's picks",
			},
			{
				Name:  "☁︎  list",
				Value: "`/civ help list`: lists all possible civs",
			},
		}
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       title,
		Description: description,
		Fields:      fields,
		Color:       constants.ColorBLUE,
	})
}

func infoCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	cs := civsession.CS
	title := "ℹ️ current civ session info"
	players := util.FormatUsers(cs.Players)
	if players == "" {
		players = "no players yet"
	}

	_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title: title,
		Color: constants.ColorGREEN,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "players",
				Value: players,
			},
			{
				Name:  "bans",
				Value: cs.FormatBans(),
			},
		},
	})

	if err != nil {
		fmt.Printf("error generating info: %+v", err)
		return
	}
}

func listCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	var fields []*discordgo.MessageEmbedField
	for _, c := range civsession.CS.Civs {
		f := &discordgo.MessageEmbedField{
			Name:  c.CivBase + " -- " + c.LeaderBase,
			Value: fmt.Sprintf("[zigzag guide >>](%s)\n", c.ZigURL),
		}
		fields = append(fields, f)
	}
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:  "☁︎  list all possible civs",
		Color:  constants.ColorGREEN,
		Fields: fields,
	})

	if err != nil {
		fmt.Printf("error listing civs: %+v", err)
		return
	}
}

func newCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	title := "🆕 starting a new game"
	description := "- whoever wants to play react with  ✋\n- someone add a  ✅ react when ready to continue \n- enter `/civ oops` at any point to completely start over"

	newSession, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       title,
		Description: description,
		Color:       constants.ColorDARKPURPLE,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "new",
		},
	})

	if err != nil {
		fmt.Println("error creating new session")
		return
	}

	// Reset the CivSession and add the two reactions needed to add players to the
	// game, and complete adding players to the game.
	civsession.CS.Reset()
	s.MessageReactionAdd(m.ChannelID, newSession.ID, "✋")
	s.MessageReactionAdd(m.ChannelID, newSession.ID, "✅")
}
