package model

// Car represents data about a single car
type Car struct {
	Color          string
	Make           string
	Plate          string
	State          string
	TowDate        string
	TowPhoneNumber string
	TowedToAddress string
}

func (car *Car) color() string {
	return car.Color
}

func (car *Car) make() string {
	return car.Make
}

func (car *Car) plate() string {
	return car.Plate
}

func (car *Car) state() string {
	return car.State
}

func (car *Car) towDate() string {
	return car.TowDate
}

func (car *Car) towPhoneNumber() string {
	return car.TowPhoneNumber
}

func (car *Car) towedToAddress() string {
	return car.TowedToAddress
}

func (car *Car) setColor(value string) {
	car.Color = value
}

func (car *Car) setMake(value string) {
	car.Make = value
}

func (car *Car) setPlate(value string) {
	car.Plate = value
}

func (car *Car) setState(value string) {
	car.State = value
}

func (car *Car) setTowDate(value string) {
	car.TowDate = value
}

func (car *Car) setTowPhoneNumber(value string) {
	car.TowPhoneNumber = value
}

func (car *Car) setTowedToAddress(value string) {
	car.TowedToAddress = value
}

// NewCar creates a new car
func NewCar() *Car {
	car := new(Car)
	car.Color = ""
	car.Make = ""
	car.Plate = ""
	car.State = ""
	car.TowDate = ""
	car.TowPhoneNumber = ""
	car.TowedToAddress = ""
	return car
}

// SearchByLicensePlate makes a call to get all of the towed cars in the last 7 days,
// Then searches the returned data for a matching license plate.
func SearchByLicensePlate(licensePlate string) *Car {

	searchedCar := new(Car)

	carData := GetTowedCarData()

	for _, car := range carData {
		if car.Plate == licensePlate {
			searchedCar.setColor(car.Color)
			searchedCar.setMake(car.Make)
			searchedCar.setPlate(car.Plate)
			searchedCar.setState(car.State)
			searchedCar.setTowDate(car.TowDate)
			searchedCar.setTowPhoneNumber(car.TowPhoneNumber)
			searchedCar.setTowedToAddress(car.TowedToAddress)
		}
	}

	return searchedCar
}
