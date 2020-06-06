package civsession

import (
	"testing"

	"github.com/ecshreve/civ-bot-go/pkg/discord"

	"github.com/stretchr/testify/assert"

	mocks "github.com/ecshreve/civ-bot-go/mocks/pkg/discord"
)

func TestBanCommandHandler(t *testing.T) {
	cs := NewCivSession()
	mockDiscordSession := &mocks.DataAccessLayer{}
	m := discord.MessageCreate{
		Message: discord.Message{
			M
		},
	}
	mockDiscordSession.On("ChannelMessageSend").Return(nil)
	cs.banCommandHandler(mockDiscordSession, nil, []string{"test"})
	mockDiscordSession.AssertExpectations(t)
	assert.NotNil(t, cs)
	assert.NotNil(t, mockDiscordSession)
}
