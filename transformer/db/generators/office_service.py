from ..models.office_service import OfficeService

import random

class OfficeServiceGenerator:
  def __init__(self):
    self.model = OfficeService
    self.names = ['Test1', 'Test2']

  def generate_one(self, office):
    pass
    # os = OfficeService.create(
    #   name=
    # )