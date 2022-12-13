import csv

from bs4 import BeautifulSoup
import requests

'''
const (
	WEAPON_TYPE_DANSHOUJIAN    = 1
	WEAPON_TYPE_SHUANGSHOUJIAN = 2
	WEAPON_TYPE_BUQIANG        = 3
	WEAPON_TYPE_GONG           = 4
	WEAPON_TYPE_LANZI          = 5
	WEAPON_TYPE_QUANZHANG      = 6
	WEAPON_TYPE_SHOUTAO        = 7
	WEAPON_TYPE_DUNPAI         = 8
	WEAPON_TYPE_ZHUA           = 9
	WEAPON_TYPE_SHIPIN         = 10
	WEAPON_TYPE_QIANGHUA       = 11
	
	JIDAO = 1
	ZHONGSHANG = 2
	JIFEI = 3
	QUANBU = 4
	WU = 5
	
	TYPE_BASIC = 1
	TYPE_LIGHT = 2
	TYPE_DARK  = 3
	TYPE_WATER = 4
	TYPE_FIRE  = 5
	TYPE_EARTH = 6

	HERO_TYPE_WARRIOR = 1
	HERO_TYPE_RANGED  = 2
	HERO_TYPE_TANKER  = 3
	HERO_TYPE_SUPPORT = 4
)'''

weapontypemap = {
    "单手剑":1,
    "双手剑":2,
    "步枪":3,
    "突击步枪":3,
    "弓":4,
    "篮子":5,
    "杖":6,
    "权杖":6,
    "魔法杖":6,
    "手套":7,
    "盾":8,
    "盾牌":8,
    "爪":9,
}

skillmap = {
    "击倒":1,
    "重伤":2,
    "击飞":3,
    "全部":4,
    "/":5,
    "无":5,
}

typemap = {
    "虚":1,
    "光":2,
    "暗":3,
    "水":4,
    "火":5,
    "土":6,
}

herotypemap = {
    "战士":1,
    "射手":2,
    "坦克":3,
    "辅助":4,
}

def get_hero_data():

    re = requests.get("https://wiki.biligame.com/gt/%E8%8B%B1%E9%9B%84%E7%AD%9B%E9%80%89%E8%A1%A8")


    bs = BeautifulSoup(re.content, "html.parser")

    # re = bs.select("#CardSelectTr")
    re = bs.select(".divsort")

    heroList = []

    import csv

    f = open("hero.csv", "w", encoding='utf8', newline='')
    writer = csv.writer(f)

    for heroInfo in re:
        curhero = []
        for attr, v in heroInfo.attrs.items():
            curhero.append(v)

        trs = heroInfo.select('td')

        name_weapon = []
        name_tr = trs[0].select('a')
        if len(name_tr) > 0:
            name_weapon.append(name_tr[0].attrs['title'])
        else:
            name_weapon.append("/")

        weapon_tr = trs[2].select('a')
        if len(weapon_tr) > 0:
            name_weapon.append(weapon_tr[0].attrs['title'])
        else:
            name_weapon.append("/")

        curhero = name_weapon + curhero[1:]

        writer.writerow(curhero)


    f.close()

def wash_hero_data():

    f1 = open("hero.csv", "r", encoding="utf8", newline='')
    reader = csv.reader(f1)

    f2 = open("hero_new.csv", "w", encoding="utf8", newline='')
    writer = csv.writer(f2)

    for line in reader:
        hero = []
        hero.append(line[0])
        hero.append(line[1])

        weapons = line[2].split(',')
        weapon_idx = []
        for weapon in weapons:
            weapon_idx.append(weapontypemap[weapon])
        hero.append(weapon_idx)

        hero.append(int(line[3]))
        hero.append(typemap[line[4]])
        hero.append(skillmap[line[5]])
        hero.append(skillmap[line[6]])
        hero.append(herotypemap[line[7]])
        writer.writerow(hero)

    f1.close()
    f2.close()

def create_hero_idx():

    start = 2000000
    idx = 1

    f1 = open("hero.csv", "r", encoding="utf8", newline='')
    reader = csv.reader(f1)

    f2 = open("hero_idx.csv", "w", encoding="utf8", newline='')
    writer = csv.writer(f2)

    for line in reader:
        hero = [start + idx]
        idx += 1

        hero.append(line[0])

        weapons = line[2].split(',')
        weapon_idx = []
        for weapon in weapons:
            weapon_idx.append(weapontypemap[weapon])
        hero.append(weapon_idx)

        hero.append(int(line[3]))
        hero.append(typemap[line[4]])
        hero.append(skillmap[line[5]])
        hero.append(skillmap[line[6]])
        hero.append(herotypemap[line[7]])
        writer.writerow(hero)

    f1.close()
    f2.close()


create_hero_idx()