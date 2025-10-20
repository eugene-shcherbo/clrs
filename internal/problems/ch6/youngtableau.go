package ch6

import (
	"errors"
	"strconv"
)

type tableauItem struct {
	val int
}

type YoungTableau struct {
	items      [][]*tableauItem
	rows, cols int
}

func NewYoungTableaue(tableau [][]string) *YoungTableau {
	if len(tableau) == 0 {
		return &YoungTableau{nil, 0, 0}
	}

	items := make([][]*tableauItem, len(tableau))

	for r, row := range tableau {
		items[r] = make([]*tableauItem, len(row))

		for c, col := range row {
			if col == "x" {
				items[r][c] = nil
			}

			parsed, err := strconv.Atoi(col)

			if err != nil {
				items[r][c] = nil
			} else {
				items[r][c] = &tableauItem{parsed}
			}
		}
	}

	return &YoungTableau{
		items,
		len(items),
		len(items[0]),
	}
}

func (t *YoungTableau) IsEmpty() bool {
	return len(t.items) == 0 || t.items[0][0] == nil
}

func (t *YoungTableau) IsFull() bool {
	return len(t.items) == 0 || t.items[t.rows-1][t.cols-1] != nil
}

func (t *YoungTableau) Peek() (int, error) {
	if t.IsEmpty() {
		return 0, errors.New("tableau is empty")
	}

	return t.items[0][0].val, nil
}

func (t *YoungTableau) PopMin() (int, error) {
	if t.IsEmpty() {
		return 0, errors.New("tableau is empty")
	}

	item := t.items[0][0]
	last := t.items[len(t.items)-1][len(t.items[0])-1]

	t.items[0][0] = last
	t.items[len(t.items)-1][len(t.items[0])-1] = nil

	i, j := 0, 0

	for {
		right := getItem(t, i, j+1)
		bottom := getItem(t, i+1, j)

		sRow, sCol := i, j

		if less(right, t.items[sRow][sCol]) {
			sRow, sCol = i, j+1
		}

		if less(bottom, t.items[sRow][sCol]) {
			sRow, sCol = i+1, j
		}

		if t.items[sRow][sCol] == t.items[i][j] {
			break
		}

		t.items[sRow][sCol], t.items[i][j] = t.items[i][j], t.items[sRow][sCol]
		i, j = sRow, sCol
	}

	return item.val, nil
}

func (t *YoungTableau) Add(value int) error {
	if t.IsFull() {
		return errors.New("tableau is full")
	}

	t.items[t.rows-1][t.cols-1] = &tableauItem{value}

	i, j := t.rows-1, t.cols-1
	for {
		top := getItem(t, i-1, j)
		left := getItem(t, i, j-1)

		bRow, bCol := i, j
		if i-1 >= 0 && less(t.items[bRow][bCol], top) {
			bRow, bCol = i-1, j
		}

		if j-1 >= 0 && less(t.items[bRow][bCol], left) {
			bRow, bCol = i, j-1
		}

		if t.items[bRow][bCol] == t.items[i][j] {
			break
		}

		t.items[bRow][bCol], t.items[i][j] = t.items[i][j], t.items[bRow][bCol]
		i, j = bRow, bCol
	}

	return nil
}

func (t *YoungTableau) Exists(item int) bool {
	if t.IsEmpty() {
		return false
	}

	for i, j := t.rows-1, 0; i >= 0 && j < t.cols; {
		if t.items[i][j] == nil || t.items[i][j].val > item {
			i = i - 1
		} else if t.items[i][j].val < item {
			j = j + 1
		} else {
			return true
		}
	}

	return false
}

func less(a, b *tableauItem) bool {
	if a == nil {
		return false
	}

	if b == nil {
		return true
	}

	return a.val < b.val
}

func getItem(tableau *YoungTableau, i, j int) *tableauItem {
	if i < 0 || j < 0 || i >= tableau.rows || j >= tableau.cols {
		return nil
	}

	return tableau.items[i][j]
}
