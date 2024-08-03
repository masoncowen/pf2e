package constants

import (
    "time"
	tea "github.com/charmbracelet/bubbletea"
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
    CancelCombat
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
    case CancelCombat:
        return "Combat Cancelled during setup"
    case EndCombat:
        return "End of Combat"
    }
    return "Unknown Event Type"
}

type Event struct {
    Campaign Campaign
    Time time.Time
    Type EventType
    Text string
    NeedsNote bool
}

func (e Event) ReturnMsg() tea.Cmd {
    return func() tea.Msg {
        return e
    }
}
