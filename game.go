package main

import (
	"fmt"
	"strings"

)


func getCurrentTile(PlayerPositon *Player) (Tile, error) {
	for _, tile := range tiles {
		if tile.X == PlayerPositon.X && tile.Y == PlayerPositon.Y {
			return tile, nil
		}
	} 
	return Tile{}, fmt.Errorf("Tile not found for position (%d, %d)", PlayerPositon.X, PlayerPositon.Y)
}


// var currentTile, err = getCurrentTile(PlayerPosition)

var playerInventory []string


func game() {

	var currentTile, err = getCurrentTile(PlayerPosition)
	if err != nil {
		fmt.Println(err)
		return
	}


	var i string

	// Print for Troubleshooting
	// fmt.Println(currentTile, currentTile.X, currentTile.Y)
	// fmt.Println(PlayerPosition, PlayerPosition.X, PlayerPosition.Y)


	fmt.Println()
	_, err = fmt.Scan(&i)
	if err != nil {
		fmt.Println(err)
		return
	}
	

	input := strings.ToLower(i)
	
	switch {
	case input == "north":
		// Handle north direction
		if currentTile.North {
			go_north()
			// if !currentTile.Visited {
			// 	fmt.Println(currentTile.Description)
			// 	currentTile.Visited = true
			// 	} else {
			// 	fmt.Println(CurrentTile.Description)
			// }
		} else {
			fmt.Println("You can not go that way.")
		}

	case input == "south":
		if currentTile.South {
			go_south()
			// if !currentTile.Visited {
			// 	fmt.Println(currentTile.Description)
			// 	currentTile.Visited = true
			// 	} else {
			// 	fmt.Println(CurrentTile.Description)
			// }
		} else {
			fmt.Println("You can not go that way.")
		}
	
	case input == "east":
		if currentTile.East {
			go_east()
			currentTile, _ = getCurrentTile(PlayerPosition)
			fmt.Println(currentTile.Description)
		} else {
			fmt.Println("You can not go that way.")
		}

	case input == "west":
		if currentTile.West {
			go_west()
			currentTile, _ = getCurrentTile(PlayerPosition)
			fmt.Println(currentTile.Description)
		} else {
			fmt.Println("You can not go that way.")
		}

	case input == "examine":
		if currentTile.Examine != "none" {
			fmt.Println(currentTile.Examine)
			if currentTile.X == 2 && currentTile.Y == 3 {
				fmt.Println("Continue to try to pass? y/n")
				var h string
				_, err = fmt.Scan(&h)
				if err != nil {
					fmt.Println(err)
					return
				}
				h = strings.ToLower(h)
				if h == "y" {
					fmt.Println(hellscape)
					alive = false
					break
				} else if h == "n" {

				} else {
					fmt.Println("Invalid input.")
				}
		} else {
			fmt.Println("There is nothing to Examine.")
		}
	}

	case input == "get":
		if currentTile.Item != "none" {
			if len(playerInventory) == 5 {
				fmt.Println("Your inventory is full.")
			} else {
				playerInventory = append(playerInventory, currentTile.Item)
				fmt.Printf("You picked up the %s.", currentTile.Item)
			}
		} else {
			fmt.Println("There is nothing to pick up here.")
		}

	case input == "inventory": 
		if len(playerInventory) == 0 {
			fmt.Println("Your inventory is empty.")
		} else {
			fmt.Println(playerInventory)
		}	

	case input == "help":
		// Handle help
		fmt.Println(help)

	case input == "place":
		if hasItem("stone") {
		place_stone(3, 3, true)
		fmt.Println(stone_placed)
		} else {
			fmt.Println("You don't have anything to place there.")
		}

	case input == "quit":
		fmt.Println("Are you sure? y/n")
		var q string
		_, err = fmt.Scan(&q)
		if err != nil {
			fmt.Println(err)
			return
		}
		q = strings.ToLower(q)
		if q == "y" {
			alive = false
			break
		} else if q == "n" {

		} else {
			fmt.Println("Invalid input.")
		}

	default:
		fmt.Println("Invalid command.")

		
	}
}