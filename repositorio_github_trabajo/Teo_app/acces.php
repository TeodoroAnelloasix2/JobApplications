<?php
session_start();
if ( isset($_SESSION['fail']) || !isset($_SESSION['user']) || !isset($_SESSION['ok'])){
    session_destroy();
    header("location:index.html");
    exit();
}
else {
$_SESSION['bonus']="bonus";
require_once('header.php');

if (isset($_SESSION['mensaje_registre'])){
    echo $_SESSION['mensaje_registre'];
    unset($_SESSION['mensaje_registre']);
    echo "<br>";
}


echo '<p style="color: #21618c;">Disfruta de la herramienta filas y columnas ofrecida por esta pàagina web!</p>';
echo "<button>";
echo "<a href=\"bonus.php\">Aplicación bonus!</a>";
echo "</button>";

echo "<br>";

echo "<br>";
echo "<button>";
echo "<a href=\"logout.php\">Logout </a>";
echo "</button>";
}
?>