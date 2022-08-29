package tradesman

import "fmt"

type Tradesman struct {
	capacity int
	speed    int
	cash     int
	items    []Item
}

type Item struct {
	weight    int
	item_type string
	quality   float32
	cost      int
}

func Affordable(t Tradesman, i Item) bool {
	if t.cash >= i.cost {
		return true
	} else {
		return false
	}
}

func EnoughCapacity(t Tradesman, i Item) bool {
	if t.capacity >= i.weight {
		return true
	} else {
		return false
	}
}

func (t *Tradesman) BuyItem(i Item) {
	if !Affordable(*t, i) {
		fmt.Println("Not enough money(")
		return
	}

	if !EnoughCapacity(*t, i) {
		fmt.Println("Too heavy(")
		return
	}

	t.cash = t.cash - i.cost
	t.capacity = -i.weight
	t.items = append(t.items, i)

	fmt.Println("Item bought!")
}

func main() {
	tradesman := Tradesman{capacity: 50, speed: 5, cash: 500}

}
