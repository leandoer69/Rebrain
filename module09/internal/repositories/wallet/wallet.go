package wallet

import "Rebrain/module09/internal/entities"

var wallets = map[int]*entities.Wallet{
	1: {ID: 1, Owner: "test_owner_1", Count: 100000},
	2: {ID: 2, Owner: "test_owner_2", Count: 200000},
	3: {ID: 3, Owner: "test_owner_3", Count: 300000},
	4: {ID: 4, Owner: "test_owner_4", Count: 400000},
	5: {ID: 5, Owner: "test_owner_5", Count: 500000},
	6: {ID: 6, Owner: "test_owner_6", Count: 600000},
	7: {ID: 7, Owner: "test_owner_7", Count: 700000},
	8: {ID: 8, Owner: "test_owner_8", Count: 800000},
}

func GetById(id int) *entities.Wallet {
	wallet, ok := wallets[id]
	if !ok {
		return nil
	}

	return wallet
}
