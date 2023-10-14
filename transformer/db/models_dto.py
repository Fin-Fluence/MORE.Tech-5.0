# Загрузка моделей
# from .models_dto.office import Office
# from .models_dto.position import Position
# from .models_dto.office_service import OfficeService
# from .models_dto.open_hours import OpenHours
from .models.atm import Atm
from .models.atm import Position

# from .db import DB

from peewee import *

def fixtures():
  pos = Position.create(latitude=12.123456, longitude=12.123456)
  pos.save()
  atm = Atm.create(address='Test addrr', all_day=True, position=pos)
  atm.save()
