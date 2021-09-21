// Package gross implements routines used for a POS system
package gross

// Units store the Gross Store unit measurement
func Units() map[string]int {
    units := map[string]int {
        "quarter_of_a_dozen": 3,
        "half_of_a_dozen": 6,
        "dozen": 12,
        "small_gross": 120,
        "gross": 144,
        "great_gross": 1728,
    }
    return units
}

// NewBill create a new bill
func NewBill() map[string]int {
    return make(map[string]int)
}

// AddItem add item to customer bill
func AddItem(bill, units map[string]int, item, unit string) bool {
    value, exists := units[unit]
    if exists {
        bill[item] = value
        return true
    }
    return false
}

// RemoveItem remove item from customer bill
func RemoveItem(bill, units map[string]int, item, unit string) bool {
    itemValue, itemExists := bill[item]
    if !itemExists {
        return false
    }
    unitValue, unitExists := units[unit]
    if !unitExists {
        return false
    }
    if itemValue < unitValue {
        return false
    } else if itemValue == unitValue {
        delete(bill, item)
    } else {
        bill[item] -= unitValue
    }
    return true
}

// GetItem return the quantity of item that the customer has in his/her bill
func GetItem(bill map[string]int, item string) (int, bool) {
    itemValue, itemExists := bill[item]
    if !itemExists {
        return 0, false
    }
    return itemValue, true
}
