import pydantic

from typing import *

class Creature(pydantic.BaseModel):
  name: Optional[str] = None
  species_name: str
  level: int
  max_health: int
  AC: int
  max_actions: int = 3

  def xp_cost(self: Self, party_level: int) -> int:
    match self.level - party_level:
      case -4:
        return 10
      case -3:
        return 15
      case -2:
        return 20
      case -1:
        return 30
      case 0:
        return 40
      case 1:
        return 60
      case 2:
        return 80
      case 3:
        return 120
      case 4:
        return 160
