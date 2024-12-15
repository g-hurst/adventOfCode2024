package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
    "math"
)

func get_input(f_name string) [][]int {
    data, err := os.ReadFile(f_name)
    if err != nil {
        panic(err)
    }

    lines := strings.Split(strings.TrimSpace(string(data)), "\n")
    nums  := make([][]int, len(lines))
    for i, num_str := range lines {
        fields := strings.Fields(num_str)
        nums[i] = make([]int, len(fields))
        for j, n := range fields {
            val, _ := strconv.Atoi(n)
            nums[i][j] = val
        }
    }

    return nums
}

func check_is_safe(row []int) int{
    var curr int
    curr, row = row[0], row[1:]
    is_safe := -1
    is_increasing := curr < row[0]
    for i, val := range row {
        if ((is_increasing != (curr < val)) || (curr == val) || (math.Abs(float64(curr - val)) > 3)) {
            is_safe = i
            break;
        }
        is_increasing = curr < val
        curr = val
    }
    return is_safe
}

func part_1(data [][]int) int {
    safe := 0
    for _, row := range data {
        is_safe := check_is_safe(row)
        if is_safe == -1 {
            safe += 1
        }
    }
    return safe
}


func part_2(data [][]int) int {
    safe := 0
    for _, row := range data {
        // check the row
        is_safe := check_is_safe(row)
        if is_safe == -1 {
            safe += 1
        } else{
            // check star and and compared to array bounds
            start := is_safe - 1
            end   := is_safe + 1
            if start < 0 {
                start = 0
            }
            if end > len(row) {
                end -= 1
            }
            // remove an element one at a time and check agian
            for i:=start; i<=end; i+=1 {
                row_copy := make([]int, len(row))
                copy(row_copy, row)
                row_copy = append(row_copy[:i], row_copy[i+1:]...)
                if check_is_safe(row_copy) == -1 {
                    safe += 1
                    break
                } 
            }
        }
    }
    return safe
}

func main() {
    ex_data := get_input("input_example.txt")
    data    := get_input("input.txt")

    // check part1 example
    example_sol_part1 := part_1(ex_data)
    example_ans_part1 := 2
    if example_sol_part1 != example_ans_part1 {
        fmt.Printf("%d != %d\n", example_sol_part1, example_ans_part1)
        panic("example solution and answer must be equal")
    }
    // get the answer to part1
    sol_part1 := part_1(data)
    fmt.Printf("part1: %d\n", sol_part1)

    // check part2 example
    example_sol_part2 := part_2(ex_data)
    example_ans_part2 := 4
    if example_sol_part2 !=example_ans_part2 {
        fmt.Printf("%d != %d\n", example_sol_part2, example_ans_part2)
        panic("example solution and answer must be equal")
    }
    // get the answer to part2
    sol_part2 := part_2(data)
    fmt.Printf("part2: %d\n", sol_part2)
}

