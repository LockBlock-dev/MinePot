package typings

type Version struct {
    Name     string `json:"name"`
    Protocol int    `json:"protocol"`
}

type PlayersInfo struct {
	Max    int           `json:"max"`
	Online int           `json:"online"`
	Sample []Player      `json:"sample"`
}

type Player struct {
    Name string `json:"name"`
    Id   string `json:"id"`
}

type ServerStatus struct {
    Version             Version       `json:"version"`
    Players             Players       `json:"players"`
    Description         string        `json:"description"`
    Favicon             string        `json:"favicon"`
    EnforcesSecureChat  bool          `json:"enforcesSecureChat"`
}

type Players struct {
    Max    int      `json:"max"`
    Online int      `json:"online"`
    Sample []Player `json:"sample"`
}
