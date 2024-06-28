from enum import Enum
from typing import *

from .equipment import ArmourWeight, WeaponTraining
from .saving_throws import SavingThrows

class ProficiencyLevel(Enum):
  Untrained = 0
  Trained = 2
  Expert = 4
  Master = 6
  Legendary = 8

type Proficiency = Union[WeaponTraining, ArmourWeight, SavingThrows]
type Proficiencies = list[tuple[Proficiency, ProficiencyLevel]]
