#!/usr/bin/env node

import { ArgLoader } from "./argloader.js";
import { Downloader } from "./downloader.js";

async function main() {
  const args = new ArgLoader().load();
  console.info(args);

  if (args.download) {
    const downloader = new Downloader();
    downloader.fetchVDLFiles();
  }
}
main();
