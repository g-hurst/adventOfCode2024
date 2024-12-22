package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
    "regexp"
)


func get_input(f_name string) [][]int{
    data, err := os.ReadFile(f_name)
    if err != nil {
        panic(err)
    }
    exp := regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)

    lines := strings.Split(strings.TrimSpace(string(data)), "\n")
    m := make([][]int, len(lines))
    for i, line := range lines {
        for _, num := range  exp.FindStringSubmatch(line)[1:] {
            val, _ := strconv.Atoi(num)
            m[i] = append(m[i], val)
        }
    }
    return m
}

func part_1(data [][]int, x_max, y_max int) int {
    var quadrants [4]int
    const seconds = 100
    for _, points := range data {
        x := (points[0] + seconds * points[2]) 
        y := (points[1] + seconds * points[3])
        x = (x % x_max + x_max) % x_max 
        y = (y % y_max + y_max) % y_max
        if x < x_max / 2 {
            if y < y_max / 2 {
                quadrants[0] += 1
            } else if y > y_max / 2 {
                quadrants[1] += 1
            }
        } else if x > x_max / 2 {
            if y < y_max / 2 {
                quadrants[2] += 1
            } else if y > y_max / 2 {
                quadrants[3] += 1
            }
        }
    }

    product := 1
    for i:=0; i<len(quadrants); i+=1 {
        product *= quadrants[i]
    }

    return product
}

func display(data [][]int, x_max, y_max int) {
    points := make(map[[2]int]bool)
    for _,d := range data {
        points[[2]int{d[0],d[1]}] = true
    }

    for y:=0; y<y_max; y+=1 {
        for x:=0; x<x_max; x+=1 {
            if points[[2]int{x,y}] {
                fmt.Printf("#")
            } else {
                fmt.Printf(" ")
            }
        }
        fmt.Println("|")
    }
}

func part_2(data [][]int, x_max, y_max int) int {
    update := func(seconds int) map[[2]int]bool {
        seen_points := make(map[[2]int]bool)
        for _, points := range data {
            x := (points[0] + seconds * points[2]) 
            y := (points[1] + seconds * points[3])
            x = (x % x_max + x_max) % x_max 
            y = (y % y_max + y_max) % y_max
            points[0] = x
            points[1] = y
            seen_points[[2]int{x,y}] = true
        }
        return seen_points
    }

    // check for all the points to be unique
    num_updates := 1
    uniques := update(1)
    for  len(uniques) != len(data) {
        uniques = update(1)
        num_updates += 1
    }
    display(data, x_max, y_max)

    return num_updates
}

func main() {
    data    := get_input("input.txt")
    ex_data := get_input("input_example.txt")

    // check part1 example
    example_sol_part1 := part_1(ex_data, 11, 7)
    example_ans_part1 := 12
    if example_sol_part1 != example_ans_part1 {
        fmt.Printf("%d != %d\n", example_sol_part1, example_ans_part1)
        panic("example solution and answer must be equal")
    }
    // get the answer to part1
    sol_part1 := part_1(data, 101, 103)
    fmt.Printf("part1: %d\n", sol_part1)

    // get the answer to part2
    sol_part2 := part_2(data, 101, 103)
    fmt.Printf("part2: %d\n", sol_part2)
}

