package main

import "github.com/gdamore/tcell/v2"

func drawBoard(s tcell.Screen, b *Board) {
	s.Clear()
	for r := 0; r < cellNums; r++ {
		for c := 0; c < cellNums; c++ {
			drawCell(s, r, c, b.Cells[r][c], b.Cursor == [2]int{r, c}, b.State == GameOver)
		}
	}
}

func drawCell(s tcell.Screen, row, col int, cell Cell, isCursor bool, isGameOver bool) {
	x := col * 2
	y := row

	// カーソル表示
	cursorStyle := tcell.StyleDefault
	if isCursor {
		cursorStyle = cursorStyle.Reverse(true)
	}

	// マス描画
	if cell.IsOpen {
		// 開いたマス
		if cell.IsMine {
			drawChar(s, x, y, '!', cursorStyle) // 地雷
		} else {
			drawChar(s, x, y, ' ', cursorStyle)
			if cell.Adjacent > 0 {
				drawChar(s, x, y, rune('0'+cell.Adjacent), cursorStyle.Foreground(tcell.ColorWhite).Background(tcell.ColorBlue))
			}
		}
	} else {
		// 未開のマス
		mark := '■'
		if cell.IsFlagged {
			mark = 'F'
		}
		// ゲームオーバになったら全て開ける
		if isGameOver && cell.IsMine {
			mark = '*'
		}
		drawChar(s, x, y, mark, cursorStyle)
	}

	// 枠線
	drawChar(s, x+1, y, '|', tcell.StyleDefault)
}

func drawChar(s tcell.Screen, x, y int, r rune, style tcell.Style) {
	s.SetContent(x, y, r, nil, style)
}

func drawMessage(s tcell.Screen, msg string) {
	// ボードの下に1行あけてメッセージを表示
	y := cellNums + 1
	for i, ch := range msg {
		s.SetContent(i, y, ch, nil, tcell.StyleDefault)
	}
}
