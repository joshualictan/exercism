package purchase

import (
	"fmt"
	"strings"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	var required = false

	if kind == "car" || kind == "truck" {
		required = true
	}

	return required
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	var result string
	compare := strings.Compare(option1, option2)
	if compare == -1 {
		result = fmt.Sprintf("%s is clearly the better choice.", option1)
	} else {
		result = fmt.Sprintf("%s is clearly the better choice.", option2)
	}
	return result
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var discountedPrice float64
	if age < 3 {
		discountedPrice = originalPrice * 0.8
	} else if age >= 10 {
		discountedPrice = originalPrice * 0.5
	} else {
		discountedPrice = originalPrice * 0.7
	}
	return discountedPrice
}
