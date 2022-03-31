package book

import (
	"fmt"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Id         int    `json:"id"`
	Name       string `json:"name"`
	PageNumber int    `json:"pageNumber"`
	Stock      int    `json:"stock"`
	Price      int    `json:"price"`
	StockCode  string `json:"stockCode"`
	ISBN       string `json:"isbn"`
	AuthorId   int    `json:"authorId"`
}

func (a *Book) ToString() string {
	return fmt.Sprintf("Id: %d \n"+"Name: %s \n"+"PageNumber: %d \n"+"Stock: %d \n"+"Price: %d \n"+
		"StockCode: %s \n"+"ISBN: %s \n"+"AuthorId: %d \n",
		a.Id, a.Name, a.PageNumber, a.Stock, a.Price, a.StockCode, a.ISBN, a.AuthorId)
}
