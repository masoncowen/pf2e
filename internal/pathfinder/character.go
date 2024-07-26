package pathfinder

type Status struct {
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
}
