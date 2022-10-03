package position

import "github.com/notnil/chess"

func FromFEN(fen string) *chess.Position {
	fn, _ := chess.FEN(fen)
	game := chess.NewGame(fn)
	return game.Position()
}
