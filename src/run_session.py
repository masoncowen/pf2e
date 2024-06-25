#!/usr/bin/env python3
import pydantic

from enum import Enum, auto
from typing import *

from stackoverflow_logging import log
from pathfinder import Party
from engines import CombatEngine, Engine
from private import HardCodedCombatEncountersPleaseChange, HardCodedPartyPleaseChange

type Arguments = Sequence[str]

class States(Enum):
  Exploration = auto()
  Break = auto()
  CombatSetup = auto()
  Combat = auto()

class CommandInfo(pydantic.BaseModel):
  text: str
  description: str
  allowed_states: Optional[set[States]] = None
  blocked_states: Optional[set[States]] = None
  arguments: Optional[Arguments] = None

class UtilityCommands(Enum):
  HELP = CommandInfo(text = 'HELP', description = "Prints this message with no arguments or provides more detail on another command.")
  TODO = CommandInfo(text = 'TODO', description = "Appends note to todo list.")
  QUIT = CommandInfo(text = 'QUIT', description = "Ignores any arguments, quits program.")
  ECHO = CommandInfo(text = 'ECHO', description = "Echoes arguments.")
  BREAK = CommandInfo(text = 'BREAK', description = "Desginates start or end of a break, must be followed by 'start' or 'stop'.")
  SAVE = CommandInfo(text = 'SAVE', description = "Saves current context to file to load later. Program loads most recent save.")
  LOAD = CommandInfo(text = 'LOAD', description = "Loads different saved context if most recent save isn't desired.")
  BEGIN = CommandInfo(text = 'BEGIN', description = "Starts states that require engines (just combat currently).", blocked_states = {States.Break})

class Context(pydantic.BaseModel):
  party: Party
  state: States
  engine: Optional[Engine] = None
  command_info: Optional[CommandInfo] = None
  arguments: Arguments
  previous_context: Optional[Self] = None

  def command_prompt(self: Self) -> str:
    match self.state:
      case States.Exploration:
        return '?: '
      case States.Break:
        return 'PAUSED: '
      case States.Combat:
        return self.engine.command_prompt()
    return 'ERROR: '

  def get_utility_command(self: Self) -> None:
    command_str: str = input(context.command_prompt())
    command_text: str = command_str.split(' ')[0]
    arguments: Arguments = command_str.split(' ')[1:]
    try:
      temp_command_info: CommandInfo = UtilityCommands[command_text.upper()].value
      self.command_info: CommandInfo = temp_command_info
      self.command_info.arguments = arguments
    except KeyError:
      self.command_info = None
      if self.engine is None:
        log.warning("Unrecognised command while no engine was spinning.")
    return None

  def main_loop(self: Self) -> None:
    self.get_utility_command()
    if self.command_info is None:
      if self.engine is not None:
        self.engine.main_loop()
        return None
    log.debug("Command identified as: {}".format(self.command_info))
    if self.can_run_command():
      return None
    log.debug("Command is runnable in this context")
    self.run_utility_command(self.command_info)
    return None

  def can_run_command(self: Self) -> bool:
    if self.command_info.blocked_states is not None:
      if self.state in self.command_info.blocked_states:
        log.warning("Command is not allowed in current state: ({}, {}).".format(self.command_info.text, self.state))
        return False
    if self.command_info.allowed_states is not None:
      if self.state not in self.command_info.allowed_states:
        log.warning("Command is not allowed in current state: ({}, {}).".format(self.command_info.text, self.state))
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
    hist_line = "> Ran '{}' with {}.".format(self.command_info.text.lower(), self.arguments)
    log.history(hist_line)

  def printHelp(self: Self) -> None:
    if self.arguments == [] or self.arguments[0].upper() == CommandList.HELP.value.text:
      log.info("Following command:")
      for command_container in CommandList:
        command_info = command_container.value
        log.info("'{}' - {}".format(self.command_info.text.lower(), self.command_info.description))
    elif len(self.arguments) == 1:
      log.info("Specific help is not yet supported.")
    return

  def addTodo(self: Self) -> None:
    todo = " ".join(self.arguments)
    with open('../todolist.md', 'a') as todo_list:
      todo_list.write('- [ ] {}'.format(todo))
    log.info("Added '{}' to todo list.".format(todo))
    return None

  def quitSession(self: Self) -> None:
    quit()
    return None

  def echoArguments(self: Self) -> None:
    log.info(self.arguments)
    return None

  def controlBreak(self: Self) -> None:
    if len(self.arguments) == 0:
      log.warning("Command must be followed by either 'start' or 'stop'")
      return None
    elif len(arguments) > 1:
      log.warning("Too many arguments")
      return None
    sub_command = self.arguments[0]
    match sub_command.upper():
      case "START":
        log.info("Beginning break")
        self.saveContext()
        self.state = States.Break
      case "STOP":
        log.info("Resuming session")
        self.loadContext()
      case _:
        log.warning("Command must be followed by either 'start' or 'stop'")
    return None

  def beginEngine(self: Self) -> None:
    if len(self.arguments) == 0:
      log.warning("Command must be followed by name of engine.")
      log.warning("Current engines: combat.")
      return None
    if len(self.arguments) > 1:
      log.warning("Too many arguments")
    match self.arguments[0].upper():
      case "COMBAT":
        self.engine = CombatEngine(self.party)

    return None

  def saveContext(self: Self) -> None:
    log.info("Feature not yet implemented")
    return None

  def loadContext(self: Self) -> None:
    log.info("Feature not yet implemented")
    return None
  
def start_session() -> Context:
  log.info("New session started")
  return Context(state = States.Exploration)

def main() -> None:
  context = start_session()
  while True:
    context.main_loop()

if __name__ == '__main__':
  main()
