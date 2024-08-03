package pathfinder

type Class int

const (
    Alchemist Class = iota
    Barbarian
    Bard
    Champion
    Cleric
    Druid
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
    case Druid:
        return "Druid"
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

func (c Class) BaseHealth() int {
    switch c {
    case Alchemist:
        return 8
    case Barbarian:
        return 12
    case Bard:
        return 8
    case Champion:
        return 10
    case Cleric:
        return 8
    case Druid:
        return 8
    case Fighter:
        return 10
    case Gunslinger:
        return 8
    case Inventor:
        return 8
    case Investigator:
        return 8
    case Kineticist:
        return 8
    case Magus:
        return 8
    case Monk:
        return 10
    case Oracle:
        return 8
    case Psychic:
        return 6
    case Ranger:
        return 10
    case Rogue:
        return 8
    case Sorcerer:
        return 6
    case Summoner:
        return 10
    case Swashbuckler:
        return 10
    case Thaumaturge:
        return 8
    case Witch:
        return 6
    case Wizard:
        return 6
    }
    return 0
}
