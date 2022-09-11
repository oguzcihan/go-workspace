package repository

import (
	"Tutorial/internal/dtos"
	. "Tutorial/internal/models"
	"fmt"
	"gorm.io/gorm"
	"math"
)

func NewPersonRepository(database *gorm.DB) *PersonRepository {
	//error olmalı
	return &PersonRepository{DB: database}
}

type PersonRepository struct {
	DB *gorm.DB
}

func (u PersonRepository) Create(person *Person) *Person {
	u.DB.Create(&person)
	return person
}

func (u PersonRepository) GetAll(pagination *dtos.Pagination) (RepositoryResult, int64) {
	var persons []Person
	var totalRows, totalPages, fromRow, toRow int64 = 0, 0, 0, 0
	//var offset int = 0
	offset := (pagination.Page * pagination.Limit) - pagination.Limit //verileri bu sonuçtan sonra almyaa başlar örn 10 dan başla almaya
	//offset = offset + pagination.Limit
	//get data with limit,offset &order
	find := u.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	//find := u.DB.Limit(pagination.Limit).Order(pagination.Sort)

	//generate where query
	searhcs := pagination.Searches
	if searhcs != nil {
		for _, value := range searhcs {
			column := value.Column
			//action := value.Action
			query := value.Query

			whereQuery := fmt.Sprintf("%s=?", column)
			find = find.Where(whereQuery, query)

			//whereQuery: endpointten gelen key
			//switch action {
			//case "equals":
			//	whereQuery := fmt.Sprintf("%s=?", column)
			//	find = find.Where(whereQuery, query)
			//	break
			//case "contains":
			//	whereQuery := fmt.Sprintf("%s LIKE ?", column)
			//	find = find.Where(whereQuery, "%"+query+"%")
			//	break
			//case "in":
			//	whereQuery := fmt.Sprintf("%s IN ?", column)
			//	queryArray := strings.Split(query, ",")
			//	find = find.Where(whereQuery, queryArray)
			//
			//}

		}

	}

	find = find.Find(&persons)

	//has error find data
	errFind := find.Error
	if errFind != nil {
		return RepositoryResult{Error: errFind}, totalPages
	}

	pagination.Rows = persons

	//count all data
	errCount := u.DB.Model(&Person{}).Count(&totalRows).Error
	if errCount != nil {
		return RepositoryResult{Error: errCount}, totalPages
	}

	pagination.TotalRows = totalRows

	// calculate total pages
	totalPages = int64(int(math.Ceil(float64(totalRows)/float64(pagination.Limit))) - 1)

	if pagination.Page == 0 {
		// set from & to row on first page
		fromRow = 1
		toRow = int64(pagination.Limit)
	} else {
		if int64(pagination.Page) <= totalPages {
			// calculate from & to row
			fromRow = int64(pagination.Page*pagination.Limit + 1)
			toRow = int64((pagination.Page + 1) * pagination.Limit)
		}
	}

	if toRow > totalRows {
		// set to row with total rows
		toRow = totalRows
	}

	pagination.FromRow = fromRow
	pagination.ToRow = toRow

	return RepositoryResult{Result: pagination}, totalPages
}
