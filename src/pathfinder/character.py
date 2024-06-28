import pydantic

from enum import Enum, auto
from typing import *

from .equipment import Equipment, Weapons, Armours, Shields
from .proficiency import ProficiencyLevel, Proficiencies, Proficiency
from .pfclass import pfClass

from utils.log import log

# This class is for temporary information
class Status(pydantic.BaseModel):
  health: int

class Ancestry(pydantic.BaseModel):  
  base_health: int

class Ancestries(Enum):
  Dwarf = Ancestry(base_health = 10)
  Elf = Ancestry(base_health = 6)
  Gnome = Ancestry(base_health = 8)
  Goblin = Ancestry(base_health = 6)
  Halfling = Ancestry(base_health = 6)
  Human = Ancestry(base_health = 8)
  Leshy = Ancestry(base_health = 8)
  Orc = Ancestry(base_health = 10)

class PlayerCharacter(pydantic.BaseModel):
  name: str
  level: int = 1
  CON: int = 2
  DEX: int = 2
  health_bonus: int = 0
  pf_class: pfClass
  ancestry: Ancestry
  status: Optional[Status] = None
  equipment: Equipment = Equipment()
  proficiencies: Optional[Proficiencies] = None
  max_actions: int = 3

  def proficiency_bonus(self: Self, checked_proficiency: Proficiency) -> int:
    max_proficiency = ProficiencyLevel.Untrained
    for proficiency, level in self.proficiencies:
      log.debug("Checking proficiency {} against proficiency {}"
                .format(proficiency,checked_proficiency))
      if proficiency is not checked_proficiency:
        continue
      log.debug("Checking level {}:{} against level {}:{}"
                .format(level, level.value,
                        max_proficiency, max_proficiency.value))
      if level.value > max_proficiency.value:
        max_proficiency = level
        log.debug("Level is now {}".format(level))
    if max_proficiency is ProficiencyLevel.Untrained:
      log.debug("Returning 0")
      return 0
    log.debug("Returning {}".format(self.level + max_proficiency.value))
    return self.level + max_proficiency.value

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
    return self.ancestry.base_health + (self.pf_class.health_per_level + self.CON) * self.level + self.health_bonus

  @pydantic.model_validator(mode='after')
  def get_status(self: Self) -> Self:
    if self.status is None:
        self.status = Status(health = self.max_health)
    return self

  @pydantic.model_validator(mode='after')
  def get_proficiencies(self: Self) -> Self:
    if self.proficiencies is None:
        self.proficiencies = self.pf_class.initial_proficiencies
    return self
