<?php

include_once('../HessianPHPv2/src/HessianService.php');

class HelloWorldService{
    public function add($a, $b){
        return $a + $b;
    }

    public function show($name, $age)
    {
        return "Hi, $name is $age old.";
    }
}

$wrapper = new HessianService(new HelloWorldService,array('displayInfo' => true));
$wrapper->handle();
