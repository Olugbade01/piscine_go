#! /bin/bash 

curl -s https://acad.learn2earn.ng/assets/superhero/all.json | jq --argjson id "$HERO_ID" '.[] | select (.id == $id) | .connections.relatives' | sed 's/^"//;s/"$//'