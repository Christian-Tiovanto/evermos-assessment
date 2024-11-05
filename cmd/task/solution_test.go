package main_test

import (
	"testing"

	solution "github.com/mughieams/evermos-assessment/cmd/task"
	"github.com/stretchr/testify/assert"
)

func Test_CountLakes(t *testing.T) {
	tests := []struct {
		name string
		grid [][]byte
		want int
	}{
		{
			name: "Sample Test Case",
			grid: [][]byte{
				{'#', '.', '#', '#', '#'},
				{'.', '.', '#', '#', '#'},
				{'#', '#', '.', '#', '#'},
				{'#', '#', '#', '#', '.'},
			},
			want: 2,
		},
		{
			name: "Empty Lake",
			grid: [][]byte{
				{'#', '#', '#', '#', '#'},
				{'#', '#', '#', '#', '#'},
				{'#', '#', '#', '#', '#'},
				{'#', '#', '#', '#', '#'},
				{'#', '#', '#', '#', '#'},
			},
			want: 0,
		},
		{
			name: "Single Lake",
			grid: [][]byte{
				{'.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.'},
			},
			want: 1,
		},
		{
			name: "Single Lake with Island",
			grid: [][]byte{
				{'.', '.', '.', '.', '.'},
				{'.', '#', '#', '#', '.'},
				{'.', '#', '#', '#', '.'},
				{'.', '#', '#', '#', '.'},
				{'.', '.', '.', '.', '.'},
			},
			want: 1,
		},
		{
			name: "Cross Lakes",
			grid: [][]byte{
				{'.', '#', '.'},
				{'#', '.', '#'},
				{'.', '#', '.'},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := solution.CountLakes(tt.grid)
			assert.Equal(t, tt.want, got)
		})
	}
}
