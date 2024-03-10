package main

import "fmt"

// weapon type
// axe
// sword
// wooden stick
// knife
type WeaponType int

// enum
const (
	Axe         WeaponType = 1
	Sword       WeaponType = 2
	WoodenStick WeaponType = 3
	Knife       WeaponType = 4
)

func getDamage(weaponType WeaponType) int {
	switch weaponType {
	case Axe:
		return 100
	case Sword:
		return 90
	case WoodenStick:
		return 1
	case Knife:
		return 40
	default:
		panic("weapon does not exist")
	}

}
func main() {
	fmt.Println(getDamage(Knife))
}
