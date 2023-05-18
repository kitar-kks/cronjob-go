package database

import (
	"github.com/dgraph-io/badger/v3"
	"log"
)

type BadgerDBDriver interface {
	Connect() *badger.DB
}

type badgerDBDriver struct {
}

func (b *badgerDBDriver) Connect() *badger.DB {
	// Open the Badger database located in the /tmp/badger directory.
	// It will be created if it doesn't exist.
	db, err := badger.Open(badger.DefaultOptions("./tmp/badger"))
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func NewBadgerDB1Driver() BadgerDBDriver {
	return &badgerDBDriver{}
}
