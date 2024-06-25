#!/usr/bin/env python3
import pydantic

from enum import Enum, auto
from typing import *

from stackoverflow_logging import log
from pathfinder import Party
from engines import Engine
# from engines import CombatEngine
from private import HardCodedCombatEncountersPleaseChange, HardCodedPartyPleaseChange

type Arguments = Sequence[str]

class CommandInfo(pydantic.BaseModel):
  text: str
  description: str
  allowed_in_break: bool = True
  arguments: Optional[Arguments] = None

class UtilityCommands(Enum):
  HELP = CommandInfo(text = 'HELP', description = "Prints this message with no arguments or provides more detail on another command.")
  TODO = CommandInfo(text = 'TODO', description = "Appends note to todo list.")
  QUIT = CommandInfo(text = 'QUIT', description = "Ignores any arguments, quits program.")
  ECHO = CommandInfo(text = 'ECHO', description = "Echoes arguments.")
  BREAK = CommandInfo(text = 'BREAK', description = "Desginates start or end of a break, must be followed by 'start' or 'stop'.")
  SAVE = CommandInfo(text = 'SAVE', description = "Saves current context to file to load later. Program loads most recent save.")
  LOAD = CommandInfo(text = 'LOAD', description = "Loads different saved context if most recent save isn't desired.")
  BEGIN = CommandInfo(text = 'BEGIN', description = "Starts states that require engines (just combat currently).", allowed_in_break = False)

class Context(pydantic.BaseModel):
  party: Party
  engine: Optional[Type[Engine]] = None
  command_info: Optional[CommandInfo] = None
  previous_context: Optional[Self] = None
  is_paused: bool = False

  def command_prompt(self: Self) -> str:
    if self.is_paused:
      return 'PAUSED: '
    if self.engine is not None:
      return self.engine.command_prompt()
    return '?: '

  def get_utility_command(self: Self) -> None:
    command_str: str = input(self.command_prompt())
    command_text: str = command_str.split(' ')[0]
    arguments: Arguments = command_str.split(' ')[1:]
    try:
      temp_command_info: CommandInfo = UtilityCommands[command_text.upper()].value
      log.debug("Command identified")
      self.command_info: CommandInfo = temp_command_info
      self.command_info.arguments = arguments
    except KeyError:
      self.command_info = None
      if self.engine is None:
        log.warning("Unrecognised command while no engine was spinning.")
    return None

  def main_loop(self: Self) -> None:
    log.debug("Main loop started")
    self.get_utility_command()
    if self.command_info is None:
      if self.engine is not None:
        self.engine.main_loop()
        return None
    log.debug("Command identified as: {}".format(self.command_info))
    if not self.can_run_command():
      return None
    log.debug("Command is runnable in this context")
    self.run_utility_command()
    return None

  def can_run_command(self: Self) -> bool:
    if self.is_paused and not self.command_info.allowed_in_break:
      log.warning("Command is not allowed during breaks: ({}).".format(self.command_info.text))
      return False
    return True

  def run_utility_command(self: Self) -> None:
    self.write_generic_history()
    match self.command_info.text:
      case "HELP":
        self.printHelp()
      case "TODO":
        self.addTodo()
      case "QUIT":
        self.quitSession()
      case "ECHO":
        self.echoArguments()
      case "BREAK":
        self.controlBreak()
      case "SAVE":
        self.saveContext()
      case "LOAD":
        self.loadContext()
      case "BEGIN":
        self.beginEngine()

  def write_generic_history(self: Self) -> None:
    hist_line = "> Ran '{}' with {}.".format(self.command_info.text.lower(), self.command_info.arguments)
    log.history(hist_line)

  def printHelp(self: Self) -> None:
    if self.command_info.arguments == [] or self.command_info.arguments[0].upper() == UtilityCommands.HELP.value.text:
      log.info("Following command:")
      for command_container in UtilityCommands:
        command_summary = command_container.value
        log.info("'{}' - {}".format(command_summary.text.lower(), command_summary.description))
    elif len(self.command_info.arguments) == 1:
      log.info("Specific help is not yet supported.")
    return

  def addTodo(self: Self) -> None:
    todo = " ".join(self.command_info.arguments)
    with open('../todolist.md', 'a') as todo_list:
      todo_list.write('- [ ] {}'.format(todo))
    log.info("Added '{}' to todo list.".format(todo))
    return None

  def quitSession(self: Self) -> None:
    quit()
    return None

  def echoArguments(self: Self) -> None:
    log.info(self.command_info.arguments)
    return None

  def controlBreak(self: Self) -> None:
    if len(self.command_info.arguments) == 0:
      log.warning("Command must be followed by either 'start' or 'stop'")
      return None
    elif len(self.command_info.arguments) > 1:
      log.warning("Too many arguments")
      return None
    sub_command = self.command_info.arguments[0]
    match sub_command.upper():
      case "START":
        log.info("Beginning break")
        self.saveContext()
        self.is_paused = True
      case "STOP":
        log.info("Resuming session")
        self.loadContext()
        self.is_paused = False
      case _:
        log.warning("Command must be followed by either 'start' or 'stop'")
    return None

  def beginEngine(self: Self) -> None:
    if len(self.command_info.arguments) == 0:
      log.warning("Command must be followed by name of engine.")
      log.warning("Current engines: combat.")
      return None
    if len(self.command_info.arguments) > 1:
      log.warning("Too many arguments")
    match self.command_info.arguments[0].upper():
      case "COMBAT":
        self.debug("Combat is to be tested")
        # self.engine = CombatEngine(party = self.party, possible_encounters = HardCodedCombatEncountersPleaseChange)
    return None

  def saveContext(self: Self) -> None:
    log.info("Feature not yet implemented")
    return None

  def loadContext(self: Self) -> None:
    log.info("Feature not yet implemented")
    return None
  
def start_session() -> Context:
  log.info("New session started")
  party = []
  for member in HardCodedPartyPleaseChange:
    party.append(member.value)
  return Context(party = party)

def main() -> None:
  context = start_session()
  log.debug("Session should be setup")
  while True:
    log.debug("Loop outside context started")
    context.main_loop()

if __name__ == '__main__':
  main()
