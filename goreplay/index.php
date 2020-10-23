<?php

$result = [
	"Code"=>200,
	"Msg" => "success",
	"Data" => [
		[
			"Name"=>"CangLaoShi",
			"Age"=>35
		]
	]
];

echo json_encode($result);