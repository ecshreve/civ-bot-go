package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// CivKey is an enum reprsenting an integer key for a civ.
type CivKey int

// Civ consts.
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

// Civ represents an individual civilization.
type Civ struct {
	Key           CivKey
	CivBase       string
	CivAliases    []string
	LeaderBase    string
	LeaderAliases []string
}

func genCivs() []*Civ {
	civs := make([]*Civ, 0)
	for k, c := range civBase {
		civ := &Civ{
			Key:        k,
			CivBase:    c,
			LeaderBase: civLeadersBase[k],
		}
		civs = append(civs, civ)
	}
	return civs
}

func listCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate, cs *CivSession) {
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title: "☁︎  list all possible civs",
		Color: cGREEN,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "All Civs",
				Value: formatCivs(cs.Civs),
			},
		},
	})

	if err != nil {
		fmt.Printf("error generating info: %+v", err)
		return
	}
}
