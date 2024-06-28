import pydantic

from typing import *
from enum import Enum, auto

from .rarity import Rarity

class ArmourWeight(Enum):
  Unarmoured = 0
  Light = auto()
  Medium = auto()
  Heavy = auto()

class Armour(pydantic.BaseModel):
  name: str
  weight: ArmourWeight
  rarity: Rarity = Rarity.Common
  AC_bonus: int = 0
  DEX_cap: int = 100
  check_penalty: int = 0
  speed_penalty: int = 0
  strength: int = 0

class Armours(Enum):
  Unarmoured = Armour(name = "Unarmoured",
                      weight = ArmourWeight.Unarmoured)
  LeatherLamellar = Armour(name = "Leather Lamellar",
                           weight = ArmourWeight.Light,
                           AC_bonus = 1, DEX_cap = 4,
                           check_penalty = 1)
  Breastplate = Armour(name = "Breastplate",
                       weight = ArmourWeight.Medium,
                       AC_bonus = 4, DEX_cap = 1,
                       check_penalty = 2,
                       speed_penalty = 5, strength = 3)
  Sankeit = Armour(name = "Sankeit",
                   weight = ArmourWeight.Light,
                   AC_bonus = 2, DEX_cap = 3, check_penalty = 1)
  ScaleMail = Armour(name = "Scale Mail",
                     weight = ArmourWeight.Medium,
                     AC_bonus = 3, DEX_cap = 2, check_penalty = 2,
                     speed_penalty = 5, strength = 2)
