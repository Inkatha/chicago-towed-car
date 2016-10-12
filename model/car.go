package model

// Car represents data about a single car
type Car struct {
	color          string
	make           string
	state          string
	towDate        string
	towPhoneNumber string
	towedToAddress string
}

func (details *Car) Color() string {
	return details.color
}

func (details *Car) Make() string {
	return details.make
}

func (details *Car) State() string {
	return details.state
}

func (details *Car) TowDate() string {
	return details.towDate
}

func (details *Car) TowPhoneNumber() string {
	return details.towPhoneNumber
}

func (details *Car) TowedToAddress() string {
	return details.towedToAddress
}

func (details *Car) SetColor(value string) {
	details.color = value
}

func (details *Car) SetMake(value string) {
	details.make = value
}

func (details *Car) SetState(value string) {
	details.state = value
}

func (details *Car) SetTowDate(value string) {
	details.towDate = value
}

func (details *Car) SetTowPhoneNumber(value string) {
	details.towPhoneNumber = value
}

func (details *Car) SetTowedToAddress(value string) {
	details.towedToAddress = value
}
