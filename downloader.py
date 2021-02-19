import vdf
import json
import requests


class Resource:
    def __init__(self, name, url, parseKV=False):
        self.name = name
        self.url = url
        self.parseKV = parseKV

    def download(self):
        print("\tDownloading {0}...".format(self.name))
        r = requests.get(self.url, allow_redirects=False)
        with open("storage/{0}.txt".format(self.name), 'wb') as outfile:
            outfile.write(r.content)

    def convert(self):
        print("\tConverting {0}...".format(self.name))
        if self.parseKV:
            d = {}
            with open('storage/{0}.txt'.format(self.name)) as infile:
                for line in infile.readlines():
                    if line[0] == "#":
                        continue
                    kv = line.split('=')
                    d[kv[0]] = kv[1].rstrip()
        else:
            with open('storage/{0}.txt'.format(self.name)) as infile:
                d = vdf.parse(infile)

        with open("storage/{0}.json".format(self.name), 'w') as outfile:
            json.dump(d, outfile, indent=4)


toFetch = [
    Resource('items_game', 'https://raw.githubusercontent.com/SteamDatabase/GameTracking-CSGO/master/csgo/scripts/items/items_game.txt'),
    Resource('items_game_cdn', 'https://raw.githubusercontent.com/SteamDatabase/GameTracking-CSGO/master/csgo/scripts/items/items_game_cdn.txt', parseKV=True),
    Resource('csgo_english', 'https://raw.githubusercontent.com/SteamDatabase/GameTracking-CSGO/master/csgo/resource/csgo_english.txt')
]


def main():
    print("Downloading game files...")
    for resource in toFetch:
        resource.download()
        resource.convert()
    print("Done downloading game files!")


if __name__ == "__main__":
    main()
