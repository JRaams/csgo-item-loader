#!/usr/bin/env node

import { ArgLoader } from "./argloader.js";
import { Downloader } from "./downloader.js";

async function main() {
  const args = new ArgLoader().load();
  const { download, parse, stitch, verbose } = args;
  verbose && console.info(args);

  if (download) {
    const downloader = new Downloader(verbose);
    downloader.fetchVDLFiles();
  }
}
main();
