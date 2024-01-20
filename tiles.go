package main

type Player struct {
	X	int
	Y   int
}

var PlayerPosition = &Player {
	X: 3,
	Y: 3,
}

type Tile struct {
	X				int
	Y				int
	Description 	string 
	Short			string
	Examine			string
	Monster 		string
	Item 			string 
	North 			bool
	South 			bool 
	East 			bool
	West 			bool
	Visited			bool

} 

var tiles = []Tile {
	{
		X:			 3,
		Y:			 2,	 
		Description: tile_32_description,
		Short:		 tile_32_short,
		Examine: 	 "none",
		Monster: 	 "none",
		Item:        "none", 
		North:		 false,
		South:		 false,
		East:        true,
		West:        false,
		Visited:     false,
	},
	{
		X:			 2,
		Y:			 3,	 
		Description: tile_23_description,
		Examine: 	 tile_23_examine,
		Short:		 tile_23_short,
		Monster: 	 "none",
		Item:        "none", 
		North:		 false,
		South:		 false,
		East:        true,
		West:        false,
		Visited:     false,
	},
	{
		X:			 3,
		Y:			 3,
		Description: tile_33_short,
		Examine: 	 tile_33_examine,
		Short:		 tile_33_short,
		Monster: 	 "none",
		Item:        "none", 
		North:		 false,
		South:		 false,
		East:        true,
		West:        true,
		Visited:     false,
	},
	{
		X:			 4,
		Y:			 3,
		Description: tile_43_description,
		Examine:	 tile_43_examine,
		Short:		 tile_43_short,
		Monster:	 "none",
		Item:        "stone",
		North:		 false, 
		South: 		 false,
		East:		 false,
		West:		 true,
		Visited:     false,
	},
	// Add more tiles if needed
}