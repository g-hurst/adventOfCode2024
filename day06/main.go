package main

import (
    "fmt"
    "os"
    "strings"
)

type Point struct {
    x int
    y int
}

type Guard struct {
    loc Point
    dir rune
}

func get_input(f_name string) string {
    data, err := os.ReadFile(f_name)
    if err != nil {
        panic(err)
    }
    return string(data)
}

func part_1(data string) int {
    grid := strings.Split(strings.TrimSpace(data), "\n")

    // get the guard and obstacle locations
    var guard Guard
    var obstacles []Point
    for y, row := range grid {
        for x, val := range row {
            if val == '#' {
                obstacles = append(obstacles, Point{x, y})
            } else if val == '^' {
                guard = Guard{Point{x,y}, 'N'}
            }
        }
    }

    // create a hasmap to keep track of the visited positions
    guard_path := make(map[Point]bool)
    guard_path[guard.loc] = true

    // loop unitl the guard goes out of bounds
    is_end := false
    for !is_end {
        // create a new potential location
        var new_loc Point
        switch guard.dir{
            case 'N': new_loc= Point{guard.loc.x,guard.loc.y - 1}
            case 'S': new_loc= Point{guard.loc.x,guard.loc.y + 1}
            case 'E': new_loc= Point{guard.loc.x + 1,guard.loc.y}
            case 'W': new_loc= Point{guard.loc.x - 1,guard.loc.y}
        }

        // check if the new point is out of bounds
        if ((new_loc.x < 0) || (new_loc.x >= len(grid[0])) ||
            (new_loc.y < 0) || (new_loc.y >= len(grid))) {
            is_end = true
        } else {
            // check if the new point is on an obstacle
            is_collision := false
            for _, obs := range obstacles {
                if new_loc == obs {
                    is_collision = true
                    break
                }
            }
            // turn on a collision or log the point as seen
            if is_collision {
                switch guard.dir {
                    case 'N': guard.dir = 'E'
                    case 'S': guard.dir = 'W' 
                    case 'E': guard.dir = 'S' 
                    case 'W': guard.dir = 'N' 

                }
            } else {
                guard_path[new_loc] = true
                guard.loc = new_loc
            }
        }
    }

    return len(guard_path)
}


func main() {
    data    := get_input("input.txt")
    ex_data := get_input("input_example.txt")

    // check part1 example
    example_sol_part1 := part_1(ex_data)
    example_ans_part1 := 41
    if example_sol_part1 != example_ans_part1 {
        fmt.Printf("%d != %d\n", example_sol_part1, example_ans_part1)
        panic("example solution and answer must be equal")
    }
    // get the answer to part1
    sol_part1 := part_1(data)
    fmt.Printf("part1: %d\n", sol_part1)

    // // check part2 example
    // example_sol_part2 := part_2(ex_data)
    // example_ans_part2 := 9
    // if example_sol_part2 !=example_ans_part2 {
    //     fmt.Printf("%d != %d\n", example_sol_part2, example_ans_part2)
    //     panic("example solution and answer must be equal")
    // }
    // // get the answer to part2
    // sol_part2 := part_2(data)
    // fmt.Printf("part2: %d\n", sol_part2)

}

