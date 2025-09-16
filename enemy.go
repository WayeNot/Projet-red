package red

import "fmt"

type enemy struct {
	ID     int
	Name   string
	Pv     int
	Atk    int
	Level  int
	IsDead bool
}

func main() {
	orcPmu := enemy{ID: 1, Name: "Orc du PMU", Pv: 170, Atk: 35 Level 7}
	dragonRsa := enemy{ID: 2, Name: "Dragon du RSA", Pv: 190, Atk: 40, Level 8}
	drhUltime := enemy{ID: 3, Name: "Le DRH ultime", Pv: 250, Atk: 50, Level 10}
	ratCdd := enemy{ID: 4, Name: "Rat géant du CDD", Pv: 130, Atk: 30, Level 5}
	gobelinBanquier := enemy{ID: 5, Name: "Gobelin banquier", Pv: 120, Atk: 45, Level 6}
	nainSyndicaliste := enemy{ID: 6, Name: "Nain syndicaliste", Pv: 100, Atk: 80, Level 2}
	planteVerte := enemy{ID: 7, Name: "Plante verte", Pv: 500, Atk: 40, Level 6}

	if orcPmu.Pv <= 0 {
		fmt.Println("Bien joué d'avoir tué un Orc du PMU !")
	}
	if dragonRsa.Pv <= 0 {
		fmt.Println("Tu as vaincu le Dragon du RSA !")
	}
	if drhUltime.Pv <= 0 {
		fmt.Println("Tu as finalement vaincu le DRH ultime et obtenu un CDI !")
	}
	if ratCdd.Pv <= 0 {
		fmt.Println("Tu as tué un gros rat !")
	}
	if gobelinBanquier.Pv <= 0 {
		fmt.Println("Félicitations, tu as tué un banquier !")
	}
	if nainSyndicaliste.Pv <= 0 {
		fmt.Println("Vous avez tué un nain syndicaliste !")
	}
	if planteVerte.Pv <= 0 {
		fmt.Println("Tu viens de tuer une plante ?!")
	}
}



                                 



