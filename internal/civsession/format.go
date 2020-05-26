package civsession

import (
	"github.com/bwmarrin/discordgo"
	"github.com/ecshreve/civ-bot-go/internal/civ"
	"github.com/ecshreve/civ-bot-go/internal/util"
)

// FormatPicks returns a string in a readable format for each player's picks.
func FormatPicks(picks map[*discordgo.User][]*civ.Civ) string {
	ret := ""
	for k, v := range picks {
		ret = ret + "\n" + util.FormatUser(k) + ":\n" + civ.FormatCivs(v) + "\n-----\n"
	}
	return ret
}

// FormatBans returns a string for the CivSession's Bans in a readable format.
func (cs *CivSession) FormatBans() string {
	bans := cs.Bans
	if bans == nil || len(bans) == 0 {
		return "no bans yet"
	}
	ret := ""
	for k, v := range bans {
		ret = ret + util.FormatUser(cs.Players[k]) + ": " + civ.FormatCiv(v) + "\n"
	}
	return ret
}
