<?php
session_start();
if(!isset($_SESSION['bonus']) || isset($_SESSION['fail'])){
    unset($_SESSION['ok']);
    unset($_SESSION['user']);
    unset($_SESSION['bonus']);
    $_SESSION['fail']="fail";
    header('location: registre.php');
    exit();
}
else {
require_once('header.php');
require_once('error.php');

$column=$_GET['column'];
$rows=$_GET['rows'];
if($column<=0 || $rows<=0){
    setcookie('column_rows_false',gestio_errors(11));
    header('location: bonus.php');
    exit();
}
else{
echo "Columnas: $column "."</br>"  ;
echo "Filas: $rows";

echo "<table border= 1>";

for ($a=1;$a<=$rows;$a++){
echo "<tr>";


for ($i=1;$i<=$column;$i++){
	echo"<td>";
        echo "$a."."$i";
        echo "</td>";
}

echo "</tr>";
}

echo "</table>";

echo "<br>";
echo "<br>";
echo "<br>";
echo "<button>";
echo "<a href=\"logout.php\">Logout </a>";
echo "</button>";

echo "<br>";
echo "<button>";
echo "<a href=\"acces.php\">Home</a>";
echo "</button>";

echo "<br>";
echo "<br>";
}
}
?>