package main

import (
    "fmt"
    "os"
    "strconv"
)

func get_input(f_name string) []int{
    data, err := os.ReadFile(f_name)
    if err != nil {
        panic(err)
    }

    ints := make([]int, 0)
    for _, char := range data {
        num, _ := strconv.Atoi(string(char))
        ints = append(ints, num)
    }
    return ints
}

type Block struct {
   val int
   sz  int
   is_free bool
}

func part_1(data []int) int {
    var blocks []Block
    // make blocks from list of ints
    for i, val := range data {
        switch i % 2 {
            case 0: blocks = append(blocks, Block{val:i/2,  sz:val, is_free:false})
            case 1: blocks = append(blocks, Block{val:-1, sz:val, is_free:true})
        }
    }

    // densify with a left, right pointer approach
    var blocks_dense []Block
    left := 0
    right := len(blocks) - 1
    for left <= right {
        if blocks[left].is_free && !blocks[right].is_free {
            if blocks[left].sz < blocks[right].sz {
                blocks_dense = append(blocks_dense,
                                     Block{val:blocks[right].val, sz:blocks[left].sz, is_free:false})
                blocks[left].is_free = false
                blocks[right].sz -= blocks[left].sz 
            } else if blocks[left].sz == blocks[right].sz  {
                blocks_dense = append(blocks_dense,
                                     Block{val:blocks[right].val, sz:blocks[left].sz, is_free:false})
                blocks[left].is_free = false
                blocks[right].is_free = true
            } else {
                blocks_dense = append(blocks_dense,
                                     Block{val:blocks[right].val, sz:blocks[right].sz, is_free:false})
                blocks[left].sz -= blocks[right].sz
                blocks[right].is_free = true
            }
        }
        if !blocks[left].is_free {
            if blocks[left].val != -1 {
                blocks_dense = append(blocks_dense, blocks[left])
            }
            left += 1
        }
        if blocks[right].is_free {
            right -= 1
        }
    }

    // get the checksum
    idx := 0
    checksum := 0
    for _, b := range blocks_dense {
        for b_idx:=0; b_idx<b.sz; b_idx+=1 {
            checksum += idx * b.val
            idx += 1
        }
    }

    return checksum
}

func part_2(data []int) int {
    var blocks []Block
    // make blocks from list of ints
    for i, val := range data {
        switch i % 2 {
            case 0: blocks = append(blocks, Block{val:i/2, sz:val, is_free:false})
            case 1: blocks = append(blocks, Block{val:-1,  sz:val, is_free:true})
        }
    }

    var blocks_dense []Block
    for len(blocks)>0 {
        if !blocks[0].is_free {
            blocks_dense = append(blocks_dense, blocks[0])
        } else {
            find_fillers:
            found := false
            for i:=len(blocks)-1; i>0; i-=1 {
                if !blocks[i].is_free && (blocks[i].sz <= blocks[0].sz) {
                    found = true
                    // add the found smaller block to free space
                    blocks_dense = append(blocks_dense,
                                          Block{val:blocks[i].val,
                                                sz:blocks[i].sz,
                                                is_free:false})
                    // mark the old location as free
                    blocks[i].is_free = true
                    blocks[i].val     = -1
                    // update remaining free space
                    if blocks[i].sz < blocks[0].sz {
                        blocks[0].sz -= blocks[i].sz
                        goto find_fillers
                    } 
                    break
                }
            }
            if !found {
                // free space cannot be filled
                blocks_dense = append(blocks_dense, blocks[0])
            }
        }
        blocks = blocks[1:]
    }

    // get the checksum
    idx := 0
    checksum := 0
    for _, b := range blocks_dense {
        for b_idx:=0; b_idx<b.sz; b_idx+=1 {
            if !b.is_free {
                checksum += idx * b.val
            }
            idx += 1
        }
    }

    return checksum
}
func main() {
    data    := get_input("input.txt")
    ex_data := get_input("input_example.txt")

    // check part1 example
    example_sol_part1 := part_1(ex_data)
    example_ans_part1 := 1928
    if example_sol_part1 != example_ans_part1 {
        fmt.Printf("%d != %d\n", example_sol_part1, example_ans_part1)
        panic("example solution and answer must be equal")
    }
    // get the answer to part1
    sol_part1 := part_1(data)
    fmt.Printf("part1: %d\n", sol_part1)

    // check part2 example
    example_sol_part2 := part_2(ex_data)
    example_ans_part2 := 2858
    if example_sol_part2 !=example_ans_part2 {
        fmt.Printf("%d != %d\n", example_sol_part2, example_ans_part2)
        panic("example solution and answer must be equal")
    }
    // get the answer to part2
    sol_part2 := part_2(data)
    fmt.Printf("part2: %d\n", sol_part2)

}

