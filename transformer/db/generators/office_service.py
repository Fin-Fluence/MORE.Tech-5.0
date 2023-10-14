from ..models.office_service import OfficeService

import random

class OfficeServiceGenerator:
  def __init__(self):
    self.model = OfficeService
    self.names = ['Test1', 'Test2']
    self.bool_variants = [True, False]

  def generate_one(self, office):
    return 
    os = OfficeService.create(
      name=random.choice(self.names),
      capability=random.choice(self.bool_variants),
      activity=random.choice(self.bool_variants),
      current_ticker=random.choice(self.bool_variants),
      last_ticker=random.choice(self.bool_variants),
      office=office
    )

    return os
