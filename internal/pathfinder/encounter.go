package pathfinder

type Attack struct {
}

type Creature struct {
    PersonalName string
    CreatureName string
    Level int
    Speed int
    MaxHealth int
    AC int
    Status Status
    Modifiers Modifiers
    Proficiencies Proficiencies
    Attacks []Attack
}

type threat int

const (
    trivial threat = iota
    low
    moderate
    severe
    extreme
)

type Encounter struct {
    Name string
    Severity threat
    Enemies []Creature
    Minions []Creature
}

func (c Creature) IsDead() bool {
    if c.Status.Health > 0 {
        return false
    }
    return true
}

func (c Creature) IsDying() bool {
    return false
}

func (c *Creature) ReduceHealth(amount int) {
    if amount > c.Status.Health {
        c.Status.Health = 0
        c.Status.Dying = 4
        return
    }
    c.Status.Health -= amount
    return
}
func (c Creature) GetName() string {
    if c.PersonalName != "" {
        return c.PersonalName
    }
    return c.CreatureName
}

func (c Creature) GetInitiative() int {
    return c.Status.Initiative
}

func (c *Creature) SetInitiative(initiative int) {
    c.Status.Initiative = initiative
}

func (c Creature) GetHealth() int {
    return c.Status.Health
}

func (c Creature) GetMaxHealth() int {
    return c.MaxHealth
}

func (c Creature) GetDying() int {
    return c.Status.Dying
}

func (c Creature) GetAC() int {
    return c.AC
}
