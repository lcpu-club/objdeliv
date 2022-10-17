<?php

$sock = fsockopen("127.0.0.1:24032");

$write = fn ($data) => @fwrite($sock, $data);

$write("CONNECT /new-object?expire=60&id=e45d7d04-52b6-473d-bc78-aa40f4b07e73 HTTP/1.1\r\n\r\n");
$write(<<<EOF
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <h1>My Website</h1>
    <p>This is a test.</p>
</body>
</html>
EOF);

var_dump(json_decode(fgets($sock), true));

fclose($sock);