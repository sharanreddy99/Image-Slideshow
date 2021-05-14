<?php

    require("../index.php");

    function isValidUser($conn,$token){

        $email = verifyToken($token);
        
        $sql = "select * from user where email='$email' and token='$token'";
        $res = $conn->query($sql);

        if($res->num_rows>0){
            return $email;
        }else{
            return null;
        }
    }
?>