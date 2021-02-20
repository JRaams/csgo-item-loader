import json
import sys


class Collection:
    def __init__(self, id, name, released, stattrak, souvenir, tag):
        self.id = id
        self.name = name
        self.released = released
        self.stattrak = stattrak
        self.souvenir = souvenir
        self.tag = tag


def parseCollections() -> dict[str, Collection]:
    # Load collections from resources
    collections = {}
    with open("resources/collections.json") as rfile:
        for obj in json.loads(rfile.read()):
            c = Collection(**obj)
            collections[c.name] = c

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


def main():
    collections = parseCollections()
    a = 5


if __name__ == "__main__":
    main()