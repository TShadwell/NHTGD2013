package database

import (
	"encoding/gob"
	"encoding/binary"
	"bufio"
	"log"
	"github.com/TShadwell/nhtgd2013/twfy"
	D "github.com/TShadwell/nhtgd2013/database"
)

var database = mustDatabase(d.Init())

func mustDatabase(d *D.Database, e error) {
	if err != nil{
		log.Fatal("Error binding database: ", e)
	}
	return d
}

const (
	byteorder = binary.LittleEndian
)

type (
	dbIndex uint
)

const (
	members dbIndex = iota
	names
)

/*
	Converts the index and the ...interface{}
	in order to one large, summed []byte,
	which is returned for use as a DB key.
*/
func dbKey(d dbIndex, i ...interface{}) (k D.Key, err error){
	buf := new(bytes.Buffer)

	for _, v := range i {
		err = binary.Write(buf, byteorder,  v)
		if err != nil{
			return
		}
	}

	k = buf.Bytes()

	return
}

/*

*/
func getFromKey(k D.Key, err error) (D.Value, error){
	if err != nil{
		return nil, err
	}

	return database.Get(k)
}

func readGobFromKey(k D.Key, i interface{}) error{
	b, e := getFromKey(k, err)
	if e != nil{
		return e
	}
	return gob.NewDecoder(bytes.NewBuffer(b)).Decode(i)
}

func writeGobToKey(k D.Key, i interface{}) error{
	
}

func GetMembers() (mems []twfy.Member, err error){
	key, err := dbKey(members)
	if err != nil{
		return nil, err
	}
	err = readGobFromKey(key, mems)
	return
}

func StoreMembers([]twfy.Member) error {
	
}
