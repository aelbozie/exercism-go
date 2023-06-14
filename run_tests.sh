#!/bin/bash
set -eo pipefail

dirs=$(ls -d -- */);

for dir in $dirs
do
    cd "$dir";
    echo -e "\e[1;32m==========testing $dir==========\e[0;0m";
    go test -v --bench . --benchmem;
    go test -coverprofile=cover.out;
    bash <(curl -Ls https://coverage.codacy.com/get.sh) report --partial --force-coverage-parser go -r cover.out;
    cd ..
done
export CODACY_PROJECT_TOKEN=$1
bash <(curl -Ls https://coverage.codacy.com/get.sh) final;
