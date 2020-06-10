package data

// Data struct
type Data struct {
	Cumun      int    `json:"cumun"`
	AutCom     string `json:"autCom"`
	ProvCode   int    `json:"provCode"`
	Province   string `json:"province"`
	Municip    string `json:"municip"`
	Population int    `json:"population"`
	Cases      int    `json:"cases"`
	Rate100k   int    `json:"rate100k"`
	Deaths     int    `json:"deaths"`
	Recoveries int    `json:"recoveries"`
}

// DataList returns a list
type DataList []Data
