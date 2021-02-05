<?php
    require("./checkvaliduserhandler.php");
    $response = array();
    $conn = null;
    
    try {
  
        $email = verifyUserHandler("Logout",$response,$conn);
        $sql = "update user set token=null where email='$email'";
        $conn->query($sql);

        echo "Successfully logged out";

    }catch(Exception $exception){
        echo json_encode($response);
    }

?>