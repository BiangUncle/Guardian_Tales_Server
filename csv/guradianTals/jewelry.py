from bs4 import BeautifulSoup
import requests
import csv


re = requests.get("https://wiki.biligame.com/gt/%E6%AD%A6%E5%99%A8")

bs = BeautifulSoup(re.content, "html.parser")

items = bs.select(".resp-tab-content")

item = items[2].contents[1]

trs = item.select('.divsort')

f = open("首饰.csv", "w", encoding='utf8', newline='')
writer = csv.writer(f)

for tr in trs:
    strs = tr.text.split('\n')
    line = []

    for str in strs:
        if str != '':
            line.append(str)

    writer.writerow(line)

f.close()
