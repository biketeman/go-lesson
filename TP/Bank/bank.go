package bank

//une struc de nos users avec un slice de leurs banques et de leurs comptes ainsi que leur comptable
type User struct {
	Id       int
	Name     string
	Banks    []Banks
	Accounts []*Account
}

type Account struct {
	Id      int
	Balance int
}

//une interface qui récupère les fonctions à appliquer aux banques
type Banks interface {
	GetData(*User, *Account) error
	StoreData(*User, *Account) error
}
