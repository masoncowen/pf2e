package sessionmenu

import (
	"fmt"
    "os"
    "path/filepath"
	tea "github.com/charmbracelet/bubbletea"
)

type BackMsg struct {}

type Model struct {
    ExistingSession bool
    sessions []string
    campaigns []string
    cursor int
}

func InitialModel(pf2eDir string) Model {
    campainsDir := filepath.Join(pf2eDir, "campaigns")
    campaigns, errCampaign := os.ReadDir(campainsDir)
    if os.IsNotExist(errCampaign) {
        return Model{}
    }
    campaignList := []string{}
    for _, currCampaign := range campaigns {
        campaignList = append(campaignList, currCampaign.Name())
    }
    sessionsDir := filepath.Join(pf2eDir, "sessions")
    sessions, errSession := os.ReadDir(sessionsDir)
    if os.IsNotExist(errSession) {
        return Model{campaigns: campaignList}
    }
    sessionList := []string{}
    for _, currSession := range sessions {
        sessionList = append(sessionList, currSession.Name())
    }
    return Model{campaigns: campaignList, sessions: sessionList}
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
    if m.ExistingSession {
        for i, session := range m.sessions {
            cursor := " "
            if m.cursor == i {
                cursor = ">"
            }

            s += fmt.Sprintf("%s %s\n", cursor, session)
        }
        return s
    }
    for i, campaign := range m.campaigns {
        cursor := " "
        if m.cursor == i {
            cursor = ">"
        }

        s += fmt.Sprintf("%s %s\n", cursor, campaign)
    }
    return s
}
