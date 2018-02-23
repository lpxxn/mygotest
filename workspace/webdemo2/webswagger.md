
## Copy SwaggerUI to a nodeserver directory

modify index.js
```
var express = require('express');
var app = express();
app.use('/static', express.static('public'));
app.get('/', function (req, res) {
    res.send('Hello World!');
});

app.listen(3111, function () {
    console.log('Example app listening on port 3111!');
});

```

modify index.html
```
  var url = window.location.href;
  console.log(url);
  var arr = url.split("/");
  var result = arr[0] + "//" + arr[2] + "/" +arr[3] + "/swagger.json"
  const ui = SwaggerUIBundle({
   // url: "http://petstore.swagger.io/v2/swagger.json",
    url: result,
```
http://localhost:3111/static/index.html


## Generate swagger.json
generate swagger.json put it to public/ directory
```
swagger generate spec -o ./nodeserver/public/swagger.json
```