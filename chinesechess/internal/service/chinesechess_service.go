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

	var strategy PieceMoveStrategy
	switch piece.Name {
	case "Red General", "Black General":
		strategy = &GeneralMoveStrategy{}
	case "Red Guard", "Black Guard":
		strategy = &GuardMoveStrategy{}
	case "Red Rook", "Black Rook":
		strategy = &RookMoveStrategy{}
	case "Red Horse", "Black Horse":
		strategy = &HorseMoveStrategy{}
	case "Red Cannon", "Black Cannon":
		strategy = &CannonMoveStrategy{}
	case "Red Elephant", "Black Elephant":
		strategy = &ElephantMoveStrategy{}
	case "Red Soldier", "Black Soldier":
		strategy = &SoldierMoveStrategy{}
	default:
		return false // Not implemented for other pieces yet
	}

	return strategy.IsLegalMove(s.board, from, to, piece.Color)
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
