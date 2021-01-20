<?php
$servername = "localhost";
$user = "root";
$pass = "";

$conn = new mysqli($servername,$user,$pass);

if($conn){
    $sql = "create database if not exists sureify";
    $conn->query($sql);

    $sql = "use sureify";
    $conn->query($sql);
}

?>