package mainmenu

import (
    "fmt"
    tea "github.com/charmbracelet/bubbletea"
)

type menuOptions int

const (
    newSession menuOptions = iota
    loadOngoingSession
    optionsMenu
    quit
)

func (o menuOptions) String() string {
    switch o {
    case newSession:
        return "New Session"
    case loadOngoingSession:
        return "Load Ongoing Session"
    case optionsMenu:
        return "Options"
    case quit:
        return "Quit"
    }
    return "ERROR"
}


type SelectSessionMsg struct {}
type ReloadSessionMsg struct {}
type OptionMsg struct {}

type Model struct {
    activeMenuOptions []menuOptions
    cursor int
}

func InitialModel() Model {
    return Model{[]menuOptions{newSession, optionsMenu, quit,}, 0}
}

func (m Model) Init() tea.Cmd {
    return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "q", "h":
            return m, tea.Quit
        case "up", "k":
            if m.cursor > 0 {
                m.cursor--
            }
        case "down", "j":
            if m.cursor < len(m.activeMenuOptions)-1 {
                m.cursor++
            }
        case "enter", " ", "l":
            switch m.activeMenuOptions[m.cursor] {
            case newSession:
                return m, func() tea.Msg { return SelectSessionMsg{} }
            case loadOngoingSession:
                return m, func() tea.Msg { return ReloadSessionMsg{} }
            case optionsMenu:
                return m, func() tea.Msg { return OptionMsg{} }
            case quit:
                return m, tea.Quit
            }
        }
    }
    return m, nil
}

func (m Model) View() string {
    s := "Pathfinder 2e Sessioner v.0.go.1\n\n"
    for i, option := range m.activeMenuOptions {
        cursor := " "
        if m.cursor == i {
            cursor = ">"
        }

        s += fmt.Sprintf("%s %s\n", cursor, option)
    }
    return s
}
