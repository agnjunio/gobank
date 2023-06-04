package database

import "github.com/agnjunio/gobank/internal/common/models"

func AddAccount(account models.Account) {
	if db == nil {
		panic("DB not initialized")
	}

	txn := db.Txn(true)
	if err := txn.Insert("accounts", account); err != nil {
		panic(err)
	}
	txn.Commit()
}

func GetAccounts() ([]models.Account, error) {
	if db == nil {
		panic("DB not initialized")
	}

	var accounts []models.Account

	txn := db.Txn(false)
	defer txn.Abort()

	it, err := txn.Get("accounts", "id")

	for obj := it.Next(); obj != nil; obj = it.Next() {
		accounts = append(accounts, obj.(models.Account))
	}

	return accounts, err
}
