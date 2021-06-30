package game

type Game struct {
	count  int
	number int
}

func New() *Game {
	return &Game{
		0,
		20,
	}
}

func (game *Game) Compare(target int) (int, int) {
	game.count++
	result := 0
	if target > game.number {
		result = -1
	} else if game.number > target {
		result = 1
	}
	return result, game.count
}
