package database

type Account struct {
	Id      int
	Name    string
	Email   string
	Balance float32
}

func AddAccount(account Account) {
	if db == nil {
		panic("DB not initialized")
	}

	txn := db.Txn(true)
	if err := txn.Insert("accounts", account); err != nil {
		panic(err)
	}
	txn.Commit()
}

func GetAccounts() ([]Account, error) {
	if db == nil {
		panic("DB not initialized")
	}

	var accounts []Account

	txn := db.Txn(false)
	defer txn.Abort()

	it, err := txn.Get("accounts", "id")

	for obj := it.Next(); obj != nil; obj = it.Next() {
		accounts = append(accounts, obj.(Account))
	}

	return accounts, err
}
