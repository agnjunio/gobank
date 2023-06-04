package database

import (
	"github.com/agnjunio/gobank/internal/common/models"
	"github.com/hashicorp/go-memdb"
)

var db *memdb.MemDB

func Init() {
	// Create database
	db = NewDatabase()

	// Seed the database
	AddAccount(models.Account{
		Id:      1,
		Name:    "Agnaldo Junior",
		Email:   "agnaldo.junior01@gmail.com",
		Balance: 1.99,
	})
	AddAccount(models.Account{
		Id:      2,
		Name:    "Marcos Schuler",
		Email:   "marbschuler@gmail.com",
		Balance: 100.5,
	})
}

func NewDatabase() *memdb.MemDB {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"accounts": {
				Name: "accounts",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.IntFieldIndex{Field: "Id"},
					},
				},
			},
		},
	}

	db, err := memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}

	return db
}
