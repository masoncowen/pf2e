package main

import (
    "fmt"
    "os"
    tea "github.com/charmbracelet/bubbletea"
    "internal/engines/mainmenu"
    "internal/engines/replacemetimer"
    "internal/engines/options"
)

type sessionState int

const (
    mainMenuView sessionState = iota
    optionsView
    timerView
    combatView
)

type model struct {
    state sessionState
    mainmenu tea.Model
    options tea.Model
    timer tea.Model
}

func initialModel() model {
	return model{
        state: mainMenuView,
        mainmenu: mainmenu.InitialModel(),
        timer: replacemetimer.InitialModel(),
        options: options.InitialModel(),
	}
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd
    switch msg := msg.(type) {
    case mainmenu.StartMsg:
        m.state = timerView
    case mainmenu.OptionMsg:
        m.state = optionsView   
    case options.BackMsg:
        m.state = mainMenuView
    case tea.KeyMsg:
        switch msg.String() {
        case "q":
            return m, tea.Quit
        }
    }
        
    switch m.state {
    case mainMenuView:
        newMainMenu, newCmd := m.mainmenu.Update(msg)
        mainMenuModel, ok := newMainMenu.(mainmenu.Model)
        if !ok {
            panic("Could not perform assertion on mainmenu model")
        }
        m.mainmenu = mainMenuModel
        cmd = newCmd
    case optionsView:
        newOptions, newCmd := m.options.Update(msg)
        optionsModel, ok := newOptions.(options.Model)
        if !ok {
            panic("Could not perform assertion on options model")
        }
        m.options = optionsModel
        cmd = newCmd
    case timerView:
        newTimer, newCmd := m.timer.Update(msg)
        timerModel, ok := newTimer.(replacemetimer.Model)
        if !ok {
            panic("Could not perform assertion on timer model")
        }
        m.timer = timerModel
        cmd = newCmd
    }
    return m, cmd
}

func (m model) View() string {
    switch m.state {
    case mainMenuView:
        return m.mainmenu.View()
    case optionsView:
        return m.options.View()
    case timerView:
        return m.timer.View()
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
