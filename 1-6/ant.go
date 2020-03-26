package ant

import (
	"math"
)

func antWalkMin(l int, ants []int) int {
	var minTime float64
	// 最小の時間を計算
	for _, ant := range ants {
		absAnt := 0
		if l%2 == 0 {
			half := l / 2
			if ant <= half {
				absAnt = ant - half
			} else if ant >= half+1 {
				absAnt = ant - half
			}
		} else {
			half := l/2 + 1
			if ant <= half {
				absAnt = half - ant
			} else if ant >= half {
				absAnt = ant - half
			}
		}
		minTime = math.Min(minTime, math.Abs(float64(absAnt)))
	}

	return l/2 - int(minTime)
}

func antWalkMax(l int, ants []int) int {
	var maxAnt float64
	// 最大の時間を計算
	for _, ant := range ants {
		absAnt := 0
		if l%2 == 0 {
			half := l / 2
			if ant <= half {
				absAnt = l - ant
			} else if ant >= half+1 {
				absAnt = ant
			}
		} else {
			half := l/2 + 1
			if ant <= half {
				absAnt = ant - l
			} else if ant >= half {
				absAnt = ant
			}
		}
		maxAnt = math.Max(maxAnt, math.Abs(float64(absAnt)))
	}

	return int(maxAnt) + 1
}
