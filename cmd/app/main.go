package main

import (
    "fmt"
    "os"
    tea "github.com/charmbracelet/bubbletea"
    "internal/engines/entry"
)

type sessionState int

const (
    setup sessionState = iota
    entryView
    combatView
)

type model struct {
    state sessionState
    entry tea.Model
}

func initialModel() model {
	return model{
        state: entryView,
        entry: entryui.InitialModel(),
	}
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "q":
            return m, tea.Quit
        }
    }
        
    switch m.state {
    case entryView:
        newEntry, newCmd := m.entry.Update(msg)
        entryModel, ok := newEntry.(entryui.Model)
        if !ok {
            panic("Could not perform assertion on entryui model")
        }
        m.entry = entryModel
        cmd = newCmd
    }
    return m, cmd
}

func (m model) View() string {
    switch m.state {
    case entryView:
        return m.entry.View()
    // case combatView:
    //     return m.combat.View()
    }
    return "EMPT"
}

func main() {
    m := initialModel()
    p := tea.NewProgram(m)
    if _, err := p.Run(); err != nil {
        fmt.Printf("You done fucked up A-A-Ron: %v", err)
        os.Exit(1)
    }
}
