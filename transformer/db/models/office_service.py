from peewee import *
from .office import Office
from .base_model import BaseModel

import uuid


class OfficeService(BaseModel):
  id = UUIDField(primary_key=True, default=uuid.uuid4)
  name = CharField(max_length=512)
  capability = BooleanField(null=True)
  activity = BooleanField(null=True)
  current_ticket = CharField(max_length=8)
  last_ticket = CharField(max_length=8)
  office = ForeignKeyField(Office, backref='office_services')
