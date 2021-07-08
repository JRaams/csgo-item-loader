import json
import sys


def parseCollections() -> dict:
    # Load collections from resources
    collections = {}
    with open("resources/collections.json") as rfile:
        for obj in json.loads(rfile.read()):
            collections[obj["name"]] = obj

    # Check if all actual collections have an entry in resources/collections
    # If not, it needs to be added manually
    with open("storage/csgo_english.json") as cefile:
        j = json.loads(cefile.read())
        for key in j["lang"]["Tokens"]:
            if "CSGO_set" not in key:
                continue
            if "_desc" in key or "_short" in key:
                continue
            if key not in collections:
                sys.exit(
                    "Error parsing collections: {0} was not found in resources/collections.json".format(
                        key
                    )
                )

    return collections


weaponTypes = {"secondary": "Pistol", "smg": "SMG", "heavy": "Heavy", "rifle": "Rifle"}


def parseWeapons() -> dict:
    weapons = {}

    with open("storage/items_game.json") as wfile:
        j = json.loads(wfile.read())
        items = j["items_game"]["items"]
        for key in items:
            item = items[key]
            if "prefab" not in item:
                continue
            if "weapon_" not in item["prefab"]:
                continue
            if "baseitem" not in item or item["baseitem"] != "1":
                continue
            if "weapon_taser" in item["name"]:
                continue

            prefab = j["items_game"]["prefabs"][item["prefab"]]
            item["prefab"] = prefab
            typeKey = "".join(c for c in item["item_sub_position"] if not c.isdigit())
            item["type"] = weaponTypes[typeKey]
            weapons[key] = item

    with open("resources/weapons.json", "w") as f:
        json.dump(weapons, f, indent=4)

    return weapons


def parseLootLists() -> dict:
    lootLists = {}

    with open("storage/items_game.json") as wfile:
        j = json.loads(wfile.read())
        items_game = j["items_game"]
        rarities = ["_" + key for key in items_game["rarities"]]
        gameLootLists = items_game["client_loot_lists"]
        for lootList in gameLootLists:
            rarity, collection = getRarity(rarities, lootList)
            if rarity is None or lootList == "set_op10_ancient":
                continue

            if collection not in lootLists:
                lootLists[collection] = {}

            loot = gameLootLists[lootList]
            lootLists[collection][rarity] = []

            for item in loot:
                # Weapon/skin
                if "[" in item:
                    paintkit, weapon = item.replace("[", "").split("]")
                    lootLists[collection][rarity].append(
                        {"weapon": weapon, "paintkit": paintkit}
                    )
                # Pin
                elif "Pin" in item:
                    pin = item.split(" - ")[1]
                    lootLists[collection][rarity].append(pin)
                # Agent
                elif "customplayer" in item:
                    lootLists[collection][rarity].append(item)
                # Sprays
                elif "spray" in item:
                    spray = item.replace("[", "").split("]")[0]
                    a = 5
                else:
                    print(item)

    with open("resources/lootlists.json", "w") as outfile:
        json.dump(lootLists, outfile, indent=4)

    return lootLists


def getRarity(rarities, lootList) -> tuple:
    for rarity in rarities:
        if rarity in lootList:
            strippedLootlist = lootList.replace(rarity, "")
            return rarity.replace("_", ""), strippedLootlist
    return None, None


def main():
    # collections = parseCollections()
    # weapons = parseWeapons()
    lootLists = parseLootLists()


if __name__ == "__main__":
    main()