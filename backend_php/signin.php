<?php
    require("./php_cors.php");
    require("./connection.php");

    $response = array();

    if($_SERVER["REQUEST_METHOD"]=="POST"){
        $email = $_POST['email'];
        $password = $_POST['password'];

        $sql = "select * from user where email='$email' and password='$password';";
        $res = $conn->query($sql);

        if($res->num_rows>0){
            $response["status"]="success";
            $response["email"]=$email;
            $response["password"]=$password;
            http_response_code(201);
        }
        else{
                 $response["status"]="failure";
                $response["ModalTitle"]="Invalid Credentials...";
                $response["ModalBody"]="The Email or Password is incorrect. Please enter valid credentials...";    
                http_response_code(200);
        }

        echo json_encode($response);
    }

?>