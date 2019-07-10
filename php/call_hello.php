<?php

require dirname(__FILE__) . '/vendor/autoload.php';

$name = '小明';

$client = new Helloworld\HelloClient('localhost:8488', [
    'credentials' => Grpc\ChannelCredentials::createInsecure(),
]);
$request = new Helloworld\HelloRequest();
$request->setName($name);
list($response, $status) = $client->SayHello($request)->wait();

echo '服务端返回状态码: ' . $status->code . PHP_EOL;
echo "服务器端回复内容：" . $response->getMessage() . PHP_EOL;
