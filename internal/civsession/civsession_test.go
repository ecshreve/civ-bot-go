package civsession_test

import (
	"testing"

	"github.com/ecshreve/civ-bot-go/internal/civsession"
	"github.com/stretchr/testify/assert"
)

func TestNewCivSession(t *testing.T) {
	cs := civsession.NewCivSession()
	assert.Equal(t, 0, len(cs.Players))
	assert.Equal(t, 43, len(cs.Civs))
	assert.Equal(t, 0, len(cs.Bans))
	assert.Equal(t, 0, len(cs.Picks))
	assert.True(t, cs.PickTime.IsZero())
	assert.Equal(t, 0, cs.RePickVotes)
}
