package database

import (
	"fmt"

	"github.com/agnjunio/gobank/internal/common/models"
)

func AddAccount(account *models.Account) error {
	if db == nil {
		return fmt.Errorf("DB not initialized")
	}

	txn := db.Txn(true)

	// Increment ID by one (needed because in-memory db does not support auto index)
	accounts, err := GetAccounts()
	if err != nil {
		return err
	}
	if len(accounts) == 0 {
		account.Id = 1
	} else {
		account.Id = accounts[len(accounts)-1].Id + 1
	}

	if err := txn.Insert("accounts", *account); err != nil {
		return err
	}
	txn.Commit()

	return nil
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
