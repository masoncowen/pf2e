package pathfinder

type Heritage int

const (
    Aiuvarin Heritage = iota
    Dromaar
    Beastkin
    HumanVersatile
    FetchlingLiminal
)

func (h Heritage) String() string {
    switch h {
    case Aiuvarin:
        return "Aiuvarin"
    case Dromaar:
        return "Dromaar"
    case Beastkin:
        return "Beastkin"
    case HumanVersatile:
        return "Versatile Human"
    case FetchlingLiminal:
        return "Liminal Fetchling"
    }
    return "Unknown"
}
