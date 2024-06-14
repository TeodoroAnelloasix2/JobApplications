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

else{

    require_once('header.php');

   

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

    if (isset($_COOKIE['column_rows_false'])){
        echo $_COOKIE['column_rows_false'];
        setcookie('column_rows_false',NULL,time()-(24*60*60));
    }
}

?>

<!DOCTYPE html>
<html>
	<head></head>
    <h3 style="color: #21618c">Inserta las columnas y las filas! </h3>
	<body>
          <form action="column_rows.php" method="GET">
          <label for="column">Column: </label>
	      <input type="number" name="column" id="column" required><br>
	      <label for="rows">Rows: </label>
	      <input type="number" name="rows" id="rows" required><br>
	      <input type="submit" value="send">
          </form>
   
	</body>
