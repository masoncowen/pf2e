import pydantic

from typing import *
from enum import Enum, Flag, auto

from .rarity import Rarity

class WeaponTraining(Enum):
  Unarmed = 0
  Simple = auto()
  Martial = auto()
  Advanced = auto()
  SimpleFirearms = auto()
  MartialFirearms = auto()
  AdvancedFirearms = auto()

class WeaponTraits(Flag):
  Agile = auto()
  Backstabber = auto()
  Finesse = auto()

class Weapon(pydantic.BaseModel):
  name: str
  traits: list[WeaponTraits] = []
  rarity: Rarity = Rarity.Common
  training_required: WeaponTraining

class Weapons(Enum):
  Unarmed = Weapon(name = "Unarmed",
                   training_required = WeaponTraining.Unarmed)
