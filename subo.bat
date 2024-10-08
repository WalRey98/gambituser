git add .
git commit -m "last Commit"
git push 
set GOOS=linux
set GOARCH=amd64
go build -o bootstrap main.go
del bootstrap.zip
tar.exe -a -cf bootstrap.zip bootstrap