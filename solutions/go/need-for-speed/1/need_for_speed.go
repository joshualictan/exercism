package speed

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{speed: speed, battery: 100, batteryDrain: batteryDrain}
}

type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{distance: distance}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	result := car

	if result.battery >= car.batteryDrain {
		result = Car{
			battery:      car.battery - car.batteryDrain,
			batteryDrain: car.batteryDrain,
			speed:        car.speed,
			distance:     car.speed + car.distance,
		}
	}
	return result
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	requiredDrives := track.distance / car.speed
	possibleDrives := car.battery / car.batteryDrain
	return requiredDrives <= possibleDrives
}
