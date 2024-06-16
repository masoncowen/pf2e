#!/usr/bin/env python3
import pydantic

from enum import Enum, auto
from typing import *

from stackoverflow_logging import log

class States(Enum):
  Exploration = auto()
  Break = auto()
  CombatSetup = auto()
  Combat = auto()

class Character(pydantic.BaseModel):
  name: str
  actions_remaining: int = 3
  max_actions: int = 3

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
  function: Callable[[Context, tuple[str]], Context]
  allowed_states: Optional[Tuple[States]] = None
  blocked_states: Optional[Tuple[States]] = None

def printHelp(context: Context, args: tuple[str]) -> Context:
  if args == [] or args[0].upper() == Commands.HELP.value.text:
    log.info("Following commands:")
    for command_container in Commands:
      command = command_container.value
      log.info("'{}' - {}".format(command.text.lower(), command.description))
  elif len(args) == 1:
    log.info("Specific help is not yet supported.")
  return context

def addTodo(context: Context, args: tuple[str]) -> None:
  todo = " ".join(args)
  with open('todolist.md', 'a') as todo_list:
    todo_list.write('- [ ] {}'.format(todo))
  log.info("Added '{}' to todo list.".format(todo))
  return context

def quitSession(context: Context, args: tuple[str]) -> Context:
  quit()
  return context

def echoArguments(context: Context, args: tuple[str]) -> Context:
  log.info(args)
  return context

def controlBreak(context: Context, args: tuple[str]) -> Context:
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

def beginCombat(context: Context, args: tuple[str]) -> Context:
  log.info("we do be fighting")
  return Context(state = States.CombatSetup)

class Commands(Enum):
  HELP = Command(text = 'HELP',
                 description = "Prints this message with no arguments or provides more detail on another command",
                 function = printHelp)
  TODO = Command(text = 'TODO',
                 description = "appends note to todo list.",
                 function = addTodo)
  QUIT = Command(text = 'QUIT',
                 description = "ignores any argumments, quits repl.",
                 function = quitSession)
  ECHO = Command(text = 'ECHO',
                 description = "echoes arguments.",
                 function = echoArguments)
  BREAK = Command(text = 'BREAK',
                  description = "desginates start or end of a break, must be followed by 'start' or 'stop'",
                  function = controlBreak)
  COMBAT = Command(text = 'COMBAT',
                   description = "begins a combat encounter.",
                   function = beginCombat,
                   blocked_states = (States.Break,))

def start_session() -> Context:
  log.info("New session started")
  return Context(state = States.Exploration)

def write_generic_command_history(command: Command, args: tuple[str]) -> None:
  hist_line = "> Ran '{}' with {}.".format(command.text.lower(), args)
  log.history(hist_line)

def main_loop(context: Context) -> Context:
  command_str: str = input(context.command_prompt())
  command_text: str = command_str.split(' ')[0]
  args: tuple[str] = command_str.split(' ')[1:]
  
  try:
    command = Commands[command_text.upper()].value
  except KeyError:
    log.warning("Unrecognised command.")
    return context

  log.debug(command.blocked_states)
  log.debug(command.allowed_states)
  log.debug(context.state)
  if command.blocked_states is not None:
    if context.state in command.blocked_states:
      log.warning("Command is not allowed in current state: ({}, {}).".format(command.text, context.state))
      return context
  if command.allowed_states is not None:
    if context.state not in command.allowed_states:
      log.warning("Command is not allowed in current state: ({}, {}).".format(command.text, context.state))
      return context
  write_generic_command_history(command, args)
  context = command.function(context, args)
  log.debug("Returning context")
  log.debug(context)
  return context

def main() -> None:
  context = start_session()
  while True:
    context = main_loop(context)

if __name__ == '__main__':
  main()
