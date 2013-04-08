package database

import (
	"github.com/jmhodges/levigo"
)

type (
	Database struct {
		*levigo.DB
		Cache   *levigo.Cache
		Options *levigo.Options
		*levigo.ReadOptions
		*levigo.WriteOptions
		Location string
	}
	/*
		A set of atomic writes and deletions;
		if one fails, none of them make a change
		to the database.

			x := new(database.Atom).Put(
				[]byte("tin"),
				[]byte("Beans"),
			).Delete(
				[]byte("washing up"),
			)
			defer x.Close()

			db.Write(x)

		or:

			db.Commit(new(databse.Atom).Put(
				[]byte("tin"),
				[]byte("beans"),
			).Delete([]byte("washing up"))

	*/
	Atom struct {
		*levigo.WriteBatch
	}
	Key   []byte
	Value []byte
)

func (d *Database) Close() {
	/*
		This is C doing this.
	*/
	d.DB.Close()
	d.Cache.Close()
	d.Options.Close()
	d.ReadOptions.Close()
	d.WriteOptions.Close()
}

func (d *Database) Get(k Key) (Value, error) {
	return d.DB.Get(d.ReadOptions, k)
}

func (d *Database) Delete(k Key) error {
	return d.DB.Delete(d.WriteOptions, k)
}

func (d *Database) Put(k Key, v Value) error {
	return d.DB.Put(d.WriteOptions, k, v)
}

func (d *Database) Write(a *Atom) error {
	return d.DB.Write(d.WriteOptions, a.WriteBatch)
}

/*
	Write an Atom to the Database, closing it afterward.
*/
func (d *Database) Commit(a *Atom) error {
	defer a.Close()
	return d.Write(a)
}

func (a *Atom) getBatch() *levigo.WriteBatch {
	if a.WriteBatch == nil {
		a.WriteBatch = levigo.NewWriteBatch()
	}
	return a.WriteBatch
}

func (a *Atom) Clear() *Atom {
	a.getBatch().Clear()
	return a
}

func (a *Atom) Close() *Atom {
	a.getBatch().Close()
	return a
}

func (a *Atom) Delete(k Key) *Atom {
	a.getBatch().Delete(k)
	return a
}

func (a *Atom) Put(k Key, v Value) *Atom {
	a.getBatch().Put(k, v)
	return a
}

const (
	Byte = 1 << (10 * iota)
	Kilobyte
	Megabyte
	Gigabyte
	Terabyte
)
