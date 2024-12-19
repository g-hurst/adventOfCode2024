package main

import (
    "fmt"
    "os"
    "strings"
    "math"
)



type Point struct {
    x int
    y int
}
func (p *Point) is_valid(x, y int) bool {
    return ( (p.x >= 0) && (p.x < x) &&
             (p.y >= 0) && (p.y < y) )
}

func get_input(f_name string) string {
    data, err := os.ReadFile(f_name)
    if err != nil {
        panic(err)
    }
    return string(data)
}

func part_1(data string) int {
    // get all the frequency points
    freq_groups := make(map[rune][]Point)
    grid := strings.Split(strings.TrimSpace(data), "\n")
    for y, row := range grid {
        for x, val := range row {
            if val != '.' {
                freq_groups[val] = append(freq_groups[val], Point{x:x, y:y})
            }
        }
    }

    x_max := len(grid[0])
    y_max := len(grid)
    antinodes := make(map[Point]bool)
    for _, points := range freq_groups {
        // look through all pairs of two
        for i, p1 := range points[:len(points)-1] {
            for _, p2 := range points[i+1:] {
                // define the antinodes for each point
                dx := int(math.Abs(float64(p1.x - p2.x)))
                dy := int(math.Abs(float64(p1.y - p2.y)))
                var anode1 Point
                var anode2 Point
                switch {
                    case (p1.x <  p2.x) && (p1.y <  p2.y):
                        anode1 = Point{x:p1.x - dx, y:p1.y - dy}
                        anode2 = Point{x:p2.x + dx, y:p2.y + dy}
                    case (p1.x >= p2.x) && (p1.y >= p2.y):
                        anode1 = Point{x:p1.x + dx, y:p1.y + dy}
                        anode2 = Point{x:p2.x - dx, y:p2.y - dy}
                    case (p1.x >= p2.x) && (p1.y <  p2.y):
                        anode1 = Point{x:p1.x + dx, y:p1.y - dy}
                        anode2 = Point{x:p2.x - dx, y:p2.y + dy}
                    case (p1.x <  p2.x) && (p1.y >= p2.y):
                        anode1 = Point{x:p1.x - dx, y:p1.y + dy}
                        anode2 = Point{x:p2.x + dx, y:p2.y - dy}
                }
                // add the antinodes if they are within the grid
                if anode1.is_valid(x_max, y_max) {
                    antinodes[anode1] = true
                }
                if anode2.is_valid(x_max, y_max) {
                    antinodes[anode2] = true
                }

            }
        }
    }
    
    return len(antinodes)
}

func part_2(data string) int {
    freq_groups := make(map[rune][]Point)
    grid := strings.Split(strings.TrimSpace(data), "\n")
    for y, row := range grid {
        for x, val := range row {
            if val != '.' {
                freq_groups[val] = append(freq_groups[val], Point{x:x, y:y})
            }
        }
    }

    x_max := len(grid[0])
    y_max := len(grid)
    antinodes := make(map[Point]bool)
    update_antinodes := func(p_in Point, dx, dy int) {
        // update the antinodes with the point itself and all res points
        antinodes[p_in] = true
        p := Point{x:p_in.x + dx, y:p_in.y + dy}
        for p.is_valid(x_max, y_max) {
            antinodes[p] = true
            p = Point{x:p.x + dx, y:p.y + dy}
        }
    }
    
    for _, points := range freq_groups {
        for i, p1 := range points[:len(points)-1] {
            for _, p2 := range points[i+1:] {
                dx := int(math.Abs(float64(p1.x - p2.x)))
                dy := int(math.Abs(float64(p1.y - p2.y)))
                switch {
                    case (p1.x <  p2.x) && (p1.y <  p2.y):
                        update_antinodes(p1, -dx, -dy)
                        update_antinodes(p2,  dx,  dy)
                    case (p1.x >= p2.x) && (p1.y >= p2.y):
                        update_antinodes(p1,  dx,  dy)
                        update_antinodes(p2, -dx, -dy)
                    case (p1.x >= p2.x) && (p1.y <  p2.y):
                        update_antinodes(p1,  dx, -dy)
                        update_antinodes(p2, -dx,  dy)
                    case (p1.x <  p2.x) && (p1.y >= p2.y):
                        update_antinodes(p1, -dx,  dy)
                        update_antinodes(p2,  dx, -dy)
                }
            }
        }
    }
    
    return len(antinodes)
}


func main() {
    data    := get_input("input.txt")
    ex_data := get_input("input_example.txt")

    // check part1 example
    example_sol_part1 := part_1(ex_data)
    example_ans_part1 := 14
    if example_sol_part1 != example_ans_part1 {
        fmt.Printf("%d != %d\n", example_sol_part1, example_ans_part1)
        panic("example solution and answer must be equal")
    }
    // get the answer to part1
    sol_part1 := part_1(data)
    fmt.Printf("part1: %d\n", sol_part1)

    // check part2 example
    example_sol_part2 := part_2(ex_data)
    example_ans_part2 := 34
    if example_sol_part2 !=example_ans_part2 {
        fmt.Printf("%d != %d\n", example_sol_part2, example_ans_part2)
        panic("example solution and answer must be equal")
    }
    // get the answer to part2
    sol_part2 := part_2(data)
    fmt.Printf("part2: %d\n", sol_part2)

}

