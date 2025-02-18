package interfaces

type Dog struct{}

func (d *Dog) MakeSound() string {
	return "Au! Au!"
}