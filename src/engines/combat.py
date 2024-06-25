import pydantic

from typing import *
from enum import Enum, auto

from pathfinder import Character, Creature, Party
from .base import Engine
type Combatant = Union[Character, Creature]

class ThreatLevel(Enum):
  Trivial  = auto()
  Low      = auto()
  Moderate = auto()
  Severe   = auto()
  Extreme  = auto()

class CombatEncounter(pydantic.BaseModel):
  boss_creature_list: list[Creature]
  creature_list: list[Creature]
  threat_level: ThreatLevel

class CombatEngine(Engine):
  encounter: CombatEncounter
  actions_remaing: int = 3
  def __init__(self: Self):
    pass

  def command_prompt(self: Self) -> str:
    return '!{} {}/{}: '.format(self.combatant.name,
                                self.combatant.actions_remaining,
                                self.combatant.max_actions)
  def main_loop(self: Self) -> None:
