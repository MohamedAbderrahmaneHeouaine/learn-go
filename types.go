package main

import "fmt"

// types
// float32 float64 string int32 int64 uint int  uint32 uint64 uint8 byte rune(single Unicode character)
type Player struct {
	name        string
	health      int
	attackPower float64
}

var (
	floatVar  float32 = 0.1
	floatVar2 float64 = 0.1
)

//	func getHealth(player Player) int {
//		return player.health
//	}
//
// functino receiver
func (player Player) getHealth() int {
	return player.health
}

// custom type
type weapon string

func getWeapon(weapon2 weapon) weapon {
	return weapon2 // or you can casting string(weapon2) add the return type will be string
}
func main() {
	player := Player{}
	player.name = "jhon"
	player.health = 5
	player.attackPower = 4.3
	fmt.Println(player.getHealth())
	//maps
	//users := map[Player]int{
	//	player: 3,
	//}
	users := make(map[Player]int)
	users[player] = 3
	score, ok := users[player]
	if ok {
		fmt.Println(score)
	} else {
		fmt.Println("player dont exist")
	}
	// to vist all map element
	for k, v := range users {
		fmt.Println(k, v)
	}
	//to delete an element from a map
	delete(users, player)
	fmt.Println(users)
	//slices
	var numbers []int
	//others := make([]int, 3)
	numbers = append(numbers, 1)
	numbers = append(numbers, 2)
	fmt.Println(numbers)

}
