#!/bin/bash
set -eo pipefail

for folder in $(ls -d -- */);
do
cd "$folder";
echo -e "\e[1;32m==========testing $folder==========\e[0;0m";
go test -v --bench . --benchmem;
go test -coverprofile=cover.out;
bash <(curl -Ls https://coverage.codacy.com/get.sh) report --force-coverage-parser go -r cover.out;
cd ..  ;
done
