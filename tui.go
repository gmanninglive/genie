package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	baseStyle  = lipgloss.NewStyle().Margin(1, 2).BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("62")).Padding(4, 4)
	focusStyle = lipgloss.NewStyle().Margin(1, 2).BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("80")).Padding(4, 4)

	filePreviewSyle = lipgloss.NewStyle().Padding(4).Background(lipgloss.Color("#101010"))
)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	list     list.Model
	preview  viewport.Model
	selected int
	focusIdx int
	tasks    []Task
	Width    int
	Height   int
}

func (m model) Init() tea.Cmd {
	return nil
}

func readFile(p string) (string, error) {
	path := filepath.Join(GENIE.BASE, p)
	f, err := os.ReadFile(path)
	return string(f), err
}

func (m model) renderPreviewContent(c string) string {
	pX, pY := baseStyle.GetFrameSize()
	applyStyling := filePreviewSyle.Width(m.preview.Width - pX).Render(c)
	return lipgloss.Place(m.Width/3*2-pX, m.Height-pY, lipgloss.Left, lipgloss.Top, applyStyling)
}

func (m model) renderListContent(c string) string {
	pX, pY := baseStyle.GetFrameSize()
	return lipgloss.Place(m.Width/3-pX, m.Height-pY, lipgloss.Left, lipgloss.Top, c)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		case "tab", "enter":
			if m.focusIdx == 0 {
				m.focusIdx = 1
			} else {
				m.focusIdx = 0
			}
			return m, cmd

		default:
			if m.focusIdx == 0 {
				batch := make([]tea.Cmd, 2)
				m.list, batch[0] = m.list.Update(msg)
				m.selected = m.list.Index()
				f, err := readFile(m.tasks[m.selected].Schedule[0].Template)

				if err != nil {
					m.preview.SetContent("Error Loading file preview...")
					m.preview, batch[1] = m.preview.Update(msg)
				} else {
					m.preview.SetContent(m.renderPreviewContent(f))
					m.preview, batch[1] = m.preview.Update(msg)
				}
				return m, tea.Batch(batch...)
			}
		}

	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height

		_, pY := baseStyle.GetFrameSize()
		m.list.SetSize(msg.Width/3, msg.Height-pY)
		m.preview.Height = msg.Height - pY
		m.preview.Width = m.Width - m.list.Width()
	}

	return m, cmd
}

func (m model) View() string {
	doc := strings.Builder{}

	var split string
	switch m.focusIdx {
	case 0:
		split = lipgloss.JoinHorizontal(lipgloss.Top,
			focusStyle.MarginRight(0).Render(m.renderListContent(m.list.View())),
			baseStyle.Render(m.preview.View()))
	case 1:
		split = lipgloss.JoinHorizontal(lipgloss.Top,
			baseStyle.MarginRight(0).Render(m.renderListContent(m.list.View())),
			focusStyle.Render(m.preview.View()))
	}

	doc.WriteString(split)
	return doc.String()
}

func initialModel(tasks []Task) model {
	items := []list.Item{}

	for _, t := range tasks {
		items = append(items, item{title: t.Title, desc: t.Description})
	}

	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0), selected: 0, tasks: tasks, focusIdx: 0}
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
