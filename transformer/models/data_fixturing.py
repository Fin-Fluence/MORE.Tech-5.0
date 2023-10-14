from peewee import *


class DataFixturing:
  def __init__(self):
    self.db = PostgresqlDatabase('my_app', user='postgres', password='test',
                           host='localhost', port=5432)
  
  def save(self):
    pass
