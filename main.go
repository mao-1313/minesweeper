package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

func main() {
	// Screen初期化
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Init(); err != nil {
		log.Fatal(err)
	}
	defer s.Fini()

	// ボード作成
	b := NewBoard()

	// 最初の描画
	drawBoard(s, b)
	// 画面更新
	s.Show()

	// ゲームのループ
	for {
		ev := s.PollEvent()
		switch ev := ev.(type) {

		case *tcell.EventKey:
			//カーソル移動
			switch ev.Key() {
			case tcell.KeyUp:
				b.MoveCursor(-1, 0)
			case tcell.KeyDown:
				b.MoveCursor(1, 0)
			case tcell.KeyLeft:
				b.MoveCursor(0, -1)
			case tcell.KeyRight:
				b.MoveCursor(0, 1)

			// キー入力の処理
			case tcell.KeyRune:
				switch ev.Rune() {
				case 'f':
					// フラグのトグル
					b.ToggleFlag()
				case ' ':
					// マスを開く
					b.Open()
				case 'q':
					// 終了
					return
				case 'r':
					// リスタート
					b = NewBoard()
				case 'w':
					b.MoveCursor(-1, 0)
				case 's':
					b.MoveCursor(1, 0)
				case 'a':
					b.MoveCursor(0, -1)
				case 'd':
					b.MoveCursor(0, 1)
				}
			}
		}
		drawBoard(s, b)
		switch b.State {
		case GameOver:
			// ゲームオーバー
			drawMessage(s, "GAME OVER!! Press 'r' to restart")
		case GameClear:
			// ゲームクリア
			drawMessage(s, "GAME CREAR!! Press 'r' to restart")
		}
		s.Show()

	}
}
