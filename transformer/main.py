from models.stats_parser import StatsParser
from models.data_transformer import DataTransformer

if __name__ == '__main__':
  StatsParser().perform()
  DataTransformer().perform()
