package database

import (
	"bitbucket.org/kardianos/osext"
	"github.com/TShadwell/go-useful/errors"
	"os"
	"github.com/jmhodges/levigo"
)

const (
	cache       = 500 * Megabyte
	writeBuffer = 100 * Megabyte
)

func (d *Database) Destroy () error {
	return os.Remove(d.Location)
}

func Init() (*Database, error) {
	path, err := osext.ExecutableFolder()
	location := path + "database"
	if err != nil {
		return nil, errors.Extend(err)
	}
	opts := levigo.NewOptions()
	opts.SetCreateIfMissing(true)
	cache := levigo.NewLRUCache(300 * Megabyte)
	opts.SetCache(cache)
	opts.SetWriteBufferSize(writeBuffer)
	db, err := levigo.Open(path, opts)
	if err != nil {
		return nil, errors.Extend(err)
	}
	return &Database{
		DB:           db,
		Cache:        cache,
		Options:      opts,
		ReadOptions:  levigo.NewReadOptions(),
		WriteOptions: levigo.NewWriteOptions(),
		Location:         location,
	}, nil
}
