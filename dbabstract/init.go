package database

//On init, reload all members
import (
	D "github.com/TShadwell/nhtgd2013/database"
	"github.com/TShadwell/nhtgd2013/secrets"
	"github.com/TShadwell/nhtgd2013/twfy"
	"log"
)

func init() {

	database, err := D.Init()

	if err != nil {
		log.Fatal("Error binding database: ", err)
	}

	API := twfy.API{
		Key: secrets.TWFYKey,
	}

	ms, err := API.GetMembers()

	if err != nil {
		log.Fatal("Error getting members: ", err)
	}

	err = storeMembers(ms)
	log.Println("Members stored.")
}
