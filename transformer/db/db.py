from peewee import *

DB = PostgresqlDatabase('postgres', user='postgres', password='test',
                        host='host.docker.internal', port=5431)
