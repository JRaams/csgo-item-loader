#!/usr/bin/env node
import { ArgLoader } from "./argloader.js";
import { Downloader } from "./downloader.js";
import { Parser } from "./parser.js";

async function main() {
  const args = new ArgLoader().load();
  const { download, parse, stitch, verbose } = args;
  verbose && console.info("Arguments:", args);

  if (download) {
    const downloader = new Downloader(verbose);
    await downloader.fetchVDLFiles();
  }

  if (parse) {
    const parser = new Parser(verbose);
    await parser.parseAssets();
    await parser.extractCollections();
    await parser.extractWeapons();
  }
}
main();
