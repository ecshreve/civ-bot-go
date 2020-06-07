package bot

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
	"testing"

	"github.com/bwmarrin/discordgo"
)

func Check(result interface{}, expected interface{}, t *testing.T) bool {
	if result != expected {
		_, fn, line, _ := runtime.Caller(1)
		fmt.Printf("[%s:%v] Expected %v but got %v\n", filepath.Base(fn), line, expected, result)
		t.Fail()
		return false
	}
	return true
}
func CheckNot(result interface{}, expected interface{}, t *testing.T) bool {
	if result == expected {
		_, fn, line, _ := runtime.Caller(1)
		fmt.Printf("[%s:%v] Unexpected result: %v\n", filepath.Base(fn), line, result)
		t.Fail()
		return false
	}
	return true
}

var mock *Mock

const (
	TestOwnerServer = iota
	TestServer
	TestRoleAdmin
	TestRoleMod
	TestRoleUser
	TestRoleMember
	TestRoleAssign
	TestRoleAssign2
	TestRoleSilence
	TestOwnerBot
	TestAdminMod
	TestAdmin
	TestMod
	TestUserNonAssign
	TestUserAssigned
	TestUserBoring
	TestUserSilence
	TestUserNew
	TestUserBot
	TestChannel
	TestChannel2
	TestChannelSpoil
	TestChannelFree
	TestChannelLog
	TestChannelMod
	TestChannelBored
	TestChannelJail
	TestChannelWelcome
)

func mockDiscordRole(role, index int) *discordgo.Role {
	name := "testrole"
	perms := 0
	switch role {
	case TestRoleAdmin:
		name = "Admin Role"
		perms = discordgo.PermissionAdministrator
	case TestRoleMod:
		name = "Mod Role"
		perms = discordgo.PermissionAllText | discordgo.PermissionManageRoles | discordgo.PermissionManageMessages | discordgo.PermissionManageChannels
	case TestRoleUser:
		name = "User Role"
		perms = discordgo.PermissionSendMessages | discordgo.PermissionReadMessages | discordgo.PermissionReadMessageHistory | discordgo.PermissionSendTTSMessages
	case TestRoleMember:
		name = "Member Role"
		perms = discordgo.PermissionSendMessages | discordgo.PermissionReadMessages
	case TestRoleAssign2:
		fallthrough
	case TestRoleAssign:
		name = "User Assignable"
		perms = discordgo.PermissionSendMessages | discordgo.PermissionReadMessages | discordgo.PermissionReadMessageHistory | discordgo.PermissionSendTTSMessages
	case TestRoleSilence:
		name = "Silent Role"
		perms = 0
	}

	return &discordgo.Role{
		ID:          strconv.Itoa(role | index),
		Name:        name,
		Mentionable: true,
		Hoist:       true,
		Color:       role,
		Position:    0,
		Permissions: perms,
	}
}

func mockDiscordMember(member, index int) *discordgo.Member {
	name := "member"
	roles := []string{}

	switch member {
	case TestOwnerBot:
		name = "Bot Owner"
	case TestOwnerServer:
		name = "Server Owner"
	case TestAdminMod:
		name = "Admin/Mod User"
		roles = append(roles, strconv.Itoa(TestRoleMod|index), strconv.Itoa(TestRoleAdmin|index))
	case TestAdmin:
		name = "Admin User"
		roles = append(roles, strconv.Itoa(TestRoleAdmin|index), strconv.Itoa(TestRoleMember|index))
	case TestMod:
		name = "Mod User"
		roles = append(roles, strconv.Itoa(TestRoleMod|index), strconv.Itoa(TestRoleMember|index))
	case TestUserNonAssign:
		name = "User With Non user-assignable Role"
		roles = append(roles, strconv.Itoa(TestRoleUser|index), strconv.Itoa(TestRoleMember|index))
	case TestUserAssigned:
		name = "User With user-assignable Role"
		roles = append(roles, strconv.Itoa(TestRoleAssign|index), strconv.Itoa(TestRoleMember|index))
	case TestUserBoring:
		name = "Boring User"
		roles = append(roles, strconv.Itoa(TestRoleMember|index))
	case TestUserSilence:
		name = "Silenced User"
		roles = append(roles, strconv.Itoa(TestRoleSilence|index))
	case TestUserNew:
		name = "New User"
	case TestUserBot:
		name = "Bot User"
	}

	return &discordgo.Member{
		GuildID: strconv.Itoa(TestServer | index),
		Nick:    name + ":" + strconv.Itoa(TestServer|index),
		User: &discordgo.User{
			ID:            strconv.Itoa(member | index),
			Email:         name + "@fake.com",
			Username:      name,
			Discriminator: strconv.Itoa(1000 + index),
			Bot:           member == TestUserBot,
		},
		Roles: roles,
	}
}

func mockDiscordChannel(channel int, index int) *discordgo.Channel {
	name := "Test"
	perms := []*discordgo.PermissionOverwrite{}

	disallowEveryone := &discordgo.PermissionOverwrite{
		ID:   strconv.Itoa(TestServer | index),
		Type: "role",
		Deny: discordgo.PermissionAllText,
	}
	allowMods := &discordgo.PermissionOverwrite{
		ID:    strconv.Itoa(TestRoleMod | index),
		Type:  "role",
		Allow: discordgo.PermissionAllText,
	}
	allowSilence := &discordgo.PermissionOverwrite{
		ID:    strconv.Itoa(TestRoleSilence | index),
		Type:  "role",
		Allow: discordgo.PermissionReadMessageHistory | discordgo.PermissionReadMessages | discordgo.PermissionSendMessages,
	}
	disallowSilence := &discordgo.PermissionOverwrite{
		ID:   strconv.Itoa(TestRoleSilence | index),
		Type: "role",
		Deny: discordgo.PermissionAllText,
	}

	switch channel {
	case TestChannel2:
		fallthrough
	case TestChannel:
		name = "Test Channel"
		perms = append(perms, disallowSilence)
	case TestChannelSpoil:
		name = "Spoiler Channel"
		perms = append(perms, disallowSilence)
	case TestChannelFree:
		name = "Free Channel"
		perms = append(perms, disallowSilence)
	case TestChannelLog:
		name = "Log Channel"
		perms = append(perms, disallowEveryone, allowMods)
	case TestChannelMod:
		name = "Mod Channel"
		perms = append(perms, disallowEveryone, allowMods)
	case TestChannelBored:
		name = "Bored Channel"
		perms = append(perms, disallowSilence)
	case TestChannelJail:
		name = "Jail Channel"
		perms = append(perms, disallowEveryone, allowSilence, allowMods)
	case TestChannelWelcome:
		name = "Welcome Channel"
		perms = append(perms, disallowEveryone, allowSilence, allowMods)
	}
	return &discordgo.Channel{
		ID:                   strconv.Itoa(channel | index),
		GuildID:              strconv.Itoa(TestServer | index),
		Name:                 name,
		Topic:                "Fake topic",
		Type:                 discordgo.ChannelTypeGuildText,
		NSFW:                 false,
		PermissionOverwrites: perms,
	}
}

func mockDiscordGuild(index int) *discordgo.Guild {
	return &discordgo.Guild{
		ID:                strconv.Itoa(TestServer),
		Name:              "Test Server",
		OwnerID:           strconv.Itoa(TestOwnerServer),
		VerificationLevel: discordgo.VerificationLevelLow,
		Large:             false,
		Unavailable:       false,
		Roles: []*discordgo.Role{
			mockDiscordRole(TestRoleAdmin, index),
			mockDiscordRole(TestRoleMod, index),
			mockDiscordRole(TestRoleUser, index),
			mockDiscordRole(TestRoleMember, index),
			mockDiscordRole(TestRoleAssign, index),
			mockDiscordRole(TestRoleAssign2, index),
			mockDiscordRole(TestRoleSilence, index),
		},
		Emojis: []*discordgo.Emoji{},
		Members: []*discordgo.Member{
			mockDiscordMember(TestOwnerBot, 0),
			mockDiscordMember(TestOwnerServer, index),
			mockDiscordMember(TestAdminMod, index),
			mockDiscordMember(TestAdmin, index),
			mockDiscordMember(TestMod, index),
			mockDiscordMember(TestUserAssigned, index),
			mockDiscordMember(TestUserNonAssign, index),
			mockDiscordMember(TestUserBoring, index),
			mockDiscordMember(TestUserSilence, index),
			mockDiscordMember(TestUserBot, index),
		},
		Presences: []*discordgo.Presence{},
		Channels: []*discordgo.Channel{
			mockDiscordChannel(TestChannel, index),
			mockDiscordChannel(TestChannel2, index),
			mockDiscordChannel(TestChannelSpoil, index),
			mockDiscordChannel(TestChannelFree, index),
			mockDiscordChannel(TestChannelLog, index),
			mockDiscordChannel(TestChannelMod, index),
			mockDiscordChannel(TestChannelBored, index),
			mockDiscordChannel(TestChannelJail, index),
			mockDiscordChannel(TestChannelWelcome, index),
		},
		VoiceStates: []*discordgo.VoiceState{},
	}
}
func mockDiscordSession() *DiscordSession {
	dg, _ := discordgo.New()
	ds := &DiscordSession{dg}
	ds.State.GuildAdd(mockDiscordGuild(1))
	return ds
}

func MockBot(t *testing.T) (*Bot, *Mock) {
	b := &Bot{
		DS:        mockDiscordSession(),
		CivConfig: NewCivConfig(),
		CivState:  NewCivState(),
		Commands:  AllCommands,
	}
	b.CommandMap = getCommandIDToCommandMap(b.Commands)
	b.ReactionMap = getCommandIDToReactionMap(AllReactions)

	mock = NewMock(t)
	return b, mock
}
