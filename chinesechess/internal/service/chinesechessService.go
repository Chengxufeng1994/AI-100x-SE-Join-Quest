package service

import (
	"github.com/Chengxufeng1994/ai-100x-se-join-quest/chinesechess/internal/domain"
)

type ChineseChessService struct {
	board *domain.Board
}

func NewChineseChessService() *ChineseChessService {
	return &ChineseChessService{
		board: domain.NewBoard(),
	}
}

func (s *ChineseChessService) SetPiece(piece domain.Piece, position domain.Position) {
	s.board.Pieces[position] = piece
}

func (s *ChineseChessService) GetPiece(position domain.Position) domain.Piece {
	return s.board.Pieces[position]
}

func (s *ChineseChessService) IsLegalMove(from, to domain.Position) bool {
	piece, exists := s.board.Pieces[from]
	if !exists {
		return false // No piece at the 'from' position
	}

	switch piece.Name {
	case "Red General", "Black General":
		return s.isLegalGeneralMove(from, to, piece.Color)
	case "Red Guard", "Black Guard":
		return s.isLegalGuardMove(from, to, piece.Color)
	case "Red Rook", "Black Rook":
		return s.isLegalRookMove(from, to, piece.Color)
	case "Red Horse", "Black Horse":
		return s.isLegalHorseMove(from, to, piece.Color)
	case "Red Cannon", "Black Cannon":
		return s.isLegalCannonMove(from, to, piece.Color)
	case "Red Elephant", "Black Elephant":
		return s.isLegalElephantMove(from, to, piece.Color)
	case "Red Soldier", "Black Soldier":
		return s.isLegalSoldierMove(from, to, piece.Color)
	default:
		return false // Not implemented for other pieces yet
	}
}

func (s *ChineseChessService) CheckGameOver(lastMovedPiece domain.Piece, to domain.Position) bool {
	// If the last moved piece captured the opponent's General, the game is over.
	// This is a simplified check for the current scenario.
	capturedPiece, exists := s.board.Pieces[to]
	if exists && capturedPiece.Name == "Black General" && lastMovedPiece.Color == "Red" {
		return true
	}
	return false
}

func (s *ChineseChessService) isLegalGeneralMove(from, to domain.Position, color string) bool {
	dx := abs(to.Col - from.Col)
	dy := abs(to.Row - from.Row)

	// Must move exactly one step orthogonally
	if !((dx == 1 && dy == 0) || (dx == 0 && dy == 1)) {
		return false
	}

	// Must remain within the palace
	var minRow, maxRow, minCol, maxCol int
	if color == "Red" {
		minRow, maxRow = 1, 3
		minCol, maxCol = 4, 6
	} else if color == "Black" {
		minRow, maxRow = 8, 10
		minCol, maxCol = 4, 6
	} else {
		return false // Invalid general color
	}

	if to.Row < minRow || to.Row > maxRow || to.Col < minCol || to.Col > maxCol {
		return false
	}

	// Simulate the move to check for general facing rule
	tempBoardPieces := make(map[domain.Position]domain.Piece)
	for pos, p := range s.board.Pieces {
		tempBoardPieces[pos] = p
	}

	// Remove the piece from its original position and place it at the new position
	delete(tempBoardPieces, from)
	tempBoardPieces[to] = s.board.Pieces[from]

	// Find generals' positions after the move
	var redGeneralPos, blackGeneralPos domain.Position
	redGeneralFound := false
	blackGeneralFound := false

	for pos, p := range tempBoardPieces {
		if p.Name == "Red General" {
			redGeneralPos = pos
			redGeneralFound = true
		} else if p.Name == "Black General" {
			blackGeneralPos = pos
			blackGeneralFound = true
		}
	}

	if redGeneralFound && blackGeneralFound && redGeneralPos.Col == blackGeneralPos.Col {
		// Generals are in the same column, check for intervening pieces
		minRow := min(redGeneralPos.Row, blackGeneralPos.Row)
		maxRow := max(redGeneralPos.Row, blackGeneralPos.Row)

		for r := minRow + 1; r < maxRow; r++ {
			if _, exists := tempBoardPieces[domain.Position{Row: r, Col: redGeneralPos.Col}]; exists {
				// There is a piece between the generals, so the move is legal in this regard
				return true
			}
		}
		// No pieces between generals, so this move is illegal
		return false
	}

	return true
}

func (s *ChineseChessService) isLegalGuardMove(from, to domain.Position, color string) bool {
	dx := abs(to.Col - from.Col)
	dy := abs(to.Row - from.Row)

	// Must move exactly one step diagonally
	if !(dx == 1 && dy == 1) {
		return false
	}

	// Must remain within the palace
	var minRow, maxRow, minCol, maxCol int
	if color == "Red" {
		minRow, maxRow = 1, 3
		minCol, maxCol = 4, 6
	} else if color == "Black" {
		minRow, maxRow = 8, 10
		minCol, maxCol = 4, 6
	} else {
		return false // Invalid guard color
	}

	if to.Row < minRow || to.Row > maxRow || to.Col < minCol || to.Col > maxCol {
		return false
	}

	return true
}

func (s *ChineseChessService) isLegalRookMove(from, to domain.Position, color string) bool {
	// Must move horizontally or vertically
	if from.Row != to.Row && from.Col != to.Col {
		return false
	}

	// Count pieces between from and to
	piecesBetween := 0
	if from.Row == to.Row { // Horizontal move
		minCol := min(from.Col, to.Col)
		maxCol := max(from.Col, to.Col)
		for c := minCol + 1; c < maxCol; c++ {
			if _, exists := s.board.Pieces[domain.Position{Row: from.Row, Col: c}]; exists {
				piecesBetween++
			}
		}
	} else { // Vertical move
		minRow := min(from.Row, to.Row)
		maxRow := max(from.Row, to.Row)
		for r := minRow + 1; r < maxRow; r++ {
			if _, exists := s.board.Pieces[domain.Position{Row: r, Col: from.Col}]; exists {
				piecesBetween++
			}
		}
	}

	// If target position is empty, must have zero pieces between
	if _, exists := s.board.Pieces[to]; !exists {
		return piecesBetween == 0
	}

	// If target position has a piece, must have exactly one piece between (the cannon platform)
	return piecesBetween == 0
}

func (s *ChineseChessService) isLegalHorseMove(from, to domain.Position, color string) bool {
	dx := abs(to.Col - from.Col)
	dy := abs(to.Row - from.Row)

	// Must move in an L shape (1 horizontal/vertical, then 2 vertical/horizontal)
	if !((dx == 1 && dy == 2) || (dx == 2 && dy == 1)) {
		return false
	}

	// Check for blocking piece
	var blockPos domain.Position
	if dy == 2 { // Horse moved 2 steps vertically, 1 step horizontally
		blockPos = domain.Position{Row: from.Row + sign(to.Row-from.Row), Col: from.Col}
	} else { // Horse moved 2 steps horizontally, 1 step vertically
		blockPos = domain.Position{Row: from.Row, Col: from.Col + sign(to.Col-from.Col)}
	}

	if _, exists := s.board.Pieces[blockPos]; exists {
		return false
	}

	return true
}

func (s *ChineseChessService) isLegalCannonMove(from, to domain.Position, color string) bool {
	// Must move horizontally or vertically
	if from.Row != to.Row && from.Col != to.Col {
		return false
	}

	// Count pieces between from and to
	piecesBetween := 0
	if from.Row == to.Row { // Horizontal move
		minCol := min(from.Col, to.Col)
		maxCol := max(from.Col, to.Col)
		for c := minCol + 1; c < maxCol; c++ {
			if _, exists := s.board.Pieces[domain.Position{Row: from.Row, Col: c}]; exists {
				piecesBetween++
			}
		}
	} else { // Vertical move
		minRow := min(from.Row, to.Row)
		maxRow := max(from.Row, to.Row)
		for r := minRow + 1; r < maxRow; r++ {
			if _, exists := s.board.Pieces[domain.Position{Row: r, Col: from.Col}]; exists {
				piecesBetween++
			}
		}
	}

	// If target position is empty, must have zero pieces between
	if _, exists := s.board.Pieces[to]; !exists {
		return piecesBetween == 0
	}

	// If target position has a piece, must have exactly one piece between (the cannon platform)
	return piecesBetween == 1
}

func (s *ChineseChessService) isLegalElephantMove(from, to domain.Position, color string) bool {
	dx := abs(to.Col - from.Col)
	dy := abs(to.Row - from.Row)

	// Must move exactly two steps diagonally
	if !(dx == 2 && dy == 2) {
		return false
	}

	// Cannot cross the river
	if color == "Red" && to.Row > 5 {
		return false
	}
	if color == "Black" && to.Row < 6 {
		return false
	}

	// Check for blocking piece at midpoint
	midRow := from.Row + sign(to.Row-from.Row)
	midCol := from.Col + sign(to.Col-from.Col)
	if _, exists := s.board.Pieces[domain.Position{Row: midRow, Col: midCol}]; exists {
		return false
	}

	return true
}

func (s *ChineseChessService) isLegalSoldierMove(from, to domain.Position, color string) bool {
	dx := abs(to.Col - from.Col)
	dy := abs(to.Row - from.Row)

	// Must move exactly one step
	if !((dx == 0 && dy == 1) || (dx == 1 && dy == 0)) {
		return false
	}

	// Red soldier moves forward (increasing row number)
	if color == "Red" {
		if to.Row < from.Row {
			return false // Cannot move backward
		}
		// Before crossing river (row 5), can only move forward (dy=1, dx=0)
		if from.Row <= 5 && dx == 1 {
			return false // Cannot move sideways before crossing river
		}
	} else if color == "Black" {
		// Black soldier moves forward (decreasing row number)
		if to.Row > from.Row {
			return false // Cannot move backward
		}
		// Before crossing river (row 6), can only move forward (dy=1, dx=0)
		if from.Row >= 6 && dx == 1 {
			return false // Cannot move sideways before crossing river
		}
	}

	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sign(x int) int {
	if x == 0 {
		return 0
	}
	if x > 0 {
		return 1
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
