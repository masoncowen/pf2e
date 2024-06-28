import pydantic

from typing import *
from enum import Enum

from .proficiency import ProficiencyLevel, Proficiency
from .equipment import ArmourWeight, WeaponTraining
from .saving_throws import SavingThrows

class pfClass(pydantic.BaseModel):
  health_per_level: int
  initial_proficiencies: list[tuple[Proficiency, ProficiencyLevel]] = []

class pfClasses(Enum):
  Alchemist = pfClass(health_per_level = 8,
                  initial_proficiencies = [
                    (SavingThrows.Fortitude, ProficiencyLevel.Expert),
                    (SavingThrows.Reflex, ProficiencyLevel.Expert),
                    (SavingThrows.Will, ProficiencyLevel.Trained),
                    (WeaponTraining.Unarmed, ProficiencyLevel.Trained),
                    (WeaponTraining.Simple, ProficiencyLevel.Trained), #TODO: Trained in alchemical bombs specifically
                    (ArmourWeight.Unarmoured, ProficiencyLevel.Trained),
                    (ArmourWeight.Light, ProficiencyLevel.Trained),
                    (ArmourWeight.Medium, ProficiencyLevel.Trained),
                    ])
  Barbarian = pfClass(health_per_level = 12,
                  initial_proficiencies = [
                    (SavingThrows.Fortitude, ProficiencyLevel.Expert),
                    (SavingThrows.Reflex, ProficiencyLevel.Trained),
                    (SavingThrows.Will, ProficiencyLevel.Expert),
                    (WeaponTraining.Unarmed, ProficiencyLevel.Trained),
                    (WeaponTraining.Simple, ProficiencyLevel.Trained),
                    (WeaponTraining.Martial, ProficiencyLevel.Trained),
                    (ArmourWeight.Unarmoured, ProficiencyLevel.Trained),
                    (ArmourWeight.Light, ProficiencyLevel.Trained),
                    (ArmourWeight.Medium, ProficiencyLevel.Trained),
                    ])
  Bard = pfClass(health_per_level = 8,
                  initial_proficiencies = [
                    (SavingThrows.Fortitude, ProficiencyLevel.Trained),
                    (SavingThrows.Reflex, ProficiencyLevel.Trained),
                    (SavingThrows.Will, ProficiencyLevel.Expert),
                    (WeaponTraining.Unarmed, ProficiencyLevel.Trained),
                    (WeaponTraining.Simple, ProficiencyLevel.Trained),
                    (WeaponTraining.Martial, ProficiencyLevel.Trained),
                    (ArmourWeight.Unarmoured, ProficiencyLevel.Trained),
                    (ArmourWeight.Light, ProficiencyLevel.Trained),
                    ])
  Champion = pfClass(health_per_level = 10,
                  initial_proficiencies = [
                    (SavingThrows.Fortitude, ProficiencyLevel.Expert),
                    (SavingThrows.Reflex, ProficiencyLevel.Trained),
                    (SavingThrows.Will, ProficiencyLevel.Expert),
                    (WeaponTraining.Unarmed, ProficiencyLevel.Trained),
                    (WeaponTraining.Simple, ProficiencyLevel.Trained),
                    (WeaponTraining.Martial, ProficiencyLevel.Trained),
                    (ArmourWeight.Unarmoured, ProficiencyLevel.Trained),
                    (ArmourWeight.Light, ProficiencyLevel.Trained),
                    (ArmourWeight.Medium, ProficiencyLevel.Trained),
                    (ArmourWeight.Heavy, ProficiencyLevel.Trained),
                    ])
  Cleric = pfClass(health_per_level = 8,
                  initial_proficiencies = [
                    (SavingThrows.Fortitude, ProficiencyLevel.Trained),
                    (SavingThrows.Reflex, ProficiencyLevel.Trained),
                    (SavingThrows.Will, ProficiencyLevel.Expert),
                    (WeaponTraining.Unarmed, ProficiencyLevel.Trained),
                    (WeaponTraining.Simple, ProficiencyLevel.Trained), #TODO: Trained in God's favoured weapon
                    (ArmourWeight.Unarmoured, ProficiencyLevel.Trained),
                    ])
  Druid = pfClass(health_per_level = 8,
                  initial_proficiencies = [
                    (SavingThrows.Fortitude, ProficiencyLevel.Trained),
                    (SavingThrows.Reflex, ProficiencyLevel.Trained),
                    (SavingThrows.Will, ProficiencyLevel.Expert),
                    (WeaponTraining.Unarmed, ProficiencyLevel.Trained),
                    (WeaponTraining.Simple, ProficiencyLevel.Trained),
                    (ArmourWeight.Unarmoured, ProficiencyLevel.Trained),
                    (ArmourWeight.Light, ProficiencyLevel.Trained),
                    (ArmourWeight.Medium, ProficiencyLevel.Trained),
                    ])
  Fighter = pfClass(health_per_level = 10,
                  initial_proficiencies = [
                    (SavingThrows.Fortitude, ProficiencyLevel.Expert),
                    (SavingThrows.Reflex, ProficiencyLevel.Expert),
                    (SavingThrows.Will, ProficiencyLevel.Trained),
                    (WeaponTraining.Unarmed, ProficiencyLevel.Expert),
                    (WeaponTraining.Simple, ProficiencyLevel.Expert),
                    (WeaponTraining.Martial, ProficiencyLevel.Expert),
                    (WeaponTraining.Advanced, ProficiencyLevel.Trained),
                    (ArmourWeight.Unarmoured, ProficiencyLevel.Trained),
                    (ArmourWeight.Light, ProficiencyLevel.Trained),
                    (ArmourWeight.Medium, ProficiencyLevel.Trained),
                    (ArmourWeight.Heavy, ProficiencyLevel.Trained),
                    ])
  Gunslinger = pfClass(health_per_level = 8,
                  initial_proficiencies = [
                    (SavingThrows.Fortitude, ProficiencyLevel.Expert),
                    (SavingThrows.Reflex, ProficiencyLevel.Expert),
                    (SavingThrows.Will, ProficiencyLevel.Trained),
                    (WeaponTraining.Unarmed, ProficiencyLevel.Trained),
                    (WeaponTraining.Simple, ProficiencyLevel.Trained),
                    (WeaponTraining.Martial, ProficiencyLevel.Trained),
                    (WeaponTraining.SimpleFirearms, ProficiencyLevel.Expert),
                    (WeaponTraining.MartialFirearms, ProficiencyLevel.Expert),
                    (WeaponTraining.AdvancedFirearms, ProficiencyLevel.Trained),
                    (ArmourWeight.Unarmoured, ProficiencyLevel.Trained),
                    (ArmourWeight.Light, ProficiencyLevel.Trained),
                    (ArmourWeight.Medium, ProficiencyLevel.Trained),
                    ])
  Inventor = pfClass(health_per_level = 8,
                  initial_proficiencies = [
                    (SavingThrows.Fortitude, ProficiencyLevel.Expert),
                    (SavingThrows.Reflex, ProficiencyLevel.Trained),
                    (SavingThrows.Will, ProficiencyLevel.Expert),
                    (WeaponTraining.Unarmed, ProficiencyLevel.Trained),
                    (WeaponTraining.Simple, ProficiencyLevel.Trained),
                    (WeaponTraining.Martial, ProficiencyLevel.Trained),
                    (ArmourWeight.Unarmoured, ProficiencyLevel.Trained),
                    (ArmourWeight.Light, ProficiencyLevel.Trained),
                    (ArmourWeight.Medium, ProficiencyLevel.Trained),
                    ])
  Investigator = pfClass(health_per_level = 8,
                  initial_proficiencies = [
                    (SavingThrows.Fortitude, ProficiencyLevel.Trained),
                    (SavingThrows.Reflex, ProficiencyLevel.Expert),
                    (SavingThrows.Will, ProficiencyLevel.Expert),
                    (WeaponTraining.Unarmed, ProficiencyLevel.Trained),
                    (WeaponTraining.Simple, ProficiencyLevel.Trained),
                    (WeaponTraining.Martial, ProficiencyLevel.Trained),
                    (ArmourWeight.Unarmoured, ProficiencyLevel.Trained),
                    (ArmourWeight.Light, ProficiencyLevel.Trained),
                    ])
  Kineticist = pfClass(health_per_level = 8,
                  initial_proficiencies = [
                    (SavingThrows.Fortitude, ProficiencyLevel.Expert),
                    (SavingThrows.Reflex, ProficiencyLevel.Expert),
                    (SavingThrows.Will, ProficiencyLevel.Trained),
                    (WeaponTraining.Unarmed, ProficiencyLevel.Trained),
                    (WeaponTraining.Simple, ProficiencyLevel.Trained),
                    (ArmourWeight.Unarmoured, ProficiencyLevel.Trained),
                    (ArmourWeight.Light, ProficiencyLevel.Trained),
                    ])
  Magus = pfClass(health_per_level = 8,
                  initial_proficiencies = [
                    (SavingThrows.Fortitude, ProficiencyLevel.Expert),
                    (SavingThrows.Reflex, ProficiencyLevel.Trained),
                    (SavingThrows.Will, ProficiencyLevel.Expert),
                    (WeaponTraining.Unarmed, ProficiencyLevel.Trained),
                    (WeaponTraining.Simple, ProficiencyLevel.Trained),
                    (WeaponTraining.Martial, ProficiencyLevel.Trained),
                    (ArmourWeight.Unarmoured, ProficiencyLevel.Trained),
                    (ArmourWeight.Light, ProficiencyLevel.Trained),
                    (ArmourWeight.Medium, ProficiencyLevel.Trained),
                    ])
  Monk = pfClass(health_per_level = 10,
                  initial_proficiencies = [
                    (SavingThrows.Fortitude, ProficiencyLevel.Expert),
                    (SavingThrows.Reflex, ProficiencyLevel.Expert),
                    (SavingThrows.Will, ProficiencyLevel.Expert),
                    (WeaponTraining.Unarmed, ProficiencyLevel.Trained),
                    (WeaponTraining.Simple, ProficiencyLevel.Trained),
                    (ArmourWeight.Unarmoured, ProficiencyLevel.Expert),
                    ])
  Oracle = pfClass(health_per_level = 8,
                  initial_proficiencies = [
                    (SavingThrows.Fortitude, ProficiencyLevel.Trained),
                    (SavingThrows.Reflex, ProficiencyLevel.Trained),
                    (SavingThrows.Will, ProficiencyLevel.Expert),
                    (WeaponTraining.Unarmed, ProficiencyLevel.Trained),
                    (WeaponTraining.Simple, ProficiencyLevel.Trained),
                    (ArmourWeight.Unarmoured, ProficiencyLevel.Trained),
                    (ArmourWeight.Light, ProficiencyLevel.Trained),
                    ])
  Psychic = pfClass(health_per_level = 6,
                  initial_proficiencies = [
                    (SavingThrows.Fortitude, ProficiencyLevel.Trained),
                    (SavingThrows.Reflex, ProficiencyLevel.Trained),
                    (SavingThrows.Will, ProficiencyLevel.Expert),
                    (WeaponTraining.Unarmed, ProficiencyLevel.Trained),
                    (WeaponTraining.Simple, ProficiencyLevel.Trained),
                    (ArmourWeight.Unarmoured, ProficiencyLevel.Trained),
                    ])
  Ranger = pfClass(health_per_level = 10,
                  initial_proficiencies = [
                    (SavingThrows.Fortitude, ProficiencyLevel.Expert),
                    (SavingThrows.Reflex, ProficiencyLevel.Expert),
                    (SavingThrows.Will, ProficiencyLevel.Trained),
                    (WeaponTraining.Unarmed, ProficiencyLevel.Trained),
                    (WeaponTraining.Simple, ProficiencyLevel.Trained),
                    (WeaponTraining.Martial, ProficiencyLevel.Trained),
                    (ArmourWeight.Unarmoured, ProficiencyLevel.Trained),
                    (ArmourWeight.Light, ProficiencyLevel.Trained),
                    (ArmourWeight.Medium, ProficiencyLevel.Trained),
                    ])
  Rogue = pfClass(health_per_level = 8,
                  initial_proficiencies = [
                    (SavingThrows.Fortitude, ProficiencyLevel.Trained),
                    (SavingThrows.Reflex, ProficiencyLevel.Expert),
                    (SavingThrows.Will, ProficiencyLevel.Expert),
                    (WeaponTraining.Unarmed, ProficiencyLevel.Trained),
                    (WeaponTraining.Simple, ProficiencyLevel.Trained),
                    (WeaponTraining.Martial, ProficiencyLevel.Trained),
                    (ArmourWeight.Unarmoured, ProficiencyLevel.Trained),
                    (ArmourWeight.Light, ProficiencyLevel.Trained),
                    ])
  Sorcerer = pfClass(health_per_level = 6,
                  initial_proficiencies = [
                    (SavingThrows.Fortitude, ProficiencyLevel.Trained),
                    (SavingThrows.Reflex, ProficiencyLevel.Trained),
                    (SavingThrows.Will, ProficiencyLevel.Expert),
                    (WeaponTraining.Unarmed, ProficiencyLevel.Trained),
                    (WeaponTraining.Simple, ProficiencyLevel.Trained),
                    (ArmourWeight.Unarmoured, ProficiencyLevel.Trained),
                    ])
  Summoner = pfClass(health_per_level = 10,
                  initial_proficiencies = [
                    (SavingThrows.Fortitude, ProficiencyLevel.Expert),
                    (SavingThrows.Reflex, ProficiencyLevel.Trained),
                    (SavingThrows.Will, ProficiencyLevel.Expert),
                    (WeaponTraining.Unarmed, ProficiencyLevel.Trained),
                    (WeaponTraining.Simple, ProficiencyLevel.Trained),
                    (ArmourWeight.Unarmoured, ProficiencyLevel.Trained),
                    ])
  Swashbuckler = pfClass(health_per_level = 10,
                  initial_proficiencies = [
                    (SavingThrows.Fortitude, ProficiencyLevel.Trained),
                    (SavingThrows.Reflex, ProficiencyLevel.Expert),
                    (SavingThrows.Will, ProficiencyLevel.Expert),
                    (WeaponTraining.Unarmed, ProficiencyLevel.Trained),
                    (WeaponTraining.Simple, ProficiencyLevel.Trained),
                    (WeaponTraining.Martial, ProficiencyLevel.Trained),
                    (ArmourWeight.Unarmoured, ProficiencyLevel.Trained),
                    (ArmourWeight.Light, ProficiencyLevel.Trained),
                    ])
  Thaumaturge = pfClass(health_per_level = 8,
                  initial_proficiencies = [
                    (SavingThrows.Fortitude, ProficiencyLevel.Expert),
                    (SavingThrows.Reflex, ProficiencyLevel.Trained),
                    (SavingThrows.Will, ProficiencyLevel.Expert),
                    (WeaponTraining.Unarmed, ProficiencyLevel.Trained),
                    (WeaponTraining.Simple, ProficiencyLevel.Trained),
                    (WeaponTraining.Martial, ProficiencyLevel.Trained),
                    (ArmourWeight.Unarmoured, ProficiencyLevel.Trained),
                    (ArmourWeight.Light, ProficiencyLevel.Trained),
                    (ArmourWeight.Medium, ProficiencyLevel.Trained),
                    ])
  Witch = pfClass(health_per_level = 6,
                  initial_proficiencies = [
                    (SavingThrows.Fortitude, ProficiencyLevel.Trained),
                    (SavingThrows.Reflex, ProficiencyLevel.Trained),
                    (SavingThrows.Will, ProficiencyLevel.Expert),
                    (WeaponTraining.Unarmed, ProficiencyLevel.Trained),
                    (WeaponTraining.Simple, ProficiencyLevel.Trained),
                    (ArmourWeight.Unarmoured, ProficiencyLevel.Trained),
                    ])
  Wizard = pfClass(health_per_level = 6,
                  initial_proficiencies = [
                    (SavingThrows.Fortitude, ProficiencyLevel.Trained),
                    (SavingThrows.Reflex, ProficiencyLevel.Trained),
                    (SavingThrows.Will, ProficiencyLevel.Expert),
                    (WeaponTraining.Unarmed, ProficiencyLevel.Trained),
                    (WeaponTraining.Simple, ProficiencyLevel.Trained),
                    (ArmourWeight.Unarmoured, ProficiencyLevel.Trained),
                    ])
