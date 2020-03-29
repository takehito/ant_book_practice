package main

func fill(field [][]string, x int, y int) [][]string {
	// 今いるところを置き換える
	field[x][y] = "."

	// 移動する８方向をループ
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			// x方向にdx y方向にdy 移動した場所をnx, ny)とする
			nx := x + dx
			ny := y + dy

			if 0 <= nx && nx < len(field) && 0 <= ny && ny < len(field[nx]) {
				if field[nx][ny] == "W" {
					fill(field, nx, ny)
				}
			}
		}
	}
	return field
}

func count(field [][]string) int {
	var result int
	for i, _ := range field {
		for j, _ := range field[i] {
			if field[i][j] == "W" {
				field = fill(field, i, j)
				result++
			}
		}
	}
	return result
}
