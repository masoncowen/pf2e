import pydantic

from typing import *

class Engine(pydantic.BaseModel):

  def run_command(self: Self):
    pass
