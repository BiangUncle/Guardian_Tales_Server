from bs4 import BeautifulSoup
import requests
import csv


re = requests.get("https://wiki.biligame.com/gt/%E6%AD%A6%E5%99%A8")
bs = BeautifulSoup(re.content, "html.parser")
items = bs.select(".resp-tab-content")
item = items[4].contents[1]


trs = item.select('.divsort')

f = open("时装.csv", "w", encoding='utf8', newline='')
writer = csv.writer(f)

for tr in trs:

    line = []
    ctns = tr.contents

    line.append(ctns[1].text[:-1])
    line.append(ctns[3].text[:-1])

    ifhavea = ctns[5].find_all("a")
    if len(ifhavea) == 0:
        line.append(ctns[5].text[:-1])
    else:
        costitem = ifhavea[0].contents[0].attrs['alt'][:-4]
        xx = ctns[5].text.split(' ')
        line.append(xx[0] + costitem + xx[1])

    if ctns[7].text == "\n":
        line.append("/")
    else:
        line.append(ctns[7].text[:-1])
    writer.writerow(line)

f.close()
