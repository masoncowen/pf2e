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
