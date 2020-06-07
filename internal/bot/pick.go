package bot

import (
	"math/rand"
	"time"

	"github.com/ecshreve/civ-bot-go/internal/civ"
	"github.com/samsarahq/go/oops"
)

// makePick returns a random Civ from the given slice of Civs that has not been
// marked as Picked. If the provided list is empty, or all Civs in the list are
// already Banned or Picked then it returns nil and an error.
func (cs *CivState) makePick(civs []*civ.Civ) (*civ.Civ, error) {
	if len(civs) == 0 {
		return nil, oops.Errorf("empty civs arg")
	}

	rand.Seed(cs.Clk.Now().Unix())
	possibles := civs
	var p *civ.Civ
	foundPick := false

	// Keep picking at random until we find a Civ that hasn't been picked yet.
	// Once we find one, mark it as picked.
	for !foundPick && len(possibles) > 0 {
		n := rand.Int() % len(possibles)
		if civs[n].Picked != true && civs[n].Banned != true {
			p = civs[n]
			p.Picked = true
			foundPick = true

		} else {
			possibles = append(possibles[:n], possibles[n+1:]...)
		}
	}

	if p == nil {
		return nil, oops.Errorf("all civs are already Banned or Picked")
	}

	return p, nil
}

// makePicks returns a slice of numPicks random Civs from the given slice of
// Civs. If the provided slice is empty or if we are unable to get numPicks
// random Civs, it returns nil and an error.
func (cs *CivState) makePicks(civs []*civ.Civ, numPicks int) ([]*civ.Civ, error) {
	if len(civs) < numPicks {
		return nil, oops.Errorf("can't make %d picks from a slice of length %d", numPicks, len(civs))
	}

	var picks []*civ.Civ
	for i := 0; i < numPicks; i++ {
		pick, err := cs.makePick(civs)
		if err != nil {
			return nil, oops.Wrapf(err, "unable to make picks")
		}
		picks = append(picks, pick)
	}

	civ.SortCivs(picks)
	return picks, nil
}

// resetPicks resets the Picked value to false for all Civs and resets the CivState
// Picks field to a nil map.
func (cs *CivState) resetPicks() {
	for _, c := range cs.Civs {
		c.Picked = false
	}

	cs.Picks = make(map[PlayerID][]*civ.Civ)
}

// makePicksWithTier returns random picks for each Player ensuring that each
// Player gets at minimum one top tier Civ. It directly alters the CivSession
// that's the pointer receiver of the function. If unable to make picks for all
// Players then it returns an error.
func (b *Bot) makePicksWithTier() error {
	cs := b.CivState
	civsByTier := civ.GetCivsByTier(cs.Civs)
	topTierCivs := append(civsByTier[1], civsByTier[2]...)

	picks := make(map[PlayerID][]*civ.Civ)
	for _, p := range cs.Players {
		picks[p.PlayerID] = []*civ.Civ{}

		// Pick a top tier Civ for this player.
		topTierPick, err := cs.makePick(topTierCivs)
		if err != nil {
			cs.resetPicks()
			return oops.Wrapf(err, "unable to pick top tier civ for player: %s", p.Username)
		}
		picks[p.PlayerID] = append(picks[p.PlayerID], topTierPick)
	}

	// Pick remaining Civs for each Player.
	if b.CivConfig.Picks > 1 {
		for _, p := range cs.Players {
			lowTierPicks, err := cs.makePicks(cs.Civs, b.CivConfig.Picks-1)
			if err != nil {
				cs.resetPicks()
				return oops.Wrapf(err, "unable to pick remaining civs for player: %s", p.Username)
			}
			picks[p.PlayerID] = append(picks[p.PlayerID], lowTierPicks...)
		}
	}
	cs.Picks = picks
	cs.PickTime = time.Now()
	return nil
}

// makePicksWithoutTier returns random picks for each Player with no guarantees
// related to the Civ tiers.
func (b *Bot) makePicksWithoutTier() error {
	picks := make(map[PlayerID][]*civ.Civ)
	for _, p := range b.CivState.Players {
		// Pick Civs for this Player.
		picksForPlayer, err := b.CivState.makePicks(b.CivState.Civs, b.CivConfig.Picks)
		if err != nil {
			return oops.Wrapf(err, "unable to make picks for player: %s", p.PlayerID)
		}
		picks[p.PlayerID] = picksForPlayer
	}
	b.CivState.Picks = picks
	b.CivState.PickTime = time.Now()

	return nil
}

func (b *Bot) Pick() {
	// Pick
}
