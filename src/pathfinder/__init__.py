from typing import *

from .character import PlayerCharacter
from .npc import NonPlayerCharacter
from .creature import Creature

type Character = Union[PlayerCharacter, NonPlayerCharacter]
type Party = list[Character]
