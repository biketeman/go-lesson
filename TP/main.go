package main

//ue struc de nos clients à savoir les comptables

import (
	"fmt"

	"./bank"
	"./bank/qonto"
)

type Accounter struct {
	Id        int
	Name      string
	Customers []*bank.User
}

func main() {

	var userList = []*bank.User{
		&bank.User{
			Id:   3,
			Name: "Michel",
			Accounts: []*bank.Account{
				{Id: 3,
					Balance: 5,
				},
				{
					Id:      6,
					Balance: 455},
			},
		},
		&bank.User{
			Id:   14,
			Name: "Aldo",
			Accounts: []*bank.Account{
				{Id: 3,
					Balance: 5,
				},
				{
					Id:      6,
					Balance: 455},
			},
		},
	}
	for _, u := range userList {
		fmt.Println("lol", u.Id)
		//for _, v := range bank.User.Accounts {
		//	fmt.Println("here are your accounts")
		//}
	}
	//on initialise l'adapter pour qonto
	myqonto, _ := qonto.ConnectToBank(qonto.Qonto{
		ApiKey:   "dfghiey9834y87y5",
		ApiToken: "ihi23y4c98y74983427yc498y37",
	})

	//on appplique la fonction GetDATA de l'interface
	myqonto.GetData(
		&bank.User{
			Id:   3,
			Name: "Jean Paul",
		},
		&bank.Account{
			Id:      14,
			Balance: 14,
		},
	)
	//on appplique la fonction StoreData de l'interface
	myqonto.StoreData(
		&bank.User{
			Id:   7,
			Name: "Paolo",
		},
		&bank.Account{
			Id:      24,
			Balance: -13,
		},
	)
}
