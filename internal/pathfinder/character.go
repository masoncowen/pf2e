package pathfinder

type Status struct {
    Initiative int
    Health int
    Dying int
    Wounded int
}

type Modifiers struct {
    STR int
    DEX int
    CON int
    INT int
    WIS int
    CHA int
}

type Character struct {
    Name string
    Class Class
    Level int
    Ancestry Ancestry
    Heritage Heritage
    Background Background
    Status Status
    Modifiers Modifiers
    Proficiencies Proficiencies
}

func (c Character) IsDead() bool {
    if c.Status.Health > 0 {
        return false
    }
    if c.Status.Dying < 4 {
        return false
    }
    return true
}

func (c Character) IsDying() bool {
    if c.Status.Dying == 0 {
        return false
    }
    if c.Status.Dying > 3 {
        return false
    }
    return true
}

func (c Character) GetName() string {
    return c.Name
}

func (c *Character) ReduceHealth(amount int) {
    if amount > c.Status.Health {
        c.Status.Health = 0
        c.Status.Dying = c.Status.Wounded + 1
        return
    }
    if amount > 2 * c.GetMaxHealth() {
        c.Status.Health = 0
        c.Status.Dying = 4
        return
    }
    if amount < 0 && c.IsDying() {
        c.Status.Wounded++
        c.Status.Dying = 0
    }
    c.Status.Health -= amount
    return
}

func (c Character) GetInitiative() int {
    return c.Status.Initiative
}

func (c *Character) SetInitiative(initiative int) {
    c.Status.Initiative = initiative
}

func (c Character) GetHealth() int {
    return c.Status.Health
}

func (c Character) GetMaxHealth() int {
    var maxHealth int = c.Ancestry.BaseHealth() + c.Class.BaseHealth() + c.Modifiers.CON
    return maxHealth
}

func (c Character) GetDying() int {
    return c.Status.Dying
}

func (c Character) GetAC() int {
    return 10 + c.Modifiers.DEX
}
