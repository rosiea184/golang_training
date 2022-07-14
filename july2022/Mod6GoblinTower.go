package july2022

import (
	"fmt"
	"math/rand"
	"time"
)

type hero struct {
	hp      int
	attack  int
	defense int
	inv     []int
	gold    int
	level   int
}
type goblin struct {
	hp      int
	attack  int
	defense int
}

func (h hero) viewStats() {
	fmt.Println("Health: ", h.hp)
	fmt.Println("Attack: ", h.attack)
	fmt.Println("Defense: ", h.defense)
	fmt.Println("Gold: ", h.gold)
	fmt.Println("Level: ", h.level)
	count := 0
	for i := 0; i < len(h.inv); i++ {
		count += h.inv[i]
	}
	fmt.Println("Number of potions: ", count)
}
func createHero() hero {
	h := hero{}
	h.hp = rand.Intn(10) + 20
	h.attack = rand.Intn(2) + 1
	h.defense = rand.Intn(4) + 1
	//h.inv = [5]{2, 2, 2, 2, 2}
	h.gold = 0
	h.level = 1
	return h
}
func createGoblin() goblin {
	g := goblin{}
	g.hp = rand.Intn(5) + 5
	g.attack = rand.Intn(1) + 2
	g.defense = rand.Intn(1) + 1
	return g
}
func (h hero) potionShop() hero {
	return h
}

func Mod6GoblinTower() {
	rand.Seed(time.Now().UnixNano())
	h := createHero()
	potionCheck := ""
	steps := 0
	if steps == 10 {
		h.level++
		steps = 0
		fmt.Print("Do you wish to visit the potion shop? yes or no: ")
		if potionCheck == "yes" {

		}
	}

}
