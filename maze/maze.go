package main

import (
	"os"
	"fmt"
)

func readMaze(filename string) [][]int{
	file, err := os.Open(filename)

	if err != nil{
		panic(err)
	}

	var row, col int

	// read file
	// 1. first line: row col
	fmt.Fscanf(file, "%d %d", &row, &col) // pass the pointer
	// 2. create a 2D slice and read the rest of file
	maze := make([][]int, row)

	for i := range maze{
		maze[i] = make([]int, col)
		for j := range maze[i]{
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	// return the maze
	return maze
}

func printMaze(maze [][]int){
	for row := range maze{
		for col := range maze[row]{
			fmt.Printf("%3d", maze[row][col])
		}
		fmt.Println()
	}
}
type point struct {
	i, j int
}

// define four directions: up, left, down, right
var dirs = [4]point{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

func (self point) add(other point) point{
	return point{self.i + other.i, self.j + other.j}
}

// get the value of a point in a grid
func (p point) at(grid [][]int) (int, bool){
	// check out of bounds
	if p.i < 0 || p.i >= len(grid){
		return -1, false
	}
	if p.j < 0 || p.j >= len(grid[0]){
		return -1, false
	}

	return grid[p.i][p.j], true
}


// BFS
func walk(maze [][] int, start point, end point) [][]int{
	// record steps
	steps := make([][] int, len(maze))
	for i := range maze{
		steps[i] = make([]int, len(maze[i]))
	}

	// create a queue
	Q := []point{start}

	for len(Q) > 0{
		// poll the first element
		cur := Q[0]
		Q = Q[1:]

		if cur == end{
			break
		}

		for _, dir := range dirs{
			next := cur.add(dir)

			// validate next:
			// 1. next point in maze can be visited(equals to 0)
			// 2. next point has not been visited(we have not marked it in steps)
			// 3. next point is not start point(because start point has value 0 in both maze and steps)
			if val, inbound := next.at(maze); !inbound || val != 0{
				continue
			}

			if val, inbound := next.at(steps); !inbound || val != 0{
				continue
			}

			if(next == start){
				continue
			}

			// now we know next is valid:
			// 1. add to queue
			Q = append(Q, next)
			// 2. fill in steps
			curStep, _ := cur.at(steps)
			steps[next.i][next.j] = curStep + 1
		}
	}
	return steps
}

func main() {
	// 1. read the maze from file
	maze := readMaze("maze/maze.in")

	// 2. print the maze to check for correctness
	printMaze(maze)
	fmt.Println()

	// 3. walk through the maze
	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})

	printMaze(steps)
}
