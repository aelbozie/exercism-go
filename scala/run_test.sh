for i in $( ls);
do  
cd $i;
echo -e "\e[1;32m==========testing $i==========";
sbt ++$TRAVIS_SCALA_VERSION test;
cd ..  ;
done
