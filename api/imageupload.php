<?php 
    require("./checkvaliduserhandler.php");
    $response = array();
    $conn = null;
    
    try {
  
      $email = verifyUserHandler("Image Upload",$response,$conn);

        $sql = "select max(fileno) fileno from images";
        $maxfileno = ($conn->query($sql)->fetch_assoc())["fileno"];

        if($maxfileno==null){
            $maxfileno = 1;
        }else{
            $maxfileno += 1;
        }

        $filename = $_FILES["imagefile"]["name"]; 
        $tempname = $_FILES["imagefile"]["tmp_name"]; 

        if($tempname==""){
            http_response_code(403);
            $response["status"] = "failure";
            $response["ModalTitle"]="Image Size Exceeded....";
            $response["ModalBody"]="Please upload image of size less than 4MB...";
            throw new Exception("403 error");
        }
        
        $imageFileType = strtolower(pathinfo($filename,PATHINFO_EXTENSION));
        $filename="image".$maxfileno.".".$imageFileType;


        // Allow certain file formats
        if($imageFileType != "jpg" && $imageFileType != "png" && $imageFileType != "jpeg") {
            http_response_code(403);
            $response["status"] = "failure";
            $response["ModalTitle"]="Invalid File Format....";
            $response["ModalBody"]="Please upload files of format jpg, jpeg and png...";
            throw new Exception("403 error");

        }else if($_FILES["imagefile"]["size"]>4000000){
            http_response_code(403);
            $response["status"] = "failure";
            $response["ModalTitle"]="Image Size Exceeded....";
            $response["ModalBody"]="Please upload image of size less than 4MB...";
            throw new Exception("403 error");
        }
        else{

            $imageData = base64_encode(file_get_contents($tempname));

            $sql = "insert into images values(null,'$filename','$imageData','$email');";
            $conn->query($sql);

            $response["status"]="success";
            $response["filename"]= $filename;
            http_response_code(200);
        }
        echo json_encode($response);

    }catch(Exception $exception){
        echo json_encode($response);
    }
?>
