package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/encoder"
	"log"
	"net/http"
	"strconv"
)

type User struct {
	id   int
	name string
}

type Game struct {
	id         int
	players    []User
	puzzles    []int
	dictionary string
}

type Puzzle struct {
	Id     int    `json:"id"`
	Phrase string `json:"phrase"`
}

type Player struct {
	id         int
	puzzle     Puzzle
	dictionary string
}

var puzzles []Puzzle

func main() {
	// puzzles := make([]Puzzle, 0, 100)
	// puzzles = append(puzzles, Puzzle{
	// 	id:     0,
	// 	phrase: "The quick brown fox jumps high over the fence.",
	// }, Puzzle{
	// 	id:     1,
	// 	phrase: "Hung and Bob are building yet another stupid hackathon project.",
	// }, Puzzle{
	// 	id:     2,
	// 	phrase: "Zoolander was a funny movie.",
	// })

	//games := []Game{
	//game := Game{
	//	id: 0,
	//	players: []Player,
	//	puzzles: []Puzzle
	//}

	m := martini.Classic()

	m.Use(func(c martini.Context, w http.ResponseWriter, r *http.Request) {
		pretty, _ := strconv.ParseBool(r.FormValue("pretty"))
		c.MapTo(encoder.JsonEncoder{PrettyPrint: pretty}, (*encoder.Encoder)(nil))
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	})

	m.Get("/", func(enc encoder.Encoder) (int, []byte) {
		puzzle := new(Puzzle)
		puzzle.Id = 5
		puzzle.Phrase = "Hello there."
		return http.StatusOK, encoder.Must(enc.Encode(puzzle))
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
		r.Get("/:userId", GetUser)
		r.Put("/:userId", UpdateUser)
		r.Delete("/:userId", DeleteUser)
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
	return 200, "Puzzle in Game Requested"
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
