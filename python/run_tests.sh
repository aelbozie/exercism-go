#!/bin/bash
set -eo pipefail

pip install pytest-cov=="2.12.1"

for folder in $(ls -d -- */ | sed "s,/$,," );
do
cd "$folder";
echo -e "\e[1;32m==========testing $folder==========\e[0;0m";

pyproj=$(echo $folder | sed "s,-,_," )
pytest --cov="$pyproj" --cov-report xml
bash <(curl -Ls https://coverage.codacy.com/get.sh) report --partial -r coverage.xml;
cd ..  ;
done
