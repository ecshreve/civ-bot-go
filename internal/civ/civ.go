package civ

import (
	"fmt"

	"github.com/schollz/closestmatch"

	"github.com/ecshreve/civ-bot-go/internal/constants"
)

// Civ represents an individual civilization.
type Civ struct {
	// Key is the CivKey enum entry for this Civ.
	Key constants.CivKey

	// CivBase is the string representation of the Civ's name.
	CivBase string

	// LeaderBase is the string representation of the Civ's Leader's name.
	LeaderBase string

	// ZigURL is the URL of the ZigZagal guide for the Civ.
	ZigURL string

	// FilthyTier is an integer [1...6] representing the Civ's tier.
	FilthyTier int

	// Banned is a boolean indicating if this Civ is currently banned.
	Banned bool

	// Picked is a boolean indicating if this Civ has been picked for a Player.
	Picked bool
}

// GenCivs generates and returns a slice of Civs for all the base values
// defined in the constants package.
func GenCivs() []*Civ {
	civs := make([]*Civ, 0)
	for _, k := range constants.CivKeys {
		civ := &Civ{
			Key:        k,
			CivBase:    constants.CivBase[k],
			LeaderBase: constants.CivLeaders[k],
			ZigURL:     constants.CivZig[k],
			FilthyTier: constants.CivFilthyTiers[k],
		}
		civs = append(civs, civ)
	}
	return civs
}

// GenCivMap generates and returns a map of CivKey to Civ for all base values
// defined in the constants package.
func GenCivMap() map[constants.CivKey]*Civ {
	civs := GenCivs()
	civMap := make(map[constants.CivKey]*Civ)

	for _, c := range civs {
		civMap[c.Key] = c
	}

	return civMap
}

// GetCivByString takes a string and returns the Civ whose name or leader most
// closely matches the input string.
func GetCivByString(s string, civs []*Civ) *Civ {
	// We want to test the input string agains all civiliation and leader names.
	strsToTest := make([]string, 0)
	for _, k := range constants.CivKeys {
		strsToTest = append(strsToTest, constants.CivBase[k])
		strsToTest = append(strsToTest, constants.CivLeaders[k])
	}

	bagSizes := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	cm := closestmatch.New(strsToTest, bagSizes)
	closest := cm.Closest(s)

	retCiv := &Civ{}
	for _, c := range civs {
		if c.CivBase == closest || c.LeaderBase == closest {
			retCiv = c
			break
		}
	}

	return retCiv
}

// FormatCiv returns a string for a single Civ in a readable format.
func FormatCiv(c *Civ) string {
	formatStr := "{ t-%d } -- [%s -- %s](%s)"
	return fmt.Sprintf(formatStr, c.FilthyTier, c.CivBase, c.LeaderBase, c.ZigURL)
}

// FormatCivs returns a string for a slice of Civs in a readable format.
func FormatCivs(cs []*Civ) string {
	ret := ""
	for _, c := range cs {
		ret = ret + "\n" + FormatCiv(c)
	}
	return ret
}
