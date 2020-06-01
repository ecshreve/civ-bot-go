package civsession

import "github.com/ecshreve/civ-bot-go/internal/civ"

func getCivsByTier(civs []*civ.Civ) map[int][]*civ.Civ {
	civsByTier := make(map[int][]*civ.Civ)
	for _, c := range civs {
		civsByTier[c.FilthyTier] = append(civsByTier[c.FilthyTier], c)
	}
	return civsByTier
}

// resetPicks resets the Picked value to false for all Civs in the CivSession
// and resets the CivSession Picks field to a nil map.
func (cs *CivSession) resetPicks() {
	for _, c := range cs.Civs {
		c.Picked = false
	}

	cs.Picks = make(map[string][]*civ.Civ)
}
