package sessionmenu

import (
	"fmt"
    "os"
    "path/filepath"
	tea "github.com/charmbracelet/bubbletea"
)

type BackMsg struct {}

type session struct {
    filepath string
    name string
}

type Model struct {
    sessions []session
    cursor int
}

func InitialModel() Model {
    pf2eDir := os.Getenv("PF2E_DIR")
    if pf2eDir == "" {
        homeDir, err := os.UserHomeDir()
        if err != nil {
            panic(err)
        }
        pf2eDir = filepath.Join(homeDir, ".pf2e")
    }
    sessionsDir := filepath.Join(pf2eDir, "sessions")
    sessions, err := os.ReadDir(sessionsDir)
    if os.IsNotExist(err) {
        return Model{}
    }
    sessionList := []session{}
    for _, currSession := range sessions {
        sessionList = append(sessionList, session{currSession.Name(), currSession.Name()})
    }
    return Model{sessions: sessionList}
}

func (m Model) Init() tea.Cmd {
    return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "q", "h", "left":
            return m, func() tea.Msg { return BackMsg{} }
        case "up", "k":
            if m.cursor > 0 {
                m.cursor--
            }
        case "down", "j":
            if m.cursor < len(m.sessions)-1 {
                m.cursor++
            }
        case "enter", " ", "l", "right":
            return m, tea.Quit
        }
    }
    return m, nil
}

func (m Model) View() string {
    s := "Pathfinder 2e Sessioner v.0.go.1\n\n"
    for i, session := range m.sessions {
        cursor := " "
        if m.cursor == i {
            cursor = ">"
        }

        s += fmt.Sprintf("%s %s\n", cursor, session)
    }
    return s
}
