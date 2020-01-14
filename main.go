package main

func main() {

}
func fuelConsuption(consuption, fuel int) int {
	const distancePerKm = 100
	const fuelReserveForEmergency = 4

	distance := distancePerKm * fuel / consuption
	distance -= fuelReserveForEmergency * distance / distancePerKm
	return distance

}
