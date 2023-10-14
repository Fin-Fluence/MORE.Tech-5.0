from db.serilizer import Serilizer
from .base.base_parser import BaseParser


class DataFixturing(BaseParser):
  def __init__(self):
    super().__init__('data/results/formatted/', '')

  def perform(self):
    Serilizer.delete()
    print('Start loading ATMS')
    data = self.load_data('atms.json')
    for entry in data:
      Serilizer.create_atm(entry)
    print('Complete')

    print('Start loading OFFICESS')
    data = self.load_data('offices.json')
    for entry in data:
      Serilizer.create_office(entry)
    print('Complete')

  def save(self):
    pass
