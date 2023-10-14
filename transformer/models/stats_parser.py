from .base.base_parser import BaseParser


# Класс для сбора статистики полей данных json
class StatsParser(BaseParser):
  def __init__(self):
    super().__init__('transformer/data/basic_data/', 'transformer/data/results/stats/')

  def get_field_uniqs(self, data, field):
    uniq_counter = {'NONE': 0}
    for el in data:
      try:
        value = el[field]
        count = uniq_counter.get(value, 0) + 1
        uniq_counter[value] = count
      except:
        uniq_counter['NONE'] += 1
    
    return uniq_counter

  def get_all_fields(self, data):
    fields = set()
    for el in data:
      for field in list(el.keys()):
        if type(field) is str:
          fields.add(field)

    return fields

  def get_uniq(self, data, fields=[], exclude=[]):
    keys = data[0].keys()
    print(keys)
    results = {}
    if fields:
      for field in fields:
        if field in keys:
          results[field] = self.get_field_uniqs(data, field)
    else:
      for key in self.get_all_fields(data).difference(set(exclude)):
        results[key] = self.get_field_uniqs(data, key)

    return results

  def check_atms(self):
    FIELDS_TO_CHECK = ['allDay']
    # FIELDS_TO_CHECK = []
    FIELDS_TO_EXCLUDE = []
    results = self.get_uniq(self.load_data('atms.json')['atms'], FIELDS_TO_CHECK, FIELDS_TO_EXCLUDE)

    self.save_results('atms.json', results)

  def check_offices(self):
    FIELDS_TO_CHECK = []
    FIELDS_TO_EXCLUDE = ['distance', 'address', 'latitude', 'longitude', 'metroStation', 'salePointName', 'officeType', 'salePointFormat']
    # FIELDS_TO_EXCLUDE = []
    results = self.get_uniq(self.load_data('offices.json'), FIELDS_TO_CHECK, FIELDS_TO_EXCLUDE)

    self.save_results('offices.json', results)

  def perform(self, modules=['atms', 'offices']):
    if 'atms' in modules:
      self.check_atms()
    if 'offices' in modules:
      self.check_offices()
