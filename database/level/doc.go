package level

import (
	"errors"
)

/*
	A number of bytes, for sizing the LRU cache.
*/
type BytesSize uint
const (
	Byte = 1 << (10 * iota)
	Kilobyte
	Megabyte
)

//These represent the interfaces to which implimentations must conform
//these will be extended and abstracted by their exported versions
type (
	//A Key, for the database
	Key []byte
	//A Value, for the database
	Value []byte
	closer interface{
		Close()
	}
	options interface{
		SetCreateIfMissing(yes bool)
		SetLRUCache(cache)
		closer
	}
	database interface{
		closer
		Delete(writeOptions, Key)
		Put(writeOptions, Key, Value)
		Write(writeOptions, writeBatch)
	}
	writeOptions interface {
		closer
		SetSync(sync bool)
	}
	readOptions interface{
		closer
		SetVerifyChecksums(yes bool)
	}
	writeBatch interface{
		closer
		Clear()
		Delete(Key)
		Put(Key, Value)
	}
	cache interface {
		closer
	}

)

func newLRUCache(capacity int) cache
func destroyDatabase(name string, o options) error
func repairDatabase(name string, o options) error
func openDatabase(name string, o options) (database, error)
func newOptions() options
func newReadOptions() readOptions
func newWriteOptions() writeOptions
func newWriteBatch() writeBatch

//Define the abtstract implimentations of the interfaces.
type (
	Options struct{
		options
	}
	Cache struct{
		cache
	}
	WriteOptions struct{
		writeOptions
	}
	ReadOptions struct{
		readOptions
	}
	Database struct{
		database
		Cache Cache
		Options Options
		*ReadOptions
		*WriteOptions
	}
)

/*
	=== Options Functions ===
*/

/*
	Gets the underlying implimentation of Options as an interface,
	creating it if it doesn't already exist.
*/
func (o *Options) Inner() options{
	if o == nil{
		o = new(Options)
	}
	if o.options == nil{
		o.options = newOptions()
	}
	return o.options
}

/*
	Function SetCreateIfMissing causes an attempt
	to open a database to also create it if it did not exist.
*/
func (o *Options) SetCreateIfMissing(b bool) *Options{
	o.Inner().SetCreateIfMissing(b)
	return o
}

/*
	Function SetCache sets the cache object for the database
*/
func (o *Options) SetCache(c *Cache) *Options{
	o.Inner().SetLRUCache(c.Inner())
	return o
}

/*
	Function SetCacheSize sets the cache object for the database to a new cache of given size.
*/
func (o *Options) SetCacheSize(size BytesSize) *Options{
	o.SetCache(new(Cache).Size(size))
	return o
}

/*
	=== Cache Functions ===
*/

/*
	Function Inner returns the underlying implimentation of the Cache.

	Unlike other Inner Functions, this may return nil, since LRUCaches must
	be created with given size.

	Therefore, Size(BytesSize) should be called before this.
*/
func (c *Cache) Inner() cache{
	return c.cache
}

/*
	Function Size sets the size of the underlying LRUCache.
*/
func (c *Cache) Size(b BytesSize) *Cache{
	c.cache = newLRUCache(int(b))
	return c
}

/*
	=== Write Options Functions ===  
*/

func (w *WriteOptions) Inner() writeOptions{
	if w.writeOptions == nil{
		w.writeOptions = newWriteOptions()
	}
	return w.writeOptions
}

/*
	Function SetSync sets whether these writes will be flushed
	immediately from the buffer cache. This slows down writes
	but has better crash semantics.
*/
func (w *WriteOptions) SetSync(b bool) *WriteOptions{
	w.Inner().SetSync(b)
	return w
}

/*
	=== Read Options Functions ===  
*/

func (r *ReadOptions) Inner() readOptions{
	if r == nil{
		r = new(ReadOptions)
	}
	if r.readOptions == nil{
		r.readOptions = newReadOptions()
	}
	return r.readOptions
}

func (r *ReadOptions) SetVerifyChecksums(b bool) *ReadOptions{
	r.Inner().SetVerifyChecksums(b)
	return r
}

/*
	=== Database Functions ===  
*/

func (d *Database) Open(location string) error{
	if d.database != nil{
		return Already_Open
	}
	dt, err := openDatabase(location, d.Options.Inner())
	d.database = dt
	return err
}

func (d *Database) Close() {
	d.database.Close()
	d.Cache.Close()
	d.Options.Close()
	d.ReadOptions.Close()
	d.WriteOptions.Close()
}

func (c *Cache) Close(){
	if c != nil && c.cache != nil{
		c.cache.Close()
	}
}

func (o *Options) Close(){
	if o != nil && o.options != nil{
		o.options.Close()
	}
}

func (r *ReadOptions) Close(){
	if r != nil && r.readOptions != nil{
		r.readOptions.Close()
	}
}

func (w *WriteOptions) Close() {
	if w != nil && w.writeOptions != nil{
		w.writeOptions.Close()
	}
}





type Atom struct{
	writeBatch
}
