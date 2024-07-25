package replacemetimer

import (
    "fmt"
    "time"
    tea "github.com/charmbracelet/bubbletea"
)

type entry struct {
  entryTime time.Time
  entryText string
}
  

type Model struct {
    entries []entry
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
        switch msg.String() {
        case "ctrl+c", "q":
            return m, tea.Quit
        case "enter", " ":
            m.entries = append(m.entries, entry{time.Now(), "Blank Message"})
          case "b":
            m.entries = append(m.entries, entry{time.Now(), "Break"})
        }
    }
    return m, nil
}

func (m Model) View() string {
    s := "Pathfinder 2e Sessioner v.0.go.1\n\n"
    for _, entry := range m.entries {
      s += fmt.Sprintf("%s\n", time.Now().String())
      s += fmt.Sprintf("%s %s\n", entry.entryTime.Format("3:4"), entry.entryText)
    }
    return s
}
