cat file.txt

if [ $? -eq 0 ]
then
    echo "the script ran ok"
    exit 0
else
    echo "the script failed" >&2
    exit 1
fi
