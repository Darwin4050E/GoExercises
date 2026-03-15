// Package purchase helps you prepare to buy a vehicle.
package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	result := option1
	if option2 < option1 {
		result = option2
	}
	return result + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	result := float64(0)
	if age < 3 {
		result = originalPrice * 0.8
	} else if age >= 3 && age < 10 {
		result = originalPrice * 0.7
	} else {
		result = originalPrice * 0.5
	}
	return result
}
