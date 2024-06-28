import pydantic

from typing import *

class Status(pydantic.BaseModel):
  health: int

class Creature(pydantic.BaseModel):
  name: Optional[str] = None
  species_name: str
  level: int
  max_health: int
  AC: int
  max_actions: int = 3
  status: Optional[Status] = None

  def copy(self: Self) -> Self:
    return Creature(species_name = self.species_name,
                    level = self.level,
                    max_health = self.max_health,
                    AC = self.AC)

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

  @pydantic.model_validator(mode='after')
  def get_name(self: Self) -> Self:
      if self.name is None:
          self.name = self.species_name
      return self

  @pydantic.model_validator(mode='after')
  def get_status(self: Self) -> Self:
    if self.status is None:
        self.status = Status(health = self.max_health)
    return self
