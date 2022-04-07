import logging
import os
from logging import handlers

from config.conf import LOG_DIR, LOG_LEVEL

if not os.path.exists(LOG_DIR):
    os.makedirs(LOG_DIR)


def get_logger(log_name: str, when='D', backup_count=5):
    formatter = logging.Formatter(
        "[%(asctime)s %(filename)s %(funcName)s line:%(lineno)d %(levelname)s]: %(message)s")

    logger = logging.getLogger(log_name)
    logger.setLevel(LOG_LEVEL)

    sh = logging.StreamHandler()
    sh.setFormatter(formatter)
    logger.addHandler(sh)

    filename = os.path.join(LOG_DIR, f'{log_name}.log')
    th = handlers.TimedRotatingFileHandler(
        filename=filename,
        when=when,
        backupCount=backup_count,
        encoding='utf-8'
    )
    th.setFormatter(formatter)
    logger.addHandler(th)
    return logger
