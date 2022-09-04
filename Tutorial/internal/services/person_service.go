package services

import (
	"Tutorial/internal/dtos"
	. "Tutorial/internal/models"
	. "Tutorial/internal/repository"
	"fmt"
	"github.com/gin-gonic/gin"
)

func NewPersonService(_repository PersonRepository) *PersonService {
	return &PersonService{repository: _repository}
}

type PersonService struct {
	repository PersonRepository
}

func (service PersonService) Create(person *Person) *Person {
	service.repository.Create(person)
	return person
}

func (service PersonService) GetAll(ctx *gin.Context, pagination *dtos.Pagination) dtos.Response {
	operationResult, totalPages := service.repository.GetAll(pagination)
	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*dtos.Pagination)

	//get current url path
	urlPath := ctx.Request.URL.Path

	searchQueryParams := "" //search query params
	for _, search := range pagination.Searches {
		//searchQueryParams += fmt.Sprintf("&%s.%s=%s", search.Column, search.Action, search.Query)
		searchQueryParams += fmt.Sprintf("&%s.%s=%s", search.Column, search.Query)
	}

	//Set first &last page pagination response
	data.FirstPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, 0, pagination.Sort) + searchQueryParams
	data.LastPage = fmt.Sprintf("%s?limit=%d&page=%d&sort=%s", urlPath, pagination.Limit, totalPages, pagination.Sort) + searchQueryParams

	if data.Page > 0 {
		//set previous page pagination response
		data.PreviousPage = fmt.Sprintf("%s?limit=%d&page=%d&sort%s", urlPath, pagination.Limit, data.Page-1, pagination.Sort) + searchQueryParams
	}

	if data.Page < int(totalPages) {
		//set next page pagination response
		data.NextPage = fmt.Sprintf("%s?limit=%d&page=%d&sort%s", urlPath, pagination.Limit, data.Page+1, pagination.Sort) + searchQueryParams

	}

	if data.Page > int(totalPages) {
		//reset previous page
		data.PreviousPage = ""
	}
	return dtos.Response{Success: true, Data: data}
}
