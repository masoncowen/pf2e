package constants

import (
    pf "internal/pathfinder"
)

type NPC struct {}
type Campaign struct {
    Name string
    Path string
    Party []pf.Character
    NPCs []NPC
    Encounters []pf.Encounter
    Creatures []pf.Creature
}
