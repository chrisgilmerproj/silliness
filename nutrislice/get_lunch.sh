#! /usr/bin/env bash

#
# Get the menu for the current day or the last day of the week
#

set -eu -o pipefail

LOOK_BACK=0

DAY_OF_WEEK=$(date +%w)
if [ "${DAY_OF_WEEK}" -eq 0 ]; then
  LOOK_BACK="tomorrow"
elif [ "${DAY_OF_WEEK}" -eq 6 ]; then   
  LOOK_BACK="1 day ago"
fi

echo "Getting lunch for:"
date -d "${LOOK_BACK}" +'%Y/%m/%d'
echo

date_dash=$(date -d "${LOOK_BACK} day ago" +"%Y-%m-%d")
date_slash=$(date -d "${LOOK_BACK} day ago" +"%Y/%m/%d")
menu_data=$(curl -s -XGET "https://bend.nutrislice.com/menu/api/weeks/school/elk-meadow-elementary/menu-type/fall-2020-menu/${date_slash}/")

jq -r ".days[] | select( .date | contains(\"${date_dash}\" )) | .menu_items[].food | select( . != null ) | .name" <<< "${menu_data}"

