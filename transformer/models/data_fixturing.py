from db.serilizer import Serilizer
from .base import BaseParser


class DataFixturing(BaseParser):
  def __init__(self):
    super().__init__('data/results/formatted', '')

  def perform(self):
    data = self.load_data('atm.json')
    for entry in data:
      print(Serilizer.create_atm(entry))

  def save(self):
    pass
