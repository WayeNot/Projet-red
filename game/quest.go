package red

type Quest struct {
	Name string
	Reward int
	RequiredItemId int
	Quantity int
	IsCompleted bool
	Menu Menu
}

func InitQuest(name string, reward int, requiredItemId int, quantity int, menu Menu) Quest {
	return Quest{name, reward, requiredItemId, quantity, false, menu}
}

func (q *Quest) MarkAsCompleted(){
	q.IsCompleted = true
}

func (q *Quest) RewardPlayer(character Character) {
	character.AddItem(q.Reward, q.Quantity)
}

func IsCompleted(q Quest) bool {
	return q.IsCompleted
}

func (q *Quest) PlayQuest(player *Character) {
	q.Menu.Display(player)
}