package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

type CellSate int

const (
	WIDTH      = 25
	HEIGHT     = 25
	SPEED      = 50
	BACKGROUND = '-'
	CELL       = '#'

	DEAD CellSate = iota
	ALIVE
)

type Cell struct {
	state CellSate
}

var grid [HEIGHT][WIDTH]Cell

func init_grid() {
	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			grid[i][j].state = DEAD
		}
	}
}

func gen_next() {
	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			alive_count := 0
			for k := -1; k <= 1; k++ {
				for l := -1; l <= 1; l++ {
					if k == 0 && l == 0 {
						continue
					}
					if i+k < HEIGHT && int(i+k) >= 0 && j+l < WIDTH && int(j+l) >= 0 {
						if grid[i+k][j+l].state == ALIVE {
							alive_count++
						}
					}
				}
			}
			switch alive_count {
			case 0, 1:
				grid[i][j].state = DEAD
			case 2, 3:
				if grid[i][j].state == DEAD && alive_count == 3 {
					grid[i][j].state = ALIVE
				}
			default:
				grid[i][j].state = DEAD
			}
		}
	}
}

func print_grid() int {
	alive_count := 0
	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			if grid[i][j].state == ALIVE {
				alive_count++
				fmt.Printf("%c", CELL)
			} else {
				fmt.Printf("%c", BACKGROUND)
			}

		}
		fmt.Print("\n")
	}

	return alive_count
}

func clearConsole() {
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func main() {

	init_grid()
	for i := 0; i < WIDTH/5; i++ {
		for j := 0; j < HEIGHT/5; j++ {
			grid[i][j].state = ALIVE
		}
	}

	for cells_alive := print_grid(); cells_alive != 0; cells_alive = print_grid() {
		time.Sleep(SPEED * time.Millisecond)
		gen_next()
		clearConsole()
	}

	fmt.Print("GAME OVER\n")

}
