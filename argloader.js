import commandLineUsage from "command-line-usage";
import commandLineArgs from "command-line-args";

export class ArgLoader {
  constructor() {
    this.optionList = [
      {
        name: "help",
        description: "Print this usage guide.",
        alias: "h",
        type: Boolean,
      },
      {
        name: "download",
        description: "Download vdf files from GameTracking-csgo repo",
        alias: "d",
        type: Boolean,
      },
      {
        name: "parse",
        description: "Parse vdf into json",
        alias: "p",
        type: Boolean,
      },
      {
        name: "stitch",
        description: "Stitch json together and output skin info",
        alias: "s",
        type: Boolean,
      },
      {
        name: "verbose",
        description: "Enable console logs with extra info",
        alias: "v",
        type: Boolean,
      },
    ];

    this.usageSections = [
      {
        header: "CSGO Item Loader",
        content:
          "Parses CS:GO weapons/skins from game files into api-ready json.",
      },
      {
        header: "Options",
        optionList: this.optionList,
      },
    ];
  }

  load() {
    const args = commandLineArgs(this.optionList);
    // If no arguments were specified, show help screen and exit
    if (Object.keys(args).length === 0 || args.help) {
      const usage = commandLineUsage(this.usageSections);
      console.log(usage);
      process.exit(1);
    }
    return args;
  }
}
