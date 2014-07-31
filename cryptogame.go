package main

import "github.com/go-martini/martini"

// Log In
// 1) POST /users/
//
// Start Game
// 1) POST /games/
//
// Find Game
// 1) GET /games/
//
// Join Game
// 1) PUT /games/{id}/players/{id}
//
// Leave Game
// 1) DELETE /games/{id}/players/{id}
//
// Log Out
// 1) DELETE /users/{id}
//
// 

func main() {
	m := martini.Classic()
	
	m.Get("/", func() (int, string) {
		return 200, "Hello World!"
	})
	
	m.Group("/games", func(r martini.Router) {
		r.Post("/", NewGame)
		r.Get("/", ListGames)
		r.Get("/:gameId", GetGame)
	})
	
	m.Group("/games/:gameId/puzzles", func(r martini.Router) {
		r.Get("/", ListPuzzlesInGame)
		r.Get("/:puzzleId", GetPuzzleInGame)
	})
	
	m.Group("/games/:gameId/players", func(r martini.Router) {
		r.Get("/", ListPlayers)
		r.Get("/:playerId", GetPlayer)
		r.Put("/:playerId", AddPlayer)
		r.Delete("/:playerId", RemovePlayer)
	})

	m.Group("/users", func(r martini.Router) {
		r.Post("/", NewUser)
		r.Get("/", ListUsers)
		r.Put("/", UpdateUser)
		r.Delete("/", DeleteUser)
	})
	
	m.Run()
}

func NewGame() (int, string) {
	return 201, "New Game Requested"
}

func ListGames() (int, string) {
	return 200, "Game List Requested"
}

func GetGame(params martini.Params) (int, string) {
	return 200, "Game " + params["gameId"] + " Requested"
}

func ListPuzzlesInGame() (int, string) {
	return 200, "Game Puzzles List Requested"
}

func GetPuzzleInGame() (int, string) {
	return 200, "Game Puzzle Requested"
}

func ListPlayers() (int, string) {
	return 200, "Player List Requested"
}

func GetPlayer() (int, string) {
	return 200, "Player Requested"
}

func AddPlayer() (int, string) {
	return 204, "Player Add Requested"
}

func RemovePlayer() (int, string) {
	return 204, "Player Remove Requested"
}

func NewUser() (int, string) {
	return 201, "New User Requested"
}

func UpdateUser() (int, string) {
	return 204, "Update User Requested"
}

func GetUser() (int, string) {
	return 200, "Get User Requested"
}

func DeleteUser() (int, string) {
	return 204, "Delete User Requested"
}

func ListUsers() (int, string) {
	return 200, "List Users Requested"
}