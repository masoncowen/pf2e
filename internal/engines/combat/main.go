package combat

import (
    "fmt"
    "time"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
    combatant []Combatant
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
        newEvent := event{}
        switch msg.String() {
        case "ctrl+c", "q":
            return m, tea.Quit
        case "enter", " ":
            newEvent = event{time.Now(), Note, "Blank Note: Need to add text"}
        case "b":
            newEvent = event{time.Now(), BeginBreak, "Begin Break: maybe add context for break"}
        case "r":
            newEvent = event{time.Now(), EndBreak, "End Break: Not sure what text could be used for here?"}
        }
        m.entries = append(m.entries, newEvent)
        log_event(newEvent)
    }
    return m, nil
}

func (m Model) View() string {
    s := "Pathfinder 2e Sessioner v.0.go.1\n\n"
    for _, entry := range m.entries {
      s += fmt.Sprintf("%s\n", time.Now().String())
      s += fmt.Sprintf("%s %s\n", entry.eventTime.Format("3:4"), entry.eventText)
    }
    return lipgloss.JoinHorizontal(lipgloss.Bottom, s, "This is a test to the right\n")
}

