package main

import (
    "fmt"
    "os"
    "strings"
    "reflect"
    "slices"
)



func get_input(f_name string) string {
    data, err := os.ReadFile(f_name)
    if err != nil {
        panic(err)
    }
    return string(data)
}


type Plot struct {
    area int
    perimeter int
    plant_type rune
}
func (p Plot) get_price() int {
    return p.area * p.perimeter
}

func part_1(data string) int {
    type Duo [2]int
    grid := strings.Split(strings.TrimSpace(data), "\n")
    seen_point := make(map[Duo]bool)

    is_valid := func (x, y int) bool {
        return (x>=0 && x<len(grid[0])) && (y>=0 && y < len(grid))
    }
    calc_plot := func (x, y int) *Plot{
        var _calc func (x, y int) Duo
        _calc = func (x, y int) Duo {
            if seen_point[Duo{x,y}] {
                return Duo{0, 0}
            }
            seen_point[Duo{x,y}] = true

            area := 1
            perimeter := 0
            __update := func (x_next, y_next int) {
                if is_valid(x_next, y_next) && (grid[y_next][x_next]==grid[y][x]) {
                    ap := _calc(x_next, y_next)
                    area      += ap[0]
                    perimeter += ap[1]
                } else {
                    perimeter += 1
                }
            }
            __update(x+1, y)
            __update(x-1, y)
            __update(x,   y-1)
            __update(x,   y+1)

            return Duo{area, perimeter}
        }
        ap := _calc(x, y)
        return &Plot{area:ap[0], perimeter:ap[1], plant_type:rune(grid[y][x])}
    }

    // calculate the area and perimeter for each square
    total_price := 0
    for y, row := range grid {
        for x, _ := range row {
            if !seen_point[Duo{x,y}] {
                plot := calc_plot(x,y)
                total_price += plot.get_price()
            }
        }
    }

    return total_price
}

func part_2(data string) int {
    const (
        NORTH = 0
        EAST  = 1
        SOUTH = 2
        WEST  = 3
    )
    type Duo [2]int
    grid := strings.Split(strings.TrimSpace(data), "\n")
    seen_point := make(map[Duo]bool)

    is_valid := func (x, y int) bool {
        return (x>=0 && x<len(grid[0])) && (y>=0 && y < len(grid))
    }
    calc_plot := func (x, y int) *Plot{
        perim_points := make(map[int][]Duo)

        // funciton to calculate the area and log all perimiter points
        var _calc func (x, y int) int
        _calc = func (x, y int) int {
            // recursive end
            if seen_point[Duo{x,y}] {
                return 0
            }
            seen_point[Duo{x,y}] = true
            // get area
            area := 1
            __update := func (x_next, y_next , perim_loc int) {
                if is_valid(x_next, y_next) && (grid[y_next][x_next]==grid[y][x]) {
                    area += _calc(x_next, y_next)
                } else {
                    // log when a perimiter point is found
                    perim_points[perim_loc] = append(perim_points[perim_loc] ,Duo{x,y})
                }
            }
            __update(x+1, y,   EAST)
            __update(x-1, y,   WEST)
            __update(x,   y-1, NORTH)
            __update(x,   y+1, SOUTH)
            return area
        }

        area := _calc(x, y)
        perimeter := 0
        // loop over each direction 
        for len(perim_points) > 0 {
            dir := reflect.ValueOf(perim_points).MapKeys()[0].Interface().(int)
            seen_perim := make(map[Duo]bool)
            // look at the points in the direction
            for _,point := range perim_points[dir] {
                if !seen_perim[point] {
                    perimeter += 1
                    queue := []Duo{point}
                    // make a queue to check new dirs
                    for len(queue) > 0 {
                        pt := queue[0]
                        queue = queue[1:]
                        if !seen_perim[pt] {
                            seen_perim[pt] = true
                            // add each adjacent point if in same direction
                            for _, adj_pt := range []Duo{Duo{pt[0]+1,pt[1]},
                                                         Duo{pt[0],  pt[1]-1},
                                                         Duo{pt[0]-1,pt[1]},
                                                         Duo{pt[0],  pt[1]+1}} {
                                if slices.Contains(perim_points[dir], adj_pt) {
                                    queue = append(queue, adj_pt)
                                }
                             }
                        }
                    }
                }
            }
            delete(perim_points, dir)
        }

        return &Plot{area:area, perimeter:perimeter, plant_type:rune(grid[y][x])}
    }

    // calculate the area and perimeter for each square
    total_price := 0
    for y, row := range grid {
        for x, _ := range row {
            if !seen_point[Duo{x,y}] {
                plot := calc_plot(x,y)
                total_price += plot.get_price()
            }
        }
    }

    return total_price
}

func main() {
    data    := get_input("input.txt")
    ex_data := get_input("input_example.txt")

    // check part1 example
    example_sol_part1 := part_1(ex_data)
    example_ans_part1 := 1930
    if example_sol_part1 != example_ans_part1 {
        fmt.Printf("%d != %d\n", example_sol_part1, example_ans_part1)
        panic("example solution and answer must be equal")
    }
    // get the answer to part1
    sol_part1 := part_1(data)
    fmt.Printf("part1: %d\n", sol_part1)

    // check part2 example
    example_sol_part2 := part_2(ex_data)
    example_ans_part2 := 1206
    if example_sol_part2 !=example_ans_part2 {
        fmt.Printf("%d != %d\n", example_sol_part2, example_ans_part2)
        panic("example solution and answer must be equal")
    }
    // get the answer to part2
    sol_part2 := part_2(data)
    fmt.Printf("part2: %d\n", sol_part2)

}

