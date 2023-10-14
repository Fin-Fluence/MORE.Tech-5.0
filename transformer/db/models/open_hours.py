from peewee import *
from .office import Office
from .base_model import BaseModel

import uuid


class OpenHours(BaseModel):
  id = UUIDField(primary_key=True, default=uuid.uuid4)
  day = CharField(max_length=50)
  hours = CharField(max_length=50)
  is_individual = BooleanField(max_length=50)
  office = ForeignKeyField(Office, backref='open_hours')
