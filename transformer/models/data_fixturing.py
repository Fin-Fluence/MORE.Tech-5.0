from db.serilizer import Serilizer
from .base.base_parser import BaseParser


class DataFixturing(BaseParser):
  def __init__(self):
    super().__init__('transformer/data/results/formatted/', '')

  def perform(self):
    data = self.load_data('atms.json')
    for entry in data:
      print(Serilizer.create_atm(entry))

    data = self.load_data('offices.json')
    for entry in data:
      print(Serilizer.create_office(entry))

  def save(self):
    pass
