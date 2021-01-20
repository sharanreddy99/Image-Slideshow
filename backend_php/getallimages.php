<?php 
  
    require("./php_cors.php");
    require("./connection.php");

    $response = array();
    
  if ($_SERVER["REQUEST_METHOD"]=="POST") { 
      $email = $_POST["email"];

      $sql = "select * from imagepaths";
      $res = $conn->query($sql);
      $imagesarray = array();
      $extensionsarray = array();
        while($row = $res->fetch_assoc()){
          $filename = "./gallery/".$row["filename"];
          $imageFileType = strtolower(pathinfo($filename,PATHINFO_EXTENSION));
          $imageData = base64_encode(file_get_contents($filename));

          array_push($imagesarray,$imageData);
          array_push($extensionsarray,$imageFileType);
          
        }
    
    $response["imagesarray"] = json_encode($imagesarray);
    $response["extensionsarray"] = json_encode($extensionsarray);

    echo json_encode($response); 
  }
?>
