package main

import "math/rand"

const mineNums = 10
const cellNums = 9

type GameState int

const (
	GamePlaying GameState = iota
	GameOver
	GameClear
)

type Cell struct {
	IsMine    bool
	IsOpen    bool
	IsFlagged bool
	Adjacent  int // 周囲の地雷数
}

type Board struct {
	Cells  [cellNums][cellNums]Cell
	Cursor [2]int // [row, col]
	State  GameState
}

func NewBoard() *Board {
	b := &Board{}
	// 地雷をランダム配置
	b.putMines()
	b.calcAdjacent()
	b.State = GamePlaying
	return b
}

// 8方向のオフセット
var dirs = [8][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func (b *Board) calcAdjacent() {
	// 全セルの周囲地雷数を計算

	for r := 0; r < cellNums; r++ {
		for c := 0; c < cellNums; c++ {
			// 地雷が設定されているなら周囲の地雷数は計算しない
			if b.Cells[r][c].IsMine {
				continue
			}

			// dirsを参照して地雷数をカウント
			count := 0
			for _, dir := range dirs {
				nr := r + dir[0]
				nc := c + dir[1]
				if 0 <= nr && nr < cellNums && 0 <= nc && nc < cellNums {
					if b.Cells[nr][nc].IsMine {
						count++
					}
				}
			}
			b.Cells[r][c].Adjacent = count

		}

	}
}

func (b *Board) putMines() {
	// 地雷をランダム配置

	for i := 0; i < mineNums; i++ {
		row := rand.Intn(cellNums)
		col := rand.Intn(cellNums)
		if b.Cells[row][col].IsMine {
			i--
			continue
		}
		b.Cells[row][col].IsMine = true
	}
}

func (b *Board) MoveCursor(dr, dc int) {
	nr := b.Cursor[0] + dr
	nc := b.Cursor[1] + dc

	// 枠外に出る場合は何もしない
	if nr < 0 || cellNums <= nr || nc < 0 || cellNums <= nc {
		return
	}
	b.Cursor[0] += dr
	b.Cursor[1] += dc
}

func (b *Board) open(r, c int) {
	cell := b.Cells[r][c]
	if cell.IsFlagged || cell.IsMine || cell.IsOpen {
		return
	}
	b.Cells[r][c].IsOpen = true

	if cell.Adjacent == 0 {
		for _, dir := range dirs {
			nr := r + dir[0]
			nc := c + dir[1]
			if 0 <= nr && nr < cellNums && 0 <= nc && nc < cellNums {
				b.open(nr, nc)
			}
		}

	}
}

func (b *Board) Open() {
	r := b.Cursor[0]
	c := b.Cursor[1]

	//
	cell := b.Cells[r][c]

	if cell.IsOpen || cell.IsFlagged {
		return
	}

	if cell.IsMine {
		b.Cells[r][c].IsOpen = true
		b.State = GameOver
		return
	}

	b.open(r, c)

	// クリア判定
	if b.isCleared() {
		b.State = GameClear
	}

}

func (b *Board) isCleared() bool {
	for r := 0; r < cellNums; r++ {

		for c := 0; c < cellNums; c++ {
			if !b.Cells[r][c].IsMine && !b.Cells[r][c].IsOpen {
				return false
			}
		}
	}
	return true
}

func (b *Board) ToggleFlag() {
	r := b.Cursor[0]
	c := b.Cursor[1]

	//
	b.Cells[r][c].IsFlagged = !b.Cells[r][c].IsFlagged
}
