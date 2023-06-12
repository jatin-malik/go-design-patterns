package structural

import (
	"errors"
	"fmt"
)

// Flyweight is a structural design pattern that allows programs to support vast quantities of objects by keeping
// their memory consumption low.

// The pattern achieves it by sharing parts of object state between multiple objects. In other words,
// the Flyweight saves RAM by caching the same data used by different objects.

// Let's take the example of players in CS-GO.

type DressFactory struct {
	dresses map[string]Dress // pool of dresses available
}

func (df *DressFactory) getDressByType(dtype string) (Dress, error) {
	if df.dresses[dtype] != nil {
		return df.dresses[dtype], nil
	}

	if dtype == "T" {
		df.dresses[dtype] = newTDress()
		return df.dresses[dtype], nil
	}

	if dtype == "CT" {
		df.dresses[dtype] = newCTDress()
		return df.dresses[dtype], nil
	}

	return nil, errors.New("unrecognised dress")
}

var dressFactory *DressFactory = &DressFactory{
	dresses: make(map[string]Dress),
}

func getDressFactory() *DressFactory {
	return dressFactory
}

type Dress interface {
	getColor() string
}

type CTDress struct {
	color string
}

func (ct *CTDress) getColor() string {
	return ct.color
}

type TDress struct {
	color string
}

func (t *TDress) getColor() string {
	return t.color
}

func newTDress() Dress {
	return &TDress{
		color: "Red",
	}
}

func newCTDress() Dress {
	return &CTDress{
		color: "Green",
	}
}

type Player struct {
	ptype string
	lat   int
	lon   int
	dress Dress // flyweight
}

func newPlayer(ptype, dtype string) *Player {
	dress, _ := getDressFactory().getDressByType(dtype)
	return &Player{
		ptype: ptype,
		dress: dress,
	}
}

func (p *Player) setLocation(lat, lon int) {
	p.lat = lat
	p.lon = lon
}

type Game struct {
	terrorists        []*Player
	counterTerrorists []*Player
}

func newGame() *Game {
	return &Game{
		terrorists:        make([]*Player, 1),
		counterTerrorists: make([]*Player, 1),
	}
}

func (g *Game) AddTerrorist() {
	player := newPlayer("T", "T")
	g.terrorists = append(g.terrorists, player)
}

func (g *Game) AddCTerrorist() {
	player := newPlayer("CT", "CT")
	g.counterTerrorists = append(g.counterTerrorists, player)
}

func RunFlyweightDemo() {
	// Make a new game
	game := newGame()

	game.AddTerrorist()
	game.AddTerrorist()
	game.AddTerrorist()
	game.AddTerrorist()

	game.AddCTerrorist()
	game.AddCTerrorist()
	game.AddCTerrorist()
	game.AddCTerrorist()

	// See our dress pool so far
	for dtype, dress := range getDressFactory().dresses {
		fmt.Printf("For dress type %s, we have dress %p with value %v\n", dtype, &dress, dress)
	}
}
