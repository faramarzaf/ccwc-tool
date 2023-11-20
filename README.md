# ccwc-tool
Demonstration of Unix wc command written in Go.

# How to build
```
go build -o wc main.go
```

# How to Run
 
```bash
# Prints the number of bytes
 ./wc -ccwc=c test.txt
 
# Prints the number of lines
 ./wc -ccwc=l test.txt
 
# Prints the number of words
 ./wc -ccwc=w test.txt

# Prints the number of characters
 ./wc -ccwc=m test.txt

# Prints the number of bytes, lines and words
./wc -ccwc= or ./wc
```

* Also you can exclude the filename and it uses `test.txt` as a default file.














