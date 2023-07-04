#!/bin/bash
COVERAGE=`go tool cover -func=coverage.out | grep total: | grep -Eo '[0-9]+\.[0-9]+'`
echo $COVERAGE
COLOR=orange
if (( $(echo "$COVERAGE < 50" | bc -l) )) ; then
    COLOR=red
    elif (( $(echo "$COVERAGE == 100" | bc -l) )); then
    COLOR='brightgreen'
    elif (( $(echo "$COVERAGE > 80" | bc -l) )); then
    COLOR=green
    elif (( $(echo "$COVERAGE > 70" | bc -l) )); then
    COLOR=yellowgreen
    elif (( $(echo "$COVERAGE > 60" | bc -l) )); then
    COLOR=yellow
fi
curl "https://img.shields.io/badge/coverage-$COVERAGE%25-$COLOR?style=flat&logo=go" > coverage.svg
