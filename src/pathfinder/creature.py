import pydantic

from typing import *

class Creature(pydantic.BaseModel):
  name: Optional[str] = None
  species_name: str
  level: int
  max_health: int
  AC: int
