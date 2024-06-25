import pydantic

from typing import *

type Arguments = Sequence[str]

class CommandInfo(pydantic.BaseModel):
  text: str
  description: str
  arguments: Optional[Arguments] = None
