
# Simple Web Server with Go Lang

A simple web server built with Go Lang.  
All the web content can dynamically called via url 
and for the project config is using [GoDotEnv](https://github.com/joho/godotenv).


## Usage/Examples

update dependency (first time pull only)
```go
go mod tidy
```

run with
```go
go run .
```
the expected output on the terminal should be like:  
![picture alt](https://github.com/williamluisan/golang_simple_webserver/blob/master/screenshot/webserver_run.png "Program running")



if success run the program, try to access via your web browser: 
* http://localhost:8081/site_a
* http://localhost:8081/site_b/other.html
* http://localhost:8081/site_b/sub_dir/sub_dir.html
* http://localhost:8081/download/sample.pdf
* http://localhost:8081/download/sample.jpg
All the website content are inside the root folder `./public_html`  
To change the configuration can do on this file `./env`


## References
* https://www.geeksforgeeks.org/how-to-build-a-simple-web-server-with-golang/
* https://blog.logrocket.com/creating-a-web-server-with-golang/
* https://stackoverflow.com/questions/26903797/how-do-i-handle-modified-time-with-http-servecontent