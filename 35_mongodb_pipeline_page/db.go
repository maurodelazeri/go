package main

import (
	"dblogin"
	"fmt"
	"gopkg.in/mgo.v2"
)

var (
	// DB bookstore database
	//DB *mgo.Database

	// OS online-store database
	OS *mgo.Database

	// Books collection
	//Books *mgo.Collection

	// Products3 collection
	Products3 *mgo.Collection

	s *mgo.Session
)

func init() {
	// S get a mongodb sessions
	s, err := mgo.Dial(dblogin.Bookstore) // mongodb://username:yourpasscode@serverip:27017/database?authSource=admin
	if err != nil {
		panic(err)
	}

	if err = s.Ping(); err != nil {
		panic(err)
	}

	//DB = s.DB("bookstore")
	//Books = DB.C("books")

	OS = s.DB("onlinestore")
	Products3 = OS.C("products3")

	fmt.Println("You connected to your mongo database.")

}
