<?php
require_once('header.php');
require_once('error.php');
$user_name=$_POST['user'];
$password=$_POST['pswd'];

$patron_nombre="/^[A-Za-z]+\s?[A-Za-z]*$/";
$es_nombre_valido=preg_match($patron_nombre,$user_name);



if (empty($user_name) || empty($password)){
    if(empty($user_name)){
        setcookie('user_empty',gestio_errors(15));
    }
    if(empty($password)){
        setcookie('password_empty',gestio_errors(13));
    }
        unset($_SESSION['ok']);
        unset($_SESSION['user']); 
        setcookie('rellena_campos',gestio_errors(2));
        $_SESSION['fail']="fail";   
        header("location: login.php");
        
}
elseif ($user_name!=$password) {
    
    unset($_SESSION['ok']);
    unset($_SESSION['user']); 
    setcookie('error_login',gestio_errors(1));
    $_SESSION['fail']="fail";   
    header("location: login.php"); 
}

elseif($es_nombre_valido==0){
    unset($_SESSION['ok']);
    unset($_SESSION['user']); 
    setcookie('false_nom_user',gestio_errors(4));
    $_SESSION['fail']="fail";   
    header("location: login.php"); 
}
else {
    unset($_SESSION['fail']); 
    $_SESSION['ok']="OK";
    $_SESSION['user']=$user_name;
    header("location:acces.php");

}
?>