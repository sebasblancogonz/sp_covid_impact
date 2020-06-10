package helpers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sebasblancogonz/sp_covid_impact/pkg/models/pagination"
)

// GeneratePaginationRequest function
func GeneratePaginationRequest(c *gin.Context) *pagination.Pagination {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))

	return &pagination.Pagination{Limit: limit, Page: page}
}
