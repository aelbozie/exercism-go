for file in $(ls -d */);
do  
cd $file;
echo -e "\e[1;32m==========testing $file==========";
sbt ++$TRAVIS_SCALA_VERSION test;
cd ..  ;
done
