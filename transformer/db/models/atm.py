from peewee import *
from .office import Position
from .base_model import BaseModel

import uuid


class Atm(BaseModel):
  id = UUIDField(primary_key=True, default=uuid.uuid4)
  address = CharField(max_length=255)
  all_day = BooleanField()
  position = ForeignKeyField(Position, related_name='position')
