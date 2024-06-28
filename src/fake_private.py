from enum import Enum

from pathfinder.character import Ancestries, PlayerCharacter
from pathfinder.pfclass import pfClasses
from pathfinder.creature import Creature
from engines.combat import CombatEncounter, ThreatLevel

class Creatures(Enum):
  RedDragon = Creature(species_name = "Red Dragon", level = -1, max_health = 8, AC = 15)
  DaveFromAccounting = Creature(name = "Dave", species_name = "Accountant", level = 3, max_health = 50, AC = 17)
  Steve = Creature(species_name = "Steve", level = 5, max_health = 400, AC = 30)

class HardCodedCombatEncountersPleaseChange(Enum):
  AngryDragon = CombatEncounter(threat_level = ThreatLevel.Low,
                                       description = "You have angered the dragon, kill it",
                                       boss_creature_list = [Creatures.RedDragon.value],
                                       filler_creature = None)
  DaveLostHisPet = CombatEncounter(threat_level = ThreatLevel.Moderate,
                                    description = "Dave has found that you killed his dragon",
                                    boss_creature_list = [Creatures.DaveFromAccounting.value],
                                    filler_creature = Creatures.RedDragon.value)
  TalkToMyManager = CombatEncounter(threat_level = ThreatLevel.Severe,
                                     description = "Dave has escalated to a higher being",
                                     boss_creature_list = [
                                       Creatures.Steve.value,
                                       Creatures.DaveFromAccounting.value
                                       ],
                                     filler_creature = Creatures.RedDragon.value)

class HardCodedPartyPleaseChange(Enum):
  Jake = PlayerCharacter(name = "Bhelmar", AC=14, CON=3, pf_class = pfClasses.Druid.value, pf_ancestry = Ancestries.Dwarf.value)
