package test

import (
	"context"
	"fmt"
	"strings"

	"github.com/Chengxufeng1994/ai-100x-se-join-quest/chinesechess/internal/domain"
	"github.com/Chengxufeng1994/ai-100x-se-join-quest/chinesechess/internal/service"
	"github.com/cucumber/godog"
)

type testContext struct {
	isLegal   bool
	isGameOver bool
	service   *service.ChineseChessService
}

func theBoardIsEmptyExceptForARedGeneralAt(ctx context.Context, row, col int) (context.Context, error) {
	tc := ctx.Value("testContext").(*testContext)
	tc.service = service.NewChineseChessService()
	tc.service.SetPiece(domain.Piece{Name: "Red General", Color: "Red"}, domain.Position{Row: row, Col: col})
	return context.WithValue(ctx, "testContext", tc), nil
}

func theBoardIsEmptyExceptForARedGuardAt(ctx context.Context, row, col int) (context.Context, error) {
	tc := ctx.Value("testContext").(*testContext)
	tc.service = service.NewChineseChessService()
	tc.service.SetPiece(domain.Piece{Name: "Red Guard", Color: "Red"}, domain.Position{Row: row, Col: col})
	return context.WithValue(ctx, "testContext", tc), nil
}

func theBoardIsEmptyExceptForARedRookAt(ctx context.Context, row, col int) (context.Context, error) {
	tc := ctx.Value("testContext").(*testContext)
	tc.service = service.NewChineseChessService()
	tc.service.SetPiece(domain.Piece{Name: "Red Rook", Color: "Red"}, domain.Position{Row: row, Col: col})
	return context.WithValue(ctx, "testContext", tc), nil
}

func theBoardIsEmptyExceptForARedHorseAt(ctx context.Context, row, col int) (context.Context, error) {
	tc := ctx.Value("testContext").(*testContext)
	tc.service = service.NewChineseChessService()
	tc.service.SetPiece(domain.Piece{Name: "Red Horse", Color: "Red"}, domain.Position{Row: row, Col: col})
	return context.WithValue(ctx, "testContext", tc), nil
}

func theBoardIsEmptyExceptForARedCannonAt(ctx context.Context, row, col int) (context.Context, error) {
	tc := ctx.Value("testContext").(*testContext)
	tc.service = service.NewChineseChessService()
	tc.service.SetPiece(domain.Piece{Name: "Red Cannon", Color: "Red"}, domain.Position{Row: row, Col: col})
	return context.WithValue(ctx, "testContext", tc), nil
}

func theBoardIsEmptyExceptForARedElephantAt(ctx context.Context, row, col int) (context.Context, error) {
	tc := ctx.Value("testContext").(*testContext)
	tc.service = service.NewChineseChessService()
	tc.service.SetPiece(domain.Piece{Name: "Red Elephant", Color: "Red"}, domain.Position{Row: row, Col: col})
	return context.WithValue(ctx, "testContext", tc), nil
}

func theBoardIsEmptyExceptForARedSoldierAt(ctx context.Context, row, col int) (context.Context, error) {
	tc := ctx.Value("testContext").(*testContext)
	tc.service = service.NewChineseChessService()
	tc.service.SetPiece(domain.Piece{Name: "Red Soldier", Color: "Red"}, domain.Position{Row: row, Col: col})
	return context.WithValue(ctx, "testContext", tc), nil
}

func redMovesTheGeneralFromTo(ctx context.Context, fromRow, fromCol, toRow, toCol int) (context.Context, error) {
	tc := ctx.Value("testContext").(*testContext)
	fromPos := domain.Position{Row: fromRow, Col: fromCol}
	toPos := domain.Position{Row: toRow, Col: toCol}
	tc.isLegal = tc.service.IsLegalMove(fromPos, toPos)
	if tc.isLegal {
		piece := tc.service.GetPiece(fromPos)
		tc.isGameOver = tc.service.CheckGameOver(piece, toPos)
	}
	return context.WithValue(ctx, "testContext", tc), nil
}

func redMovesTheGuardFromTo(ctx context.Context, fromRow, fromCol, toRow, toCol int) (context.Context, error) {
	tc := ctx.Value("testContext").(*testContext)
	fromPos := domain.Position{Row: fromRow, Col: fromCol}
	toPos := domain.Position{Row: toRow, Col: toCol}
	tc.isLegal = tc.service.IsLegalMove(fromPos, toPos)
	if tc.isLegal {
		piece := tc.service.GetPiece(fromPos)
		tc.isGameOver = tc.service.CheckGameOver(piece, toPos)
	}
	return context.WithValue(ctx, "testContext", tc), nil
}

func redMovesTheRookFromTo(ctx context.Context, fromRow, fromCol, toRow, toCol int) (context.Context, error) {
	tc := ctx.Value("testContext").(*testContext)
	fromPos := domain.Position{Row: fromRow, Col: fromCol}
	toPos := domain.Position{Row: toRow, Col: toCol}
	tc.isLegal = tc.service.IsLegalMove(fromPos, toPos)
	if tc.isLegal {
		piece := tc.service.GetPiece(fromPos)
		tc.isGameOver = tc.service.CheckGameOver(piece, toPos)
	}
	return context.WithValue(ctx, "testContext", tc), nil
}

func redMovesTheHorseFromTo(ctx context.Context, fromRow, fromCol, toRow, toCol int) (context.Context, error) {
	tc := ctx.Value("testContext").(*testContext)
	fromPos := domain.Position{Row: fromRow, Col: fromCol}
	toPos := domain.Position{Row: toRow, Col: toCol}
	tc.isLegal = tc.service.IsLegalMove(fromPos, toPos)
	if tc.isLegal {
		piece := tc.service.GetPiece(fromPos)
		tc.isGameOver = tc.service.CheckGameOver(piece, toPos)
	}
	return context.WithValue(ctx, "testContext", tc), nil
}

func redMovesTheCannonFromTo(ctx context.Context, fromRow, fromCol, toRow, toCol int) (context.Context, error) {
	tc := ctx.Value("testContext").(*testContext)
	fromPos := domain.Position{Row: fromRow, Col: fromCol}
	toPos := domain.Position{Row: toRow, Col: toCol}
	tc.isLegal = tc.service.IsLegalMove(fromPos, toPos)
	if tc.isLegal {
		piece := tc.service.GetPiece(fromPos)
		tc.isGameOver = tc.service.CheckGameOver(piece, toPos)
	}
	return context.WithValue(ctx, "testContext", tc), nil
}

func redMovesTheElephantFromTo(ctx context.Context, fromRow, fromCol, toRow, toCol int) (context.Context, error) {
	tc := ctx.Value("testContext").(*testContext)
	fromPos := domain.Position{Row: fromRow, Col: fromCol}
	toPos := domain.Position{Row: toRow, Col: toCol}
	tc.isLegal = tc.service.IsLegalMove(fromPos, toPos)
	if tc.isLegal {
		piece := tc.service.GetPiece(fromPos)
		tc.isGameOver = tc.service.CheckGameOver(piece, toPos)
	}
	return context.WithValue(ctx, "testContext", tc), nil
}

func redMovesTheSoldierFromTo(ctx context.Context, fromRow, fromCol, toRow, toCol int) (context.Context, error) {
	tc := ctx.Value("testContext").(*testContext)
	fromPos := domain.Position{Row: fromRow, Col: fromCol}
	toPos := domain.Position{Row: toRow, Col: toCol}
	tc.isLegal = tc.service.IsLegalMove(fromPos, toPos)
	if tc.isLegal {
		piece := tc.service.GetPiece(fromPos)
		tc.isGameOver = tc.service.CheckGameOver(piece, toPos)
	}
	return context.WithValue(ctx, "testContext", tc), nil
}

func theMoveIsLegal(ctx context.Context) error {
	tc := ctx.Value("testContext").(*testContext)
	if !tc.isLegal {
		return fmt.Errorf("expected move to be legal, but it was illegal")
	}
	return nil
}

func theMoveIsIllegal(ctx context.Context) error {
	tc := ctx.Value("testContext").(*testContext)
	if tc.isLegal {
		return fmt.Errorf("expected move to be illegal, but it was legal")
	}
	return nil
}

func redWinsImmediately(ctx context.Context) error {
	tc := ctx.Value("testContext").(*testContext)
	if !tc.isGameOver {
		return fmt.Errorf("expected game to be over, but it was not")
	}
	return nil
}

func theGameIsNotOverJustFromThatCapture(ctx context.Context) error {
	tc := ctx.Value("testContext").(*testContext)
	if tc.isGameOver {
		return fmt.Errorf("expected game to not be over, but it was")
	}
	return nil
}

func theBoardHas(ctx context.Context, table *godog.Table) (context.Context, error) {
	tc := ctx.Value("testContext").(*testContext)
	tc.service = service.NewChineseChessService()

	for _, row := range table.Rows[1:] {
		pieceName := row.Cells[0].Value
		positionStr := row.Cells[1].Value

		var pieceColor string
		if strings.Contains(pieceName, "Red") {
			pieceColor = "Red"
		} else if strings.Contains(pieceName, "Black") {
			pieceColor = "Black"
		}

		var row, col int
		fmt.Sscanf(positionStr, "(%d, %d)", &row, &col)

		tc.service.SetPiece(domain.Piece{Name: pieceName, Color: pieceColor}, domain.Position{Row: row, Col: col})
	}
	return context.WithValue(ctx, "testContext", tc), nil
}

func InitializeGeneralScenario(ctx *godog.ScenarioContext) {
	tc := &testContext{}

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		return context.WithValue(ctx, "testContext", tc), nil
	})

	ctx.Step(`^the board is empty except for a Red General at \((\d+), (\d+)\)$`, theBoardIsEmptyExceptForARedGeneralAt)
	ctx.Step(`^the board is empty except for a Red Guard at \((\d+), (\d+)\)$`, theBoardIsEmptyExceptForARedGuardAt)
	ctx.Step(`^the board is empty except for a Red Rook at \((\d+), (\d+)\)$`, theBoardIsEmptyExceptForARedRookAt)
	ctx.Step(`^the board is empty except for a Red Horse at \((\d+), (\d+)\)$`, theBoardIsEmptyExceptForARedHorseAt)
	ctx.Step(`^the board is empty except for a Red Cannon at \((\d+), (\d+)\)$`, theBoardIsEmptyExceptForARedCannonAt)
	ctx.Step(`^the board is empty except for a Red Elephant at \((\d+), (\d+)\)$`, theBoardIsEmptyExceptForARedElephantAt)
	ctx.Step(`^the board is empty except for a Red Soldier at \((\d+), (\d+)\)$`, theBoardIsEmptyExceptForARedSoldierAt)
	ctx.Step(`^Red moves the General from \((\d+), (\d+)\) to \((\d+), (\d+)\)$`, redMovesTheGeneralFromTo)
	ctx.Step(`^Red moves the Guard from \((\d+), (\d+)\) to \((\d+), (\d+)\)$`, redMovesTheGuardFromTo)
	ctx.Step(`^Red moves the Rook from \((\d+), (\d+)\) to \((\d+), (\d+)\)$`, redMovesTheRookFromTo)
	ctx.Step(`^Red moves the Horse from \((\d+), (\d+)\) to \((\d+), (\d+)\)$`, redMovesTheHorseFromTo)
	ctx.Step(`^Red moves the Cannon from \((\d+), (\d+)\) to \((\d+), (\d+)\)$`, redMovesTheCannonFromTo)
	ctx.Step(`^Red moves the Elephant from \((\d+), (\d+)\) to \((\d+), (\d+)\)$`, redMovesTheElephantFromTo)
	ctx.Step(`^Red moves the Soldier from \((\d+), (\d+)\) to \((\d+), (\d+)\)$`, redMovesTheSoldierFromTo)
	ctx.Step(`^the move is legal$`, theMoveIsLegal)
	ctx.Step(`^the move is illegal$`, theMoveIsIllegal)
	ctx.Step(`^Red wins immediately$`, redWinsImmediately)
	ctx.Step(`^the game is not over just from that capture$`, theGameIsNotOverJustFromThatCapture)
	ctx.Step(`^the board has:$`, theBoardHas)

	ctx.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		return ctx, nil
	})
}