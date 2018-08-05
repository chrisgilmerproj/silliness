# Cleanup Slack

Original code modified from @betaorbust at https://gist.github.com/betaorbust/9a164694d0bdf976f8010740bcefc1ae

## Usage

To use you need to set up `SLACK_TOKEN` in your env vars.
Get the token from the [Legacy Tokens](https://api.slack.com/custom-integrations/legacy-tokens)
on Slack.

```sh
export SLACK_TOKEN=<my unique token from slack>
```

Install Node and install requirements:

```sh
brew install node
npm install
```

Then use the script

```sh
SLACK_TOKEN=$SLACK_TOKEN ./cleanup.js
```

You can run it endlessly until it stops making output this way:

```sh
while true; do SLACK_TOKEN=$SLACK_TOKEN ./cleanup.js; done
```
