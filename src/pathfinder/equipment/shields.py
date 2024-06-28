import pydantic

from enum import Enum
from typing import *

class Shield(pydantic.BaseModel):
  hardness: int

class Shields(Enum):
  pass
