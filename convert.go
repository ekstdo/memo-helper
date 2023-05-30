package main

import (
	"os"
	"fmt"
	"strings"
	"bufio"
)

type config_options struct {
	replace runenode
	convert map[rune]rune
}

func load_config() config_options {
	replace := runenode { value: 0 }
	var convert = make(map[rune]rune)
	file, _ := os.Open("./config.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)


	is_replace := true
	fmt.Printf("loading config")
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		} else if line == "convert:" {
			is_replace = false
			continue
		} else if is_replace {
			split_line := strings.Split(line, ":")
			replace.add(split_line[0], split_line[1])
		} else {
			split_line := strings.Split(line, ":")
			var to rune
			for _, run := range split_line[1] {
				to = run 
				break
			}
			for _, run := range split_line[0] {
				convert[run] = to 
			}
		}

	}
	fmt.Printf("done")
	return config_options {replace: replace, convert: convert}
}

func (c config_options) convert_string(str string) string {
	var b strings.Builder
	for _, ch := range str {
		to := c.convert[ch]
		if to != 0 {
			b.WriteRune(to)
		}
	}
	return b.String()
}
 
func (c config_options) apply(str string) string {
	return c.convert_string(c.replace.run(strings.ToLower(str)))
}

// func main(){
// 	config := load_config()
// 	fmt.Printf("%s\n", config.apply("Taschen"))
// }
