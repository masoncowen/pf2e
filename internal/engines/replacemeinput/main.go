package replacemeinput

import (
    "fmt"
    "time"
	tea "github.com/charmbracelet/bubbletea"
    "internal/engines/replacemetimer"
)

type Model struct {
    Type replacemetimer.EventType
    Text string
}

func InitialModel() Model {
    return Model{}
}

func (m Model) Init() tea.Cmd {
    return nil
}

func (m Model) returnFinalString() tea.Cmd {
    return func() tea.Msg {
        return replacemetimer.Event{Type: m.Type,
            Text: m.Text,
            Time: time.Now(),
        }
    }
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case replacemetimer.Event:
        m.Type = msg.Type
        m.Text = ""
    case tea.KeyMsg:
        switch msg.String() {
        case "enter":
            return m, m.returnFinalString()
        case "backspace":
            m.Text = m.Text[:len(m.Text)-1]
            return m, nil
        }
        m.Text = m.Text + msg.String()
    }
    return m, nil
}

func (m Model) View() string {
    s := "Pathfinder 2e Sessioner v.0.go.1\nType message for note.\n\n"
    s += fmt.Sprintf("> (%s) %s", m.Type, m.Text)
    return s
}
