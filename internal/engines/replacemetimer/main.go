package replacemetimer

import (
	"fmt"
	"time"
    "internal/constants"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
    campaign constants.Campaign
    entries []constants.Event
}

func InitialModel() Model {
    return Model{}
}

func (m Model) Init() tea.Cmd {
    return nil
}


func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case constants.NewSessionMsg:
        m.campaign = msg.Campaign
    case constants.Event:
        m.entries = append(m.entries[:len(m.entries)-1], msg)
    case tea.KeyMsg:
        newEvent := constants.Event{Time: time.Now()}
        switch msg.String() {
        case "ctrl+c", "q":
            newEvent.Type = constants.Quit
            newEvent.Text = "Manually Quit Program"
        case "enter", " ", "n":
            newEvent.Type = constants.Note
            newEvent.NeedsNote = true
        case "b":
            newEvent.Type = constants.BeginBreak
            newEvent.NeedsNote = true
        case "r":
            newEvent.Type = constants.EndBreak
        case "c":
            newEvent.Type = constants.BeginCombat
            newEvent.Campaign = m.campaign
        case "t":
            newEvent.Type = constants.ToDoItem
            newEvent.NeedsNote = true
        }
        m.entries = append(m.entries, newEvent)
        return m, newEvent.ReturnMsg()
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
