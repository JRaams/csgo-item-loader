import fetch from "node-fetch";
import fs from "fs";

export const ASSETS_TO_FETCH = [
  {
    fileName: "items_game.txt",
    url: "https://raw.githubusercontent.com/SteamDatabase/GameTracking-CSGO/master/csgo/scripts/items/items_game.txt",
    type: "vdf",
  },
  {
    fileName: "items_game_cdn.txt",
    url: "https://raw.githubusercontent.com/SteamDatabase/GameTracking-CSGO/master/csgo/scripts/items/items_game_cdn.txt",
    type: "rawkeyvalue",
  },
  {
    fileName: "csgo_english.txt",
    url: "https://raw.githubusercontent.com/SteamDatabase/GameTracking-CSGO/master/csgo/resource/csgo_english.txt",
    type: "vdf",
  },
];

export const ASSETS_FOLDER_PATH = "./assets";

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
    this.isVerbose && console.info("Downloader createAssetsFolder: start");
    try {
      await fs.promises.mkdir(ASSETS_FOLDER_PATH);
      this.isVerbose && console.info("Downloader createAssetsFolder: success");
    } catch (error) {
      this.isVerbose && console.info("Downloader createAssetsFolder: error");
    }
    this.isVerbose && console.info("Downloader createAssetsFolder: done");
  }

  async fetchFile({ fileName, url }) {
    this.isVerbose &&
      console.info("Downloader fetchFile: start", fileName, url);
    const response = await fetch(url).catch((e) => {
      console.info(e);
    });
    const data = await response.text();
    await fs.promises.writeFile(`${ASSETS_FOLDER_PATH}/${fileName}`, data);
    this.isVerbose && console.info("Downloader fetchFile: done");
  }
}
