package botutil

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func genPlayersForTest() map[string]*discordgo.User {
	var players map[string]*discordgo.User
	for i := 1; i <= 10; i++ {
		p := &discordgo.User{
			ID:       fmt.Sprintf("p%dID", i),
			Email:    fmt.Sprintf("p%dEmail@devnull.com", i),
			Username: fmt.Sprintf("p%dUsername", i),
		}
		players[p.ID] = p
	}
	return players
}

// CivBotTestData contains data to use in automated tests.
type CivBotTestData struct {
	Players map[string]*discordgo.User
}

// Data is an instance of CivBotTestData that contains initialized test data.
var Data = &CivBotTestData{
	Players: genPlayersForTest(),
}
