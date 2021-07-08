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
}
