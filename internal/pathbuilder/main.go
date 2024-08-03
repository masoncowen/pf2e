package pathbuilder

import (
    "errors"
    "encoding/json"
    pf "internal/pathfinder"
)

type PathbuilderAbilities struct {
    STR int `json:"str"`
    DEX int `json:"dex"`
    CON int `json:"con"`
    INT int `json:"int"`
    WIS int `json:"wis"`
    CHA int `json:"cha"`
    breakdown map[string]interface{}
}

type PathbuilderProficiencies struct {
    ClassDC int `json:"classDC"`
    Perception int `json:"perception"`
    Fortitude int `json:"fortitude"`
    Reflex int `json:"reflex"`
    Will int `json:"will"`
    Heavy int `json:"heavy"`
    Medium int `json:"medium"`
    Light int `json:"light"`
    Unarmored int `json:"unarmored"`
    Advanced int `json:"advanced"`
    Martial int `json:"martial"`
    Simple int `json:"simple"`
    Unarmed int `json:"unarmed"`
    CastingArcane int `json:"castingArcane"`
    CastingDivine int `json:"castingDivine"`
    CastingOccult int `json:"castingOccult"`
    CastingPrimal int `json:"castingPrimal"`
    Acrobatics int `json:"acrobatics"`
    Arcana int `json:"arcana"`
    Athletics int `json:"athletics"`
    Crafting int `json:"crafting"`
    Deception int `json:"deception"`
    Diplomacy int `json:"diplomacy"`
    Intimidation int `json:"intimidation"`
    Medicine int `json:"medicine"`
    Nature int `json:"nature"`
    Occultism int `json:"occultism"`
    Performance int `json:"performance"`
    Religion int `json:"religion"`
    Society int `json:"society"`
    Stealth int `json:"stealth"`
    Survival int `json:"survival"`
    Thievery int `json:"thievery"`
}

type PathbuilderCharacter struct {
    Name string `json:"name"`
    Class string `json:"class"`
    DualClass *string `json:"dualClass"`
    Level int `json:"level"`
    Ancestry string `json:"ancestry"`
    Heritage string `json:"heritage"`
    Background string `json:"background"`
    Alignment string `json:"alignment"`
    Gender string `json:"gender"`
    Age string `json:"age"`
    Deity string `json:"deity"`
    Size int `json:"size"`
    SizeName string `json:"sizeName"`
    Keyability string `json:"keyability"`
    Languages []string `json:"languages"`
    Rituals []string `json:"rituals"`
    Resistances []string `json:"resistances"`
    InventorMods []string `json:"inventorMods"`
    Attributes map[string]interface{} `json:"attributes"`
    Abilities *PathbuilderAbilities `json:"abilities"`
    Proficiencies *PathbuilderProficiencies `json:"proficiencies"`
    Mods map[string]interface{} `json:"mods"`
    Feats [][]interface{} `json:"feats"`
    Specials []string `json:"specials"`
    Lores [][]interface{} `json:"lores"`
    EquipmentContainers map[string]interface{} `json:"equipmentContainers"`
    Equipment []interface{} `json:"equipment"`
    SpecificProficiences map[string]interface{} `json:"specificProficiencies"`
    Money map[string]interface{} `json:"money"`
    Armor []interface{} `json:"armor"`
    SpellCasters []interface{}  `json:"spellCasters"`
    FocusPoints int `json:"focusPoints"`
    Focus map[string]interface{} `json:"focus"`
    Formula []interface{} `json:"formula"`
    AcTotal map[string]interface{} `json:"acTotal"`
    Pets []interface{} `json:"pets"`
    Familiars []interface{} `json:"familiars"`
}

type PathbuilderJSON struct {
    Success bool
    Build *PathbuilderCharacter
}

func (pbc *PathbuilderJSON) UnmarshallJSON(bytes []byte) error {
    err := json.Unmarshal(bytes, &pbc)
    if err != nil {
        return err
    }
    if !pbc.Success {
        return errors.New("Pathbuilder JSON indicates unsuccessful build")
    }
    return nil
}

func (pbc PathbuilderCharacter) ConvertClass() (pf.Class, error) {
    switch pbc.Class {
    case "Alchemist":
        return pf.Alchemist, nil
    case "Barbarian":
        return pf.Barbarian, nil
    case "Bard":
        return pf.Bard, nil
    case "Champion":
        return pf.Champion, nil
    case "Cleric":
        return pf.Cleric, nil
    case "Fighter":
        return pf.Fighter, nil
    case "Gunslinger":
        return pf.Gunslinger, nil
    case "Inventor":
        return pf.Inventor, nil
    case "Investigator":
        return pf.Investigator, nil
    case "Kineticist":
        return pf.Kineticist, nil
    case "Magus":
        return pf.Magus, nil
    case "Monk":
        return pf.Monk, nil
    case "Oracle":
        return pf.Oracle, nil
    case "Psychic":
        return pf.Psychic, nil
    case "Ranger":
        return pf.Ranger, nil
    case "Rogue":
        return pf.Rogue, nil
    case "Sorcerer":
        return pf.Sorcerer, nil
    case "Summoner":
        return pf.Summoner, nil
    case "Swashbuckler":
        return pf.Swashbuckler, nil
    case "Thaumaturge":
        return pf.Thaumaturge, nil
    case "Witch":
        return pf.Witch, nil
    case "Wizard":
        return pf.Wizard, nil
    }
    return pf.Alchemist, errors.New("Unrecognised Class")
}

func (pbc PathbuilderCharacter) ConvertAncestry() (pf.Ancestry, error) {
    switch pbc.Ancestry {
    case "Dwarf":
        return pf.Dwarf, nil
    case "Elf":
        return pf.Elf, nil
    case "Gnome":
        return pf.Gnome, nil
    case "Goblin":
        return pf.Goblin, nil
    case "Halfling":
        return pf.Halfling, nil
    case "Human":
        return pf.Human, nil
    case "Leshy":
        return pf.Leshy, nil
    case "Orc":
        return pf.Orc, nil
    case "Fetchling":
        return pf.Fetchling, nil
    }
    return pf.Halfling, errors.New("Unrecognised Ancestry")
}

func (pbc PathbuilderCharacter) ConvertHeritage() (pf.Heritage, error) {
    switch pbc.Heritage {
    case "Aiuvarin":
        return pf.Aiuvarin, nil
    case "Dromaar":
        return pf.Dromaar, nil
    case "Beastkin":
        return pf.Beastkin, nil
    case "Versatile Human":
        return pf.HumanVersatile, nil
    case "Liminal Fetchling":
        return pf.FetchlingLiminal, nil
    case "Nomadic Halfling":
        return pf.HalflingNomadic, nil
    }
    return pf.Aiuvarin, errors.New("Unrecognised Heritage")
}

func (pbc PathbuilderCharacter) ConvertBackground() (pf.Background, error) {
    switch pbc.Background {
    case "Crystal Healer":
        return pf.CrystalHealer, nil
    case "Magical Experiment (Enhanced Senses)":
        return pf.MagicalExperimentEnhancedSenses, nil
    case "Bounty Hunter":
        return pf.BountyHunter, nil
    case "Criminal":
        return pf.Criminal, nil
    }
    return pf.CrystalHealer, errors.New("Unrecognised Background")
}

func (pbc PathbuilderCharacter) Convert() (pf.Character, error) {
    newChar := pf.Character{}
    newChar.Name = pbc.Name
    newChar.Level = pbc.Level
    newChar.Modifiers.STR = pbc.Abilities.STR/2 - 5
    newChar.Modifiers.DEX = pbc.Abilities.DEX/2 - 5
    newChar.Modifiers.CON = pbc.Abilities.CON/2 - 5
    newChar.Modifiers.INT = pbc.Abilities.INT/2 - 5
    newChar.Modifiers.WIS = pbc.Abilities.WIS/2 - 5
    newChar.Modifiers.CHA = pbc.Abilities.CHA/2 - 5

    newClass, err := pbc.ConvertClass()
    if err != nil {
        return pf.Character{}, err
    }
    newChar.Class = newClass

    newAncestry, err := pbc.ConvertAncestry()
    if err != nil {
        return pf.Character{}, err
    }
    newChar.Ancestry = newAncestry

    newHeritage, err := pbc.ConvertHeritage()
    if err != nil {
        return pf.Character{}, err
    }
    newChar.Heritage = newHeritage

    newBackground, err := pbc.ConvertBackground()
    if err != nil {
        return pf.Character{}, err
    }
    newChar.Background = newBackground

    return newChar, nil
}
