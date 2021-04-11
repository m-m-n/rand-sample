package main

import (
    "fmt"
    "time"
    "math/rand"
    "errors"
    "flag"
    "strconv"
)

func main() {
    flag.Parse()
    args := flag.Args()

    s := "18"
    if len(args) > 0 {
        s = args[1]
    }
    n, _ := strconv.ParseUint(s, 10, 8)

    if n < 5 {
        fmt.Println("5以上18以下の数値を与えてください")
        return
    }

    rand.Seed(time.Now().UnixNano())
    result := make([]int, n, n)

    var ranking []int
    for {
        result = makeRand(int(n), result)
        var err error
        ranking, err = makeRanking(result)
        if err == nil {
            break
        }
    }

    for key, value := range ranking[0:5] {
        fmt.Println(key + 1, value + 1)
    }
}

func makeRand(n int, slice []int) []int {
    a := rand.Intn(n)
    slice[a]++
    return slice
}

func makeRanking(slice []int) ([]int, error) {
    ranking := make([]int, 6, 6)
    skipIndex := make([]bool, 18, 18)
    for i := 0; i < 6; i++ {
        max := -1
        skip := -1
        for key, value := range slice {
            if skipIndex[key] {
                continue
            }
            if max < value {
                max = value
                skip = key
            }
        }
        if (skip < 0) {
            return nil, errors.New("予期せぬエラー")
        }
        skipIndex[skip] = true
        ranking[i] = skip
        if i > 0 && slice[skip] == slice[ranking[i - 1]] {
            return nil, errors.New("同率順位があります")
        }
    }

    return ranking, nil
}
