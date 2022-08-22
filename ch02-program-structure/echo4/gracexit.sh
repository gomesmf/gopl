echo "executing..."
cat doesnotexist.txt 2>/dev/null || exit 0
