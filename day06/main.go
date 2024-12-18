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
func (p *Point) is_valid(x, y int) bool {
    return ( (p.x >= 0) && (p.x < x) &&
             (p.y >= 0) && (p.y < y) )
}
func (p Point) is_same_loc(p2 Point) bool {
    return (p.x==p2.x) && (p.y==p2.y)
}

type Guard struct {
    loc Point
    dir rune
}
func (g Guard) get_next_pt() Point {
    var new_loc Point
    switch g.dir{
        case 'N': new_loc=Point{x:g.loc.x,     y:g.loc.y - 1}
        case 'S': new_loc=Point{x:g.loc.x,     y:g.loc.y + 1}
        case 'E': new_loc=Point{x:g.loc.x + 1, y:g.loc.y}
        case 'W': new_loc=Point{x:g.loc.x - 1, y:g.loc.y}
    }
    return new_loc
}
func (g *Guard) turn() {
    switch g.dir {
        case 'N': g.dir = 'E'
        case 'S': g.dir = 'W' 
        case 'E': g.dir = 'S' 
        case 'W': g.dir = 'N' 
    }
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

    // make a hashmap with all the obstacles
    // and find the guard 
    var guard Guard
    obstacles := make(map[Point]bool)
    for y, row := range grid {
        for x, val := range row {
            if val == '#' {
                obstacles[Point{x:x, y: y}] = true
            } else if val == '^' {
                guard = Guard{Point{x:x, y:y}, 'N'}
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
        new_loc := guard.get_next_pt()
        if !new_loc.is_valid(len(grid[0]), len(grid)) {
            // check if the new point is out of bounds
            is_end = true
        } else if obstacles[new_loc] {
            // turn on a collision
            guard.turn()
        } else {
            // log the point as seen
            guard_path[new_loc] = true
            guard.loc = new_loc
        }
    }

    return len(guard_path)
}

func part_2(data string) int {
    grid := strings.Split(strings.TrimSpace(data), "\n")
    // make a hashmap with all the obstacles
    // and find the guard 
    var guard Guard
    obstacles := make(map[Point]bool)
    for y, row := range grid {
        for x, val := range row {
            if val == '#' {
                obstacles[Point{x:x, y: y}] = true
            } else if val == '^' {
                guard = Guard{Point{x:x, y:y}, 'N'}
            }
        }
    }

    guard_path := make(map[Point]bool)

    is_circular := func(g Guard, obs map[Point]bool) Point{
        circle := g.get_next_pt()
        dummy  := Point{-1,-1}
        if !circle.is_valid(len(grid[0]), len(grid)) || guard_path[circle] {
            // bad if the circular obstacle point is in the previously treaded path
            return dummy
        }
        obs[circle] = true
        g.turn()
        path := make(map[Guard]bool)
        path[g] = true
        is_end := false
        for !is_end {
            new_loc := g.get_next_pt()
            if !new_loc.is_valid(len(grid[0]), len(grid)) {
                is_end = true
            } else if obs[new_loc] {
                g.turn()
            } else {
                g.loc = new_loc
                if path[g] {
                    return circle
                }
                path[g] = true
            }
        }
        return dummy
    }

    is_end := false
    loop_points := make(map[Point]bool)
    guard_path[guard.loc] = true
    for !is_end {
        new_loc := guard.get_next_pt()
        if !new_loc.is_valid(len(grid[0]), len(grid)) {
            is_end = true
        } else if obstacles[new_loc] {
            guard.turn()
        } else {
            // make a copy of obstacles and check for circular path
            obs := make(map[Point]bool)
            for k, v := range obstacles {
                obs[k] = v
            }
            if pt:=is_circular(guard, obs); pt.x!=-1 {
                loop_points[pt] = true
            }
            guard_path[new_loc] = true
            guard.loc = new_loc
        }
    }

    return len(loop_points)
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

    // check part2 example
    example_sol_part2 := part_2(ex_data)
    example_ans_part2 := 6
    if example_sol_part2 !=example_ans_part2 {
        fmt.Printf("%d != %d\n", example_sol_part2, example_ans_part2)
        panic("example solution and answer must be equal")
    }
    // get the answer to part2
    sol_part2 := part_2(data)
    fmt.Printf("part2: %d\n", sol_part2)

}

