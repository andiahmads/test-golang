package game

type DataGameFormater struct {
	ID      int    `json:"id"`
	Mdate   string `json:"mdate"`
	Stadium string `json:"stadium"`
	Team1   string `json:"team_1"`
	Team2   string `json:"team_2"`
}

func FormatDataGame(data DataGame) DataGameFormater {
	formater := DataGameFormater{
		ID:      data.ID,
		Mdate:   data.Mdate,
		Stadium: data.Stadium,
		Team1:   data.Team1,
		Team2:   data.Team2,
	}
	return formater
}
