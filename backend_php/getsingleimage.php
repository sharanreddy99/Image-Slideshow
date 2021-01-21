<?php 
  
    require("./php_cors.php");
    require("./connection.php");

    $response = array();
    
  if ($_SERVER["REQUEST_METHOD"]=="POST") { 

    $filename = "../backend_golang/gallery/".$_POST["filename"];
    $imageData = base64_encode(file_get_contents($filename));
    $response["imagedata"] = $imageData;
    echo json_encode($response); 
  }
?>
