package utils

import (
	. "ExampleProject/internal/dtos"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	errMessage   = "parse_error"
	notConverted = NewError("parse_error", http.StatusBadRequest)
)

func GeneratePaginationRequest(context *gin.Context) (*Pagination, error) {
	//default limit,page & sort parameter
	var err error
	limit := 10
	page := 1
	sort := "asc"
	start := 0
	orderBy := "created_at"

	var filters []Filter
	query := context.Request.URL.Query()

	for key, value := range query {
		queryValue := value[len(value)-1]

		switch key {
		case "limit":
			limit, err = strconv.Atoi(queryValue)
			if err != nil {
				Logger.Error(errMessage)
				return nil, notConverted
			}
			break
		case "page":
			page, err = strconv.Atoi(queryValue)
			if err != nil {
				Logger.Error(errMessage)
				return nil, notConverted
			}
			break
		case "start":
			start, err = strconv.Atoi(queryValue)
			if err != nil {
				Logger.Error(errMessage)
				return nil, notConverted
			}
			break
		case "order_by":
			orderBy = queryValue
			break
		case "sort_by":
			sort = queryValue
			break

		default:
			search := Filter{Column: key, Query: queryValue}
			filters = append(filters, search)
		}

	}

	return &Pagination{Limit: limit, Page: page, Sort: sort, Filters: filters, Start: start, OrderBy: orderBy}, nil

}
