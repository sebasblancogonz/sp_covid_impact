package data

// Data struct
type Data struct {
	Cumun      int
	AutCom     string
	ProvCode   int
	Province   string
	Municip    string
	Population int
	Cases      int
	Rate100k   int
	Deaths     int
	Recoveries int
}

// DataList returns a list
type DataList []Data
