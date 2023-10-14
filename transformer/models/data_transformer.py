from .base.base_parser import BaseParser


# Класс для преобразования данных json в более удобный формат работы
class DataTransformer(BaseParser):
  def __init__(self):
    super().__init__('data/basic_data/', 'data/results/formatted/')

  def change_fields(self, data, fields_change):
    if type(data) is list:
      for el in data:
        for field in fields_change:
          el[fields_change[field]] = el[field]
          del el[field]
    elif type(data) is dict:
      for el_key in data.keys():
        el = data[el_key]
        for field in fields_change:
          el[fields_change[field]] = el[field]
          del el[field]
    
    return data

  def change_bools(self, data, fields):
    keys_to_change = list(fields.keys())
    if type(data) is list:
      for el in data:
        for key in keys_to_change:
          val = el[key]
          if val == fields[key][0]:
            el[key] = True
          elif val == fields[key][1]:
            el[key] = False
          else:
            el[key] = 'null'
    elif type(data) is dict:
      for el_key in data.keys():
        el = data[el_key]
        for key in keys_to_change:
          val = el[key]
          if val == fields[key][0]:
            el[key] = True
          elif val == fields[key][1]:
            el[key] = False
          else:
            el[key] = 'null'

    return data

  def change_dict(self, data, fields):
    keys_to_change = list(fields.keys())

    for el in data:
      for key in keys_to_change:
        val = el[key]
        if val == fields[key][0]:
          el[key] = True
        elif val == fields[key][1]:
          el[key] = False
        else:
          el[key] = 'null'



  def offices(self):
    FIELDS = {'rko': ['есть РКО', 'нет РКО'], 'hasRamp': ['Y', 'N'], 'suoAvailability': ['Y', 'N']}

    results = self.change_bools(self.load_data('offices.json'), FIELDS)

    FIELDS_TO_CHANGE = {'officeType': 'type'}
    results = self.change_fields(results, FIELDS_TO_CHANGE)

    self.save_results('offices.json', results)

  def atms(self):
    data = self.load_data('atms.json')['atms']

    for atm in data: 
      services = atm['services']

      FIELDS = {'serviceCapability': ['SUPPORTED', 'UNSUPPORTED'], 'serviceActivity': ['AVAILABLE', 'UNAVAILABLE']}

      results = self.change_bools(services, FIELDS)

      FIELDS_TO_CHANGE = {'serviceCapability': 'capability', 'serviceActivity': 'activity'}
      results = self.change_fields(results, FIELDS_TO_CHANGE)

      atm['services'] = results

    self.save_results('atms.json', data)

  def perform(self):
    self.offices()
    self.atms()
