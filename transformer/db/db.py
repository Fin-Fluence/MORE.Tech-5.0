from peewee import *

DB = PostgresqlDatabase('my_app', user='postgres', password='test',
                           host='localhost', port=5432)
