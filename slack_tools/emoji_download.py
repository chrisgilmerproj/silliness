#! /usr/bin/env python3

import os
import sys

import requests


def main():
    slack_token = os.environ.get('SLACK_TOKEN')
    r = requests.get("https://slack.com/api/emoji.list?token={}".format(slack_token))
    if not r.ok:
        sys.exit(1)
    for emoji_name, emoji_url in r.json()['emoji'].items():
        try:
            image_r = requests.get(emoji_url, allow_redirects=True)
            image_name = "{}{}".format(emoji_name, os.path.splitext(emoji_url)[1])
            print(image_name)
            open(os.path.join('./img', image_name), 'wb').write(image_r.content)
        except Exception:
            continue


if __name__ == "__main__":
    main()
