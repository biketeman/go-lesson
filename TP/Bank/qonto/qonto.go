package qonto

import (
	"fmt"

	"../../bank"
)

type Qonto struct {
	ApiKey   string
	ApiToken string
}

func ConnectToBank(qonto Qonto) (bank.Banks, error) {
	return &Qonto{
		ApiKey:   qonto.ApiKey,
		ApiToken: qonto.ApiToken,
	}, nil
}

func (qonto *Qonto) GetData(user *bank.User, account *bank.Account) error {
	fmt.Println("get all data from User", user.Id, "with account", account.Id)
	return nil
}

func (qonto *Qonto) StoreData(user *bank.User, account *bank.Account) error {
	fmt.Println("User number ", user.Id, "now has a balance of", account.Balance)
	return nil
}
