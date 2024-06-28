import pydantic

from typing import *
from enum import Enum, auto

from pathfinder import Character, Creature, Party
from .base import Engine
from utils.types import CommandInfo
from utils.log import log

type Combatant = Union[Character, Creature]

class CombatCommandInfo(CommandInfo):
  placeholder: bool = True

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
    log.debug("xp_budget: {}".format(xp_budget))
    party_levels = [mem.level for mem in party]
    party_level = sum(party_levels) / len(party_levels)
    log.debug("party_level: {}".format(party_level))
    if self.boss_creature_list is not None:
      for boss in self.boss_creature_list:
        xp_budget -= boss.xp_cost(party_level)
    log.debug("xp_budget: {}".format(xp_budget))
    xp_cost = self.filler_creature.xp_cost(party_level)
    log.debug("xp_cost: {}".format(xp_cost))
    count = int(xp_budget / xp_cost) 
    log.debug("count: {}".format(count))
    return count

class InitativeTracker(pydantic.BaseModel):
  combatant_dict: Dict[int, list[Combatant]] = {}
  actions_remaining: int = 3
  initiative: int = 0
  max_initiative: int = 0
  same_initiative_index: int = 0
  holding: list[Combatant] = []
  delayed_combatant: Optional[Combatant] = None
  delay_counter: int = 0

  def __init__(self: Self, party: Party, encounter: CombatEncounter):
    super().__init__()
    for member in party:
      self.add_to_initiative(member)

    if encounter.boss_creature_list is not None:
      for creature in encounter.boss_creature_list:
        self.add_to_initiative(creature)

    if (minion := encounter.filler_creature) is not None:
      minion_count = encounter.calculate_minion_count(party)
      self.add_to_initiative(minion, count=minion_count)

    for init in self.combatant_dict:
      if init > self.max_initiative:
        self.max_initiative = init
    self.initiative = self.max_initiative

  def add_to_initiative(self: Self, combatant: Combatant, count: int = 1, initiative: int = 0) -> None:
    if count == 0:
      return
    if initiative == 0:
      possible_init = input("Initiative for {}?".format(combatant.name))
      try:
        initiative = int(possible_init)
      except Exception as e:
        log.info("Empty or invalid initiative, setting to 1")
        initiative = 1
    if initiative in self.combatant_dict:
      self.combatant_dict[initiative].append(combatant)
    else:
      self.combatant_dict[initiative] = [combatant]
    if count > 1:
      for i in range(count - 1):
        self.combatant_dict[initiative].append(combatant)

  def first(self: Self) -> Combatant:
    return self.combatant_dict[self.max_initiative][0]

  def current(self: Self) -> Combatant:
    if self.delayed_combatant is None:
      return self.combatant_dict[self.initiative][self.same_initiative_index]
    return self.delayed_combatant

  def skip(self: Self) -> None:
    if self.actions_remaining < self.current().max_actions:
      log.warning("Cannot skip a turn once an action has been used.")
    self.next()

  def delay(self: Self) -> None:
    if self.actions_remaining < self.current().max_actions:
      log.warning("Cannot skip a turn once an action has been used.")
    self.holding.append(self.current())
    self.delay_counter += 1
    self.next()

  def next(self: Self) -> None:
    if len(self.holding) > 0 and self.delay_counter == 0:
      self.delayed_combatant = self.holding.pop(0)
      return None
    if len(self.holding) == 0 and self.delayed_combatant is not None:
      self.delayed_combatant = None
    if self.delay_counter > 0:
      self.delay_counter -= 1
    if len(current_initiative_combatants := self.combatant_dict[self.initiative]) - 1 > self.same_initiative_index:
      log.debug("More combatants at same initiative")
      self.same_initiative_index += 1
      return None
    log.debug("Moving onto next initiative count.")
    self.same_initiative_index = 0
    next_init = 0
    for init in self.combatant_dict:
      if self.initiative > init > next_init:
        next_init = init
    self.initiative = next_init
    if next_init == 0:
      self.initiative = self.max_initiative

  def set_new_combatant_initiative(self: Self, new_initiative) -> None:
    combatant = (self.combatant_dict[self.initiative]
                     .pop(self.same_initiative_index))
    self.add_to_initiative(combatant, initiative = new_initiative)
    self.next()

  def list_order(self: Self) -> None:  
    for i in range(self.max_initiative, 0, -1):
      if i in self.combatant_dict:
        for combatant in self.combatant_dict[i]:
          log.info("Initiative {}: {}".format(i, combatant.name))
    return None

class CombatEngine(Engine):
  encounter: Optional[CombatEncounter] = None
  initiative_tracker: Optional[InitativeTracker] = None
  possible_commands: tuple[CombatCommandInfo] = (
      CombatCommandInfo(text = "SKIP", description = "Skips combatant until next round."),
      CombatCommandInfo(text = "DELAY", description = "Delays combatant's turn until after next combatant."),
      CombatCommandInfo(text = "SHOW", description = "Shows value of variable."),
      CombatCommandInfo(text = "SET", description = "Changes variables to value."),
      CombatCommandInfo(text = "ADJUST", description = "Adjust variables by value."),
      # CombatCommandInfo(text = "ATTACK", description = "A
      )

  def can_run_command(self: Self) -> bool:
    return True

  def run_command(self: Self) -> None:
    self.write_generic_history()
    match self.command_info.text:
      case "SKIP":
        self.initiative_tracker.skip()
      case "DELAY":
        self.initiative_tracker.delay()
      case "SHOW":
        self.showInformation()
      case "SET":
        self.setInformation()
      case "ADJUST":
        self.adjustInformation()
    return None

  def command_prompt(self: Self) -> str:
    if self.initiative_tracker is None:
      return "!!: "
    current = self.initiative_tracker.current()
    return '!{} ({}/{}) AC:{} HP:{}/{}: '.format(current.name,
                                 self.initiative_tracker.actions_remaining,
                                 current.max_actions,
                                 current.AC,
                                 current.status.health,
                                 current.max_health)

  def select_encounter(self: Self, possible_encounters: Type[Enum]) -> None:
    log.info("Select encounter:")
    possible_idx = []
    for idx, encounter in enumerate(possible_encounters):
      log.info(" {} - {}".format(idx, encounter.value.description))
      possible_idx.append(idx)
    option = input("?: ")
    if int(option) not in possible_idx:
      log.warning("Invalid option")
      return self.select_encounter(possible_encounters)
    for idx, encounter in enumerate(possible_encounters):
      if idx == int(option):
        self.encounter = encounter.value
        return None

  def setup(self: Self, party: Party, possible_encounters: Type[Enum]) -> None:
    self.select_encounter(possible_encounters)
    self.initiative_tracker = InitativeTracker(party, self.encounter)
    return None

  def showInformation(self: Self) -> None:
    args = self.command_info.arguments
    if args is None or len(args) < 1:
      return None
    match self.command_info.arguments[0].upper():
      case "ORDER":
        self.initiative_tracker.list_order()
      case "HEALTH":
        current_combatant = self.initiative_tracker.current()
        name = current_combatant.name
        chealth = current_combatant.status.health
        mhealth = current_combatant.max_health
        log.info("{}'s current health is: {}/{}".format(name, chealth, mhealth))
      case "AC":
        current_combatant = self.initiative_tracker.current()
        log.info("{}'s AC is: {}".format(current_combatant.name, current_combatant.AC))
      case "INITIATIVE":
        current_combatant = self.initiative_tracker.current()
        log.info("{}'s iniative is: {}".format(current_combatant.name, self.initiative_tracker.initiative))

  def adjustInformation(self: Self):
    match self.command_info.arguments[0].upper():
      case "HEALTH":
        try:
          hp_change = int(self.command_info.arguments[1])
          self.initiative_tracker.current().status.health += hp_change
          self.showInformation()
        except Exception as e:
          log.debug(e)
          log.info(e)

  def setInformation(self: Self):
    match self.command_info.arguments[0].upper():
      case "HEALTH":
        try:
          new_health = int(self.command_info.arguments[1])
          self.initiative_tracker.current().status.health = new_health
          self.showInformation()
        except Exception as e:
          log.debug(e)
          return None
      case "INITIATIVE":
        try:
          new_initiative = int(self.command_info.arguments[1])
        except Exception as e:
          log.debug(e)
          return None
        self.initiative_tracker.set_new_combatant_initiative(new_initiative)
        return None

