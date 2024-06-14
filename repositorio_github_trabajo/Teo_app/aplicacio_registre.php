<?php
if(isset($_SESSION['user']) || isset($_SESSION['ok'])){
    unset($_SESSION['ok']);
    unset($_SESSION['user']);
}
require_once('header.php');
require_once('error.php');
$mail=$_POST['mail'];
$user_name=$_POST['nom_user'];//user_id
$pass=$_POST['pwd'];
$nombre_usuario=$_POST['Nombre'];//nombre de la persona
$apellido=$_POST['cognom'];


//validar mail 
$patron_email="/^[A-Z0-9a-z]+\@[A-Z0-9a-z]+\.[A-Za-z]+\s?$/";

$es_mail_valido=preg_match($patron_email,$mail);


//validar nombre y apellido del usuario
$patron_nombre="/^[A-Za-z]+\s?[A-Za-z]*/";
$es_nombre_valido=preg_match($patron_nombre,$nombre_usuario);

$patron_apellido="/^[A-Za-z]+\s?[A-Za-z]*/";
$es_apellido_valido=preg_match($patron_apellido,$apellido);

$patron_id_usuario="/^[A-Za-z0-9\-\_]+\s?[A-Za-z0-9\-\_]+\s?$/";

$es_id_usuario_valido=preg_match($patron_id_usuario,$user_name);


if( empty($mail) || empty($user_name) || empty($pass) ||  empty($nombre_usuario) || empty($apellido) || $user_name==" " ){
    unset($_SESSION['ok']);
    unset($_SESSION['user']); 
    setcookie('campos_vacios',gestio_errors(2));
    $_SESSION['fail']="fail";

    if (empty($mail)){
        setcookie('mail_empty',gestio_errors(12));
    }
    if (empty($user_name)){
        setcookie('user_empty',gestio_errors(15));
    }
    if (empty($pass)){
        setcookie('pass_empty',gestio_errors(13));
    }
    if (empty($nombre_usuario)){
        setcookie('nom_empty',gestio_errors(25));
    }
    if (empty($apellido)){
        setcookie('apellido_empty',gestio_errors(26));
    }
    header("location:registre.php");
    exit();
}
elseif ($user_name!=$pass) {
    unset($_SESSION['ok']);
    unset($_SESSION['user']); 
    setcookie('error_login',gestio_errors(1));
    $_SESSION['fail']="fail";
    header("location:registre.php");
    exit();
}
elseif ($es_mail_valido==0){ 
    unset($_SESSION['ok']);
    unset($_SESSION['user']); 
    setcookie('email_false',gestio_errors(3));
    $_SESSION['fail']="fail";
    header("location:registre.php");
    exit();
}
elseif ($es_nombre_valido==0 || $es_apellido_valido==0) {
    unset($_SESSION['ok']);
    unset($_SESSION['user']); 
    setcookie('false_nom_user',gestio_errors(5));
    header("location:registre.php");
    exit();
}
elseif($es_id_usuario_valido==0){
    unset($_SESSION['ok']);
    unset($_SESSION['user']);
    setcookie('id_user_registre_false',gestio_errors(6));
    header("location:registre.php");
    exit();
}
else{
    $_SESSION['ok']="OK";
    $_SESSION['user']=$user_name;
    $_SESSION['mensaje_registre']="Usuario Registrado Correctamente!";
    unset($_SESSION['fail']);
    header("location:acces.php");   
}
?>