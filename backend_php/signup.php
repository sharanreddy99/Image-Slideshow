<?php
    require("./php_cors.php");
    require("./connection.php");

    $response = array();

    if($_SERVER["REQUEST_METHOD"]=="POST"){
        $firstname = $_POST['firstname'];
        $lastname = $_POST['lastname'];
        $email = $_POST['email'];
        $mobile = $_POST['mobile'];
        $password = $_POST['password'];

        $sql = "create table if not exists user(id int(10) auto_increment primary key,firstname varchar(50),lastname varchar(50),email varchar(50),mobile varchar(10),password varchar(50));";
        $conn->query($sql);

        $sql = "select * from user where email='$email';";
        $res = $conn->query($sql);

        if($res->num_rows>0){
            $response["ModalTitle"]="Email already registered...";
            $response["ModalBody"]="A User has already registered with the specified Email..Please try again with another Email...";
            http_response_code(200);
        }
        else{
            $sql = "insert into user values(NULL,'$firstname','$lastname','$email','$mobile','$password');";
            if($conn->query($sql)){
                $response["ModalTitle"]="Registration Successful...";
                $response["ModalBody"]="Registration has been completed successfully..Please Login...";
                http_response_code(201);
            }else{
                $response["ModalTitle"]="Registration Failed...";
                $response["ModalBody"]="Error occured during registration..Please try again later...";    
                http_response_code(200);
            }
        }

        echo json_encode($response);
    }

?>