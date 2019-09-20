/**
 * This file is our electron entrypoint.
 **/
"use strict";

const astilectron = require("astilectron");
const join = require("path").join;
const child_process = require("child_process");

const port = "8080";
// Start go-astilectron
const cmd = child_process.exec(`${join("app", "main")} ${port}`);

cmd.stdout.pipe(process.stdout);
cmd.stderr.pipe(process.stderr);

cmd.stdout.on("data", data => {
  /**
   * Wait for go-astilectron to print a JSON object to stdout
   * that contains the address of the TCP server.
   **/
  if (data.match("Start")) {
    console.log("Starting astilectron");
    // Start astilectron.
    astilectron.start(`127.0.0.1:${port}`);
  }
});
