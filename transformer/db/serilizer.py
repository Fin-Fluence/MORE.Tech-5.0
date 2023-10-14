from .models.position import Position
from .models.atm import Atm
from .models.service import Service
from .models.open_hours import OpenHours
from .models.office import Office
from .models.office_service import OfficeService

class Serilizer:
  def __init__(self):
    pass
  
  @staticmethod
  def create_atm(json):
    position = Position.create(
      latitude=json.get('latitude', 0),
      longitude=json.get('longitude', 0),
      metro_station=json.get('metroStation', None)
    )
    position.save()

    atm = Atm.create(
      address=json['address'],
      position=position,
      all_day=json.get('allDay', 'false'),
    )
    atm.save()

    for ser_name in json['services'].keys():
      ser_json = json['ser_name']
      service = Service.create(
        name=ser_name,
        capability=ser_json.get('capability', None),
        activity=ser_json.get('activity', None),
        atm=atm
      )
      service.save()
    
    return atm

  @staticmethod
  def create_office(json):
    position = Position.create(
      latitude=json.get('latitude', 0),
      longitude=json.get('longitude', 0),
      metro_station=json.get('metroStation', None)
    )
    position.save()

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
    )

    office.save()

    for hours in json.get("openHours", []):
      open_hours = OpenHours.create(
        days=hours.get('days', ''),
        hours=hours.get('hours', ''),
        office=office
      )
      open_hours.save

    for hours in json.get("openHoursIndividual", []):
      open_hours_ind = OpenHours.create(
        days=hours.get('days', 'Нет данных'),
        hours=hours.get('hours', 'Нет данных'),
        is_individual=True,
        office=office
      )
      open_hours_ind.save

    return office
