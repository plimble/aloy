#!/bin/sh

set -e

giturl="https://$3:$4@$1.com/$2"

if ! (git clone $giturl -q app) then
  exit 1
fi

cd app

glide -q init --non-interactive && glide -q install
if [ $? -gt 0 ]; then
    echo "Cannot glide '$giturl'" >&2
    exit 2
fi

rm -rf coverage.out

echo 'mode: count' > coverage.out

for d in $(go list ./... | grep -v vendor); do
    ret=$(go test -tags=$5 -coverprofile=coverage-tmp.out -covermode=count $d)
    if [ $? -gt 0 ]; then
        echo $ret >&2
        exit 3
    fi

    if [ -f coverage-tmp.out ]; then
        tail -n +2 coverage-tmp.out >> coverage.out
        rm -rf coverage-tmp.out
    fi
done

go tool cover -func=coverage.out -o=report.out
if [ $? -gt 0 ]; then
    echo "Cannot make func coverage" >&2
    exit 4
fi

go tool cover -html=coverage.out -o=/dev/stdout
if [ $? -gt 0 ]; then
    echo "Cannot make html coverage" >&2
    exit 5
fi

number=$(cat report.out | grep 'total' | grep -o '[0-9]*[.][0-9]')

echo "<!-- cov:$number -->"
