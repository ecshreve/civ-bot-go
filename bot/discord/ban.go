package discord

// banCiv does a fuzzy match on the given string, if it finds a match it sets that
// Civ's Banned value to true and updates the CivSession's slice of Bans.
func banCiv(civToBan string, uid string) *Civ {
	cs := Session
	c := cs.getCivByString(civToBan)
	if c == nil || c.Banned == true {
		return nil
	}

	// If this player had previously banned a Civ then unban the previous Civ.
	if _, ok := cs.Bans[uid]; ok {
		cs.Bans[uid].Banned = false
	}

	c.Banned = true
	cs.Bans[uid] = c

	return c
}
