import pydantic

from enum import Enum, auto
from typing import Optional

# This class is for temporary information
class Status(pydantic.BaseModel):
  initiative: int
  actions_remaining: int = 3

class pfClass(Enum):
  Alchemist = auto()
  Barbarian = auto()
  Bard = auto()
  Champion = auto()
  Cleric = auto()
  Druid = auto()
  Fighter = auto()
  Gunslinger = auto()
  Inventor = auto()
  Investigator = auto()
  Kineticist = auto()
  Magus = auto()
  Monk = auto()
  Oracle = auto()
  Psychic = auto()
  Ranger = auto()
  Rogue = auto()
  Sorcerer = auto()
  Summoner = auto()
  Swashbuckler = auto()
  Thaumaturge = auto()
  Witch = auto()
  Wizard = auto()

class Character(pydantic.BaseModel):
  name: str
  level: int
  AC: int #TODO: calculate AC = 10 + min(DEX, Armor.Dex_cap) + proficiency bonus (level + 2,4,6,8) + armour's item bonus + other bonuses + penalties
  pf_class: pfClass
  status: Optional[Status] = None
