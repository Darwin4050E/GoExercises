// Package gross implements a point of sale (POS) system for a store called Gross Store.
package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	value_item, exists_item := bill[item]
	value_unit, exists_unit := units[unit]
	if !exists_unit {
		return false
	}
	if !exists_item {
		bill[item] = value_unit
	} else {
		bill[item] = value_item + value_unit
	}
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	value_item, exists_item := bill[item]
	value_unit, exists_unit := units[unit]
	if !exists_item || !exists_unit {
		return false
	}
	new_value := value_item - value_unit
	if new_value < 0 {
		return false
	} else if new_value == 0 {
		delete(bill, item)
	} else {
		bill[item] = new_value
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	value, exist := bill[item]
	if !exist {
		return 0, false
	}
	return value, true
}
