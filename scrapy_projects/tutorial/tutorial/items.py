# Define here the models for your scraped items
#
# See documentation in:
# http://doc.scrapy.org/topics/items.html

from scrapy.item import Field
from scrapy.item import Item


class DmozItem(Item):
    title = Field()
    link = Field()
    desc = Field()
