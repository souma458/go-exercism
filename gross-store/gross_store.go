package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	unitsMap := make(map[string]int)
	unitsMap["quarter_of_a_dozen"] = 3
	unitsMap["half_of_a_dozen"] = 6
	unitsMap["dozen"] = 12
	unitsMap["small_gross"] = 120
	unitsMap["gross"] = 144
	unitsMap["great_gross"] = 1728
	return unitsMap
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	amount, exists := units[unit]
	if !exists {
		return false
	}
	bill[item] += amount
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	currentAmount, itemExits := bill[item]
	amount, unitExits := units[unit]
	newAmount := currentAmount - amount
	if !itemExits || !unitExits || newAmount < 0 {
		return false
	}

	if newAmount == 0 {
		delete(bill, item)
	} else {
		bill[item] = newAmount
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	amount, exists := bill[item]
	return amount, exists
}
