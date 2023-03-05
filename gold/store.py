import datetime
from corn_monitor.db.db import DBStore

class GoldStore(DBStore):

    def save(self, price:float):
        self._insert("insert into gold (`price`, `date`) values (?,?)", price, datetime.date.today())

    def load(self, size=10)->list:
        if size == -1:
            return self._select_many("select * from gold")
        return self._select_many("select * from gold order by `date` asc limit ?", size)
