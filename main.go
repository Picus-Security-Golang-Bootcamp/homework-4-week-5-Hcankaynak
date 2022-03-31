package main

import (
	postgres "homework4-week5/common/db"
	_ "homework4-week5/domain/author"
	_ "homework4-week5/domain/book"
	"homework4-week5/server"
	"log"
)

func main() {
	// loading db
	_, err := postgres.LoadDB()
	if err != nil {
		log.Fatal(err)
	}

	//// creating book repository
	//bookRepo, err := book.NewBookRepository(db)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// creating author repository
	//authorRepo, err := author.NewAuthorRepository(db)
	//if err != nil {
	//	log.Fatal(err)
	//}

	server.StartServer()

}
