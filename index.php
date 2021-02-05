<?php

require __DIR__.'/vendor/autoload.php';
use \Firebase\JWT\JWT;

$dotenv = Dotenv\Dotenv::createImmutable(__DIR__);
$dotenv->load();

function generateRandomString($length = 10) {
    $characters = '0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ';
    $charactersLength = strlen($characters);
    $randomString = '';
    for ($i = 0; $i < $length; $i++) {
        $randomString .= $characters[rand(0, $charactersLength - 1)];
    }
    return $randomString;
}

function generateToken($email){
    $key = $_ENV["SECRETKEY"];
    $string = generateRandomString();

    $payload = array(
        "iat" => time(),
        "jti" => $string,
        "iss" => "localhost",
        "email" => $email,
        'exp' => time()+60*60
    );

    $jwt = JWT::encode($payload, $key);
    return $jwt;
}

function verifyToken($token){
    $key = $_ENV["SECRETKEY"];

    try{
        $decoded = JWT::decode($token, $key, array('HS256'));
        return $decoded->email;
    }catch(Exception $e){
        return "";
    }
}

?>