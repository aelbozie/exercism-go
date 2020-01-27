#!/bin/bash
set -eo pipefail

for folder in $(ls -d -- */);
do  
cd "$folder";
echo -e "\e[1;32m==========testing $folder==========";
sbt ++"$TRAVIS_SCALA_VERSION" test;
cd ..  ;
done
