var http = require('http');
var url = require('url');
var fs = require('fs');
var server = http.createServer(function(request, response) {

    var path = url.parse(request.url).pathname;
    switch (path) {
        case '/page1.html':
            fs.readFile(__dirname + path, function(error, data) {
                if (error) {
                    response.writeHead(404);
                    response.write(error.toString());
                    response.end();
                } else {
                  response.writeHead(200, {
    'Content-Type': 'text/html',
    'Access-Control-Allow-Origin' : '*',
    'Access-Control-Allow-Methods': 'GET,PUT,POST,DELETE'
});
                    response.write(data.toString());
                    response.end();
                }
            });
            break;

            case '/page2.html':
                fs.readFile(__dirname + path, function(error, data) {
                    if (error) {
                        response.writeHead(404);
                        response.write(error.toString());
                        response.end();
                    } else {
                      response.writeHead(200, {
        'Content-Type': 'text/html',
        'Access-Control-Allow-Origin' : '*',
        'Access-Control-Allow-Methods': 'GET,PUT,POST,DELETE'
    });
                        response.write(data.toString());
                        response.end();
                    }
                });
                break;
    }
});
server.listen(3000);
