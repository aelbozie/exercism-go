for i in $( ls);
do  
cd $i;
echo -e "\e[1;32m==========testing $i==========";
stack  test;
cd ..  ;
done
