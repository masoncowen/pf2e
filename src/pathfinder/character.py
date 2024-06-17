class Character(pydantic.BaseModel):
  name: str
  actions_remaining: int = 3
  max_actions: int = 3
