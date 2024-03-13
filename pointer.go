package main

import "fmt"

type PlayerB struct {
	HP int
}

// if player is not a  pointer we are adjusting the copy of the player not the actual player
func TakeDamage(player *PlayerB, amount int) {
	fmt.Println("player taking damage", amount)
	player.HP -= amount
	fmt.Println("player new hp is", player.HP)
}

func pointer() {
	player := &PlayerB{
		HP: 100,
	}
	TakeDamage(player, 10)
	TakeDamage(player, 10)
	fmt.Println(player.HP)
}
