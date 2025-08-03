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
	val, exists := units[unit]
    if !exists{
        return false
    }

    val2, exists2 := bill[item]
    if exists2{
        bill[item]  = val+val2
    }else{
         bill[item]  = val
    }
    return true
    
}

func RemoveItem(bill, units map[string]int, item, unit string) bool {
	unitValue, unitExists := units[unit]
	if !unitExists {
		return false
	}

	itemQuantity, itemExists := bill[item]
	if !itemExists {
		return false
	}

	newQuantity := itemQuantity - unitValue
	if newQuantity < 0 {
		return false
	}
    
	if newQuantity == 0 {
		delete(bill, item)
	} else {
		bill[item] = newQuantity
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	itemQuantity, itemExists := bill[item]
	if !itemExists {
		return 0,false
	}
	return itemQuantity,true
}
