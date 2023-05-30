package main

import (
	"math"
)


func has_match(needle []byte, haystack []byte) bool {
	index := 0
	if len(needle) == 0 {
		return true
	}
	if len(haystack) == 0 {
		return false
	}
	for ind_needle, nchr := range needle {
		for haystack[index] != nchr {
			index += 1
			if index >= len(haystack) {
				return false
			}
		}
		if nchr != haystack[index] {
			return false
		}
		index += 1
		if index >= len(haystack) {
			return ind_needle == len(needle) - 1
		}
	}
	return true
}

var MATCH int = 3
var MISMATCH int = -2
var GAP int = -5
var GAPEXTEND int = -1
var MIN int = math.MinInt + 1000

func match_score(x byte, y byte) int {
	if x == y {
		return MATCH
	} else {
		return MISMATCH
	}
}

func gap_cost(length int) int {
	return GAP + GAPEXTEND * (length - 1)
}

func max(vars ...int) int {
	m := vars[0]

	for _, i := range vars {
		if m < i {
			m = i
		}
	}

	return m
}

func min(vars ...int) int {
	m := vars[0]

	for _, i := range vars {
		if m > i {
			m = i
		}
	}

	return m
}

func calc_row_score(u []byte, v []byte, i int, a []int, a_prev []int, b []int, b_prev []int, c []int, c_prev []int) {
	a[0] = MIN;
	b[0] = gap_cost(i);
	c[0] = MIN;
	for j := 1; j <= len(v); j++ {
		a[j] = max(a_prev[j - 1], b_prev[j - 1], c_prev[j - 1]) + match_score(u[i - 1], v[j - 1])
		b[j] = max(a_prev[j] + GAP, b_prev[j] + GAPEXTEND, c_prev[j] + GAP)
		c[j] = max(a[j - 1] + GAP, b[j - 1] + GAP, c[j - 1] + GAPEXTEND)
	}
	// fmt.Printf("%v\t%v\t%v\n", a, b, c) // debug
}

func score_gotoh_str(u string, v string) int {
	return score_gotoh([]byte(u), []byte(v))
}

func score_gotoh(u []byte, v []byte) int {
	u_len := len(u)
	v_len := len(v)
	a_prev := make([]int, v_len + 1)
	a := make([]int, v_len + 1)
	b_prev := make([]int, v_len + 1)
	b := make([]int, v_len + 1)
	c_prev := make([]int, v_len + 1)
	c := make([]int, v_len + 1)
	a_prev[0] = 0
	b_prev[0] = 0
	c_prev[0] = 0
	swapped := false
	for j := 1; j <= v_len; j++ {
		a_prev[j] = MIN 
		b_prev[j] = MIN
		c_prev[j] = gap_cost(j)
	}

	for i := 1; i <= u_len; i++ {
		if swapped {
			calc_row_score(u, v, i, a_prev, a, b_prev, b, c_prev, c)
		} else {
			calc_row_score(u, v, i, a, a_prev, b, b_prev, c, c_prev)
		}
		swapped = !swapped
	}


	if !swapped {
		return max(a_prev[v_len], b_prev[v_len], c_prev[v_len])
	}
	return max(a[v_len], b[v_len], c[v_len])
}

