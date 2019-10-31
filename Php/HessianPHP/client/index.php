<?php

include_once('../HessianPHPv2/src/HessianClient.php');

$uri = 'http://localhost:8001/index.php';
//$options = new HessianOptions();
//$proxy = new HessianClient($uri,$options);
$proxy = new HessianClient($uri);

echo $proxy->add(3,12);

echo "<br/><br/><br/>";

echo $proxy->show('Donovan.Xu','33');
