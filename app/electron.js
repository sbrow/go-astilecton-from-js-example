"use strict";

const astilectron = require("astilectron");
const join = require("path").join;
const child_process = require("child_process");

const cmd = child_process.exec(join("app", "main.exe"));

cmd.stdout.pipe(process.stdout);
cmd.stderr.pipe(process.stderr);

cmd.stdout.on("data", data => {
  if (data[0] === "{") {
    try {
      const obj = JSON.parse(data);
      if (obj.addr) {
        const { addr } = obj;
        console.log(`Starting astilectron with "${addr}"`);
        astilectron.start(addr);
      }
    } catch (err) {
      console.error(err);
    }
  }
});
