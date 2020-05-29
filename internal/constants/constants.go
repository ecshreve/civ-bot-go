package constants

// Color Enum.
const (
	ColorDEFAULT           = 0
	ColorAQUA              = 1752220
	ColorGREEN             = 3066993
	ColorBLUE              = 3447003
	ColorPURPLE            = 10181046
	ColorGOLD              = 15844367
	ColorORANGE            = 15105570
	ColorRED               = 15158332
	ColorGREY              = 9807270
	ColorDARKERGREY        = 8359053
	ColorNAVY              = 3426654
	ColorDARKAQUA          = 1146986
	ColorDARKGREEN         = 2067276
	ColorDARKBLUE          = 2123412
	ColorDARKPURPLE        = 7419530
	ColorDARKGOLD          = 12745742
	ColorDARKORANGE        = 11027200
	ColorDARKRED           = 10038562
	ColorDARKGREY          = 9936031
	ColorLIGHTGREY         = 12370112
	ColorDARKNAVY          = 2899536
	ColorLUMINOUSVIVIDPINK = 16580705
	ColorDARKVIVIDPINK     = 12320855
)

// NumEmojiMap is a map of integers to their string emoji equivalents.
var NumEmojiMap = map[int]string{
	0: "0️⃣",
	1: "1️⃣",
	2: "2️⃣",
	3: "3️⃣",
	4: "4️⃣",
	5: "5️⃣",
}

// EmojiNumMap is a map of string emojis to their integer equivalents.
var EmojiNumMap = map[string]int{
	"0️⃣": 0,
	"1️⃣": 1,
	"2️⃣": 2,
	"3️⃣": 3,
	"4️⃣": 4,
	"5️⃣": 5,
}

// CivKey represents an integer key for a Civ.
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

// CivKeys is a slice of the defined CivKey values, this makes it easier for us
// to iterate over all CivKeys.
var CivKeys = []CivKey{
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

// CivBase is a map of CivKey to string name of that Civ.
var CivBase = map[CivKey]string{
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

// CivLeaders is a map of CivKey to string representation of that Civ's leader.
var CivLeaders = map[CivKey]string{
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

// CivZig is a map of CivKey to the URL of the ZigZagal guide for the Civ.
var CivZig = map[CivKey]string{
	AMERICA:     "http://steamcommunity.com/sharedfiles/filedetails/?id=180689240",
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
