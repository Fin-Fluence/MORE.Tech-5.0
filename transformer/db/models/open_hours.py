from peewee import *
from .office import Office
from .base_model import BaseModel

import uuid


class OpenHours(BaseModel):
  id = UUIDField(primary_key=True, default=uuid.uuid4)
  day = CharField(max_length=50)
  hours = CharField(max_length=50)
  is_individual = BooleanField(default=False)
  office = ForeignKeyField(Office, backref='open_hours', on_delete='CASCADE')

  class Meta:
    table_name = 'open_hours'
