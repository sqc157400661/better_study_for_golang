<?php

$result = [
	"Code"=>200,
	"Msg" => "success",
	"Data" => [
		[
		    "Age"=>"35",
			"Name"=>"CangLaoShi",
		],
	]
];

@file_put_contents('/website/goreplay/phpdemo.log',"ok".PHP_EOL,FILE_APPEND);

echo json_encode($result);