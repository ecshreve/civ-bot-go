package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// CivConfig stores configuratoin information for an instance of the Bot.
type CivConfig struct {
	Bans     int
	Picks    int
	RePicks  int
	UseTiers bool
}

// DefaultCivConfig defines the default CivConfig used when creating a new Bot.
var DefaultCivConfig = CivConfig{
	Bans:     1,
	Picks:    3,
	RePicks:  3,
	UseTiers: false,
}

// NewCivConfig simply returns a pointer to a copy of the DefaultCivConfig.
func NewCivConfig() *CivConfig {
	cfg := DefaultCivConfig
	return &cfg
}

// GetEmbedFields returns a slice of MessageEmbedFields for the given CivConfig.
func (c *CivConfig) GetEmbedFields() []*discordgo.MessageEmbedField {
	return []*discordgo.MessageEmbedField{
		{
			Name:  "`Bans` -- the number of Civs each player gets to ban",
			Value: fmt.Sprintf("**%d**", c.Bans),
		},
		{
			Name:  "`Picks` -- the number of Civs each player gets to choose from",
			Value: fmt.Sprintf("**%d**", c.Picks),
		},
		{
			Name:  "`RePicks` -- the max number of times allowed to re-pick Civs",
			Value: fmt.Sprintf("**%d**", c.RePicks),
		},
		{
			Name:  "`UseTiers` -- make picks based on FilthyRobot's tier list -- setting this to `true` ensures that each Player gets at minimum one t1/t2 Civ in their list of Picks",
			Value: fmt.Sprintf("**%v**", c.UseTiers),
		},
	}
}
