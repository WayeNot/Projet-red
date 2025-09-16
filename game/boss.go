package red

import (
	"fmt"
)

type Boss struct {
	Name string
	PV int
	MaxPV int
	Dead bool
	Attacks []func(*Character)
}

func InitBoss(name string, pv int, maxpv int) Boss{
	return Boss{name, pv, maxpv, false, []func(*Character){}}
}

func (b *Boss) AddPV(pv int) {
	if b.GetPV() + pv > b.GetMaxPV() {
		fmt.Println("Action impossible, PV limités à ", b.GetMaxPV())
	} else {
		b.PV += pv
	}
}

func (b *Boss) RemovePV(pv int) {
	if b.GetPV() - pv <= 0 {
		b.SetPV(0)
		b.Dead = true
	} else {
		b.SetPV(b.GetPV() - pv)
	}
}

func (b *Boss) SetPV(pv int) {
	b.PV = pv
}

func (b *Boss) GetPV() int{
	return b.PV
}

func (b *Boss) GetMaxPV() int{
	return b.MaxPV
}

func (b *Boss) GetName() string {
	return b.Name
}

func (b *Boss) AddAttacks(f func(*Character)) {
	b.Attacks = append(b.Attacks, f)
}

func (b *Boss) GetAttacks() []func(*Character) {
	return b.Attacks
}

func (b *Boss) IsDead() bool {
	if b.Dead == false{
		return false
	}
	return true
}