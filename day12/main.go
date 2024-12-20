package main

import (
    "fmt"
    "os"
    "strings"
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

            area      := 1
            perimeter := 0
            dirs := Duo{-1,1}
            for _, d := range dirs {
                // check points in y direction
                if is_valid(x, y+d) && (grid[y+d][x]==grid[y][x]) {
                    ap        := _calc(x, y+d)
                    area      += ap[0]
                    perimeter += ap[1]
                } else {
                    perimeter += 1
                }
                // check points in x direction
                if is_valid(x+d, y) && (grid[y][x+d]==grid[y][x]) {
                    ap        := _calc(x+d, y)
                    area      += ap[0]
                    perimeter += ap[1]
                } else {
                    perimeter += 1
                }
            }
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

            area      := 1
            perimeter := 0
            dirs := Duo{-1,1}
            for _, d := range dirs {
                // check points in y direction
                if is_valid(x, y+d) && (grid[y+d][x]==grid[y][x]) {
                    ap        := _calc(x, y+d)
                    area      += ap[0]
                    perimeter += ap[1]
                } else {
                    perimeter += 1
                }
                // check points in x direction
                if is_valid(x+d, y) && (grid[y][x+d]==grid[y][x]) {
                    ap        := _calc(x+d, y)
                    area      += ap[0]
                    perimeter += ap[1]
                } else {
                    perimeter += 1
                }
            }
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
                fmt.Printf("%v| %v %v -> %v \n",
                                string(plot.plant_type),
                                plot.area,
                                plot.perimeter,
                                plot.get_price())
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

