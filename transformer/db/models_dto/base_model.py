from peewee import *
from ..db import DB


class BaseModel(Model):

  class Meta:
      database = DB