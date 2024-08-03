package pathfinder

type Ancestry int

const (
    Dwarf Ancestry = iota
    Elf
    Gnome
    Goblin
    Halfling
    Human
    Leshy
    Orc
    Fetchling
)

func (a Ancestry) String() string {
    switch a {
    case Dwarf:
        return "Dwarf"
    case Elf:
      return "Elf"
    case Gnome:
      return "Gnome"
    case Goblin:
      return "Goblin"
    case Halfling:
      return "Halfling"
    case Human:
      return "Human"
    case Leshy:
      return "Leshy"
    case Orc:
      return "Orc"
    case Fetchling:
      return "Fetchling"
    }
    return "Unknown"
}

func (a Ancestry) BaseHealth() int {
    switch a {
    case Dwarf:
        return 10
    case Elf:
      return 6
    case Gnome:
      return 8
    case Goblin:
      return 6
    case Halfling:
      return 6
    case Human:
      return 8
    case Leshy:
      return 8
    case Orc:
      return 10
    case Fetchling:
      return 8
    }
    return 0
}

