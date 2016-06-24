from scrapy.selector import HtmlXPathSelector
from scrapy.spider import BaseSpider
from tutorial.items import DmozItem


class DmozSpider(BaseSpider):
    name = "dmoz"
    allowed_domains = ["dmoz.org"]
    start_urls = [
        "http://www.dmoz.org/Computers/Programming/Languages/Python/Books/",
        "http://www.dmoz.org/Computers/Programming/Languages/Python/Resources/"
    ]

    def parse(self, response):
        hxs = HtmlXPathSelector(response)
        sites = hxs.select('//ul/li')
        items = []
        for site in sites:
            item = DmozItem()
            item['title'] = filter(bool, [r.strip() for r in site.select('a/text()').extract()])
            item['link'] = site.select('a/@href').extract()
            item['desc'] = filter(bool, [r.strip() for r in site.select('text()').extract()])
            items.append(item)
        return items
