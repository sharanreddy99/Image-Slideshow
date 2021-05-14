<?php
    require("./config/php_cors.php");
    require("./config/connection.php");
    require("../index.php");

    $response = array();
    try {

        $conn = (new DatabaseConnection())->getConnection();

        if(!$conn){
            http_response_code(503);
            $response["ModalTitle"]="Service Unavailable...";
            $response["ModalBody"]="Signin Service is unavailable right now... Please try again later";
            throw new Exception("503 Error"); 
        }

        if($_SERVER["REQUEST_METHOD"]!="POST"){
            http_response_code(503);
            $response["ModalTitle"]="Method Not Allowed...";
            $response["ModalBody"]="Specified request type is not allowed";
            throw new Exception("503 Error"); 
        }

        if($_SERVER["REQUEST_METHOD"]=="POST"){
            $email = $_POST['email'];
            $password = $_POST['password'];

            $sql = "select * from user where email='$email' and password='$password';";
            $res = $conn->query($sql);

            if($res->num_rows>0){

                $token = generateToken($email);

                $sql = "update user set token = '$token' where email='$email' and password='$password';";    
                $conn->query($sql);
                                
                $response["status"]="success";
                $response["email"]=$email;
                $response["token"]=$token;
                http_response_code(200);
            }
            else{
                    $response["status"]="failure";
                    $response["ModalTitle"]="Invalid Credentials...";
                    $response["ModalBody"]="The Email or Password is incorrect. Please enter valid credentials...";    
                    http_response_code(401);
                    throw new Exception("401 Error");
            }

            echo json_encode($response);
        }

    }catch(Exception $exception){
        echo json_encode($response);
    }
?>