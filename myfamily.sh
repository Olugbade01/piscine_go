#!/bin/bash
curl -s "https://acad.learn2earn.ng/assets/superhero/all.json" | jq -r '.[] | select(.id == (env.HERO_ID | tonumber)) | .connections.relatives | gsub ("\n"; "\\n")'
