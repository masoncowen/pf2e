package replacemetimer

import (
	"fmt"
    "time"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type eventType int

const (
    BeginBreak eventType = iota
    EndBreak
    Note
)


type event struct {
  eventTime time.Time
  eventType eventType
  eventText string
}
  

type Model struct {
    entries []event
}

func InitialModel() Model {
    return Model{}
}

func (m Model) Init() tea.Cmd {
    return nil
}

func log_event(e event) error {
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
    title := "Pathfinder 2e Sessioner v.0.go.1"
    var titleStyle = lipgloss.NewStyle().
    Align(lipgloss.Center).
    BorderStyle(lipgloss.NormalBorder()).BorderBottom(true)
    times := ""
    notes := ""
    for _, entry := range m.entries {
      times += fmt.Sprintf("%s\n", entry.eventTime.Format("3:4"))
      notes += fmt.Sprintf("%s\n", entry.eventText)
    }
    return lipgloss.JoinVertical(lipgloss.Center, titleStyle.Render(title), lipgloss.JoinHorizontal(lipgloss.Top, times, notes))
}
