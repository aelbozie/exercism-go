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
	valueUnit, okUnit := units[unit]

	if !okUnit {
		return false
	} else {
		bill[item] += valueUnit
		return true
	}
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill map[string]int, units map[string]int, item string, unit string) bool {
	valueUnit, okUnit := units[unit]
	valueBill, okBill := bill[item]
	newValueBill := valueBill - valueUnit
	if !okUnit || !okBill || newValueBill < 0 {
		return false
	}

	bill[item] = newValueBill

	if bill[item] == 0 {
		delete(bill, item)
	}
	return true

}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	valueBill, okBill := bill[item]
	if !okBill {
		return valueBill, false
	}

	return valueBill, true

}
