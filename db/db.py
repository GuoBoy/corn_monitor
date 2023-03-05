import sqlite3
from typing import Any
import pathlib

db_path = "database.db"

class DBStore:
    def __init__(self):
        self.__con = sqlite3.connect(db_path)

    def __init_table(self):
        """初始化表结果"""
        for sql in pathlib.Path().glob("db/sqls/*.sql"):
            with open(sql, "r", encoding="utf-8") as f:
                self.__con.execute(f.read())
                self.__con.commit()

    def _insert(self, sql:str, *args):
        self.__con.execute(sql, args)
        self.__con.commit()

    def save(self, *args):
        pass

    def load(self, *args):
        pass

    def _select_many(self, sql:str, *args)->list:
        return self.__con.execute(sql, args).fetchall()

    def _select_one(self, sql:str, *args)->Any:
        return self.__con.execute(sql, args).fetchone()

    def __exit__(self, exc_type, exc_val, exc_tb):
        self.__con.close()
        super().__exit__(exc_type, exc_val, exc_tb)