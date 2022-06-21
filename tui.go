package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var listStyle = lipgloss.NewStyle().Margin(1, 2).BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("62")).Padding(4, 4)
var previewStyle = lipgloss.NewStyle().Margin(1, 2).BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("62")).Padding(4, 4)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	focusIndex int
	cursorMode textinput.CursorMode
	list       list.Model
	preview    viewport.Model
	selected   int
	tasks      []Task
	inputs     []textinput.Model
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func readFile(p string) (string, error) {

	path := filepath.Join(GENIE.BASE, p)
	f, err := os.ReadFile(path)
	return string(f), err
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "enter":
			var cmd tea.Cmd
			m.selected = m.list.Index()
			f, err := readFile(m.tasks[m.selected].Schedule[0].Template)
			if err != nil {
				m.preview.SetContent("Error Loading file preview...")
				m.preview, cmd = m.preview.Update(msg)
			} else {
				m.preview.SetContent(f)
				m.preview, cmd = m.preview.Update(msg)
			}
			return m, cmd
		}

	case tea.WindowSizeMsg:
		w, h := listStyle.GetFrameSize()
		m.list.SetSize(msg.Width-w, msg.Height-h)
		m.preview.Height = msg.Height - h
	}

	cmd := make([]tea.Cmd, 2)
	m.preview, cmd[0] = m.preview.Update(msg)
	m.list, cmd[1] = m.list.Update(msg)
	return m, tea.Batch(cmd...)
}

func (m model) View() string {
	doc := strings.Builder{}
	split := lipgloss.JoinHorizontal(lipgloss.Top,
		listStyle.Render(m.list.View()),
		previewStyle.Render(m.preview.View()))

	doc.WriteString(split)
	return doc.String()
}

func initialModel(tasks []Task) model {
	items := []list.Item{}
	inputs := make([]textinput.Model, len(tasks))

	for i, t := range tasks {
		items = append(items, item{title: t.Title, desc: "I have ’em all over my house"})
		for _, v := range t.Vars {
			ti := textinput.New()
			ti.Placeholder = v
			inputs[i] = ti
		}

	}
	f, _ := readFile(tasks[0].Schedule[0].Template)
	w, h := previewStyle.GetFrameSize()
	preview := viewport.New(w, h)
	preview.SetContent(f)

	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0), preview: preview, selected: 0, tasks: tasks}
	m.list.Title = "React Examples"

	return m
}
func TuiRun(tasks []Task) Task {

	p := tea.NewProgram(initialModel(tasks), tea.WithAltScreen())

	run, err := p.StartReturningModel()
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	// Assert the final tea.Model to our local model and print the choice.
	m, ok := run.(model)
	if !ok {
		panic("Error selecting task! 🔥")
	}
	return tasks[m.selected]
}

//func setVars(t Task) Task {
//t.Vars = make(TplVars, len(t.Params))

//for _, label := range t.Params {
//t.Vars[label] = prompt(label)
//}
//return t
//}
