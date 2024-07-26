package pathfinder

type Background int

const (
    CrystalHealer Background = iota
    MagicalExperimentEnhancedSenses
    BountyHunter
    Criminal
)

func (b Background) String() string {
    switch b {
    case CrystalHealer:
        return "Crystal Healer"
    case MagicalExperimentEnhancedSenses:
        return "Magical Experiment (Enhanced Senses)"
    case BountyHunter:
        return "Bounty Hunter"
    case Criminal:
        return "Criminal"
    }
    return "Unknown"
}
