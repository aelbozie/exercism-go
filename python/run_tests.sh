#!/bin/bash
set -eo pipefail

pip install pytest-cov=="2.12.1";
dirs=$(ls -d -- */ | sed "s,/$,,");

for dir in $dirs
do
    cd "$dir";
    echo -e "\e[1;32m==========testing $dir==========\e[0;0m";
    pyproj=$(echo "$dir" | sed "s,-,_," )
    pytest --cov="$pyproj" --cov-report xml
    bash <(curl -Ls https://coverage.codacy.com/get.sh) report --partial -r coverage.xml;
    cd ..
done

bash <(curl -Ls https://coverage.codacy.com/get.sh) final;
