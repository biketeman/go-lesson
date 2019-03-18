package animals

type Animal struct {
	Name string
	Age  int
}

func Addanimals(animal Animal) (*Animal, error) {

	return &Animal{
		Name: animal.Name,
		Age:  animal.Age,
	}, nil
}
