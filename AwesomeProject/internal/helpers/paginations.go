package helpers

import (
	"AwesomeProject/internal/dtos"
	"net/http"
	"strconv"
)

var (
	errMessage   = "parse_error"
	notConverted = NewError("parse_error", http.StatusBadRequest)
)

func GeneratePaginationRequest(w http.ResponseWriter, r *http.Request) (*dtos.Pagination, error) {
	//default limit,page & sort parameter
	var err error
	limit := 10
	page := 1
	sort := "asc"
	start := 0
	orderBy := "created_at"

	var searches []dtos.Search
	query := r.URL.Query()

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
			search := dtos.Search{Column: key, Query: queryValue}
			searches = append(searches, search)
		}

		//if strings.Contains(key, ".") {
		//	//split query parameter key by dot
		//	searchkeys := strings.Split(key, ".")
		//
		//	//create seaRCH OBJECT
		//	search := Search{Column: searchkeys[0], Query: queryValue}
		//
		//	//add search object to search array
		//	searches = append(searches, search)
		//}

	}

	return &dtos.Pagination{Limit: limit, Page: page, Sort: sort, Searches: searches, Start: start, OrderBy: orderBy}, nil

}
