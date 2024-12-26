package main

import (
    "fmt"
    "os"
    "strings"
)

type Duo [2]int

func get_input(f_name string) ([]string, []string) {
    data, err := os.ReadFile(f_name)
    if err != nil {
        panic(err)
    }

    split := strings.Split(strings.TrimSpace(string(data)), "\n\n")
    graph := strings.Split(strings.TrimSpace(split[0]),    "\n")

    var instructions []string
    for _, line := range strings.Split(strings.TrimSpace(split[1]), "\n") {
        for _, char := range strings.Split(line, "") {
            instructions = append(instructions, char)
        }
    }

    return graph, instructions
}

// func display(points map[Duo]string, x_max, y_max int) {
//     for y:=0; y<y_max; y+=1 {
//         fmt.Printf("|")
//         for x:=0; x<x_max; x+=1 {
//             fmt.Printf("%v", points[Duo{x,y}])
//         }
//         fmt.Printf("|\n")
//     }
// }

func part_1(graph []string, instructions []string) int {
    points := make(map[Duo]string)
    var robot_loc Duo
    for y, line := range graph {
        for x, char := range strings.Split(line, "") {
            pt := Duo{x,y}
            points[pt] = char
            if char == "@" {
                robot_loc = pt
            }
        }
    }

    var update func (start Duo, ins string) Duo
    update = func (start Duo, ins string) Duo {
        var new_pt Duo
        switch ins {
            case "^": new_pt = Duo{start[0],start[1]-1}
            case "v": new_pt = Duo{start[0],start[1]+1}
            case ">": new_pt = Duo{start[0]+1,start[1]}
            case "<": new_pt = Duo{start[0]-1,start[1]}
        }
        switch points[new_pt] {
            case ".":
                points[new_pt] = points[start]
            case "#":
                new_pt = start
            case "O":
                if next_new_pt := update(new_pt, ins); next_new_pt != new_pt {
                    points[next_new_pt] = points[new_pt]
                    points[new_pt] = points[start]
                } else {
                    new_pt = start
                }
        }

        if start != new_pt && points[start] == "@" {
            points[start] = "."
        }
        return new_pt
    }

    for _, ins := range instructions {
        robot_loc = update(robot_loc, ins)
    }

    sum := 0
    for pt, char := range points {
        if char == "O" {
            sum += pt[0] + 100 * pt[1]
        }
    }

    return sum
}

func main() {
    graph, instructions       := get_input("input.txt")
    ex_graph, ex_instructions := get_input("input_example.txt")

    // check part1 example
    example_sol_part1 := part_1(ex_graph, ex_instructions)
    example_ans_part1 := 10092
    if example_sol_part1 != example_ans_part1 {
        fmt.Printf("%d != %d\n", example_sol_part1, example_ans_part1)
        panic("example solution and answer must be equal")
    }
    // get the answer to part1
    sol_part1 := part_1(graph, instructions)
    fmt.Printf("part1: %d\n", sol_part1)

    // // check part2 example
    // example_sol_part2 := part_2(ex_graph, ex_instructions)
    // example_ans_part2 := 1206
    // if example_sol_part2 !=example_ans_part2 {
    //     fmt.Printf("%d != %d\n", example_sol_part2, example_ans_part2)
    //     panic("example solution and answer must be equal")
    // }
    // // get the answer to part2
    // sol_part2 := part_2(graph, instructions)
    // fmt.Printf("part2: %d\n", sol_part2)

}

