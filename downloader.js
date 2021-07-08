import fetch from "node-fetch";
import fs from "fs";

const ASSETS_TO_FETCH = [
  {
    fileName: "items_game",
    url: "https://raw.githubusercontent.com/SteamDatabase/GameTracking-CSGO/master/csgo/scripts/items/items_game.txt",
  },
  {
    fileName: "items_game_cdn",
    url: "https://raw.githubusercontent.com/SteamDatabase/GameTracking-CSGO/master/csgo/scripts/items/items_game_cdn.txt",
  },
  {
    fileName: "csgo_english",
    url: "https://raw.githubusercontent.com/SteamDatabase/GameTracking-CSGO/master/csgo/resource/csgo_english.txt",
  },
];

const ASSETS_FOLDER_PATH = "./assets";

export class Downloader {
  constructor(isVerbose = false) {
    this.isVerbose = isVerbose;
  }

  async fetchVDLFiles() {
    await this.createAssetsFolder();

    for (let asset of ASSETS_TO_FETCH) {
      await this.fetchFile(asset);
    }
  }

  async createAssetsFolder() {
    this.isVerbose && console.info("Downloader createAssetsFolder start");
    try {
      await fs.promises.mkdir(ASSETS_FOLDER_PATH);
      this.isVerbose && console.info("Downloader createAssetsFolder success");
    } catch (error) {
      this.isVerbose && console.info("Downloader createAssetsFolder error");
    }
    this.isVerbose && console.info("Downloader createAssetsFolder done");
  }

  async fetchFile({ fileName, url }) {
    this.isVerbose && console.info("Downloader fetchFile start", fileName, url);
    const response = await fetch(url).catch((e) => {
      console.info(e);
    });
    const data = await response.text();
    await fs.promises.writeFile(`${ASSETS_FOLDER_PATH}/${fileName}.txt`, data);
    this.isVerbose && console.info("Downloader fetchFile done");
  }
}
