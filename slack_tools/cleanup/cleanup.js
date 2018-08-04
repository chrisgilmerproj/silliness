#! /usr/bin/env node

var fetch = require("node-fetch");

async function wait(ms) {
  return new Promise(resolve => {
    setTimeout(resolve, ms);
  });
}

async function deleteFile(file) {
  const url = `https://slack.com/api/files.delete?token=${token}&file=${file.id}`;
  const response = await fetch(url);
  const data = await response.json();
  if (data.ok) {
    console.log(`Deleted ${file.title}`);
  } else {
    console.error(`Deleting ${file.title} failed.\n${data.error}`);
  }
}

/**
 * Deletes the largest `numToDelete` files older than `minAgeDays` on your slack.
 * Can only do 1000 at a time because I didn't want to deal with pagination.
 * YOU NEED TO BE AN ADMIN FOR THIS TO DELETE THINGS THAT YOU DIDN'T UPLOAD!
 * @param {string} token       A legacy slack token. Get yours at https://api.slack.com/custom-integrations/legacy-tokens
 * @param {number} minAgeDays  The number of days old a file can be to be deleted.
 * @param {number} numToDelete The number of files to delete total.
 */
async function deleteOldLargeFiles(token, minAgeDays, numToDelete) {
  // Slack does everything in seconds, apparently.
  const earliestDate = Date.now() / 1000 - 60 * 60 * 24 * minAgeDays;

  const filesListUrl = `https://slack.com/api/files.list?token=${token}&count=1000&ts_to=${earliestDate}`
  const listResponse = await fetch(filesListUrl);
  const listData = await listResponse.json();
  const files = listData.files.slice().sort((a, b) => b.size - a.size);

  for (let index = 0; index < Math.min(numToDelete, files.length); index++) {
    const file = files[index];
    await deleteFile(file);
    await wait(250);
  }
}

var token = process.env.SLACK_TOKEN;
var minAgeDays = 60;
var numToDelete = 1000;
deleteOldLargeFiles(token, minAgeDays, numToDelete);
