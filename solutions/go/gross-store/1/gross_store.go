package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	foo := map[string]int{"quarter_of_a_dozen": 3, "half_of_a_dozen": 6, "dozen": 12, "small_gross": 120, "gross": 144, "great_gross": 1728}

    return foo
}

// NewBill creates a new bill.
func NewBill() map[string]int {
    bill := make(map[string]int)
	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool { 
	unitValue, exists := units[unit]

    if !exists {
        return false
    }
    currentQuantity, itemExists := bill[item]

    if itemExists{
        bill[item] = currentQuantity + unitValue
    }else{
        bill[item] = unitValue
    }

    return true

}

// RemovcurrentQuantityeItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	currentQuantity, itemExists := bill[item]
    amountToRemove, existsUnit := units[unit]
    
    if !itemExists || !existsUnit{
        return false
    }

    newQuantity := currentQuantity - amountToRemove

    if newQuantity < 0 {
        return false
    }

    if newQuantity == 0 {
        delete(bill, item)
        return true
    }

    bill[item] = newQuantity

    return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	currentQuantity, itemExists := bill[item]

    if !itemExists{
        return 0, false
    }
    return currentQuantity, itemExists
}
