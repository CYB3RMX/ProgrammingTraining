from strutils import parseInt

echo "What is your age?: "
let age: int = parseInt(readLine(stdin))

if age > 50:
    echo "Oldd"
elif age < 50:
    echo "Young"
else:
    echo "Age: ", age