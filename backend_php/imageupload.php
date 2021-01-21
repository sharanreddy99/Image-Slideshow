<?php 
  
    require("./php_cors.php");
    require("./connection.php");

    $response = array();


  // If upload button is clicked ... 
  if ($_SERVER["REQUEST_METHOD"]=="POST") { 

    $sql = "select max(fileno) fileno from imagepaths";
    $maxfileno = ($conn->query($sql)->fetch_assoc())["fileno"];

    if($maxfileno==null){
        $maxfileno = 1;
    }else{
        $maxfileno += 1;
    }

    $email = $_POST["email"];
    $filename = $_FILES["imagefile"]["name"]; 
    $tempname = $_FILES["imagefile"]["tmp_name"]; 
    $targetdir = "../backend_golang/gallery/";
    $imageFileType = strtolower(pathinfo($filename,PATHINFO_EXTENSION));
    $targetpath = $targetdir."image".$maxfileno.".".$imageFileType;
    $filename="image".$maxfileno.".".$imageFileType;


    // Allow certain file formats
    if($imageFileType != "jpg" && $imageFileType != "png" && $imageFileType != "jpeg") {
        $response["status"] = "failure";
        $response["ModalTitle"]="Invalid File Format....";
        $response["ModalBody"]="Please upload files of format jpg, jpeg and png...";
        
    }else{

        move_uploaded_file($tempname, $targetpath);

        $sql = "insert into imagepaths values(null,'$filename','$email');";
        $conn->query($sql);

        $response["status"]="success";
        $response["filename"]= $filename;
    }
    echo json_encode($response);
  }
?>
