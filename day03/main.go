package main

import (
    "fmt"
    "os"
    "regexp"
    "strings"
    "strconv"
    "math"
)

func get_input(f_name string) string {
    data, err := os.ReadFile(f_name)
    if err != nil {
        panic(err)
    }
    return string(data)
}

func parse_product(str string) int {
    // takes a string of mul(\d+,\d+) and returns the product
    product := 1
    str = str[4:]
    str = str[:len(str)-1]

    for _, num := range strings.Split(str, ",") {
        v, _ := strconv.Atoi(num)
        product *= v
    }
    return product
}

func part_1(data string) int {
    exp := regexp.MustCompile(`mul\((\d+,\d+\))`)
    matches := exp.FindAllString(data, -1)
    sum := 0
    for _, match := range matches{
        sum += parse_product(match)
    }
    return sum
}

func part_2(data string) int {
    // regex find all the do/dont occurances
    exp_mul := regexp.MustCompile(`mul\((\d+,\d+\))`)
    exp_do  := regexp.MustCompile(`do\(\)`)
    exp_ndo := regexp.MustCompile(`don't\(\)`)
    matches_mul := exp_mul.FindAllStringIndex(data, -1)
    matches_do  := exp_do.FindAllStringIndex(data, -1)
    matches_ndo := exp_ndo.FindAllStringIndex(data, -1)

    sum := 0
    for len(matches_mul) > 0 {
        // get the current do/dont status
        var is_do bool
        var idx   int
        if len(matches_do)>0 && len(matches_ndo)>0 {
            idx_do  := matches_do[len(matches_do)-1][0]
            idx_ndo := matches_ndo[len(matches_ndo)-1][0]
            idx     = int(math.Max(float64(idx_do), float64(idx_ndo)))
            is_do   = idx_do > idx_ndo
            if is_do {
                matches_do = matches_do[:len(matches_do)-1]
            } else {
                matches_ndo = matches_ndo[:len(matches_ndo)-1]
            }
        } else if len(matches_ndo) > 0 {
            idx = matches_ndo[len(matches_ndo)-1][0]
            is_do   = false
            matches_ndo = matches_ndo[:len(matches_ndo)-1]
        } else {
            idx = 0
            is_do  = true
        }


        // remove matches if dont, sum if do
        for (len(matches_mul)>0) && (matches_mul[len(matches_mul)-1][0] > idx) {
            if is_do {
                match := matches_mul[len(matches_mul)-1]
                sum += parse_product(data[match[0]:match[1]])
            }
            matches_mul = matches_mul[:len(matches_mul)-1]
        }
    }

    return sum
}



func main() {
    data    := get_input("input.txt")

    // check part1 example
    ex_data := get_input("input_example_1.txt")
    example_sol_part1 := part_1(ex_data)
    example_ans_part1 := 161
    if example_sol_part1 != example_ans_part1 {
        fmt.Printf("%d != %d\n", example_sol_part1, example_ans_part1)
        panic("example solution and answer must be equal")
    }
    // get the answer to part1
    sol_part1 := part_1(data)
    fmt.Printf("part1: %d\n", sol_part1)

    // check part2 example
    ex_data = get_input("input_example_2.txt")
    example_sol_part2 := part_2(ex_data)
    example_ans_part2 := 48
    if example_sol_part2 !=example_ans_part2 {
        fmt.Printf("%d != %d\n", example_sol_part2, example_ans_part2)
        panic("example solution and answer must be equal")
    }
    // get the answer to part2
    sol_part2 := part_2(data)
    fmt.Printf("part2: %d\n", sol_part2)

}

