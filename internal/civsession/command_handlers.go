package civsession

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/ecshreve/civ-bot-go/internal/civ"
	"github.com/ecshreve/civ-bot-go/internal/constants"
	"github.com/ecshreve/civ-bot-go/internal/util"
)

func (cs *CivSession) banCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) == 1 {
		s.ChannelMessageSend(m.ChannelID, util.ErrorMessage("ban missing", "🤔  "+util.FormatUser(m.Author)+" you have to actually ban someone"))
		return
	}

	c := cs.banCiv(args[1], m.Author.ID)
	if c == nil {
		s.ChannelMessageSend(m.ChannelID, util.ErrorMessage("invalid ban", "🤔  "+util.FormatUser(m.Author)+" can you pick a valid civ to ban?"))
		return
	}

	// TODO make this a generic helper func.
	var embedFields []*discordgo.MessageEmbedField
	for k, v := range cs.Bans {
		f := &discordgo.MessageEmbedField{
			Name:  cs.Players[k].Username,
			Value: civ.FormatCivs(v),
		}
		embedFields = append(embedFields, f)
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:  "🍌 current bans",
		Color:  constants.ColorRED,
		Fields: embedFields,
	})

	// If all players have entered the number of bans defined in cs.Config then
	// pick Civs for all players.
	if len(cs.Bans) == len(cs.Players) {
		for _, b := range cs.Bans {
			if len(b) < cs.Config.NumBans {
				return
			}
		}
		cs.pick(s, m)
	}
}

// helpCommandHandler sends a help message showing the available commands.
func helpCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	var title string
	var description string
	var fields []*discordgo.MessageEmbedField

	title = "ℹ️  civ-bot commands"
	description = "---"
	fields = []*discordgo.MessageEmbedField{
		{
			Name:  "⚙️  config",
			Value: "`/civ config`: enter the session configuration menu",
		},
		{
			Name:  "🆕  new",
			Value: "`/civ new`: start a new civ-bot session",
		},
		{
			Name:  "🤷‍♀️  oops",
			Value: "`/civ oops`: abandon current session state and start over",
		},
		{
			Name:  "ℹ️  info",
			Value: "`/civ info`: show the current state of the session",
		},
		{
			Name:  "🍌  ban",
			Value: "`/civ ban`: ban a civ so it can't be part of a player's picks in this session",
		},
		{
			Name:  "☁︎  list",
			Value: "`/civ list`: lists all possible civs",
		},
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       title,
		Description: description,
		Fields:      fields,
		Color:       constants.ColorBLUE,
	})
}

func (cs *CivSession) infoCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
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

func (cs *CivSession) listCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	var fields []*discordgo.MessageEmbedField
	for _, c := range cs.Civs {
		f := &discordgo.MessageEmbedField{
			Name:  fmt.Sprintf("{ t-%d } -- %s -- %s", c.FilthyTier, c.CivBase, c.LeaderBase),
			Value: fmt.Sprintf("[zigzag guide >>](%s)\n", c.ZigURL),
		}
		fields = append(fields, f)
	}
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:  "☁︎  list all possible civs (in no particular order)",
		Color:  constants.ColorGREEN,
		Fields: fields,
	})

	if err != nil {
		fmt.Printf("error listing civs: %+v", err)
		return
	}
}

func (cs *CivSession) configHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	title := "⚙️ configuration"
	description := "here's the current game config\nselect 🛠 to change config\nselect ✅ to accept config"
	fields := cs.getConfigEmbedFields()

	configMsg, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       title,
		Description: description,
		Fields:      fields,
		Color:       constants.ColorDARKGREY,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "config",
		},
	})
	if err != nil {
		fmt.Println("error sending config embed")
		return
	}

	s.MessageReactionAdd(m.ChannelID, configMsg.ID, "🛠")
	s.MessageReactionAdd(m.ChannelID, configMsg.ID, "✅")
}

func (cs *CivSession) newCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	title := "🆕 starting a new civ picker session"
	description := "- whoever wants to play react with  ✋\n- someone add a  ✅ react when ready to continue \n\n- enter `/civ config` to view / update the configuration \n- enter `/civ oops` at any point to completely start over\n- enter `/civ help` to see a list of available commands"

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
	cs.Reset()
	s.MessageReactionAdd(m.ChannelID, newSession.ID, "✋")
	s.MessageReactionAdd(m.ChannelID, newSession.ID, "✅")
}
