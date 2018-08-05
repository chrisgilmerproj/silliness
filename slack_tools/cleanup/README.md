# Cleanup Slack

Original code modified from @betaorbust at https://gist.github.com/betaorbust/9a164694d0bdf976f8010740bcefc1ae

## Usage

Install Node and install requirements:

```sh
brew install node
npm install
```

Make a Slack Token and copy it to your `.bashrc` as:

```sh
SLACK_TOKEN=XXXXXNOTATOKENXXXXX
```

Then use the script

```sh
source .bashrc
SLACK_TOKEN=$SLACK_TOKEN ./cleanup.js
```

You can run it endlessly until it stops making output this way:

```sh
while true; do SLACK_TOKEN=$SLACK_TOKEN ./cleanup.js; done
```
