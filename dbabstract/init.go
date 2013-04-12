package database

//On init, reload all members
import (
	"bitbucket.org/kardianos/osext"
	"github.com/TShadwell/NHTGD2013/database/level"
	"log"
)

func init() {
	path, err := osext.ExecutableFolder()
	database, err = new(level.Database).SetOptions(
		new(level.Options).SetCreateIfMissing(
			true,
		).SetCacheSize(
			500 * level.Megabyte,
		),
	).OpenDB(path + "/leveldb/")

	if err != nil {
		log.Fatal("Error binding database: ", err)
	}
}
