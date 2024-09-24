package main

import "fmt"

func Ft_missing(nums []int) int {
	max := len(nums)
	var val int = 67
	for i := 0; i < max; i++ {
		for range nums {
			if i != nums[i] {
				val = i
			}
		}
	}
	if val == 67 {
		val = max
	}
	return val
}

func Ft_profit(prices []int) int {
	size := len(prices)
	var val1 int = size + 1
	var val2 int
	var val3 int
	for i := 0; i < size; i++ {
		for range prices {
			if prices[i] <= val1 {
				val1 = prices[i]
				val3 = i
			}
		}
	}
	for j := val3; j < size; j++ {
		for range prices {
			if prices[j] >= val2 {
				val2 = prices[j]
			}
		}
	}
	return val2 - val1
}

func Ft_max_substring(s string) int {
	var tab []string
	var tempvar string
	check := false
	for i := 0; i < len(s); i++ {
		tempvar = string(s[i])
		for range s {
			for j := 0; j < len(tab); j++ {
				for range tab {
					if tempvar == tab[j] {
						check = true
					}
				}
			}
			if check == false {
				tab = append(tab, tempvar)
			}

		}
		check = false
	}
	return len(tab)
}

func Ft_coin(coins []int, amount int) int {
	var tab []int
	for j := 0; j < amount+1; j++ {
		tab = append(tab, j)
	}
	for i := 1; i <= amount; i++ {
		tab[i] = amount + 1
	}
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i >= coin {
				tab[i] = min(tab[i], tab[i-coin]+1)
			}
		}
	}
	if tab[amount] == amount+1 {
		return -1
	}
	return tab[amount]
}

func Ft_non_overlap(intervals [][]int) int {
	compt := 0
	tab := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < tab {
			compt++
		} else {
			tab = intervals[i][1]
		}
	}

	return compt
}

func Ft_min_window(s string, t string) string {
	n := len(s)
	minLen := n + 1
	result := ""

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			sub := s[i : j+1]
			if check(sub, t) {
				if len(sub) < minLen {
					minLen = len(sub)
					result = sub
				}
			}
		}
	}

	return result
}

func check(s string, t string) bool {
	for i := 0; i < len(t); i++ {
		found := false
		for j := 0; j < len(s); j++ {
			if t[i] == s[j] {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Ft_missing([]int{3, 1, 2}))                   // resultat : 0
	fmt.Println(Ft_missing([]int{0, 1}))                      // resultat : 2
	fmt.Println(Ft_missing([]int{9, 6, 4, 2, 3, 5, 7, 0, 1})) // resultat : 8
	fmt.Print("\n")

	fmt.Println(Ft_profit([]int{7, 1, 5, 3, 6, 4})) // resultat : 5
	// si on achète au jour 1, nous payons 1,
	// et si nous le vendons au 4eme jour, nous gagnons 6, le bénéfice est 6-1
	fmt.Println(Ft_profit([]int{7, 6, 4, 3, 1})) // resultat : 0
	fmt.Print("\n")

	fmt.Println(Ft_max_substring("abcabcbb")) // resultat : 3
	// "abc" est la plus grande sous chaine composé de caractère diffèrent
	fmt.Println(Ft_max_substring("bbbbb")) // resultat : 1
	// "b" est la plus grande sous chaine
	fmt.Print("\n")

	fmt.Println(Ft_coin([]int{1, 2, 5}, 11)) // resultat : 3 car (11 == 5 + 5 + 1)
	fmt.Println(Ft_coin([]int{2}, 3))        // resultat : -1
	fmt.Println(Ft_coin([]int{1}, 0))        // resultat : 0
	fmt.Print("\n")

	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}})) // resultat : 1
	// pour que les intervalles soient tous des intervalles qui ne se superpose pas,
	// il suffit de d'enlever qu'un seul intervalle, [1,3]
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {2, 3}}))         // resultat : 0
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {1, 2}, {1, 2}})) // resultat : 2
	fmt.Print("\n")

	fmt.Println(Ft_min_window("ADOBECODEBANC", "ABC")) // resultat : "BANC"
	fmt.Println(Ft_min_window("a", "aa"))              // resultat : ""

}
