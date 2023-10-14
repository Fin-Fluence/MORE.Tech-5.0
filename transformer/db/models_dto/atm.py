from peewee import *
from .office import Position
from .base_model import BaseModel

import uuid


class Atm(BaseModel):
  id = UUIDField(primary_key=True, default=uuid.uuid4)
  capability = BooleanField(null=True)
  activity = BooleanField(null=True)
  position = ForeignKeyField(Position, backref='position')
