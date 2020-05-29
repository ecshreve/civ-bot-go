package civsession

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ecshreve/civ-bot-go/internal/civ"
	"github.com/ecshreve/civ-bot-go/internal/constants"
)

func TestGetCivsByTier(t *testing.T) {
	civs := civ.GenCivs()
	civsByTier := getCivsByTier(civs)

	for k, v := range civsByTier {
		for _, c := range v {
			assert.Equal(t, constants.CivFilthyTiers[c.Key], k)
		}
	}
}
