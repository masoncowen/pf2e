import pydantic

from typing import *
from enum import Enum, auto

from stackoverflow_logging import log
from pathfinder import Character, Creature, Party
from .base import Engine
type Combatant = Union[Character, Creature]

class ThreatLevel(Enum):
  Trivial  = 40
  Low      = 60
  Moderate = 80
  Severe   = 120
  Extreme  = 160

  def character_adjustment(self: Self) -> int:
    match self.value:
      case 40:
        return 10
      case 60 | 80:
        return 20
      case 120:
        return 30
      case 160:
        return 40

class CombatEncounter(pydantic.BaseModel):
  description: str
  boss_creature_list: Optional[list[Creature]]
  filler_creature: Optional[Creature]
  threat_level: ThreatLevel

  def calculate_minion_count(self: Self, party: Party):
    if self.filler_creature is None:
      return 0
    xp_budget = self.threat_level.value + (len(party) - 4) * self.threat_level.character_adjustment()
    party_levels = [mem.level for mem in party]
    party_level = sum(party_levels) / len(party_levels)
    xp_cost = self.filler_creature.xp_cost(party_level)
    return int(xp_budget / xp_cost)

class CombatEngine(Engine):
  actions_remaining: int = 3
  current_combatant: Optional[Combatant] = None
  combatant_dict: Optional[Dict[int, list[Combatant]]] = None

  def command_prompt(self: Self) -> str:
    if self.current_combatant is None:
      return "!!: "
    return '!{} {}/{}: '.format(self.current_combatant.name,
                                self.actions_remaining,
                                self.current_combatant.max_actions)

  def setup(self: Self, party: Party, possible_encounters: Type[Enum]):
    log.info("Select encounter:")
    for idx, encounter in enumerate(possible_encounters):
      log.info(" {} - {}".format(idx, encounter.value.description))
    option = input("?: ")
    chosen_encounter = None
    for idx, encounter in enumerate(possible_encounters):
      if idx == int(option):
        chosen_encounter = encounter.value
        
    self.combatant_dict = {}
    for member in party:
      initiative = int(input("Initiative for {}?".format(member.name)))
      if initiative in self.combatant_dict:
        self.combatant_dict[initiative].append(member)
        continue
      self.combatant_dict[initiative] = [member]

    if chosen_encounter.boss_creature_list is not None:
      for creature in chosen_encounter.boss_creature_list:
        initiative = int(input("Initiative for {}?".format(creature.species_name)))
        if initiative in self.combatant_dict:
          self.combatant_dict[initiative].append(creature)
          continue
        self.combatant_dict[initiative] = [creature]

    if minion := chosen_encounter.filler_creature is not None:
      minion_count = chosen_encounter.calculate_minion_count(party)
      if minion_count == 0:
        self.current_combatant = self.get_first_combatant()
        return
      initiative = int(input("Initiative for {}?".format(creature.species_name)))
      if initiative in self.combatant_dict:
        self.combatant_dict[initiative].append(minion)
      else:
        self.combatant_dict[initiative] = [minion]
      for i in range(minion_count - 1):
        self.combatant_dict[initiative].append(minion)
    self.current_combatant = self.get_first_combatant()
    return
  
  def get_first_combatant(self: Self):
    max_init = 0
    for init in self.combatant_dict:
      if init > max_init:
        max_init = init
    return self.combatant_dict[max_init][0]

  def main_loop(self: Self) -> None:
    if self.encounter is None:
      self.setup()
