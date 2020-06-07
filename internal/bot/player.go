package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/ecshreve/civ-bot-go/internal/civ"
	"github.com/samsarahq/go/oops"
)

// PlayerID is a wrapper around a discordgo.User.ID
type PlayerID string

// Player stores information for a User that is participating in an instance Bot.
type Player struct {
	PlayerID
	*discordgo.User
	BannedCivs []*civ.Civ
}

// NewPlayer returns a Player for the given discordgo.User.
func NewPlayer(u *discordgo.User) *Player {
	return &Player{
		PlayerID(u.ID),
		u,
		make([]*civ.Civ, 0),
	}
}

func (p *Player) BanCiv(b *Bot, toBan string) (*civ.Civ, error) {
	if toBan == "" {
		return nil, oops.Errorf("empty toBan argument for player: %s", p.PlayerID)
	}

	// If we didn't find a match, or the matched Civ is already banned then
	// return nil and an error.
	c := civ.GetCivByString(toBan, b.CivState.Civs)
	if c == nil {
		return nil, oops.Errorf("unable to get toBan by string: %s", toBan)
	}
	if c.Banned {
		return nil, oops.Errorf("found toBan is already banned: %+v", c)
	}

	// If this player had previously banned the max number of Civs as defined by
	// the cs.Config, then unban the oldest one.
	if len(p.BannedCivs) == b.CivConfig.Bans {
		p.BannedCivs[0].Banned = false
		p.BannedCivs = p.BannedCivs[1:]
	}

	c.Banned = true
	p.BannedCivs = append(p.BannedCivs, c)
	b.CivState.Bans[p.PlayerID] = p.BannedCivs

	return c, nil
}

// GetPlayerIDToPlayerMap returns a map of [PlayerID]*Player for the given slice of Players.
func GetPlayerIDToPlayerMap(players []*Player) map[PlayerID]*Player {
	playerIDToPlayerMap := make(map[PlayerID]*Player)
	for _, player := range players {
		playerIDToPlayerMap[player.PlayerID] = player
	}

	return playerIDToPlayerMap
}

// FormatPlayer returns the string representation for the given Player.
func FormatPlayer(player *Player) string {
	return fmt.Sprintf("<@%s>", player.PlayerID)
}

// FormatPlayers returns the string representation for the given slice of Players.
func FormatPlayers(players []*Player) string {
	ret := ""
	for _, p := range players {
		ret = ret + FormatPlayer(p) + "\n"
	}
	return ret
}
