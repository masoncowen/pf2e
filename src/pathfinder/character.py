import pydantic

from enum import Enum, auto
from typing import *

# This class is for temporary information
class Status(pydantic.BaseModel):
  initiative: int
  actions_remaining: int = 3
  health: int

class pfClass(pydantic.BaseModel):
  health_per_level: int

class pfClasses(Enum):
  Alchemist = pfClass(health_per_level = 8)
  Barbarian = pfClass(health_per_level = 12)
  Bard = pfClass(health_per_level = 8)
  Champion = pfClass(health_per_level = 10)
  Cleric = pfClass(health_per_level = 8)
  Druid = pfClass(health_per_level = 8)
  Fighter = pfClass(health_per_level = 10)
  Gunslinger = pfClass(health_per_level = 8)
  Inventor = pfClass(health_per_level = 8)
  Investigator = pfClass(health_per_level = 8)
  Kineticist = pfClass(health_per_level = 8)
  Magus = pfClass(health_per_level = 8)
  Monk = pfClass(health_per_level = 10)
  Oracle = pfClass(health_per_level = 8)
  Psychic = pfClass(health_per_level = 6)
  Ranger = pfClass(health_per_level = 10)
  Rogue = pfClass(health_per_level = 8)
  Sorcerer = pfClass(health_per_level = 6)
  Summoner = pfClass(health_per_level = 10)
  Swashbuckler = pfClass(health_per_level = 10)
  Thaumaturge = pfClass(health_per_level = 8)
  Witch = pfClass(health_per_level = 6)
  Wizard = pfClass(health_per_level = 6)

class PlayerCharacter(pydantic.BaseModel):
  name: str
  level: int = 1
  AC: int = 15 #TODO: calculate AC = 10 + min(DEX, Armor.Dex_cap) + proficiency bonus (level + 2,4,6,8) + armour's item bonus + other bonuses + penalties
  CON: int = 2
  health_bonus: int = 0
  pf_class: pfClass
  status: Optional[Status] = None
  max_actions: int = 3

  @pydantic.computed_field
  @property
  def max_health(self: Self) -> int:
    return (self.pf_class.health_per_level + self.CON) * self.level + self.health_bonus
