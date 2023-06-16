package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := map[string]int{}

	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill map[string]int, units map[string]int, item string, unit string) bool {
	if units[unit] == 0 {
		return false
	} else {
		bill[item] += units[unit]
		return true
	}
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill map[string]int, units map[string]int, item string, unit string) bool {
	if bill[item] == 0 || units[unit] == 0 || bill[item] < units[unit] {
		return false
	}

	bill[item] -= units[unit]

	if bill[item] == 0 {
		delete(bill, item)
	}
	return true

}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	quantity := bill[item]
	if quantity == 0 {
		return quantity, false
	}

	return quantity, true

}
