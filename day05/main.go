package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
)

func get_input(f_name string) ([][]int, [][]int) {
    data, err := os.ReadFile(f_name)
    if err != nil {
        panic(err)
    }
    split := strings.Split(strings.TrimSpace(string(data)), "\n\n")

    parse := func(seperator string, idx int) [][]int {
        parsed := make([][]int, 0)
        for i, line := range strings.Split(split[idx], "\n") {
            parsed = append(parsed, make([]int, 0))
            for _, n := range strings.Split(line, seperator) {
                val, _ := strconv.Atoi(n)
                parsed[i] = append(parsed[i], val)
            }
        }
        return parsed
    }

    orders := parse("|", 0)
    lists  := parse(",", 1)

    return orders, lists
}

func make_orders_map(orders [][]int) map[int]map[int]bool {
    orders_map := make(map[int]map[int]bool)
    for _, order := range orders {
        if orders_map[order[0]] == nil {
            orders_map[order[0]] = make(map[int]bool)
        }
        orders_map[order[0]][order[1]] = true
    }
    return orders_map
}

func part_1(orders, lists [][]int) int {
    // create a hashmap from the orders list
    orders_map := make_orders_map(orders)

    middles := 0
    for _, list := range lists {
        is_valid := true
        // reverse through the list, and search for violated rules
        for i:=len(list)-1; i>0 && is_valid; i-=1 {
            for j:=i-1; j>=0 && is_valid; j-=1 {
                if orders_map[list[i]][list[j]] {
                    is_valid = false
                }
            }
        }
        // if the list is valid, add middle element to the sum
        if is_valid {
            middles += list[len(list)/2]
        }
    }
    return middles
}


func part_2(orders, lists [][]int) int {
    // create a hashmap from the orders list
    orders_map := make_orders_map(orders)

    middles := 0
    for _, list := range lists {
        is_reordered := false
        // reverse through the list, and search for violated rules
        for i:=len(list)-1; i>0; i-=1 {
            for j:=i-1; j>=0; j-=1 {
                // if a rule is violated, modify the list
                // [..., j, ..., i, ...] -> [..., i, j, ...]
                if orders_map[list[i]][list[j]] {
                    is_reordered = true

                    cpy := make([]int, len(list))
                    copy(cpy, list)
                    violator := cpy[i]
                    list = append(list[:i], list[i+1:]...)
                    list = append(append(cpy[:j], violator), list[j:]...)

                    i += 1
                    break
                }
            }
        }
        if is_reordered {
            middles += list[len(list)/2]
        }
    }
    return middles
}

func main() {
    orders, lists       := get_input("input.txt")
    ex_orders, ex_lists := get_input("input_example.txt")

    // check part1 example
    example_sol_part1 := part_1(ex_orders, ex_lists)
    example_ans_part1 := 143
    if example_sol_part1 != example_ans_part1 {
        fmt.Printf("%d != %d\n", example_sol_part1, example_ans_part1)
        panic("example solution and answer must be equal")
    }
    // get the answer to part1
    sol_part1 := part_1(orders, lists)
    fmt.Printf("part1: %d\n", sol_part1)

    // check part2 example
    example_sol_part2 := part_2(ex_orders, ex_lists)
    example_ans_part2 := 123
    if example_sol_part2 !=example_ans_part2 {
        fmt.Printf("%d != %d\n", example_sol_part2, example_ans_part2)
        panic("example solution and answer must be equal")
    }
    // get the answer to part2
    sol_part2 := part_2(orders, lists)
    fmt.Printf("part2: %d\n", sol_part2)

}

