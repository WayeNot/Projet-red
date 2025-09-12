package red

import "fmt"

func (i *Character) addItem(itemId int, itemQuantity int) {
	invItems := i.Inventory
	invItems = append(invItems, Inventory{
		Id: itemId,
		Quantity: itemQuantity,
	})
	for _,v := range allItems {
		fmt.Println(v, " - ")
	}
}
