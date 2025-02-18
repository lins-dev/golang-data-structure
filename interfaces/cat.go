package interfaces

type Cat struct{}

func (d *Cat) MakeSound() string {
	return "Miau!"
}