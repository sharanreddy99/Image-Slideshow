<?php
    require("./config/php_cors.php");
    require("./config/connection.php");

    $response = array();

    try {

        $conn = (new DatabaseConnection())->getConnection();
        
        if(!$conn){
            http_response_code(503);
            $response["ModalTitle"]="Service Unavailable...";
            $response["ModalBody"]="Signup Service is unavailable right now... Please try again later";
            throw new Exception("503 Error"); 
        }

        if($_SERVER["REQUEST_METHOD"]!="POST"){
            http_response_code(503);
            $response["ModalTitle"]="Method Not Allowed...";
            $response["ModalBody"]="Specified request type is not allowed";
            throw new Exception("503 Error"); 
        }

        $firstname = $_POST['firstname'];
        $lastname = $_POST['lastname'];
        $email = $_POST['email'];
        $mobile = $_POST['mobile'];
        $password = $_POST['password'];


        $sql = "select * from user where email='$email';";
        $res = $conn->query($sql);

        if($res->num_rows>0){
            http_response_code(403);
            $response["ModalTitle"]="Email already registered...";
            $response["ModalBody"]="A User has already registered with the specified Email..Please try again with another Email...";
            throw new Exception("403 Error");
        }

        $sql = "insert into user values(NULL,'$firstname','$lastname','$email','$mobile','$password',NULL);";
        if($conn->query($sql)){
            $response["ModalTitle"]="Registration Successful...";
            $response["ModalBody"]="Registration has been completed successfully..Please Login...";
            http_response_code(201);
        }else{
            $response["ModalTitle"]="Registration Failed...";
            $response["ModalBody"]="Error occured during registration..Please try again later...";    
            http_response_code(403);
            throw new Exception("403 Error");    
        }

        echo json_encode($response);
    
    }catch(Exception $exception){
        echo json_encode($response);
    }
?>