package bot

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
)

func TestValidateMessageHandlerArgs(t *testing.T) {
	b, mock := MockBot(t)

	testChannelID := "testChannelID"
	user := &discordgo.User{
		ID: "otheruser",
	}

	testcases := []struct {
		description     string
		message         *discordgo.Message
		mockCalls       func()
		expectedCommand Command
	}{
		{
			description: "message from bot",
			message: &discordgo.Message{
				Author: b.DS.State.User,
			},
			expectedCommand: nil,
		},
		{
			description: "message without /civ prefix",
			message: &discordgo.Message{
				Author:  user,
				Content: "blah blah blah",
			},
			expectedCommand: nil,
		},
		{
			description: "message without args",
			message: &discordgo.Message{
				ChannelID: testChannelID,
				Author:    user,
				Content:   "/civ",
			},
			mockCalls: func() {
				mock.Expect(b.DS.ChannelMessageSend, testChannelID, MockAny{})
			},
			expectedCommand: nil,
		},
		{
			description: "message with invalid command",
			message: &discordgo.Message{
				ChannelID: testChannelID,
				Author:    user,
				Content:   "/civ thisisaninvalidcommand",
			},
			mockCalls: func() {
				mock.Expect(b.DS.ChannelMessageSend, testChannelID, MockAny{})
			},
			expectedCommand: nil,
		},
		{
			description: "message with valid command",
			message: &discordgo.Message{
				ChannelID: testChannelID,
				Author:    user,
				Content:   "/civ help",
			},
			expectedCommand: &helpCommand{},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.description, func(t *testing.T) {
			if testcase.mockCalls != nil {
				testcase.mockCalls()
			}

			c := b.validateMessageHandlerArgs(testcase.message)

			if testcase.expectedCommand == nil {
				assert.Nil(t, c)
			} else {
				assert.Equal(t, testcase.expectedCommand, c)
			}
		})
	}
}
