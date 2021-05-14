<?php 
  
  require("./checkvaliduserhandler.php");
  $response = array();
  $conn = null;
  
  try {

    $email = verifyUserHandler("Fetch All Images",$response,$conn);

    $sql = "select * from images where email = '$email';";
    $res = $conn->query($sql);
    $imagesarray = array();
    $extensionsarray = array();
    $filenamesarray = array();

    while($row = $res->fetch_assoc()){
      $filename = $row["filename"];
      $imageFileType = strtolower(pathinfo($filename,PATHINFO_EXTENSION));
      $imageData = $row["imagedata"];

      array_push($imagesarray,$imageData);
      array_push($extensionsarray,$imageFileType);
      array_push($filenamesarray,$filename);
    }
    
    $response["imagesarray"] = json_encode($imagesarray);
    $response["extensionsarray"] = json_encode($extensionsarray);
    $response["filenamesarray"] = json_encode($filenamesarray);

    echo json_encode($response); 
    
  }catch(Exception $exception){
    echo json_encode($response);
}
?>
