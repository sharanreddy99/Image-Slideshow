<?php 

  require("./config/php_cors.php");
  require("./config/connection.php");
  require("./config/validuser.php");
  function verifyUserHandler($message,&$response,&$conn){
  
      $conn = (new DatabaseConnection())->getConnection();
  
      if(!$conn){
        http_response_code(503);
        $response["ModalTitle"]="Service Unavailable...";
        $response["ModalBody"]="$message Service is unavailable right now... Please try again later";
        throw new Exception("503 Error"); 
      }
  
      if($_SERVER["REQUEST_METHOD"]!="POST"){
          http_response_code(503);
          $response["ModalTitle"]="Method Not Allowed...";
          $response["ModalBody"]="Specified request type is not allowed";
          throw new Exception("503 Error"); 
      }

      $token = $_POST["token"];
      $email = isValidUser($conn,$token);
  
      if ($email == null){
          $response["ModalTitle"]="Not Authorized...";
          $response["ModalBody"]="You are not authorized...";    
          http_response_code(401);
          throw new Exception("401 Error");
      }
      
      http_response_code(200);
      return $email;
  }
  
?>