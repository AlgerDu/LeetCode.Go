package main

func main() {}

type spoint struct {
	w int
	h int
}

func maximalSquare(matrix [][]byte) int {

	var mh = len(matrix)

	if mh == 0 {
		return 0
	}

	var mw = len(matrix[0])

	points := make([][]spoint, mh)

	for i := 0; i < mh; i++ {
		points[i] = make([]spoint, mw)
	}

	for i := mh - 1; i >= 0; i-- {
		count := 0
		for j := mw - 1; j >= 0; j-- {
			if matrix[i][j] == 0x31 {
				count++
			} else {
				count = 0
			}

			points[i][j].w = count
		}
	}

	for i := mw - 1; i >= 0; i-- {
		count := 0
		for j := mh - 1; j >= 0; j-- {
			if matrix[j][i] == 0x31 {
				count++
			} else {
				count = 0
			}

			points[j][i].h = count
		}
	}

	max := 0

	for i := 0; i < mh; i++ {
		for j := 0; j < mw; j++ {
			p := points[i][j]

			if p.w > max && p.h > max {
				tmp := p.w

				for ; tmp > max; tmp-- {

					ok := true

					if p.h < tmp {
						tmp = p.h
					}

					for k := 1; k < tmp; k++ {
						if points[i][j+k].h < tmp {
							ok = false
							break
						}
					}

					if ok {
						max = tmp
						break
					}
				}
			}
		}
	}

	return max * max
}
