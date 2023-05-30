package main

import (
	"strings"
)

type runenode struct {
	value rune
	nodes []*runenode
	replace string
}

func (r *runenode) add(from string, to string) {
	current := r
	out:
	for _, run := range from {
		// fmt.Printf("add: %v\n", current)
		for _, i := range (*current).nodes {
			if (*i).value == run {
				current = i
				continue out
			}
		}
		current.nodes = append(current.nodes, &runenode { value: run })
		current = current.nodes[0]
	}
	// fmt.Printf("add: %v\n", current)
	(*current).replace = to
}

func (r runenode) match(from string) (string, int, int) {
	current := r
	count := 0
	for ind, run := range from {
		broken := false
		for _, i := range current.nodes {
			if (*i).value == run {
				current = *i
				count += 1
				broken = true
				break
			}
		}
		if !broken {
			return current.replace, count, ind
		}
	}
	return current.replace, count, 0
}

func (r runenode) run(from string) string {
	var sb strings.Builder
	last_index := -1
	for index, run := range from {
		if last_index >= index {
			continue
		}
		to_tmp, num, bytes := r.match(from[index:])

		if bytes == 0 || to_tmp == "" {
			sb.WriteRune(run)
			last_index = index
		} else {
			sb.WriteString(to_tmp)
			last_index += num 
		}
	}
	return sb.String()
}




// func main() {
// 	r := runenode { value: 0 }
// 	r.add("hi", "hello")
// 	r.add("hi2", "hello")
// 	fmt.Printf("%v\n", r);
// 	to, c, _ := r.match("ao")
// 	fmt.Printf("%s %d\n", to, c)
// 	fmt.Printf("%s", r.run("hihallo"))
// }
