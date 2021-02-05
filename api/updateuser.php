<?php 
  
  require("./checkvaliduserhandler.php");
  $response = array();
  $conn = null;
  
  try {

    $email = verifyUserHandler("Update User",$response,$conn);
    $firstname = $_POST["firstname"];
    $lastname = $_POST["lastname"];
    $mobile = $_POST["mobile"];
    $password = $_POST["password"];
    $newemail = $_POST["email"];

    if($email!=$newemail){
      $sql = "select * from user where email='$newemail';";
      $res = $conn->query($sql);
  
      if($res->num_rows>0){
          http_response_code(403);
          $response["ModalTitle"]="Email already registered...";
          $response["ModalBody"]="A User has already registered with the specified Email..Please try again with another Email...";
          throw new Exception("403 Error");
      }
    }

    $sql = "update user set firstname = '$firstname',lastname='$lastname',email='$newemail',password='$password',mobile='$mobile' where email = '$email';";

    if($conn->query($sql)){
      $sql = "update images set email='$newemail' where email='$email'";
      $conn->query($sql);
      http_response_code(200);
      $response["ModalTitle"]="User Updated...";
      $response["ModalBody"]="User details have been updated successfully.Please login again...";
        
    }else{
      http_response_code(403);
      $response["ModalTitle"]="User Updation Failed...";
      $response["ModalBody"]="Updating user has failed..Please try again...";
      throw new Exception("403 Error");
    }
    
    echo json_encode($response); 

  }catch(Exception $exception){
    echo json_encode($response);
}
?>
