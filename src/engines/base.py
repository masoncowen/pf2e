import pydantic

from typing import *

from utils.types import Arguments, CommandInfo
from utils.log import log

class Engine(pydantic.BaseModel):
  possible_commands: tuple[Type[CommandInfo]]
  command_info: Optional[Type[CommandInfo]] = None

  def can_run_command(self: Self) -> bool:
    pass

  def run_command(self: Self) -> None:
    pass

  def get_command(self: Self, arguments: Arguments) -> None:
    command_text = arguments.pop()
    for command in self.possible_commands:
      if command_text.upper() == command.text:
        self.command_info = command
        self.command_info.arguments = arguments
    return None

  def main_loop(self: Self, arguments: Arguments) -> None:
    self.get_command(arguments) 
    if self.command_info is None:
      log.debug("Engine command not identified")
      return None
    log.debug("Command identified as: {}".format(self.command_info))
    if not self.can_run_command():
      return None
    log.debug("Command is runnable")
    self.run_command()
    return None

  def write_generic_history(self: Self) -> None:
    hist_line = "> Ran '{}' with {}.".format(self.command_info.text.lower(), self.command_info.arguments)
    log.history(hist_line)
