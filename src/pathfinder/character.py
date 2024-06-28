import pydantic

from enum import Enum, auto
from typing import *

from .equipment import Equipment, Weapons, Armours, Shields
from .pfclass import pfClass
# This class is for temporary information
class Status(pydantic.BaseModel):
  health: int


class pfAncestry(pydantic.BaseModel):  
  base_health: int

class pfAncestries(Enum):
  Dwarf = pfAncestry(base_health = 10)
  Elf = pfAncestry(base_health = 6)
  Gnome = pfAncestry(base_health = 8)
  Goblin = pfAncestry(base_health = 6)
  Halfling = pfAncestry(base_health = 6)
  Human = pfAncestry(base_health = 8)
  Leshy = pfAncestry(base_health = 8)
  Orc = pfAncestry(base_health = 10)

class PlayerCharacter(pydantic.BaseModel):
  name: str
  level: int = 1
  AC: int = 15 #TODO: calculate AC = 10 + min(DEX, Armor.Dex_cap) + proficiency bonus (level + 2,4,6,8) + armour's item bonus + other bonuses + penalties
  CON: int = 2
  health_bonus: int = 0
  pf_class: pfClass
  pf_ancestry: pfAncestry
  status: Optional[Status] = None
  equipment: Equipment = Equipment()
  max_actions: int = 3

  @pydantic.computed_field
  @property
  def AC(self: Self) -> int:
    armour = self.equipment.armour
    log.debug(armour)
    effective_DEX = min(self.DEX, armour.DEX_cap)
    log.debug(effective_DEX)
    AC_bonus = armour.AC_bonus
    log.debug(AC_bonus)
    proficiency_bonus = self.proficiency_bonus(armour.weight)
    log.debug(proficiency_bonus)
    return 10 + effective_DEX + AC_bonus + proficiency_bonus
  @pydantic.computed_field
  @property
  def max_health(self: Self) -> int:
    return self.pf_ancestry.base_health + (self.pf_class.health_per_level + self.CON) * self.level + self.health_bonus

  @pydantic.model_validator(mode='after')
  def get_status(self: Self) -> Self:
    if self.status is None:
        self.status = Status(health = self.max_health)
    return self
