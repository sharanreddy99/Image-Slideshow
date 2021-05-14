<?php 
  
  require("./checkvaliduserhandler.php");
  $response = array();
  $conn = null;
  
  try {

    $email = verifyUserHandler("Fetching User",$response,$conn);

    $sql = "select * from user where email = '$email';";
    $res = $conn->query($sql); 
    
    if($res->num_rows>0){
      $row = $res->fetch_assoc();
      $response["firstname"] = $row["firstname"];
      $response["lastname"] = $row["lastname"];
      $response["mobile"] = $row["mobile"];
      $response["password"] = $row["password"];
      $response["email"] = $email;
    }
    
    echo json_encode($response); 

  }catch(Exception $exception){
    echo json_encode($response);
}
?>
