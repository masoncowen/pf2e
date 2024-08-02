package replacemetimer

import (
	"fmt"
    "time"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type eventType int

const (
    Quit eventType = iota
    BeginBreak
    EndBreak
    Note
    BeginCombat
    EndCombat
)

func (e eventType) String() string {
    switch e {
    case Quit:
        return "Quit Application"
    case BeginBreak:
        return "Beginning of Break"
    case EndBreak:
        return "End of Break"
    case Note:
        return "Note"
    case BeginCombat:
        return "Start of Combat"
    case EndCombat:
        return "End of Combat"
    }
    return "Unknown Event Type"
}

type Event struct {
  Time time.Time
  Type eventType
  Text string
}

type Model struct {
    logFile string
    entries []Event
}

func InitialModel(logFile string) Model {
    return Model{logFile: logFile}
}

func (m Model) Init() tea.Cmd {
    return nil
}


func returnEvent(e Event) tea.Cmd {
    return func() tea.Msg {
        return e
    }
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        newEvent := Event{Time: time.Now()}
        switch msg.String() {
        case "ctrl+c", "q":
            newEvent.Type = Quit
            newEvent.Text = "Manually Quit Program"
        case "enter", " ", "n":
            newEvent.Type = Note
        case "b":
            newEvent.Type = BeginBreak
        case "r":
            newEvent.Type = EndBreak
        case "c":
            newEvent.Type = BeginCombat
        }
        m.entries = append(m.entries, newEvent)
        return m, returnEvent(newEvent)
    }
    return m, nil
}

func (m Model) View() string {
    title := "Pathfinder 2e Sessioner v.0.go.1\n(b)reak, (r)esume, (c)ombat, (n)ote"
    var titleStyle = lipgloss.NewStyle().
    Align(lipgloss.Center).
    BorderStyle(lipgloss.NormalBorder()).BorderBottom(true)
    padding := " "
    times := ""
    types := ""
    notes := ""
    for _, entry := range m.entries {
      times += fmt.Sprintf("[%s]\n", entry.Time.Format("3:4"))
      types += fmt.Sprintf("%s\n", entry.Type)
      notes += fmt.Sprintf("%s\n", entry.Text)
    }
    return lipgloss.JoinVertical(lipgloss.Center, titleStyle.Render(title), lipgloss.JoinHorizontal(lipgloss.Top, times, padding, types, padding, notes))
}
