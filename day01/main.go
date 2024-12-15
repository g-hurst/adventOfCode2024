package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
    "slices"
    "math"

)

func get_input(f_name string) ([]int, []int) {
    data, err := os.ReadFile(f_name)
    if err != nil {
        panic(err)
    }

    lines := strings.Split(strings.TrimSpace(string(data)), "\n")
    left := []int{}
    right := []int{}

    for _, num_str := range lines {
        fields := strings.Fields(num_str)
        for i, n := range fields {
            val, _ := strconv.Atoi(n)
            if i % 2 == 0 {
                left  = append(left, val)
            } else {
                right = append(right, val)
            }

        }
    }
    return left, right
}

func part_1(left_in, right_in []int) float64 {
    // copy and sort the slices
    left  := make([]int, len(left_in))
    right := make([]int, len(right_in))
    copy(left, left_in)
    copy(right, right_in)
    slices.Sort(left)
    slices.Sort(right)

    // take the differences
    var sum float64 = 0
    for i:=0; i<len(left); i+=1 {
        sum += math.Abs(float64(left[i] - right[i]))
    }

    return sum
}

func part_2(left, right []int) float64 {
    // calculate the occurances of each num
    counts := make(map[int]int)
    for _, v := range right {
        counts[v] += 1
    }

    // get the total sum
    var sum float64 = 0
    for _, v := range left {
        sum += float64(v * counts[v])
    }
    return sum
}

func main() {
    ex_left, ex_right     := get_input("input_example.txt")
    nums_left, nums_right := get_input("input.txt")

    // check part1 example
    example_sol_part1 := part_1(ex_left, ex_right)
    if example_sol_part1 != 11 {
        fmt.Printf("%.0f != 11\n", example_sol_part1)
        panic("example must be equal to 11")
    }
    // get the answer to part1
    sol_part1 := part_1(nums_left, nums_right)
    fmt.Printf("part1: %.0f\n", sol_part1)

    // check part2 example
    example_sol_part2 := part_2(ex_left, ex_right)
    if example_sol_part2 != 31 {
        fmt.Printf("%.0f != 31\n", example_sol_part2)
        panic("example must be equal to 31")
    }
    // get the answer to part2
    sol_part2 := part_2(nums_left, nums_right)
    fmt.Printf("part2: %.0f\n", sol_part2)

}

