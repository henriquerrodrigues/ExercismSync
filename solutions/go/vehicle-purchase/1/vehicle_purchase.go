package purchase
import "strings"
// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car"{
        return true
    }else{
        if kind == "truck"{
            return true
        }else{
            return false
        }
    }
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
    v1:= " is clearly the better choice."
    resumeopt1 := strings.ReplaceAll(option1, " ", "")
    resumeopt2 := strings.ReplaceAll(option2, " ", "")
	if resumeopt1 < resumeopt2 {
        return option1 + v1
    }else{
        return option2 + v1
    }
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
    var resellPrice float64
	if age < 3 {
        resellPrice= originalPrice * 0.8 
    }else{
        if age >=3 && age < 10 {
            resellPrice = originalPrice * 0.7
        }else{
            resellPrice = originalPrice * 0.5
        }
    }
    return resellPrice
}
