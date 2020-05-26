package discord

import (
	"github.com/schollz/closestmatch"
)

// Civ represents an individual civilization.
type Civ struct {
	Key        CivKey
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
	for k, c := range CivBase {
		civ := &Civ{
			Key:        k,
			CivBase:    c,
			LeaderBase: CivLeaders[k],
			ZigURL:     CivZig[k],
		}
		civs = append(civs, civ)
	}
	return civs
}

// getCivByString takes a string and returns the Civ whose name or leader most
// closely matches the input string.
func (cs *CivSession) getCivByString(s string) *Civ {
	strsToTest := make([]string, 0)
	for _, k := range CivKeys {
		strsToTest = append(strsToTest, CivBase[k])
		strsToTest = append(strsToTest, CivLeaders[k])
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
