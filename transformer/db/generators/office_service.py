from ..models.office_service import OfficeService

import random

class OfficeServiceGenerator:
  def __init__(self, office, names=['Test1', 'Test2']):
    self.names = names
    self.bool_variants = [True, False]
    self.office = office

  def generate_all(self):
    for name in self.names:
      self.generate_one(name)

  def generate_one(self, name):
    os = OfficeService.create(
      name=name,
      capability=random.choice(self.bool_variants),
      activity=random.choice(self.bool_variants),
      current_ticket='AA' + format(str(random.randint(0, 999)), '*>3s'),
      last_ticket='AA' + format(str(random.randint(0, 999)), '*>3s'),
      office=self.office
    )

    return os
