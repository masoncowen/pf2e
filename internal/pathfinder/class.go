package pathfinder

type Class int

const (
    Alchemist Class = iota
    Barbarian
    Bard
    Champion
    Cleric
    Fighter
    Gunslinger
    Inventor
    Investigator
    Kineticist
    Magus
    Monk
    Oracle
    Psychic
    Ranger
    Rogue
    Sorcerer
    Summoner
    Swashbuckler
    Thaumaturge
    Witch
    Wizard
)

func (c Class) String() string {
    switch c {
    case Alchemist:
        return "Alchemist"
    case Barbarian:
        return "Barbarian"
    case Bard:
        return "Bard"
    case Champion:
        return "Champion"
    case Cleric:
        return "Cleric"
    case Fighter:
        return "Fighter"
    case Gunslinger:
        return "Gunslinger"
    case Inventor:
        return "Inventor"
    case Investigator:
        return "Investigator"
    case Kineticist:
        return "Kineticist"
    case Magus:
        return "Magus"
    case Monk:
        return "Monk"
    case Oracle:
        return "Oracle"
    case Psychic:
        return "Psychic"
    case Ranger:
        return "Ranger"
    case Rogue:
        return "Rogue"
    case Sorcerer:
        return "Sorcerer"
    case Summoner:
        return "Summoner"
    case Swashbuckler:
        return "Swashbuckler"
    case Thaumaturge:
        return "Thaumaturge"
    case Witch:
        return "Witch"
    case Wizard:
        return "Wizard"
    }
    return "Unknown"
}
