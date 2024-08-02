package options

import (
    "fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type configOptions int

const (
    quit configOptions = iota  
    optionPlaceholder
)

func (o configOptions) String() string {
    switch o {
    case quit:
        return "Return"
    case optionPlaceholder:
        return "Placeholder"
    }
    return "ERROR"
}

var activeMenuOptions = []configOptions{
    optionPlaceholder,
    quit,
}
    
type BackMsg struct {}

type Options struct {
    optionPlaceholder bool
}

type Model struct {
    cursor int
    options Options
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
        case "h", "left":
            return m, func() tea.Msg { return BackMsg{} }
        case "enter", " ", "l", "right":
            switch activeMenuOptions[m.cursor] {
            case optionPlaceholder:
                m.options.optionPlaceholder = !m.options.optionPlaceholder
            case quit:
                return m, func() tea.Msg { return BackMsg{} }
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
        if option == quit {
            s += fmt.Sprintf("%s %s\n", cursor, option)
            continue
        }
        enabled := false
        switch option {
        case optionPlaceholder:
            enabled = m.options.optionPlaceholder
        }
        value := "Disabled"
        if enabled {
            value = "Enabled"
        }
        s += fmt.Sprintf("%s %s %s\n", cursor, option, value)
    }
    return s
}
