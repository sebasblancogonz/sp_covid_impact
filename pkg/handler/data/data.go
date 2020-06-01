package data

import (
	"encoding/csv"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	dataObject "github.com/sebasblancogonz/sp_covid_impact/pkg/models/data"
)

// GetAllData will return all data
func GetAllData(c *gin.Context) {
	csvURL, err := http.Get("https://static.dwcdn.net/data/VaMvK.csv")

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error getting all data",
		})
		return
	}

	defer csvURL.Body.Close()

	r := csv.NewReader(csvURL.Body)

	data := dataObject.DataList{}

	csvLines, err := r.ReadAll()

	for _, line := range csvLines[1:] {
		data = append(data, dataObject.Data{
			Cumun:      convertToInt(line[0], c),
			AutCom:     line[1],
			ProvCode:   convertToInt(line[2], c),
			Province:   line[3],
			Municip:    line[4],
			Population: convertToInt(line[5], c),
			Cases:      convertToInt(line[6], c),
			Rate100k:   convertToInt(line[7], c),
			Deaths:     convertToInt(line[8], c),
			Recoveries: convertToInt(line[9], c),
		})
	}

	c.JSON(200, gin.H{
		"data": data,
	})

	return
}

func convertToInt(s string, c *gin.Context) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}

	return i
}
