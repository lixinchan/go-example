import logging
import os

from dynaconf import settings as dynaconf_settings

PROJECT_NAME = 'learning-quant'

BASE_DIR = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))

ENV_LOCAL = 'local'
ENV_DEV = 'dev'
ENV_PROD = 'prod'

ENV = dynaconf_settings.current_env

if dynaconf_settings.current_env == ENV_LOCAL:
    LOG_LEVEL = logging.DEBUG
    LOG_SPIDER_LEVEL = logging.DEBUG
elif dynaconf_settings.current_env == ENV_DEV:
    LOG_LEVEL = logging.INFO
    LOG_SPIDER_LEVEL = logging.INFO
else:
    LOG_LEVEL = logging.INFO
    LOG_SPIDER_LEVEL = logging.ERROR

LOG_DIR = os.path.join(os.path.join(BASE_DIR, 'logs'), 'learning-quant')


