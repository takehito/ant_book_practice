package main

const INF = 100000000

var (
	dx = []int{1, 0, -1, 0}
	dy = []int{0, 1, 0, -1}
)

type pair struct {
	x int
	y int
}

type maze struct {
	field [][]rune
	que   []pair
	d     [][]int
	goal  pair
	start pair
}

func newMaze(field [][]rune) maze {
	var start pair
	var goal pair
	var que []pair

	d := make([][]int, len(field))
	//全ての点をINFで初期化
	for i, x := range field {
		for j, y := range x {
			switch {
			case y == 'S':
				start = pair{
					x: i,
					y: j,
				}
				d[i] = append(d[i], 0)
				que = append(que, pair{x: i, y: j})
				continue
			case y == 'G':
				if field[i][j] == 'G' {
					goal = pair{
						x: i,
						y: j,
					}
				}
			}
			d[i] = append(d[i], INF)
		}
	}

	return maze{
		field: field,
		start: start,
		goal:  goal,
		d:     d,
		que:   que,
	}
}

func (m maze) search() int {
	for len(m.que) > 0 {
		p := m.que[0]
		if len(m.que) == 1 {
			m.que = []pair{}
		} else {
			m.que = m.que[1:]
		}

		if p.x == m.goal.x && p.y == m.goal.y {
			break
		}

		for i := 0; i < 4; i++ {
			// 移動した後の天を(nx, ny)とする
			nx := p.x + dx[i]
			ny := p.y + dy[i]

			// 移動が可能化の判定とすでに訪れたことがあるかの判定(d[nx][ny] != INFなら訪れたことがある)
			if 0 <= nx && nx < len(m.d) && 0 <= ny && ny < len(m.d[nx]) && m.field[nx][ny] != '#' && m.d[nx][ny] == INF {
				// 移動できるならキューに入れ、その点の距離をpからの距離+1で確定する
				m.que = append(m.que, pair{nx, ny})
				m.d[nx][ny] = m.d[p.x][p.y] + 1
			}
		}
	}

	return m.d[m.goal.x][m.goal.y]
}
