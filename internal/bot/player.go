package bot

import (
	"github.com/bwmarrin/discordgo"
)

type PlayerID string

type Player struct {
	PlayerID
	*discordgo.User
}

func GetPlayerIDToPlayerMap(players []*Player) map[PlayerID]*Player {
	playerIDToPlayerMap := make(map[PlayerID]*Player)
	for _, player := range players {
		playerIDToPlayerMap[player.PlayerID] = player
	}

	return playerIDToPlayerMap
}
