package models

import (
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

//Model içerisine db işlemleri yapılmamalı
//init fonksiyonunda sadece bir kere yapmamız gereken işleri yapabiliriz
//func init() {
//	config.ConnectDB()
//	db = config.GetDB()
//	err := db.AutoMigrate(&Book{})
//	if err != nil {
//		return
//	}
//}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func (b *Book) GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func (b *Book) GetBookById(id int) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("id=?", id).Find(&getBook)
	return &getBook, db
}

func (b *Book) DeleteBook(id int) *Book {
	var book Book
	db.Where("id=?", id).Delete(&book)
	return &book
}
