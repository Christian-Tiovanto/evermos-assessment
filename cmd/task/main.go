// Package main provides a solution to the task.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Read the map dimensions
	widthStr, _ := reader.ReadString('\n')
	width, _ := strconv.Atoi(strings.TrimSpace(widthStr))
	if width < 1 || width > 100 {
		fmt.Println("Invalid width. Please enter a value between 1 and 100.")
		return
	}

	heightStr, _ := reader.ReadString('\n')
	height, _ := strconv.Atoi(strings.TrimSpace(heightStr))
	if height < 1 || height > 100 {
		fmt.Println("Invalid height. Please enter a value between 1 and 100.")
		return
	}

	// Read the map grid
	grid := make([][]byte, height)
	for i := 0; i < height; i++ {
		rowStr, _ := reader.ReadString('\n')
		row := []byte(strings.TrimSpace(rowStr))
		if len(row) != width {
			fmt.Println("Invalid row length. Please enter exactly", width, "characters.")
			return
		}
		grid[i] = row
	}

	// Count the lakes
	lakes := CountLakes(grid)
	fmt.Println(lakes)
}
