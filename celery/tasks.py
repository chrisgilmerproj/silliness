# /usr/local/bin/python

from celery import Celery

celery = Celery()
celery.config_from_object('celeryconfig')


@celery.task
def add(x, y):
    return x + y
