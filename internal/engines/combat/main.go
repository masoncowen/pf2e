package combat

import (
    "fmt"
    "sort"
	"strconv"
    "time"

	"internal/constants"
	pf "internal/pathfinder"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

type preparednessState int

const (
    selectingEncounter preparednessState = iota
    inputtingInitiative
    duringCombat
)


type Model struct {
    state preparednessState
    cursor int
    enteringText bool
    text string
    campaign constants.Campaign
    encounter pf.Encounter
    currentInitiative int
    actionsLeft int
    attacking bool
    attackCursor int
    combatants []ICombatant
}

func InitialModel() Model {
    return Model{}
}

func (m Model) Init() tea.Cmd {
    return nil
}

func (m Model) generateCombatants() []ICombatant {
    var combatants = []ICombatant{}
    for _, member := range m.campaign.Party {
        combatants = append(combatants, &member)
    }
    for _, enemy := range m.encounter.Enemies {
        enemy.Status.Health = enemy.MaxHealth
        combatants = append(combatants, &enemy)
    }
    return combatants
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case constants.Event:
        m.campaign = msg.Campaign
    case tea.KeyMsg:
        switch m.state {
        case selectingEncounter:
            switch msg.String() {
            case "ctrl+c", "q", "h", "left":
                return m, func() tea.Msg { return constants.Event{Type: constants.CancelCombat} }
            case "up", "k":
                if m.cursor > 0 {
                    m.cursor--
                }
            case "down", "j":
                if m.cursor < len(m.campaign.Encounters)-1 {
                    m.cursor++
                }
            case "enter", " ", "l", "right":
                m.encounter = m.campaign.Encounters[m.cursor]
                m.state = inputtingInitiative
                m.combatants = m.generateCombatants()
            }
        case inputtingInitiative:
            switch msg.String() {
            case "ctrl+c", "q", "h", "left":
                if !m.enteringText {
                    return m, func() tea.Msg { return constants.Event{Time: time.Now(), Type: constants.CancelCombat} }
                }
            case "up", "k":
                if m.cursor > 0 && !m.enteringText {
                    m.cursor--
                }
            case "down", "j":
                if (m.cursor < len(m.combatants)) && !m.enteringText { 
                    m.cursor++
                }
            case "enter", " ", "l", "right":
                if m.enteringText {
                    m.enteringText = false
                    tempInit, err := strconv.Atoi(m.text)
                    if err == nil {
                        m.combatants[m.cursor].SetInitiative(tempInit)
                        m.text = ""
                    }
                } else {
                    if m.cursor == len(m.combatants) {
                        allReady := true
                        for _, c := range m.combatants {
                            currInit := c.GetInitiative()
                            if currInit == 0 {
                                allReady = false
                                break
                            }
                            if currInit > m.currentInitiative {
                                m.currentInitiative = currInit
                            }
                        }
                        if allReady {
                            m.state = duringCombat
                            m.cursor = 0
                            m.actionsLeft = 3
                            sort.Slice(m.combatants, func(i, j int) bool {
                                return m.combatants[i].GetInitiative() > m.combatants[j].GetInitiative()
                            })
                        }
                    } else {
                        m.enteringText = true
                    }
                }
            case "backspace":
                if m.enteringText {
                    m.text = m.text[:len(m.text)-1]
                }
            default:
                if m.enteringText {
                    m.text = m.text + msg.String()
                }
            }
        case duringCombat:
            switch msg.String() {
            case "ctrl+c", "q":
                if !m.enteringText {
                    return m, func() tea.Msg { return constants.Event{Time: time.Now(), Type: constants.EndCombat} }
                }
            case "enter", " ":
                if m.enteringText {
                    m.enteringText = false
                    damage, err := strconv.Atoi(m.text)
                    if err == nil {
                        m.combatants[m.attackCursor].ReduceHealth(damage)
                        m.text = ""
                        m.actionsLeft--
                    }
                } else if m.attacking {
                    m.enteringText = true
                    m.attacking = false
                } else {
                    if m.actionsLeft > 0 {
                        m.actionsLeft--
                    } else if m.cursor < len(m.combatants)-1 {
                        m.actionsLeft = 3
                        m.cursor++
                    } else {
                        m.actionsLeft = 3
                        m.cursor = 0
                    }
                }
            case "tab":
                if !m.enteringText && !m.attacking {
                    if m.cursor < len(m.combatants)-1 {
                        m.actionsLeft = 3
                        m.cursor++
                    } else {
                        m.actionsLeft = 3
                        m.cursor = 0
                    }
                }
            case "up", "k":
                if m.attackCursor > 0 && m.attacking {
                    m.attackCursor--
                }
            case "down", "j":
                if (m.attackCursor < len(m.combatants)-1) && m.attacking { 
                    m.attackCursor++
                }
            case "a":
                if !m.enteringText {
                    m.attacking = true
                }
            case "backspace":
                if m.enteringText {
                    m.text = m.text[:len(m.text)-1]
                }
            default:
                if m.enteringText {
                    m.text = m.text + msg.String()
                }
            }
        }
    }
    return m, nil
}

func (m Model) View() string {
    s := "Pathfinder 2e Sessioner v.0.go.1\n\n"
    switch m.state {
    case selectingEncounter:
        for i, encounter := range m.campaign.Encounters {
            cursor := " "
            if m.cursor == i {
                cursor = ">"
            }

            s += fmt.Sprintf("%s %s\n", cursor, encounter.Name)
        }
        return s
    case inputtingInitiative:
        for i, combatant := range m.combatants {
            cursor := " "
            typing := ""
            if m.cursor == i {
                cursor = ">"
                if m.enteringText { 
                    typing = lipgloss.JoinHorizontal(lipgloss.Top, ": ", m.text, lipgloss.NewStyle().Blink(true).Render("|"))
                }
            }
            if m.enteringText && m.cursor == i{
                s += fmt.Sprintf("%s   %s%s\n", cursor, combatant.GetName(), typing)
            } else {
                if combatant.GetInitiative() != 0 {
                    s += fmt.Sprintf("%s %s %s: %d\n", cursor, "✔", combatant.GetName(), combatant.GetInitiative())
                } else {
                    s += fmt.Sprintf("%s   %s\n", cursor, combatant.GetName())
                }
            }
        }
        if m.cursor == len(m.combatants) {
            s += fmt.Sprintln("> Confirm")
        } else {
            s += fmt.Sprintln("  Confirm")
        }
        return s
    case duringCombat:
        t := table.New()
        for i, c := range m.combatants {
            cursor := " "
            if i == m.cursor {
                cursor = fmt.Sprintf("> (%d/3)", m.actionsLeft)
            }
            healthStatus := ""
            if m.attacking {
                if i == m.attackCursor {
                    healthStatus = ">> "
                } else {
                    healthStatus = "   "
                }
            } else if  m.enteringText {
                healthStatus = lipgloss.JoinHorizontal(lipgloss.Top, m.text, lipgloss.NewStyle().Blink(true).Render("|"), " >> ")
            }

            if c.IsDead() {
                healthStatus += "DEAD"
            } else if c.IsDying() {
                healthStatus += fmt.Sprintf("DYING: %d", c.GetDying())
            } else {
                healthStatus += fmt.Sprintf("%d / %d", c.GetHealth(), c.GetMaxHealth())
            }

            t.Row(cursor, strconv.Itoa(c.GetInitiative()), c.GetName(), strconv.Itoa(c.GetAC()), healthStatus)
        }
        return lipgloss.JoinVertical(lipgloss.Center, s, t.Render())
    }
    return s
}
