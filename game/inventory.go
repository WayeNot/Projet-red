package red

type Inventory struct {
	Id       int
	Quantity int
}

type Item struct {
	Id    				int
	Name  				string
	Price 				int
	Icon  				string
	addHealth 			int
	removeEnemyHealth 	int
	giveInventory		int
	IsForgeron 			bool
	IsUsable 			bool
}

var allItems = map[int]Item{
	// Potions ↓

	1: {Id: 1, Name: "Potion de vie", Price: 3, Icon: "➕", IsForgeron: false, addHealth: 15},
	2: {Id: 2, Name: "Potion de poison", Price: 6, Icon: "🤢", IsForgeron: false, removeEnemyHealth: 10},

	4: {Id: 4, Name: "Sac à dos", Price: 30, Icon: "🎒", IsForgeron: false, giveInventory: 10},
}