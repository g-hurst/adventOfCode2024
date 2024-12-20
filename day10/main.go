package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
)


func get_input(f_name string) [][]int {
    data, err := os.ReadFile(f_name)
    if err != nil {
        panic(err)
    }

    lines := strings.Split(strings.TrimSpace(string(data)), "\n") 
    m := make([][]int, len(lines))
    for i, line := range lines {
        m[i] = make([]int, len(line))
        for j, n := range strings.Split(strings.TrimSpace(line), "") {
            val, _ := strconv.Atoi(n)
            m[i][j] = val
        }
    }
    return m
}

func part_1(data [][]int) int {
    is_valid := func (x, y int) bool {
        return (x>=0 && x<len(data[0])) && (y>=0 && y <len(data))
    }
    get_trail_score := func (x, y int) int {
        // keep track of all points seen 
        seen_points := make(map[[2]int]bool)

        // recursive function to get a trail score
        var _get_trail_score func (x, y int) int
        _get_trail_score = func (x, y int) int {
            // return 0 if the point has been seen
            point := [2]int{x,y}
            if seen_points[point] {
                return 0
            }
            seen_points[point] = true
            // if the point is 9, return 1 to indicate end was found
            if data[y][x] == 9 {
                return 1
            }
            // recurse down the x and y directions, whenever a difference of
            // 1 is found between the current and new points
            dirs := [2]int{-1,1}
            trail_score := 0
            for _, dy := range dirs {
                if is_valid(x, y+dy) && ((data[y+dy][x]-data[y][x]) == 1) {
                    trail_score += _get_trail_score(x, y+dy)
                }
            }
            for _, dx := range dirs {
                if is_valid(x+dx, y)  && ((data[y][x+dx]-data[y][x]) == 1) {
                    trail_score += _get_trail_score(x+dx, y)
                }
            }
            return trail_score
        }
        return _get_trail_score(x,y)
    }

    // get the sum of each trail score
    sum := 0
    for y, line := range data{
        for x, val := range line {
            if val == 0 {
                sum += get_trail_score(x, y)
            }
        }
    }

    return sum
}


func part_2(data [][]int) int {
    is_valid := func (x, y int) bool {
        return (x>=0 && x<len(data[0])) && (y>=0 && y <len(data))
    }
    var get_trail_score func (x, y int) int 
    get_trail_score = func (x, y int) int {
        // if the point is 9, return 1 to indicate end was found
        if data[y][x] == 9 {
            return 1
        }
        // recurse down the x and y directions, whenever a difference of
        // 1 is found between the current and new points
        dirs := [2]int{-1,1}
        trail_score := 0
        for _, dy := range dirs {
            if is_valid(x, y+dy) && ((data[y+dy][x]-data[y][x]) == 1) {
                trail_score += get_trail_score(x, y+dy)
            }
        }
        for _, dx := range dirs {
            if is_valid(x+dx, y)  && ((data[y][x+dx]-data[y][x]) == 1) {
                trail_score += get_trail_score(x+dx, y)
            }
        }
        return trail_score
    }

    // get the sum of each trail score
    sum := 0
    for y, line := range data{
        for x, val := range line {
            if val == 0 {
                sum += get_trail_score(x, y)
            }
        }
    }

    return sum
}

func main() {
    data    := get_input("input.txt")
    ex_data := get_input("input_example.txt")

    // check part1 example
    example_sol_part1 := part_1(ex_data)
    example_ans_part1 := 36
    if example_sol_part1 != example_ans_part1 {
        fmt.Printf("%d != %d\n", example_sol_part1, example_ans_part1)
        panic("example solution and answer must be equal")
    }
    // get the answer to part1
    sol_part1 := part_1(data)
    fmt.Printf("part1: %d\n", sol_part1)

    // check part2 example
    example_sol_part2 := part_2(ex_data)
    example_ans_part2 := 81
    if example_sol_part2 !=example_ans_part2 {
        fmt.Printf("%d != %d\n", example_sol_part2, example_ans_part2)
        panic("example solution and answer must be equal")
    }
    // get the answer to part2
    sol_part2 := part_2(data)
    fmt.Printf("part2: %d\n", sol_part2)

}

