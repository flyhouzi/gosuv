package badger

import (
	"log"
	"testing"

	badger "github.com/dgraph-io/badger"
)

func TestBager(t *testing.T) {
	db, err := badger.Open(badger.DefaultOptions(".gosuv/badger"))
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
