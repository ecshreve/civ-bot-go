package civsession

import "github.com/ecshreve/civ-bot-go/internal/civ"

func getCivsByTier(civs []*civ.Civ) map[int][]*civ.Civ {
	civsByTier := make(map[int][]*civ.Civ)
	for _, c := range civs {
		civsByTier[c.FilthyTier] = append(civsByTier[c.FilthyTier], c)
	}
	return civsByTier
}
