#!/bin/bash
URL="https://acad.learn2earn.ng/assets/superhero/all.json" 
curl -s "$URL" | jq  '.[] | select(.id == 70) | .name'
