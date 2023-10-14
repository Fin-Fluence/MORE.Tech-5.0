from .models.position import Position
from .models.atm import Atm
from .models.service import Service
from .models.open_hours import OpenHours
from .models.office import Office
from .models.office_service import OfficeService

from .generators.office_service import OfficeServiceGenerator

class Serilizer:
  def __init__(self):
    pass
  
  @staticmethod
  def delete():
    Atm.delete().execute()
    Office.delete().execute()
    Service.delete().execute()
    OpenHours.delete().execute()
    Position.delete().execute()
    OfficeService.delete().execute()

  @staticmethod
  def create_atm(json):
    position = Position.create(
      latitude=json.get('latitude', 0),
      longitude=json.get('longitude', 0),
      metro_station=json.get('metroStation', None)
    )

    atm = Atm.create(
      address=json['address'],
      position=position,
      all_day=json.get('allDay', 'false'),
    )

    for ser_name in json['services'].keys():
      # puts ser_name 
      ser_json = json['services'][ser_name]
      service = Service.create(
        name=ser_name,
        capability=ser_json.get('capability', None),
        activity=ser_json.get('activity', None),
        atm=atm
      )
    
    return atm

  @staticmethod
  def create_office(json):
    position = Position.create(
      latitude=json.get('latitude', 0),
      longitude=json.get('longitude', 0),
      metro_station=json.get('metroStation', None)
    )

    office = Office.create(
      sale_point_name=json.get('salePointName', 'Нет данных'),
      address=json.get('address', 'Нет данных'),
      status=json.get('status', 'закрытая'),
      rko=json.get('rko', ''),
      sale_point_format=json.get('salePointFormat', 'Нет данных'),
      suo_availability=json.get('suoAvailability', False),
      has_ramp=json.get('suoAvailability', False),
      distance=json.get('distance', -1),
      kep=json.get('kep', False),
      my_branch=json.get('myBranch', False),
      type=json.get('type', 'Нет данных'),
      position=position
    )

    for hours in json.get("openHours", []):
      hours_data = hours.get('hours', 'Нет данных')
      if not hours_data:
        hours_data = 'Нет данных'
      open_hours = OpenHours.create(
        day=hours.get('days', 'Нет данных'),
        hours=hours_data,
        office=office
      )

    for hours in json.get("openHoursIndividual", []):
      hours_data = hours.get('hours', 'Нет данных')
      if not hours_data:
        hours_data = 'Нет данных'
      open_hours_ind = OpenHours.create(
        day=hours.get('days', 'Нет данных'),
        hours=hours_data,
        is_individual=True,
        office=office
      )

    OfficeServiceGenerator(office).generate_all()

    return office
