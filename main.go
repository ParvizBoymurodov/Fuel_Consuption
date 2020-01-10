package main

func main() {

}
func fuelConsuption(consuption, fuel int) int {
	const distancePerKm = 100
	const FuelReserveForEmergency = 4

	distance := distancePerKm * fuel / consuption
	distance -= FuelReserveForEmergency * distance / distancePerKm
	return distance

}
