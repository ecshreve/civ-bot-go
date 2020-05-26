package discord

import (
	"github.com/ecshreve/civ-bot-go/constants"
	"github.com/ecshreve/civ-bot-go/util"
	"github.com/schollz/closestmatch"
)

// Civ represents an individual civilization.
type Civ struct {
	Key        util.CivKey
	CivBase    string
	LeaderBase string
	ZigURL     string
	Banned     bool
	Picked     bool
}

// genCivs generates and returns a slice of Civs based on the base values in the
// constants file.
func genCivs() []*Civ {
	civs := make([]*Civ, 0)
	for k, c := range constants.CivBase {
		civ := &Civ{
			Key:        k,
			CivBase:    c,
			LeaderBase: constants.CivLeaders[k],
			ZigURL:     constants.CivZig[k],
		}
		civs = append(civs, civ)
	}
	return civs
}

// getCivByString takes a string and returns the Civ whose name or leader most
// closely matches the input string.
func (cs *CivSession) getCivByString(s string) *Civ {
	strsToTest := make([]string, 0)
	for _, k := range constants.CivKeys {
		strsToTest = append(strsToTest, constants.CivBase[k])
		strsToTest = append(strsToTest, constants.CivLeaders[k])
	}

	bagSizes := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	cm := closestmatch.New(strsToTest, bagSizes)
	closest := cm.Closest(s)

	retCiv := &Civ{}
	for _, c := range cs.Civs {
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
