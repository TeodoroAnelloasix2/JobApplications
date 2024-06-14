<?php
require_once('header.php');
require_once('error.php');

if (isset($_COOKIE['campos_vacios'])){
    echo '<p style="color: red;"> ERROR GRAVE: </p>';
    echo $_COOKIE['campos_vacios'];
    echo '<br>';
    setcookie('campos_vacios',NULL,time()-(24*60*60));

    if(isset($_COOKIE['mail_empty'])){
        echo $_COOKIE['mail_empty'];
        setcookie('mail_empty',NULL,time()-(24*60*60));
        echo '<br>';
    }

    if(isset($_COOKIE['user_empty'])){
        echo $_COOKIE['user_empty'];
        setcookie('user_empty',NULL,time()-(24*60*60));
        echo '<br>';
    }
    if(isset($_COOKIE['pass_empty'])){
        echo $_COOKIE['pass_empty'];
        setcookie('pass_empty',NULL,time()-(24*60*60));
        echo '<br>';
    }
    if(isset($_COOKIE['nom_empty'])){
        echo $_COOKIE['nom_empty'];
        setcookie('nom_empty',NULL,time()-(24*60*60));
        echo '<br>';
    }
    if(isset($_COOKIE['apellido_empty'])){
        echo $_COOKIE['apellido_empty'];
        setcookie('apellido_empty',NULL,time()-(24*60*60));
        echo '<br>';
    }
}
elseif (isset($_COOKIE['email_false'])){

    echo $_COOKIE['email_false'];
    echo "<br>";
    echo '<p style="color: red;"> Ejemplo email: user1@gmail.com </p>';
    setcookie('email_false',NULL,time()-(24*60*60));
}
elseif (isset($_COOKIE['error_login'])){
    echo $_COOKIE['error_login'];
    echo "<br>";
    echo '<p style="color: red;">Error en el proceso del registro! </p>';
    setcookie('error_login',NULL,time()-24*60*60);
}
elseif (isset($_COOKIE['false_nom_user'])){
    echo $_COOKIE['false_nom_user'];
    echo "<br>";
    echo '<p style="color: red;">Ejemplo de Nombres y Apellidos validos: Teodoro Anello, Carlos Leon! </p>'; 
    echo "<br>";
    echo '<p style="color: red;">Solo se acepta 1 espacio en medio!</p>';
    setcookie('false_nom_user',NULL,time()-24*60*60);
}
elseif(isset($_COOKIE['id_user_registre_false'])){

    echo $_COOKIE['id_user_registre_false'];
    echo "<br>";
    echo '<p style="color: red;">Ejemplo de id usuario: teo96_asix </p>'; 
    echo "<br>";
    setcookie('id_user_registre_false',NULL,time()-(24*60*60));
}
?>
<br>
<h3>REGISTRATE<h3>
<br>
<form action="aplicacio_registre.php" method="POST">

<label for="user">User Name: </label>
<input id="user" type="text" name="nom_user"><br><br>
<label for="passwd">Password: </label>
<input id="passwd" type="password" name="pwd"><br><br>
<label for="email">Email: </label>
<input id="email" type="text" name="mail"><br><br>
<label for="nom">Tu Nombre: </label>
<input id="nom" type="text" name="Nombre"><br><br>
<label for="apellido">Tu Apellido: </label>
<input id="apellido" type="text" name="cognom"><br><br>
<input type="submit" value="Send Data">
</form>
<br>
<br>
<p style="color: #10056f">Si ya tienes cuenta puedes iniciar sesiòn aqui</p>
<button><a href="login.php">Login</a></button>