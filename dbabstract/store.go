package database

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	D "github.com/TShadwell/NHTGD2013/database"
	"github.com/TShadwell/NHTGD2013/markov"
	"github.com/TShadwell/NHTGD2013/twfy"
)

var database *D.Database

var (
	byteorder = binary.LittleEndian
)

type (
	dbIndex uint
)

const (
	members dbIndex = iota
	chains
)

/*
	Converts the index and the ...interface{}
	in order to one large, summed []byte,
	which is returned for use as a DB key.
*/
func dbKey(d dbIndex, i ...interface{}) (k D.Key, err error) {
	var buf bytes.Buffer

	for _, v := range i {
		err = binary.Write(&buf, byteorder, v)
		if err != nil {
			return
		}
	}

	k = buf.Bytes()

	return
}

func getFromKey(k D.Key, err error) (D.Value, error) {
	if err != nil {
		return nil, err
	}

	return database.Get(k)
}

func readGobFromKey(k D.Key, i interface{}) error {
	b, e := getFromKey(k, nil)
	if e != nil {
		return e
	}

	if b == nil {
		return nil
	}
	return gob.NewDecoder(bytes.NewBuffer(b)).Decode(i)
}

func writeGobToKey(k D.Key, i interface{}) (err error) {
	var buf bytes.Buffer
	err = gob.NewEncoder(&buf).Encode(i)
	if err != nil {
		return
	}
	err = database.Put(k, buf.Bytes())
	return

}

func GetMembers() (mems []twfy.Member, err error) {
	key, err := dbKey(members)
	if err != nil {
		return nil, err
	}
	err = readGobFromKey(key, mems)
	return
}

func storeMembers(ms []twfy.Member) (err error) {
	key, err := dbKey(members)
	if err != nil {
		return
	}

	return writeGobToKey(key, ms)
}

/*
	StoreChain stores a markov chain corresponding
	to a member in our K/V store.
*/
func StoreChain(m markov.Chain, p twfy.PersonID) (err error) {
	key, err := dbKey(
		chains,
		p,
	)
	if err != nil {
		return err
	}
	return writeGobToKey(key, m)
}

/*
	Retrieves a stored markov chain from the database.
	A nil chain is returned if it does not exist.
*/
func RetrieveChain(p twfy.PersonID) (m *markov.Chain, err error) {
	key, err := dbKey(
		chains,
		p,
	)

	if err != nil {
		return
	}

	err = readGobFromKey(key, &m)
	return
}
