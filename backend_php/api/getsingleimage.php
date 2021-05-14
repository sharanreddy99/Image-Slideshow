<?php 
  
  require("./checkvaliduserhandler.php");
  $response = array();
  $conn = null;
  
  try {

    $email = verifyUserHandler("Fetch Single Image",$response,$conn);
    $filename = $_POST["filename"];
    
    $sql = "select imagedata from images where email = '$email' and filename='$filename';";
    $res = $conn->query($sql); 
    
    if($res->num_rows>0){
      $row = $res->fetch_assoc();
      $imageData = $row["imagedata"];
      $response["imagedata"] = $imageData;
    }
    
    echo json_encode($response); 

  }catch(Exception $exception){
    echo json_encode($response);
}
?>
