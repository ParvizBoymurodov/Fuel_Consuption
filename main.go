package main

func main() {

}
func fuelConsuption(consuption, fuel int) int {
	const distancePer_km = 100
	const percentageOfFuel = 4

	distance := distancePer_km * fuel / consuption
	distance -= percentageOfFuel * distance / distancePer_km
	return distance

}
