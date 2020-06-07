package bot_test

import (
	"testing"
	"time"

	"github.com/ecshreve/civ-bot-go/internal/constants"

	"github.com/ecshreve/civ-bot-go/internal/bot"
	"github.com/stretchr/testify/assert"
)

func TestNewState(t *testing.T) {
	state := bot.NewCivState()
	assert.NotNil(t, state)

	assert.Nil(t, state.Players)
	assert.Nil(t, state.PlayerMap)

	assert.Equal(t, len(constants.CivKeys), len(state.Civs))
	assert.Equal(t, len(constants.CivKeys), len(state.CivMap))

	assert.Nil(t, state.Bans)
	assert.Nil(t, state.Picks)

	assert.NotNil(t, state.PickState)
	assert.Equal(t, time.Time{}, state.PickState.PickTime)
	assert.Equal(t, 0, state.PickState.RePickVotes)
	assert.Equal(t, bot.DefaultCivConfig.RePicks, state.PickState.RePicksRemaining)
}
