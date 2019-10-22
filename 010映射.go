package main

import "fmt"

func main() {
	var player = make(map[string]int)
	player["001"] = 001
	player["002"] = 002
	player["003"] = 003
	player["004"] = 004
	player["005"] = 005
	player["006"] = 005
	player["007"] = 005
	player["008"] = 005
	player["009"] = 005
	fmt.Println(player)
	delete(player, "009")
	fmt.Println(player)
}
