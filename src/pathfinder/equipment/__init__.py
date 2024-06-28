import pydantic

from typing import *

from .weapons import Weapons, Weapon, WeaponTraining
from .armours import Armours, Armour, ArmourWeight
from .shields import Shields, Shield

class Equipment(pydantic.BaseModel):  
  weapon: Weapon = Weapons.Unarmed.value
  armour: Armour = Armours.Unarmoured.value
  shield: Optional[Shield] = None
