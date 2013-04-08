package database

import (
	"testing"
	"bytes"
	"github.com/TShadwell/go-useful/errors"
)

var (
	keyone = []byte("Alpha")
	keytwo = []byte("Beta")
	valueone = []byte("x")
	valuetwo = []byte("y")
)

func TestDatabase(t *testing.T) {
	db, err := Init()
	if err != nil{
		t.Fatal("Error whilst loading DB: ", errors.Extend(err))
	}

	err = db.Commit(
		new(Atom).Put(
			keyone,
			valueone,
		).Put(
			keytwo,
			valuetwo,
		),
	)

	if err != nil{
		t.Fatal("Error performing atomic DB write: ", errors.Extend(err))
	}

	v, err := db.Get(
		keyone,
	)
	if err != nil{
		t.Fatal("Error retrieving key one: ", errors.Extend(err))
	}

	if !bytes.Equal(v, valueone){
		t.Fatal("Values stored and retrived are not the same!")
	}

	v, err = db.Get(
		keytwo,
	)

	if err != nil {
		t.Fatal("Error retrieving key two: ", errors.Extend(err))
	}

	if !bytes.Equal(v, valuetwo) {
		t.Fatal("Values stored and retrived are not the same!")
	}

	//Delete the values from the DB.

	err = db.Commit(
		new(Atom).Delete(
			keyone,
		).Delete(
			keytwo,
		),
	)

	if err != nil {
		t.Fatal("Could not delete added keys: ", err)
	}

}
