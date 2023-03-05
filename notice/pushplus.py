import requests
import json
from corn_monitor.db.log import Log
from corn_monitor.config import cfg

def send(title: str, content: str):
    log = Log()
    url = 'http://www.pushplus.plus/send'
    data = {
        "token": cfg['pushplus']['token'],
        "title": title,
        "content": content
    }
    try:
        res = requests.post(url, data=data)
        log.add(f"推送消息:{json.dumps(data)}")
        res.raise_for_status()
    except Exception as ret:
        log.add(f"pushplus推送:{json.dumps(data)}失败;错误原因：{ret}")
