from peewee import *
from .position import Position
from .base_model import BaseModel

import uuid


class Office(BaseModel):
  id = UUIDField(primary_key=True, default=uuid.uuid4)
  sale_point_name = CharField(max_length=255)
  address = CharField(max_length=400)
  status = CharField(max_length=255)
  rko = BooleanField(null=True)
  type = CharField(max_length=255)
  sale_point_format = CharField(max_length=255)
  suo_availability = BooleanField(null=True)
  has_ramp = BooleanField(null=True)
  distance = IntegerField()
  kep = BooleanField(null=True)
  my_branch = BooleanField()
  position = ForeignKeyField(Position, related_name='position')
