package bot

import (
	"testing"

	"github.com/samsarahq/go/snapshotter"
)

func TestNewState(t *testing.T) {
	snap := snapshotter.New(t)
	defer snap.Verify()

	output := NewCivState()
	snap.Snapshot("default civ state", output)
}

func TestReset(t *testing.T) {
	snap := snapshotter.New(t)
	defer snap.Verify()

	b, _ := MockBot(t)

	// Validate the default CivState and CivConfig.
	b.fixupCivStateForSnapshot()
	snap.Snapshot("new state", b.CivState)
	snap.Snapshot("new config", b.CivConfig)

	// Change some values in the CivState and CivConfig.
	b.CivState.RePickVotes = 99
	b.CivState.RePicksRemaining = 0
	b.CivConfig.Bans = 77
	b.CivConfig.Picks = 88
	b.CivConfig.RePicks = 99

	// Validate the altered CivState and CivConfig.
	b.fixupCivStateForSnapshot()
	snap.Snapshot("altered state", b.CivState)
	snap.Snapshot("altered config", b.CivConfig)

	// Reset maintaining the CivConfig.
	b.Reset(true)
	b.fixupCivStateForSnapshot()
	snap.Snapshot("state after reset - keepConfig true", b.CivState)
	snap.Snapshot("config after reset - keepConfig true", b.CivConfig)

	// Reset without maintaining the CivConfig.
	b.Reset(false)
	b.fixupCivStateForSnapshot()
	snap.Snapshot("state after reset - keepConfig false", b.CivState)
	snap.Snapshot("config after reset - keepConfig false", b.CivConfig)
}

// fixupCivStateForSnapshot nils out the CivState Civs slice and CivMap map to
// make the Snapshots easier to validate.
func (b *Bot) fixupCivStateForSnapshot() {
	b.CivState.Civs = nil
	b.CivState.CivMap = nil
}
