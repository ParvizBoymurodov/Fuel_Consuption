package main

func main() {

}
func fuelConsuption(consuption, fuel int) int {
	const distancePerKm = 100
	const fuelReserveForEmergency = 4

	distance := distancePerKm * fuel / consuption
	reserve := fuelReserveForEmergency * distance / distancePerKm
	distance -= reserve
	return distance

}
