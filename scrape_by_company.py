import urllib.request as request
import re
import time
from pprint import pprint


def scrape(url="http://leetcode.liangjiateng.cn/leetcode/Amazon/algorithm"):
    r = request.urlopen(url)
    raw = r.read().decode("utf")
    current = []
    for line in raw.split('\n'):
        line = line.strip()
        if line.startswith('href="/leetcode/'):
            current.append(line)
        elif 'label ' in line:
            current.append(line)
        elif line.startswith('style="width:'):
            current.append(line)
        if len(current) == 4:
            yield current
            current = []


def scrape_all(company='Amazon', lastPage=12):
    # type: (str) -> list[list[str]]
    data = []
    urls = [f"http://leetcode.liangjiateng.cn/leetcode/{company}/algorithm?page={i}" for i in range(1, lastPage+1)]
    for url in urls:
        for challenge in scrape(url):
            data.append(challenge)
        time.sleep(1)
    return data


class Processor:

    def __init__(self, data):
        res = []
        for title, vip, level, popularity in data:
            title = self.process_url(title)
            vip = self.process_data(vip)
            vip = self.translate_vip(vip)
            level = self.process_data(level)
            popularity = self.process_num(popularity)
            res.append([title, vip, level, round(float(popularity), 2)])
        self.data = res

    @staticmethod
    def process_url(s):
        left = s.index('">') + 2
        right = s.index('</a')
        return s[left:right]

    @staticmethod
    def translate_vip(s):
        return "free" if s == "Normal" else "locked"

    @staticmethod
    def process_num(s):
        return re.sub("[^0-9.]", "", s)

    @staticmethod
    def process_data(s):
        right = s.index("</span>")
        i = right - 1
        while i >= 0:
            if s[i] == '>':
                return s[i + 1:right]
            i -= 1
        return ''


if __name__ == '__main__':
    raw = scrape_all()
    p = Processor(raw)
    pprint(p.data)
