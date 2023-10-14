from ..models.office_service import OfficeService

import random

class OfficeServiceGenerator:
  def __init__(self, office, names=['Test1', 'Test2']):
    # self.names = names
    self.names = ['BrokerService', 'ConversionFL', 'DebitCardIssuance', 'CreditCardIssuance', 'Mortgage', 'ExrowAccount', 'DepositsAndAccountsFL']
    self.bool_variants = [True, False]
    self.office = office

  def generate_all(self):
    for name in self.names:
      self.generate_one(name)

  def generate_one(self, name):
    cur_ticker = random.randint(0, 990)
    os = OfficeService.create(
      name=name,
      capability=random.choice(self.bool_variants),
      activity=random.choice(self.bool_variants),
      current_ticket='AA' + format(str(cur_ticker), '*>3s'),
      last_ticket='AA' + format(str(cur_ticker + random.randint(0, 4)), '*>3s'),
      office=self.office
    )

    return os
