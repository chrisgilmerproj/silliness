# Slack Tools

## Emoji Download

To use you need to set up `SLACK_TOKEN` in your env vars.
Get the token from the [Legacy Tokens](https://api.slack.com/custom-integrations/legacy-tokens)
on Slack.

```sh
export SLACK_TOKEN=<my unique token from slack>
```

Install dependencies:

```
$ pip install requests
```

Then run this script:

```sh
$ ./emoji_download.py
```

You should now have emoji's in a directory named `./img/` in this directory.
