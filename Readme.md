# Http server from scratch in Go

## 🧩 Overview
- Using net package from go std lib
- Unix syscalls (low level) completely from scratch


## ✅ Some factors to remember
- windows isn't unix based so it wont work using unix, you can use [golang.org/x/sys/windows](golang.org/x/sys/windows) or if you want to stay unixified use WSL tool
- on unix based systems it will work as exepected