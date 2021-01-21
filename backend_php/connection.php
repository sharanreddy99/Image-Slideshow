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
    
    $sql = "create table if not exists user(id int(10) auto_increment primary key,firstname varchar(50),lastname varchar(50),email varchar(50),mobile varchar(10),password varchar(50));";
    $conn->query($sql);
    
    $sql = "create table if not exists imagepaths(fileno int(10) auto_increment primary key,filename varchar(50),email varchar(50));";
    $conn->query($sql);
}

?>