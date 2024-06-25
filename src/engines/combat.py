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
  boss_creature_list: Optional[list[Creature]]
  creature_list: Optional[list[Creature]]
  threat_level: ThreatLevel

class CombatEngine(Engine):
  party: Party
  possible_encounters: Type[Enum]
  encounter: Optional[CombatEncounter] = None
  actions_remaing: int = 3

  # def __init__(self: Self, party: Party, possible_encounters: Type[Enum]):
  #   for encounter in possible_encounters:
  #     print(encounter)
  #     print(encounter.value)
  #   super().__init__(party=party, encounter=encounter)

  def command_prompt(self: Self) -> str:
    return '!{} {}/{}: '.format(self.combatant.name,
                                self.combatant.actions_remaining,
                                self.combatant.max_actions)
  def main_loop(self: Self) -> None:
    if self.encounter is None:
      print("HELLO")
