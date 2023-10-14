import json
import io

# Базовый класс для работы с данными json-формата
class BaseParser:
  def __init__(self, data_path, save_path):
    self.data_path = data_path
    self.save_path = save_path

  def load_data(self, file_name):
    with io.open(self.data_path + file_name, encoding='utf-8') as json_file:
      return json.load(json_file)

  def save_results(self, file_name, data):
    with io.open(self.save_path + file_name, 'w', encoding='utf-8') as file:
      json.dump(data, file, ensure_ascii=False)
