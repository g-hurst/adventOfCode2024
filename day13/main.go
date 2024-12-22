package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
    "regexp"
    "github.com/draffensperger/golp"
)


func get_input(f_name string) [][][]int{
    data, err := os.ReadFile(f_name)
    if err != nil {
        panic(err)
    }
    exp := regexp.MustCompile(`X.(\d+), Y.(\d+)`)

    chunks := strings.Split(strings.TrimSpace(string(data)), "\n\n") 
    m := make([][][]int, len(chunks))
    for i, chunk := range chunks {
        lines := strings.Split(strings.TrimSpace(chunk), "\n")
        m[i] = make([][]int, len(lines))
        for j, line := range lines {
            for _, num := range  exp.FindStringSubmatch(line)[1:] {
                val, _ := strconv.Atoi(num)
                m[i][j] = append(m[i][j], val)
            }
        }
    }
    return m
}

func part_1(data [][][]int) int {
    tokens := 0
    for _, equasions := range data {
        // run a linear program to solve eq
        lp := golp.NewLP(0, 2)
        lp.AddConstraint([]float64{float64(equasions[0][0]), float64(equasions[1][0])}, golp.EQ, float64(equasions[2][0]))
        lp.AddConstraint([]float64{float64(equasions[0][1]), float64(equasions[1][1])}, golp.EQ, float64(equasions[2][1]))
        lp.SetObjFn([]float64{3,1})
        lp.SetInt(0, true)
        lp.SetInt(1, true)
        if sol_type:=lp.Solve(); sol_type==golp.OPTIMAL {
            vars := lp.Variables()
            tokens += 3 * int(vars[0]) + int(vars[1])
        }
    }
    return tokens
}

func part_2(data [][][]int) int{
    // the float64 required to run the lp cannot handle the solution, but it
    // turns out every eq can be solved with creamer's rule... how convienient
    tokens := 0
    const diff = 10000000000000
    for _, equasions := range data {
        x1 := equasions[0][0]
        y1 := equasions[0][1]
        x2 := equasions[1][0]
        y2 := equasions[1][1] 
        z1 := equasions[2][0]  + diff
        z2 := equasions[2][1]  + diff
        if det := (x1*y2 - y1*x2); det != 0 {
            a := (z1*y2 - z2*x2) / det 
            b := (z2*x1 - z1*y1) / det
            if (a*x1+b*x2)==z1 && (a*y1+b*y2)==z2 {
                tokens += 3*a + b
            }
        }
    }
    return int(tokens)
}


func main() {
    data    := get_input("input.txt")
    ex_data := get_input("input_example.txt")

    // check part1 example
    example_sol_part1 := part_1(ex_data)
    example_ans_part1 := 480
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

