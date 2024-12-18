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

    m := make([][]int, 0)
    for i, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
        line_split := strings.Split(line, ": ")
        key, _     := strconv.Atoi(line_split[0])
        m       = append(m, []int{key})
        for _, n := range strings.Split(strings.TrimSpace(line_split[1]), " ") {
            val, _ := strconv.Atoi(n)
            m[i] = append(m[i], val)
        }
    }
    return m
}


func reduce (nums []int, op rune) []int {
    new_nums := append([]int{}, nums[1:]...) 
    switch op {
    case '+':
        new_nums[0] = nums[0] + nums[1]
    case '*':
        new_nums[0] = nums[0] * nums[1]
    case '|':
        val, _ := strconv.Atoi(strconv.Itoa(nums[0])+strconv.Itoa(nums[1]))
        new_nums[0] = val
    }
    return new_nums
}

func search (target int, nums []int, ops []rune) bool {
    if len(nums) == 1 {
        return target == nums[0]
    }
    for _, op := range ops {
        nums_cpy := reduce(nums, op)
        if search(target, nums_cpy, ops) {
            return true
        }
    }
    return false
}

func part_1(data [][]int) int {
    ops := []rune{'+', '*'}
    sum := 0
    for _, d := range data {
        target := d[0]
        nums   := d[1:]
        if search(target, nums, ops) {
            sum += target
        }
    }
    return sum
}

func part_2(data [][]int) int {
    ops := []rune{'+', '*', '|'}
    sum := 0
    for _, d := range data {
        target := d[0]
        nums   := d[1:]
        if search(target, nums, ops) {
            sum += target
        }
    }
    return sum
}

func main() {
    data    := get_input("input.txt")
    ex_data := get_input("input_example.txt")

    // check part1 example
    example_sol_part1 := part_1(ex_data)
    example_ans_part1 := 3749
    if example_sol_part1 != example_ans_part1 {
        panic("example solution and answer must be equal")
    }
    // get the answer to part1
    sol_part1 := part_1(data)
    fmt.Printf("part1: %d\n", sol_part1)

    // check part2 example
    example_sol_part2 := part_2(ex_data)
    example_ans_part2 := 11387
    if example_sol_part2 !=example_ans_part2 {
        fmt.Printf("%d != %d\n", example_sol_part2, example_ans_part2)
        panic("example solution and answer must be equal")
    }
    // get the answer to part2
    sol_part2 := part_2(data)
    fmt.Printf("part2: %d\n", sol_part2)

}

