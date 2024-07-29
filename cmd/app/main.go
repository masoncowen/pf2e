package main

import (
    "fmt"
    "os"
    "path/filepath"
    tea "github.com/charmbracelet/bubbletea"
    "internal/engines/mainmenu"
    "internal/engines/replacemetimer"
    "internal/engines/options"
    "internal/engines/sessionmenu"
)

type sessionState int

const (
    mainMenuView sessionState = iota
    optionsView
    sessionMenuView
    sessionMenuViewExistingSession
    timerView
    combatView
)

type model struct {
    state sessionState
    mainmenu tea.Model
    options tea.Model
    sessionmenu tea.Model
    timer tea.Model
}

func initialModel() model {
    pf2eDir := os.Getenv("PF2E_DIR")
    if pf2eDir == "" {
        homeDir, err := os.UserHomeDir()
        if err != nil {
            panic(err)
        }
        pf2eDir = filepath.Join(homeDir, ".pf2e")
    }
	return model{
        state: mainMenuView,
        mainmenu: mainmenu.InitialModel(pf2eDir),
        options: options.InitialModel(),
        sessionmenu: sessionmenu.InitialModel(pf2eDir),
        timer: replacemetimer.InitialModel(),
	}
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd
    switch msg := msg.(type) {
    case mainmenu.SelectSessionMsg:
        m.state = sessionMenuView
    case mainmenu.ReloadSessionMsg:
        m.state = sessionMenuViewExistingSession
    case mainmenu.OptionMsg:
        m.state = optionsView
    case options.BackMsg, sessionmenu.BackMsg:
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
    case sessionMenuView:
        newSessionSelector, newCmd := m.sessionmenu.Update(msg)
        sessionsModel, ok := newSessionSelector.(sessionmenu.Model)
        sessionsModel.ExistingSession = false
        if !ok {
            panic("Could not perform assertion on sessions model")
        }
        m.sessionmenu = sessionsModel
        cmd = newCmd
    case sessionMenuViewExistingSession:
        newSessionSelector, newCmd := m.sessionmenu.Update(msg)
        sessionsModel, ok := newSessionSelector.(sessionmenu.Model)
        sessionsModel.ExistingSession = true
        if !ok {
            panic("Could not perform assertion on sessions model")
        }
        m.sessionmenu = sessionsModel
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
    case sessionMenuView:
        return m.sessionmenu.View()
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
