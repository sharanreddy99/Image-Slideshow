<?php

require realpath(__DIR__.'/../../vendor/autoload.php');

class DatabaseConnection {

    private $connection;
    private $db_host;
    private $db_user;
    private $db_pass;
    
    public function __construct(){
        $this->db_host = $_ENV["IMAGE_VIEWER_MYSQL_HOST"];
        $this->db_user = $_ENV["IMAGE_VIEWER_MYSQL_USER"];
        $this->db_pass = $_ENV["IMAGE_VIEWER_MYSQL_PASSWORD"];
        $this->db_database = $_ENV["IMAGE_VIEWER_MYSQL_DATABASE"];
    }

    public function getConnection(){
        $this->connection = null;
        try {

            $this->connection = new mysqli($this->db_host,$this->db_user,$this->db_pass);
            $sql = "create database if not exists ". $this->db_database;
            $this->connection->query($sql);

            $sql = "use ".$this->db_database;
            $this->connection->query($sql);
            
            $sql = "create table if not exists user(id int(10) auto_increment primary key,firstname varchar(50),lastname varchar(50),email varchar(50),mobile varchar(10),password varchar(50),token tinytext default NULL);";
            $this->connection->query($sql);
            
            $sql = "create table if not exists images(fileno int(10) auto_increment primary key,filename varchar(50),imagedata mediumtext,email varchar(50));";
            $this->connection->query($sql);

        }catch(Exception $e){
            $this->connection = null;
        }
        return $this->connection;

    }
}
?>