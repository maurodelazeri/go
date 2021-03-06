package main

import (
	"dblogin"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Dog is just a type
type Dog struct {
	Name             string
	Drools           bool
	AlwaysHungry     bool
	HoursSleptPerDay int
}

// main is the entry point for the program.
func main() {

	// Create a mgo session
	session, err := mgo.Dial(dblogin.Userpass) // mongodb://username:yourpasscode@serverip:27017/database?authSource=admin
	if err != nil {
		panic(err)
	}

	// Close the session when the function ends
	defer session.Close()

	fmt.Println("You connected to your mongo database.")

	// Run a command
	var serverStatus = &bson.M{}
	if err := session.Run("serverStatus", serverStatus); err != nil {
		panic(err)
	} else {
		fmt.Println(*serverStatus)
	}

	// Run a command with an argument
	var startupWarnings = &bson.M{}
	if err := session.Run(bson.D{{Name: "getLog", Value: "startupWarnings"}}, startupWarnings); err != nil {
		panic(err)
	} else {
		fmt.Println(*startupWarnings)
	}

	// Insert a document
	var testDoc = bson.M{}
	testDoc["name"] = "John Doe"
	var testID = bson.NewObjectId()
	testDoc["testID"] = testID
	testDoc["slice"] = []string{"one", "two", "three", "four"} // This will be saved as array in database
	session.DB("test").C("people").Insert(&testDoc)

	// Insert a document using marshalling
	session.DB("test").C("people").Insert(&Dog{Name: "Loo", Drools: false, AlwaysHungry: true, HoursSleptPerDay: 18})

	// Read a single document
	var testResultDoc = bson.M{}
	if err := session.DB("test").C("people").Find(bson.M{"testID": testID}).One(&testResultDoc); err != nil {
		if err == mgo.ErrNotFound {
			fmt.Println("The document was not found")
		} else {
			panic(err)
		}
	} else {
		if testResultDoc != nil {
			//fmt.Println(testResultDoc)
		}
	}

	// Read multiple documents
	iteration := session.DB("test").C("people").Find(nil).Iter()
	var result = &bson.M{}
	for iteration.Next(&result) {
		//fmt.Printf("Result: %v\n", *result)
	}
	if iteration.Err() != nil {
		panic(iteration.Err())
	}

}
