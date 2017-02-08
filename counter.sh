# increment a counter

path_file="/app/data/cpt.txt"
cpt=0
if [ -e "$path_file" ]
then
    cpt=`more $path_file`
fi

let "cpt=cpt+1"
echo $cpt > $path_file
echo $cpt