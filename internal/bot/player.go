package bot

import (
	"github.com/bwmarrin/discordgo"
)

// PlayerID is a wrapper around a discordgo.User.ID
type PlayerID string

// Player stores information for a User that is participating in an instance Bot.
type Player struct {
	PlayerID
	*discordgo.User
}

// GetPlayerIDToPlayerMap returns a map of [PlayerID]*Player for the given slice of Players.
func GetPlayerIDToPlayerMap(players []*Player) map[PlayerID]*Player {
	playerIDToPlayerMap := make(map[PlayerID]*Player)
	for _, player := range players {
		playerIDToPlayerMap[player.PlayerID] = player
	}

	return playerIDToPlayerMap
}
