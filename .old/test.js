import vdf from "simple-vdf";
import fs from "fs";

const items_game_data = fs.readFileSync("./storage/items_game.txt", "utf-8");
const items_game_vdf = vdf.parse(items_game_data);
const items_game = vdf.stringify(items_game_vdf);

fs.writeFileSync(
  "./resources/items_game.json",
  JSON.stringify(items_game_vdf, null, 4)
);
