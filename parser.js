import fs from "fs";
import vdf from "simple-vdf";
import { ASSETS_FOLDER_PATH, ASSETS_TO_FETCH } from "./downloader.js";

export class Parser {
  constructor(isVerbose = false) {
    this.isVerbose = isVerbose;
  }

  async parseAssets() {
    this.isVerbose && console.info("Parser parseAssets: start");

    for (const asset of ASSETS_TO_FETCH) {
      const { fileName, type } = asset;
      this.isVerbose &&
        console.info(`Parser parseAssets: parse ${fileName} (${type})`);
      const txtPath = `${ASSETS_FOLDER_PATH}/${fileName}`;
      const fileData = await fs.promises.readFile(txtPath, "utf-8");

      const parsedData = await this.parseFileData(type, fileData);
      const jsonPath = `${ASSETS_FOLDER_PATH}/${fileName.replace(
        "txt",
        "json"
      )}`;
      this.isVerbose &&
        console.info(`Parser parseAssets: saving to ${jsonPath}`);
      await fs.promises.writeFile(jsonPath, parsedData);
    }

    this.isVerbose && console.info("Parser parseAssets: end");
  }

  async parseFileData(type, data) {
    switch (type) {
      case "vdf": {
        return this.parseVdf(data);
      }
      case "rawkeyvalue": {
        return this.parseRawKeyValue(data);
      }
      default: {
        console.info(
          `Parser parseFileData: unexpected type '${type}', exiting.`
        );
        process.exit(1);
      }
    }
  }

  async parseVdf(data) {
    const vdfdata = vdf.parse(data);
    const stringified = JSON.stringify(vdfdata, null, 2);
    return stringified;
  }

  async parseRawKeyValue(data) {
    const result = {};

    const lines = data.split("\n");
    for (const line of lines) {
      if (line[0] === "#" || line.trim().length === 0) continue;

      const [key, value] = line.split("=");
      result[key] = value;
    }

    return JSON.stringify(result, null, 2);
  }

  // 1. Read items_game.json
  // 1.1 Read collections.meta.json
  // 2. Get 'item_sets'
  // 3. For each 'item_set'
  //      check 'resources/collections.meta.json' for released/stattrak/souvenir data since it isnt available in game files
  // 4. Export to assets/collections.json
  async extractCollections() {
    this.isVerbose && console.info("Parser extractCollections: start");

    // 1 Read items_game.json
    this.isVerbose &&
      console.info("Parser extractCollections: loading items_game");
    const items_game_txt = await fs.promises.readFile(
      "./assets/items_game.json",
      "utf-8"
    );
    const items_game = JSON.parse(items_game_txt);

    // 1.1 Read collections.meta.json
    this.isVerbose &&
      console.info("Parser extractCollections: loading collections.meta");
    const collections_meta_txt = await fs.promises.readFile(
      "./resources/collections.meta.json",
      "utf-8"
    );
    const collections_meta = JSON.parse(collections_meta_txt);

    // 2 Get item_sets
    const item_sets = items_game.items_game.item_sets;

    // 3. For each 'item_set'
    //      check 'resources/collections.meta.json' for released/stattrak/souvenir data since it isnt available in game files
    const result = {};
    for (const [name, obj] of Object.entries(item_sets)) {
      const collection = {
        name,
        tag_token: obj.name,
        desc_token: obj.set_description,
      };

      // Add meta
      const meta = collections_meta[name];
      if (!meta) {
        console.info(
          `Parser extractCollections: error: no metadata found in resources/collections.meta.json for '${name}'`
        );
        process.exit(1);
      }
      Object.assign(collection, meta);

      // Add items
      const items = [];
      for (const vdfItem of Object.keys(obj.items)) {
        const [paintKit, weapon] = vdfItem.replace("[", "").split("]");
        items.push({
          paintKit,
          weapon,
        });
      }
      collection.items = items;

      result[name] = collection;
    }

    const exportPath = "./assets/collections.json";
    this.isVerbose &&
      console.info(
        `Parser extractCollections: writing collections to ${exportPath}`
      );
    await fs.promises.writeFile(exportPath, JSON.stringify(result, null, 2));

    this.isVerbose && console.info("Parser extractCollections: end");
  }
}
