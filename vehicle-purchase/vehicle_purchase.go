// package purchase implements three tasks related to purchase of a vehicle.
package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
    return (kind == "car") || (kind == "truck")
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
    betterChoice := option1
    if option1 > option2 {
        betterChoice = option2
    }
    return betterChoice + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
    rate := 0.7;
    if age < 3 {
        rate = 0.8;
    } else if age >= 10 {
        rate = 0.5;
    }
    return originalPrice * rate; 
}
