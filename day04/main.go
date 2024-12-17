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


func count_str(haystack, needle string) int {
    found := 0
    for n:=0; n<len(haystack); n+=1 {
        if (haystack[n]==needle[0])  || (haystack[n]==needle[len(needle)-1]) {
            is_foreward := (haystack[n:][0] == needle[0])
            needle_ptr := 1
            for i:=1; i<len(haystack[n:]); i+=1 {
                if (i == (len(needle)-1)) && (
                   ( is_foreward && (haystack[n:][i]==needle[len(needle)-1])) ||
                   (!is_foreward && (haystack[n:][i]==needle[0]))) {
                    found += 1
                    break
                } else if (!is_foreward && (haystack[n:][i] != needle[len(needle)-needle_ptr-1])) ||
                          (is_foreward  && (haystack[n:][i] != needle[needle_ptr])) {
                    break
                } else {
                    needle_ptr += 1
                }
            }
        }
    }

    return found
}

func part_1(data string) int {
    sum := 0

    rows := strings.Split(data, "\n")
    columns   := make([][]rune, len(rows[0]))
    diags_neg := make([][]rune, len(rows)+len(columns)-1) // negitive slope diagonals
    diags_pos := make([][]rune, len(diags_neg))           // positive slope diagonals

    // check the rows and make the comumns
    for i, row := range rows {
        sum += count_str(row, "XMAS")

        for j, char := range row {
            columns[j] = append(columns[j], char)
            diags_neg[i+j] = append(diags_neg[i+j], char)
            diags_pos[(i-j)+(len(columns)-1)] = append(diags_pos[(i-j)+(len(columns)-1)], char)
        }

    }

    check := func(strs [][]rune) {
        for _, str := range strs{
            sum += count_str(string(str), "XMAS")
        }

    }

    // check the columns
    check(columns)
    check(diags_neg)
    check(diags_pos)

    return sum
}

func part_2(data string) int {
    return 0
}


func main() {
    data    := get_input("input.txt")
    ex_data := get_input("input_example.txt")

    // check part1 example
    example_sol_part1 := part_1(ex_data)
    example_ans_part1 := 18
    if example_sol_part1 != example_ans_part1 {
        fmt.Printf("%d != %d\n", example_sol_part1, example_ans_part1)
        panic("example solution and answer must be equal")
    }
    // get the answer to part1
    sol_part1 := part_1(data)
    fmt.Printf("part1: %d\n", sol_part1)

    // check part2 example
    example_sol_part2 := part_2(ex_data)
    example_ans_part2 := 9
    if example_sol_part2 !=example_ans_part2 {
        fmt.Printf("%d != %d\n", example_sol_part2, example_ans_part2)
        panic("example solution and answer must be equal")
    }
    // get the answer to part2
    sol_part2 := part_2(data)
    fmt.Printf("part2: %d\n", sol_part2)

}

