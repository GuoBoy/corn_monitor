import time

from corn_monitor.db.db import DBStore

class Log(DBStore):
    """日志储存"""

    def add(self, msg:str):
        """添加日志"""
        if msg:
            self._insert("insert into log (`message`, `created_at`) values (?,?)", msg, time.asctime())

    def load(self, size=10)->list:
        if size == -1:
            return self._select_many("select * from log order by `created_at` asc")
        return self._select_many("select * from log order by `created_at` asc limit ?", size)
