package sessionmenu

import (
	"fmt"
    "os"
    "path/filepath"
	tea "github.com/charmbracelet/bubbletea"
)

type BackMsg struct {}
type NewSessionMsg struct {
    CampaignPath string
}
type ReloadSessinMsg struct {
    SessionPath string
}

type Model struct {
    ExistingSession bool
    sessions []string
    campaigns []string
    cursor int
}

func InitialModel(pf2eDir string) Model {
    campaignsDir := filepath.Join(pf2eDir, "campaigns")
    campaigns, errCampaign := os.ReadDir(campaignsDir)
    if os.IsNotExist(errCampaign) {
        os.MkdirAll(campaignsDir, os.ModePerm)
        err := os.WriteFile(filepath.Join(campaignsDir, "foo.json"), []byte{}, 0600)
        if err != nil {
            panic(err)
        }
        campaigns, _ = os.ReadDir(campaignsDir)
    }
    campaignList := []string{}
    for _, currCampaign := range campaigns {
        campaignList = append(campaignList, currCampaign.Name())
    }
    sessionsDir := filepath.Join(pf2eDir, "sessions")
    sessions, errSession := os.ReadDir(sessionsDir)
    if os.IsNotExist(errSession) {
        os.MkdirAll(sessionsDir, os.ModePerm)
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
            if m.ExistingSession {
                return m, func() tea.Msg { return ReloadSessinMsg{m.sessions[m.cursor]}}
            }
            return m, func() tea.Msg { return NewSessionMsg{m.campaigns[m.cursor]}}
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
