package main


func go_north() {
	PlayerPosition.Y -= 1
}

func go_south() {
	PlayerPosition.Y += 1
}

func go_east() {
	PlayerPosition.X += 1
}

func go_west() {
	PlayerPosition.X -= 1
}

func place_stone(x int, y int, newValue bool) {
	for i, tile := range tiles{
		if tile.X ==  3 && tile.Y == 3 {
			tiles[i].North = true
		}
	}
}

func hasItem(item string) bool {
	for _, inventoryItem := range playerInventory{
		if inventoryItem == item {
			return true
		}
	}
	return false
}

