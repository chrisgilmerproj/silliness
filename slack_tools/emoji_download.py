#! /usr/bin/env python3

import os
import sys

import requests


def main():
    slack_token = os.environ.get('SLACK_TOKEN')
    r = requests.get("https://slack.com/api/emoji.list?token={}".format(slack_token))

    # Don't continue if bad request
    if not r.ok:
        sys.exit(1)

    for emoji_name, emoji_url in r.json()['emoji'].items():
        # Skip any emoji records that are aliased
        if emoji_url.startswith('alias:'):
            continue

        try:
            image_name = "{}{}".format(emoji_name, os.path.splitext(emoji_url)[1])
            img_loc = os.path.join('./img', image_name)
            # Don't re-download
            if not os.path.isfile(img_loc):
                print(image_name, emoji_url)
                image_r = requests.get(emoji_url, allow_redirects=True)
                open(img_loc, 'wb').write(image_r.content)
        except Exception:
            print("Error: Could not parse {} from URL {}".format(emoji_name, emoji_url))
            continue


if __name__ == "__main__":
    main()
