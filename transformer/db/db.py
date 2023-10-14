from peewee import *
import os

uri = os.environ['POSTGRES_URI']

_, uri = uri.split('://')
uri, _ = uri.split('?')

fp, sp = uri.split('@')
username, password = fp.split(':')
host, tp = sp.split(':')
port, db = tp.split('/')

DB = PostgresqlDatabase(db, user=username, password=password,
                        host=host, port=port)
