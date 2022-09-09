package utils

import (
	. "ExampleProject/internal/dtos"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GeneratePaginationRequest(context *gin.Context) *UserList {
	//default limit,page & sort parameter
	limit := 10
	page := 1
	sort := "created_at asc"

	var searches []Search
	query := context.Request.URL.Query()

	for key, value := range query {
		queryValue := value[len(value)-1]

		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort_by":
			sort = queryValue
			break

		default:
			search := Search{Column: key, Query: queryValue}
			searches = append(searches, search)
		}

	}

	return &UserList{Limit: limit, Page: page, Sort: sort, Searches: searches}

}
