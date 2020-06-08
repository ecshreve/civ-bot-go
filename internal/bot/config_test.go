package bot_test

import (
	"testing"

	"github.com/ecshreve/civ-bot-go/internal/bot"

	"github.com/samsarahq/go/snapshotter"
)

func TestNewCivConfig(t *testing.T) {
	snap := snapshotter.New(t)
	defer snap.Verify()

	output := bot.NewCivConfig()
	snap.Snapshot("default civ config", output)
}

func TestGetEmbedFields(t *testing.T) {
	snap := snapshotter.New(t)
	defer snap.Verify()

	cfg := bot.NewCivConfig()
	output := cfg.GetEmbedFields()
	snap.Snapshot("default civ config embed fields", output)
}
