package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"sort"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/textinput"
)


type model struct {
	wordlist []wordconv
	textInput textinput.Model
	fittingWords []string
	selectedWord int
	err error
}


type wordconv struct{ name string; code string }
type wordscored struct{ name string; code string; score int }

func filter(list []wordconv, sth string) []wordconv {
	result := make([]wordconv, 0)
	if len(sth) > 3 {
		for _, i := range list {
			if has_match([]byte(sth), []byte(i.code)) {
				result = append(result, i)
			} 
		}
	} else {
		for _, i := range list {
			if sth == i.code {
				result = append(result, i)
			} 
		}
	}
	return result
}

func score_each(list []wordconv, sth string) []wordscored {
	result := make([]wordscored, 0)
	for _, i := range list {
		result = append(result, wordscored { code: i.code, name: i.name, score: score_gotoh([]byte(sth), []byte(i.code)) })
	}
	return result
}

func display_wordconv(list []wordconv) string {
	var b strings.Builder
	for _, i := range list {
		b.WriteString(fmt.Sprintf("%s \t\t %s\n", i.name, i.code))
	}
	return b.String()
}

func display_wordscores(list []wordscored) string {
	var b strings.Builder
	for _, i := range list {
		b.WriteString(fmt.Sprintf("%s \t\t %s \t %d\n", i.name, i.code, i.score))
	}
	return b.String()
}

func initialModel () model {

	config := load_config()

	file, err := os.Open("./german.txt")
	defer file.Close()
	var wordlist []wordconv

	scanner := bufio.NewScanner(file)

	limit := 2000000
	counter := 0
	fmt.Printf("loading in wordlist\n")
	for scanner.Scan() {
		line := scanner.Text()
		wordlist = append(wordlist, wordconv { name: line, code: config.apply(line),})
		counter += 1
		if counter > limit {
			break
		}
	}
	fmt.Printf("done\n")



	ti := textinput.New()
	ti.Placeholder = "314159265358973"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return model {
		selectedWord: 0,
		err: err,
		textInput: ti,
		wordlist: wordlist,
	}
}

func (m model) Init() tea.Cmd {
	return nil;
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmdText tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		}
	case error:
		m.err = msg
		return m, nil
	}
	m.textInput, cmdText = m.textInput.Update(msg)
	// return m, tea.Batch(cmdList, cmdText)
	return m, cmdText
}

func (m model) View() string {
	start := "Somebody once told me some number to memorize:\nThis tool tries to find words for the phonetic codes, you're about to type in"
	text := m.textInput.Value()
	filtered := filter(m.wordlist, text)
	if len(text) > 3 { // scoring the result
		scores := score_each(filtered, text)
		sort.Slice(scores, func(i, j int) bool { return scores[i].score > scores[j].score })

		return fmt.Sprintf("%s: \n\n%s\n\n%v\n\n%d\n\n", start, m.textInput.View(), display_wordscores(scores[:min(len(scores), 35)]), len(scores))
	} else {
		return fmt.Sprintf("%s: \n\n%s\n\n%v\n\n%d\n\n", start, m.textInput.View(), display_wordconv(filtered[:min(len(filtered), 35)]), len(filtered))
	}
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
