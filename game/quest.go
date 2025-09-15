package red

type Quest struct {
	Name string
	Reward Item
	Quantity int
	IsCompleted bool
	Menu Menu
}

func InitQuest(name string, reward Item, quantity int, menu Menu) Quest {
	return Quest{name, reward, quantity, false, menu}
}

func (q *Quest) MarkAsCompleted(){
	q.IsCompleted = true
}

func (q *Quest) RewardPlayer(character Character) {
	character.AddItem(q.Reward.Id, q.Quantity)
}

func IsCompleted(q Quest) bool {
	return q.IsCompleted
}