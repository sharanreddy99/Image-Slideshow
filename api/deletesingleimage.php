<?php 

  require("./checkvaliduserhandler.php");
  $response = array();
  $conn = null;

  try {

    $email = verifyUserHandler("Delete Image",$response,$conn);
    $filename = $_POST["filename"];


    $sql = "delete from images where email = '$email' and filename='$filename';";
    
    if($conn->query($sql)){
      $response["ModalTitle"] = "Operation Successful...";
      $response["ModalBody"]  = "image ($filename) has been deleted successfully...";
      http_response_code(200);
    }else{
      $response["ModalTitle"] = "Operation Failed...";
      $response["ModalBody"]  = "Image can't be deleted...";
      http_response_code(403);
      throw new Exception("403 Error");  
    }
    
    echo json_encode($response); 

  }catch(Exception $exception){
    echo json_encode($response);
}
?>
