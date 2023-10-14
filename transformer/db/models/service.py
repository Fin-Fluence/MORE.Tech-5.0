from peewee import *
from .atm import Atm
from .base_model import BaseModel

import uuid


class Service(BaseModel):
  id = UUIDField(primary_key=True, default=uuid.uuid4)
  name = CharField(max_length=512)
  capability = BooleanField(null=True)
  activity = BooleanField(null=True)
  atm = ForeignKeyField(Atm, backref='atms')
