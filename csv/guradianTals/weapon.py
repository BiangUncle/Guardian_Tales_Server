from bs4 import BeautifulSoup
import requests
import csv


re = requests.get("https://wiki.biligame.com/gt/%E6%AD%A6%E5%99%A8")


bs = BeautifulSoup(re.content, "html.parser")

# re = bs.select("#CardSelectTr")
# items = bs.select("#CardSelectTr")
items = bs.select(".resp-tab-content")
#


for item in items:

    ch = item.contents[1].select('#CardSelectTr')[0]
    trs = ch.select('tr')

    attr_list = []
    attr = trs[0].text.split('\n')
    for att in attr:
        if att != '':
            attr_list.append(att)
    type_info = attr_list[0]
    f = open(type_info + ".csv", "w", encoding='utf8', newline='')
    writer = csv.writer(f)

    for tr in trs[1:]:
        strs = tr.text.split('\n')
        line = []

        for str in strs:
            if str != '':
                line.append(str)

        writer.writerow(line)

    f.close()

    # writer.writerow(curhero)


# f.close()