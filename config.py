import json


def _get_config()->dict:
    """获取配置"""
    with open("config.json", "r", encoding="utf-8") as f:
        return json.load(f)


cfg = _get_config()
