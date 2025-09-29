package main

import (
    "fmt"
)

func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return []string{}
    }

    ref := map[string]string{
        "2": "abc",
        "3": "def",
        "4": "ghi",
        "5": "jkl",
        "6": "mno",
        "7": "pqrs",
        "8": "tuv",
        "9": "wxyz",
    }
	ans := []string{""}
	for _, d := range digits {
		temp := []string{}
		for _, r := range ref[string(d)] {
			for _, a := range ans {
				temp = append(temp, a+string(r))
			}
		}
		ans = temp
	}

    return ans
}

func main() {
    fmt.Println(letterCombinations("23"))
}