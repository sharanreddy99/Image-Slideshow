<?php 

require("./checkvaliduserhandler.php");
  $response = array();
  $conn = null;
  
  try {

    $email = verifyUserHandler("Authorization",$response,$conn);
    }catch(Exception $exception){
      echo json_encode($response);
    }
?>