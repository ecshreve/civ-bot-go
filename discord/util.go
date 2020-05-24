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

var civZig = map[CivKey]string{
	AMERICA:     "https://steamcommunity.com/sharedfiles/filedetails/?id=180689240",
	ARABIA:      "http://steamcommunity.com/sharedfiles/filedetails/?id=253052919",
	ASSYRIA:     "http://steamcommunity.com/sharedfiles/filedetails/?id=172318070",
	AUSTRIA:     "http://steamcommunity.com/sharedfiles/filedetails/?id=326390646",
	AZTECS:      "http://steamcommunity.com/sharedfiles/filedetails/?id=245376405",
	BABYLON:     "http://steamcommunity.com/sharedfiles/filedetails/?id=288312726",
	BRAZIL:      "http://steamcommunity.com/sharedfiles/filedetails/?id=175266266",
	BYZANTIUM:   "http://steamcommunity.com/sharedfiles/filedetails/?id=212340768",
	CARTHAGE:    "http://steamcommunity.com/sharedfiles/filedetails/?id=167304884",
	CELTS:       "http://steamcommunity.com/sharedfiles/filedetails/?id=224148988",
	CHINA:       "http://steamcommunity.com/sharedfiles/filedetails/?id=307785412",
	DENMARK:     "http://steamcommunity.com/sharedfiles/filedetails/?id=195456207",
	EGYPT:       "http://steamcommunity.com/sharedfiles/filedetails/?id=309519240",
	ENGLAND:     "http://steamcommunity.com/sharedfiles/filedetails/?id=265598080",
	ETHIOPIA:    "http://steamcommunity.com/sharedfiles/filedetails/?id=292079843",
	FRANCE:      "http://steamcommunity.com/sharedfiles/filedetails/?id=247998433",
	GERMANY:     "http://steamcommunity.com/sharedfiles/filedetails/?id=289435385",
	GREECE:      "http://steamcommunity.com/sharedfiles/filedetails/?id=228474478",
	HUNS:        "http://steamcommunity.com/sharedfiles/filedetails/?id=160159907",
	INCA:        "http://steamcommunity.com/sharedfiles/filedetails/?id=178525583",
	INDIA:       "http://steamcommunity.com/sharedfiles/filedetails/?id=326395281",
	INDONESIA:   "http://steamcommunity.com/sharedfiles/filedetails/?id=189621929",
	IROQUOIS:    "http://steamcommunity.com/sharedfiles/filedetails/?id=233028579",
	JAPAN:       "http://steamcommunity.com/sharedfiles/filedetails/?id=326386343",
	KOREA:       "http://steamcommunity.com/sharedfiles/filedetails/?id=218395239",
	MAYANS:      "http://steamcommunity.com/sharedfiles/filedetails/?id=171480391",
	MONGOLIA:    "http://steamcommunity.com/sharedfiles/filedetails/?id=226647277",
	MOROCCO:     "http://steamcommunity.com/sharedfiles/filedetails/?id=169423225",
	NETHERLANDS: "http://steamcommunity.com/sharedfiles/filedetails/?id=184478987",
	OTTOMANS:    "http://steamcommunity.com/sharedfiles/filedetails/?id=170395256",
	PERSIA:      "http://steamcommunity.com/sharedfiles/filedetails/?id=301762318",
	POLAND:      "http://steamcommunity.com/sharedfiles/filedetails/?id=199444074",
	POLYNESIA:   "http://steamcommunity.com/sharedfiles/filedetails/?id=166561870",
	PORTUGAL:    "http://steamcommunity.com/sharedfiles/filedetails/?id=250413300",
	ROME:        "http://steamcommunity.com/sharedfiles/filedetails/?id=243501208",
	RUSSIA:      "http://steamcommunity.com/sharedfiles/filedetails/?id=323707819",
	SHOSHONE:    "http://steamcommunity.com/sharedfiles/filedetails/?id=295870476",
	SIAM:        "http://steamcommunity.com/sharedfiles/filedetails/?id=237640016",
	SONGHAI:     "http://steamcommunity.com/sharedfiles/filedetails/?id=259764070",
	SPAIN:       "http://steamcommunity.com/sharedfiles/filedetails/?id=163278284",
	SWEDEN:      "http://steamcommunity.com/sharedfiles/filedetails/?id=164655675",
	VENICE:      "http://steamcommunity.com/sharedfiles/filedetails/?id=197613117",
	ZULUS:       "http://steamcommunity.com/sharedfiles/filedetails/?id=161335260",
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
