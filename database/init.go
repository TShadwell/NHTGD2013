package database

import (
	"bitbucket.org/kardianos/osext"
	"github.com/TShadwell/go-useful/errors"
	"github.com/jmhodges/levigo"
)

const (
	cache       = 500 * Megabyte
	writeBuffer = 100 * Megabyte
)

func Init() (*Database, error) {
	path, err := osext.ExecutableFolder()
	if err != nil {
		return nil, errors.Extend(err)
	}
	opts := levigo.NewOptions()
	opts.SetCreateIfMissing(true)
	cache := levigo.NewLRUCache(300 * Megabyte)
	opts.SetCache(cache)
	opts.SetWriteBufferSize(writeBuffer)
	db, err := levigo.Open(path+"/database", opts)
	if err != nil {
		return nil, errors.Extend(err)
	}
	return &Database{
		DB:           db,
		Cache:        cache,
		Options:      opts,
		ReadOptions:  levigo.NewReadOptions(),
		WriteOptions: levigo.NewWriteOptions(),
	}, nil
}
