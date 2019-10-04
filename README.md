# NodeJS
Webpages on NodeJS server connected to a Golang Api on a Go server.

serverexample.js is the JS code for setting up the node in node.js as a server.
page1.html is html page.
page2.html is html page.


I've made the button work using JS.

To run :

NodeJS :
1. In the terminal run (keep terminal in the NodeJS folder): node serverexample.js
2. Open browser : localhost:3000/page.html
3. The port can be changed in the last line of serverexample.js 


Go :

1. terminal : go get github.com/gorilla/handlers
2. terminal : go get github.com/gorilla/mux
3. goapi2.go returns json objects.
4. goapi.go return text response.
5. go run goapi.go / go run goapi2.go
6. Check using http://localhost: port-number/ name

JS part :
1. Parser for text.
2. Parser for Json.

Clone the git, unzip the node_modules for it to work.
