<!DOCTYPE html>
<html>
<?php
require_once('header.php');
require_once('error.php');
if( isset($_COOKIE['error_login'])){
   echo $_COOKIE['error_login'];
   echo "<br>";
   echo '<p style="color: red;">Recarga la pàgina para borrar este mensaje! </p>';
   setcookie('error_login',NULL,time()-(24*60*60));
   
}
elseif( isset($_COOKIE['false_nom_user'])){
    echo $_COOKIE['false_nom_user'];
    echo '<br>';
    echo '<p style="color: red;">Ejemplo: Teodoro Anello</p>';
    echo '<br>';
    echo '<p style="color: red;">Solo se permite 1 espacio en medio!!</p>';
    setcookie('false_nom_user',NULL,time()-(24*60*60));
    
}
elseif( isset($_COOKIE['rellena_campos'])){
    echo $_COOKIE['rellena_campos'];
    echo '<br>';
    echo '<p style="color: red;">Se han detectados campos vacios!</p>';
    echo '<br>';   
    setcookie('rellena_campos',NULL,time()-(24*60*60));
    if (isset($_COOKIE['user_empty'])){
        echo $_COOKIE['user_empty']; 
        echo '<br>';
        setcookie('user_empty',NULL,time()-(24*60*60));
    }
    if (isset($_COOKIE['password_empty'])){
        echo $_COOKIE['password_empty']; 
        echo '<br>';
        setcookie('password_empty',NULL,time()-(24*60*60));
    }
    
}
?>
<br>
<h3>LOGIN<h3>
<br>
<p>Los campos con * son obligatorios! </p>
<form action="aplicacio_login.php" method="POST">
    <label for="username">* User name : </label>
    <input name="user" type="text" id="username"><br>

    <label for="passwd">* Password : </label>
    <input name="pswd" type="password" id="passwd"><br>
    <input type="submit" value="Send Data">

<form>
<br>
<br>
<p style="color: #10056f">Si no tienes cuenta puedes registrate aqui</p>   
<button><a href="registre.php">Registre</a></button>

</body>
</html>