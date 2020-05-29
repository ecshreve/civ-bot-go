package civ

import (
	"fmt"

	"github.com/ecshreve/civ-bot-go/internal/constants"
	"github.com/schollz/closestmatch"
)

// Civ represents an individual civilization.
type Civ struct {
	Key        constants.CivKey
	CivBase    string
	LeaderBase string
	ZigURL     string
	FilthyTier int
	Banned     bool
	Picked     bool
}

// GenCivs generates and returns a slice of Civs based on the base values in the
// constants file.
func GenCivs() []*Civ {
	civs := make([]*Civ, 0)
	for k, c := range constants.CivBase {
		civ := &Civ{
			Key:        k,
			CivBase:    c,
			LeaderBase: constants.CivLeaders[k],
			ZigURL:     constants.CivZig[k],
			FilthyTier: constants.CivFilthyTiers[k],
		}
		civs = append(civs, civ)
	}
	return civs
}

// GetCivByString takes a string and returns the Civ whose name or leader most
// closely matches the input string.
func GetCivByString(s string, civs []*Civ) *Civ {
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
		if c.CivBase == closest {
			retCiv = c
			break
		}
		if c.LeaderBase == closest {
			retCiv = c
			break
		}
	}
	return retCiv
}

// FormatCiv returns a string for a single Civ in a readable format.
func FormatCiv(c *Civ) string {
	formatStr := "[%s -- %s](%s)"
	return fmt.Sprintf(formatStr, c.CivBase, c.LeaderBase, c.ZigURL)
}

// FormatCivs returns a string for a slice of Civs in a readable format.
func FormatCivs(cs []*Civ) string {
	ret := ""
	for _, c := range cs {
		ret = ret + "\n" + FormatCiv(c)
	}
	return ret
}
