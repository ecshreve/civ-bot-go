package bot_test

import (
	"testing"

	"github.com/samsarahq/go/snapshotter"

	"github.com/ecshreve/civ-bot-go/internal/bot"
)

func TestNewState(t *testing.T) {
	snap := snapshotter.New(t)
	defer snap.Verify()

	output := bot.NewCivState()
	snap.Snapshot("default civ state", output)
}
