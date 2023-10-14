from peewee import *

DB = PostgresqlDatabase('postgres', user='postgres', password='test',
                           host='localhost', port=5431)
