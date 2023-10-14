from peewee import *
from .base_model import BaseModel

import uuid


class Position(BaseModel):
  id = UUIDField(primary_key=True, default=uuid.uuid4)
  latitude = DecimalField(decimal_places=6)
  longitude = DecimalField(decimal_places=6)
  metro_station = CharField(max_length=511, null=True)
