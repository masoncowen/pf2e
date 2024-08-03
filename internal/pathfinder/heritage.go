package pathfinder

type Heritage int

const (
    Aiuvarin Heritage = iota
    Dromaar
    Beastkin
    HumanVersatile
    FetchlingLiminal
    HalflingNomadic
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
    case HalflingNomadic:
        return "Nomadic Halfling"
    }
    return "Unknown"
}
