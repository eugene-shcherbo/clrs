package ch6_test

import (
	"clrs/internal/problems/ch6"
	"clrs/test/assert"
	"testing"
)

func TestPopMin(t *testing.T) {
	tableData := [][]string{
		{"-10", "-5", "3", "12", "25"},
		{"12", "13", "14", "15", "26"},
		{"123", "124", "124", "x", "x"},
	}

	table := ch6.NewYoungTableaue(tableData)
	want := []int{-10, -5, 3, 12, 12, 13, 14, 15, 25, 26, 123, 124, 124}

	got := []int{}
	for !table.IsEmpty() {
		minimum, _ := table.PopMin()
		got = append(got, minimum)
	}

	assert.SlicesEqual(t, want, got)
}

func TestAdd(t *testing.T) {
	table := ch6.NewYoungTableaue([][]string{
		{"x", "x", "x", "x", "x"},
		{"x", "x", "x", "x", "x"},
		{"x", "x", "x", "x", "x"},
	})

	insertions := []int{124, 123, 15, 26, 3, -5, 25, 12, 13, -10, 12, 14, 124}
	want := []int{124, 123, 15, 15, 3, -5, -5, -5, -5, -10, -10, -10, -10}

	got := []int{}
	for _, item := range insertions {
		_ = table.Add(item)

		minimum, _ := table.Peek()
		got = append(got, minimum)
	}

	assert.SlicesEqual(t, want, got)
}

func TestExists(t *testing.T) {
	table := ch6.NewYoungTableaue([][]string{
		{"-10", "-5", "3", "12", "25"},
		{"12", "13", "14", "15", "26"},
		{"123", "124", "124", "x", "x"},
	})

	assert.True(t, table.Exists(26))
	assert.True(t, table.Exists(-5))
	assert.True(t, table.Exists(124))
	assert.True(t, table.Exists(-10))
	assert.False(t, table.Exists(-9))
	assert.False(t, table.Exists(125))
}
