package entryui

import (
    "fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type menuOptions int

const (
    newSession menuOptions = iota
    loadPreviousSession
    options
    quit
)

func (o menuOptions) String() string {
    switch o {
    case newSession:
        return "New Session"
    case loadPreviousSession:
        return "Load Previous Session"
    case options:
        return "Options"
    case quit:
        return "Quit"
    }
    return "ERROR"
}

var activeMenuOptions = []menuOptions{
    newSession,
    options,
    quit,
}

type Model struct {
    cursor int
    printMessage string
}

func InitialModel() Model {
    return Model{}
}

func (m Model) Init() tea.Cmd {
    return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "q":
            return m, tea.Quit
        case "up", "k":
            if m.cursor > 0 {
                m.cursor--
            }
        case "down", "j":
            if m.cursor < len(activeMenuOptions)-1 {
                m.cursor++
            }
        case "enter", " ":
            switch activeMenuOptions[m.cursor] {
            case newSession:
                m.printMessage = "Start"
            case loadPreviousSession:
                m.printMessage = "Prev"
            case options:
                m.printMessage = "Opt"
            case quit:
                return m, tea.Quit
            }
        }
    }
    return m, nil
}

func (m Model) View() string {
    s := "Pathfinder 2e Sessioner v.0.go.1\n\n"
    for i, option := range activeMenuOptions {
        cursor := " "
        if m.cursor == i {
            cursor = ">"
        }

        s += fmt.Sprintf("%s %s\n", cursor, option)
    }
    return s
}
