package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// PlayerID is a wrapper around a discordgo.User.ID
type PlayerID string

// Player stores information for a User that is participating in an instance Bot.
type Player struct {
	PlayerID
	*discordgo.User
}

// NewPlayer returns a Player for the given discordgo.User.
func NewPlayer(u *discordgo.User) *Player {
	return &Player{
		PlayerID(u.ID),
		u,
	}
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
