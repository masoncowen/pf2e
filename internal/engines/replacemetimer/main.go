package replacemetimer

import (
	"fmt"
    "time"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type EventType int

const (
    Unknown EventType = iota
    Quit
    ToDoItem
    Note
    BeginBreak
    EndBreak
    BeginCombat
    EndCombat
)

func (e EventType) String() string {
    switch e {
    case Quit:
        return "Quit Application"
    case ToDoItem:
        return "Added ToDo Item for Application"
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
  Type EventType
  Text string
  NeedsNote bool
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
    case Event:
        m.entries = append(m.entries[:len(m.entries)-1], msg)
    case tea.KeyMsg:
        newEvent := Event{Time: time.Now()}
        switch msg.String() {
        case "ctrl+c", "q":
            newEvent.Type = Quit
            newEvent.Text = "Manually Quit Program"
        case "enter", " ", "n":
            newEvent.Type = Note
            newEvent.NeedsNote = true
        case "b":
            newEvent.Type = BeginBreak
            newEvent.NeedsNote = true
        case "r":
            newEvent.Type = EndBreak
        case "c":
            newEvent.Type = BeginCombat
        case "t":
            newEvent.Type = ToDoItem
            newEvent.NeedsNote = true
        }
        m.entries = append(m.entries, newEvent)
        return m, returnEvent(newEvent)
    }
    return m, nil
}

func (m Model) View() string {
    title := "Pathfinder 2e Sessioner v.0.go.1\n(b)reak, (r)esume, (c)ombat, (n)ote, (t)odo, (q)uit"
    var titleStyle = lipgloss.NewStyle().
    Align(lipgloss.Center).
    BorderStyle(lipgloss.NormalBorder()).BorderBottom(true)
    padding := " "
    times := ""
    types := ""
    notes := ""
    for _, entry := range m.entries {
      times += fmt.Sprintf("[%s]\n", entry.Time.Format("15:04"))
      types += fmt.Sprintf("%s\n", entry.Type)
      notes += fmt.Sprintf("%s\n", entry.Text)
    }
    return lipgloss.JoinVertical(lipgloss.Center, titleStyle.Render(title), lipgloss.JoinHorizontal(lipgloss.Top, times, padding, types, padding, notes))
}
