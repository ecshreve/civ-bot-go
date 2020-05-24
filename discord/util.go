package discord

import "github.com/bwmarrin/discordgo"

// Color Enum.
const (
	cDEFAULT           = 0
	cAQUA              = 1752220
	cGREEN             = 3066993
	cBLUE              = 3447003
	cPURPLE            = 10181046
	cGOLD              = 15844367
	cORANGE            = 15105570
	cRED               = 15158332
	cGREY              = 9807270
	cDARKERGREY        = 8359053
	cNAVY              = 3426654
	cDARKAQUA          = 1146986
	cDARKGREEN         = 2067276
	cDARKBLUE          = 2123412
	cDARKPURPLE        = 7419530
	cDARKGOLD          = 12745742
	cDARKORANGE        = 11027200
	cDARKRED           = 10038562
	cDARKGREY          = 9936031
	cLIGHTGREY         = 12370112
	cDARKNAVY          = 2899536
	cLUMINOUSVIVIDPINK = 16580705
	cDARKVIVIDPINK     = 12320855
)

// CivKey is an enum reprsenting an integer key for a civ.
type CivKey int

// Consts defining the CivKeys this bot works with.
const (
	AMERICA CivKey = iota
	ARABIA
	ASSYRIA
	AUSTRIA
	AZTECS
	BABYLON
	BRAZIL
	BYZANTIUM
	CARTHAGE
	CELTS
	CHINA
	DENMARK
	EGYPT
	ENGLAND
	ETHIOPIA
	FRANCE
	GERMANY
	GREECE
	HUNS
	INCA
	INDIA
	INDONESIA
	IROQUOIS
	JAPAN
	KOREA
	MAYANS
	MONGOLIA
	MOROCCO
	NETHERLANDS
	OTTOMANS
	PERSIA
	POLAND
	POLYNESIA
	PORTUGAL
	ROME
	RUSSIA
	SHOSHONE
	SIAM
	SONGHAI
	SPAIN
	SWEDEN
	VENICE
	ZULUS
)

// civKeys is a slice of the defined CivKey values, this makes it easier for us
// to iterate over all CivKeys.
var civKeys = []CivKey{
	AMERICA,
	ARABIA,
	ASSYRIA,
	AUSTRIA,
	AZTECS,
	BABYLON,
	BRAZIL,
	BYZANTIUM,
	CARTHAGE,
	CELTS,
	CHINA,
	DENMARK,
	EGYPT,
	ENGLAND,
	ETHIOPIA,
	FRANCE,
	GERMANY,
	GREECE,
	HUNS,
	INCA,
	INDIA,
	INDONESIA,
	IROQUOIS,
	JAPAN,
	KOREA,
	MAYANS,
	MONGOLIA,
	MOROCCO,
	NETHERLANDS,
	OTTOMANS,
	PERSIA,
	POLAND,
	POLYNESIA,
	PORTUGAL,
	ROME,
	RUSSIA,
	SHOSHONE,
	SIAM,
	SONGHAI,
	SPAIN,
	SWEDEN,
	VENICE,
	ZULUS,
}

// civBase is a map of CivKey to string representation of that Civ.
var civBase = map[CivKey]string{
	AMERICA:     "america",
	ARABIA:      "arabia",
	ASSYRIA:     "assyria",
	AUSTRIA:     "austria",
	AZTECS:      "aztecs",
	BABYLON:     "babylon",
	BRAZIL:      "brazil",
	BYZANTIUM:   "byzantium",
	CARTHAGE:    "carthage",
	CELTS:       "celts",
	CHINA:       "china",
	DENMARK:     "denmark",
	EGYPT:       "egypt",
	ENGLAND:     "england",
	ETHIOPIA:    "ethiopia",
	FRANCE:      "france",
	GERMANY:     "germany",
	GREECE:      "greece",
	HUNS:        "huns",
	INCA:        "inca",
	INDIA:       "india",
	INDONESIA:   "indonesia",
	IROQUOIS:    "iroquois",
	JAPAN:       "japan",
	KOREA:       "korea",
	MAYANS:      "mayans",
	MONGOLIA:    "mongolia",
	MOROCCO:     "morocco",
	NETHERLANDS: "netherlands",
	OTTOMANS:    "ottomans",
	PERSIA:      "persia",
	POLAND:      "poland",
	POLYNESIA:   "polynesia",
	PORTUGAL:    "portugal",
	ROME:        "rome",
	RUSSIA:      "russia",
	SHOSHONE:    "shoshone",
	SIAM:        "siam",
	SONGHAI:     "songhai",
	SPAIN:       "spain",
	SWEDEN:      "sweden",
	VENICE:      "venice",
	ZULUS:       "zulu",
}

// civLeaderBase is a map of CivKey to string representation of that Civ's leader.
var civLeadersBase = map[CivKey]string{
	AMERICA:     "washington",
	ARABIA:      "harun al-rashad",
	ASSYRIA:     "ashurbanipal",
	AUSTRIA:     "maria theresa",
	AZTECS:      "montezuma",
	BABYLON:     "nebuchadnezzar ii",
	BRAZIL:      "pedro ii",
	BYZANTIUM:   "theodora",
	CARTHAGE:    "dido",
	CELTS:       "boudicca",
	CHINA:       "wu zetian",
	DENMARK:     "harald bluetooth",
	EGYPT:       "ramesse ii",
	ENGLAND:     "elizabeth",
	ETHIOPIA:    "haile selassie",
	FRANCE:      "napolean",
	GERMANY:     "bismarck",
	GREECE:      "alexander",
	HUNS:        "attila",
	INCA:        "pachacuti",
	INDIA:       "gandhi",
	INDONESIA:   "gajah mada",
	IROQUOIS:    "hiawatha",
	JAPAN:       "oda nobunaga",
	KOREA:       "sejong",
	MAYANS:      "pacal",
	MONGOLIA:    "genghis khan",
	MOROCCO:     "ahmad al-mansur",
	NETHERLANDS: "william",
	OTTOMANS:    "suleiman",
	PERSIA:      "darius i",
	POLAND:      "casimir",
	POLYNESIA:   "kamehameha",
	PORTUGAL:    "maria i",
	ROME:        "augustus",
	RUSSIA:      "catherine",
	SHOSHONE:    "pocatello",
	SIAM:        "ramkhamhaeng",
	SONGHAI:     "askia",
	SPAIN:       "isabella",
	SWEDEN:      "gustavus adolphus",
	VENICE:      "enrico dandolo",
	ZULUS:       "shaka",
}

// isBotReaction checks if users reaction is one preset by the bot.
func isBotReaction(s *discordgo.Session, reactions []*discordgo.MessageReactions, emoji *discordgo.Emoji) bool {
	for _, r := range reactions {
		if r.Emoji.Name == emoji.Name && r.Me {
			return true
		}
	}

	return false
}
