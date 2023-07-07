# Web development with Go

* This part of Github will be filled with code and tutorials from this amazing book.
* I really like this book and want to finish it as soon as possible. 
* Sometimes I will make some notes with H4 to leave some important notice.

### Chapter 1

*

### Chapter 2
**Basic Web Application**

*

### Chapter 3
**Adding New Pages**

* We need to install _GorillaMux Router_ for this: `go get -u github.com/gorilla/mux`
* And because we in China we need to use proxy:  
`go env -w GO111MODULE=on`  
`go env -w GOPROXY=https://goproxy.cn,direct`

### Chapter 4
**A Brief Introdaction To Templates**

* `text/templates` and `html/templates` encoding in different ways, so it is more safe to use second one to prevent javascript insertions
* `p.46-47` quite usefull links about `html` templates

### Chapter 5
**Understanding MVC**

*

### Chapter 6
**Creating our first views**

*

### Chapter 7
**Creating a sign up page**

*

### Chapter 11
**Remembering users**

* Got some problems with HMAC pkg and hash function and also remembering tockens and that stuff
* `https://winitpro.ru/index.php/2019/10/25/ustanovka-nastrojka-postgresql-v-windows/`
* `D:`, `cd PostgreSQL\bin` , 
* For opening DB on Windows we should tipe `chcp 1251` to change keyset
* And also we should type `psql -U postgres`

### Chapter 12
**Validating and normalizing data**

* Box 12.12 something really difficult here happend 
* p.380 autodocumentation of programm