# 每日金价
import re
import requests
from gold.store import GoldStore
from corn_monitor.db.log import Log
from corn_monitor.notice import pushplus

api = "http://www.dyhjw.com/hjtd"

headers = {
        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/110.0.1587.63",
    }
# 阈值
Threshold = 385

def run()->float:
    """运行"""
    log = Log()
    gold_store = GoldStore()
    # 尝试5次获取
    for i in range(5):
        log.add(f"this is the {i} times load")
        res = _req()
        if res:
            res = float(res[0])
            gold_store.save(res)
            if res <= Threshold:
                pushplus.send("黄金提醒", f"现在黄金价格为：{res} 小于阈值 {Threshold}")
            return res


def _req()->list:
    """获取数据"""
    res = requests.get(api, headers=headers)
    res.encoding = "gbk"
    return re.findall("浠婂紑锛�<font>([\d.]+)<", res.text)
