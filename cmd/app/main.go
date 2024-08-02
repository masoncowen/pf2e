package main

import (
    "fmt"
    "os"
    "time"
    "path/filepath"
    tea "github.com/charmbracelet/bubbletea"
    "internal/engines/mainmenu"
    "internal/engines/options"
    "internal/engines/sessionmenu"
    "internal/engines/replacemetimer"
    "internal/engines/combat"
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
    logFile string
    mainmenu tea.Model
    options tea.Model
    sessionmenu tea.Model
    timer tea.Model
    combat tea.Model
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
    logDir := filepath.Join(pf2eDir, "logs")
    logFileName := time.Now().Format("06-01-02.log")
    logFile := filepath.Join(logDir, logFileName)
	return model{
        state: mainMenuView,
        logFile: logFile,
        mainmenu: mainmenu.InitialModel(pf2eDir),
        options: options.InitialModel(),
        sessionmenu: sessionmenu.InitialModel(pf2eDir),
        timer: replacemetimer.InitialModel(logFile),
        combat: combat.InitialModel(),
	}
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) log_event(e replacemetimer.Event) error {
    eventString := fmt.Sprintf("[%s](%s) %s\n", e.Time.Format("03:04:05"), e.Type, e.Text)
    f, err := os.OpenFile(m.logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer f.Close()
    if _, err = f.WriteString(eventString); err != nil {
        return err
    }
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
    case sessionmenu.NewSessionMsg:
        m.state = timerView
    case sessionmenu.ReloadSessinMsg:
        m.state = timerView
    case replacemetimer.Event:
        m.log_event(msg)
        switch msg.Type {
        case replacemetimer.Quit:
            return m, tea.Quit
        case replacemetimer.BeginCombat:
            m.state = combatView
        }

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
    case combatView:
        newCombat, newCmd := m.combat.Update(msg)
        combatModel, ok := newCombat.(combat.Model)
        if !ok {
            panic("Could not perform assertion on combat model")
        }
        m.combat = combatModel
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
    case combatView:
        return m.combat.View()
    }
    return "Invalid Model has been selected"
}

func main() {
    m := initialModel()
    p := tea.NewProgram(m)
    if _, err := p.Run(); err != nil {
        fmt.Printf("You done fucked up A-A-Ron: %v", err)
        os.Exit(1)
    }
}
