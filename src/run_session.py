#!/usr/bin/env python3
import pydantic

from enum import Enum, auto
from typing import *

from stackoverflow_logging import log
from pathfinder import Character

type Arguments = Sequence[str]

class States(Enum):
  Exploration = auto()
  Break = auto()
  CombatSetup = auto()
  Combat = auto()

class Context(pydantic.BaseModel):
  state: States
  character: Optional[Character] = None
  previous_context: Optional[Self] = None

  def command_prompt(self: Self) -> str:
    match self.state:
      case States.Exploration:
        return '?: '
      case States.Break:
        return 'PAUSED: '
      case States.Combat:
        return '!{} {}/{}: '.format(self.character.name,
                                    self.character.actions_remaining,
                                    self.character.max_actions)
    return 'ERROR: '

class Command(pydantic.BaseModel):
  text: str
  description: str
  function: Callable[[Context, Arguments], Context]
  allowed_states: Optional[set[States]] = None
  blocked_states: Optional[set[States]] = None

  def can_run_in_context(self: Self, context: Context) -> bool:
    if self.blocked_states is not None:
      if context.state in self.blocked_states:
        log.warning("Command is not allowed in current state: ({}, {}).".format(self.text, context.state))
        return False
    if self.allowed_states is not None:
      if context.state not in self.allowed_states:
        log.warning("Command is not allowed in current state: ({}, {}).".format(self.text, context.state))
        return False
    return True

  def write_generic_history(self: Self, args: Arguments) -> None:
    hist_line = "> Ran '{}' with {}.".format(self.text.lower(), args)
    log.history(hist_line)

def printHelp(context: Context, args: Arguments) -> Context:
  if args == [] or args[0].upper() == Commands.HELP.value.text:
    log.info("Following commands:")
    for command_container in Commands:
      command = command_container.value
      log.info("'{}' - {}".format(command.text.lower(), command.description))
  elif len(args) == 1:
    log.info("Specific help is not yet supported.")
  return context

def addTodo(context: Context, args: Arguments) -> Context:
  todo = " ".join(args)
  with open('../todolist.md', 'a') as todo_list:
    todo_list.write('- [ ] {}'.format(todo))
  log.info("Added '{}' to todo list.".format(todo))
  return context

def quitSession(context: Context, args: Arguments) -> Context:
  quit()
  return context

def echoArguments(context: Context, args: Arguments) -> Context:
  log.info(args)
  return context

def controlBreak(context: Context, args: Arguments) -> Context:
  if len(args) == 0:
    log.warning("Command must be followed by either 'start' or 'stop'")
    return context
  elif len(args) > 1:
    log.warning("Too many arguments")
    return context
  sub_command = args[0]
  match sub_command.upper():
    case "START":
      log.info("Beginning break")
      return Context(state = States.Break,
                     previous_context = context)
    case "STOP":
      log.info("Resuming session")
      return context.previous_context
    case _:
      log.warning("Command must be followed by either 'start' or 'stop'")
      return context

def beginCombat(context: Context, args: Arguments) -> Context:
  log.info("Feature not yet implemented")
  return context

def saveContext(context: Context, args: Arguments) -> Context:
  log.info("Feature not yet implemented")
  return context

def loadContext(context: Context, args: Arguments) -> Context:
  log.info("Feature not yet implemented")
  return context

class Commands(Enum):
  HELP = Command(text = 'HELP',
                 description = "Prints this message with no arguments or provides more detail on another command.",
                 function = printHelp)
  TODO = Command(text = 'TODO',
                 description = "Appends note to todo list.",
                 function = addTodo)
  QUIT = Command(text = 'QUIT',
                 description = "Ignores any argumments, quits program.",
                 function = quitSession)
  ECHO = Command(text = 'ECHO',
                 description = "Echoes arguments.",
                 function = echoArguments)
  BREAK = Command(text = 'BREAK',
                  description = "Desginates start or end of a break, must be followed by 'start' or 'stop'.",
                  function = controlBreak)
  COMBAT = Command(text = 'COMBAT',
                   description = "Begins a combat encounter.",
                   function = beginCombat,
                   blocked_states = {States.Break})
  SAVE = Command(text = 'SAVE',
                 description = "Saves current context to file to load later. Program loads most recent save.",
                 function = saveContext)
  LOAD = Command(text = 'LOAD',
                 description = "Loads different saved context if most recent save isn't desired.",
                 function = loadContext)

def start_session() -> Context:
  log.info("New session started")
  return Context(state = States.Exploration)

def get_command(context: Context) -> tuple[Optional[Command], Arguments]:
  command_str: str = input(context.command_prompt())
  command_text: str = command_str.split(' ')[0]
  args: tuple[str] = command_str.split(' ')[1:]
  try:
    command: Command =  Commands[command_text.upper()].value
    return (command, args)
  except KeyError:
    log.warning("Unrecognised command.")
    return (None, args)

def main_loop(context: Context) -> Context:
  command, args = get_command(context)
  if command is None:
    return context
  log.debug("Command identified as: {}".format(command))

  if not command.can_run_in_context(context):
    return context
  log.debug("Command is runnable in this context")

  command.write_generic_history(args)
  context = command.function(context, args)
  return context

def main() -> None:
  context = start_session()
  while True:
    context = main_loop(context)

if __name__ == '__main__':
  main()
