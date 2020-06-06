package bot

import "fmt"

type Config struct {
	Bans     int
	Picks    int
	RePicks  int
	UseTiers bool
}

var DefaultConfig = &Config{
	Bans:     1,
	Picks:    3,
	RePicks:  3,
	UseTiers: false,
}

func NewConfig() *Config {
	return DefaultConfig
}

func (b *Bot) GetConfigEmbedFields() []*EmbedField {
	banEmbedField := NewEmbedField(
		"`Bans` -- the number of Civs each player gets to ban",
		fmt.Sprintf("**%d**", b.Config.Bans),
	)
	picksEmbedField := NewEmbedField(
		"`Picks` -- the number of Civs each player gets to choose from",
		fmt.Sprintf("**%d**", b.Config.Picks),
	)
	rePicksEmbedField := NewEmbedField(
		"`RePicks` -- the max number of times allowed to re-pick Civs",
		fmt.Sprintf("**%d**", b.Config.RePicks),
	)
	useTiersEmbedField := NewEmbedField(
		"`UseTiers` -- make picks based on FilthyRobot's tier list -- setting this to `true` ensures that each Player gets at minimum one t1/t2 Civ in their list of Picks",
		fmt.Sprintf("**%v**", b.Config.UseTiers),
	)
	return []*EmbedField{banEmbedField, picksEmbedField, rePicksEmbedField, useTiersEmbedField}
}
