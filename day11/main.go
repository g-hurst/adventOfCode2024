package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
    "math"
)

func get_input(f_name string) []int{
    data, err := os.ReadFile(f_name)
    if err != nil {
        panic(err)
    }

    ints := make([]int, 0)
    for _, char := range strings.Split(strings.TrimSpace(string(data)), " ") {
        num, _ := strconv.Atoi(string(char))
        ints = append(ints, num)
    }
    return ints
}

func blink (old_stones map[int]int) map[int]int{
    new_stones := make(map[int]int)
    for key, val := range old_stones {
        digits := int(math.Log10(float64(key))) + 1

        if key == 0 {
            new_stones[1] += val
        } else if digits%2 == 0 {
            split := int(math.Pow(float64(10),float64(digits/2)))
            new_stones[key / split] += val 
            new_stones[key % split] += val

        } else {
            new_stones[key*2024] += val
        }
    }
    return new_stones
}

func part_1(data []int) int {
    // create a hashmap that counts all the stones
    stones := make(map[int]int)
    for _, val := range data {
        stones[val] += 1
    }

    // update the stones for each blink
    num_blinks := 25
    for i:=0; i<num_blinks; i+=1 {
        stones = blink(stones)
    }

    // sum the stones for each key
    total_stones := 0
    for _, num_stones := range stones {
        total_stones += num_stones
    }

    return total_stones
}

func part_2(data []int) int {
    stones := make(map[int]int)
    for _, val := range data {
        stones[val] += 1
    }

    num_blinks := 75
    for i:=0; i<num_blinks; i+=1 {
        stones = blink(stones)
    }

    total_stones := 0
    for _, num_stones := range stones {
        total_stones += num_stones
    }

    return total_stones
}
func main() {
    data    := get_input("input.txt")
    ex_data := get_input("input_example.txt")

    // check part1 example
    example_sol_part1 := part_1(ex_data)
    example_ans_part1 := 55312
    if example_sol_part1 != example_ans_part1 {
        fmt.Printf("%d != %d\n", example_sol_part1, example_ans_part1)
        panic("example solution and answer must be equal")
    }
    // get the answer to part1
    sol_part1 := part_1(data)
    fmt.Printf("part1: %d\n", sol_part1)

    // get the answer to part2
    sol_part2 := part_2(data)
    fmt.Printf("part2: %d\n", sol_part2)

}

