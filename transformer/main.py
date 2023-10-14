from models.stats_parser import StatsParser
from models.data_transformer import DataTransformer
from models.data_fixturing import DataFixturing
from db.models_dto import fixtures

if __name__ == '__main__':
  StatsParser().perform()
  DataTransformer().perform()
  
  DataFixturing().perform()

