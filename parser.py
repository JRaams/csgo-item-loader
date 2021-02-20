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


def main():
    collections = parseCollections()
    weapons = parseWeapons()
    a = 5


if __name__ == "__main__":
    main()