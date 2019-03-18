package animal

import "fmt"

type Dolphin struct {
	Name   string
	Region string
	Age    int
}

func NewDolphin(dolphin Dolphin) (*Dolphin, error) {
	if dolphin.Name == "" {
		return &Dolphin{}, fmt.Errorf("dolphin should have a name")
	}

	if dolphin.Region == "" {
		return &Dolphin{}, fmt.Errorf(dolphin.Name + " comes from nowhere")
	}

	if dolphin.Age < 0 {
		return &Dolphin{}, fmt.Errorf(dolphin.Name + " isn't born yet")
	}

	return &Dolphin{
		Name:   dolphin.Name,
		Region: dolphin.Region,
		Age:    dolphin.Age,
	}, nil
}
