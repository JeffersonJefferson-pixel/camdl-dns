#!/bin/sh

. ./.env

# echo $CLOUDFLARE_API_TOKEN
# echo $ZONE_ID

devp2p dns to-cloudflare --zoneid "$ZONE_ID" data/