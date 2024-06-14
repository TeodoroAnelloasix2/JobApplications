<?php
session_start();

if (isset($_SESSION['ok']) || isset($_SESSION['user'])){
    
    echo "<h1> Welcome  ". $_SESSION['user']. "</h1>";
}
else {
    $_SESSION['fail']="fail";
    echo "<h1> Welcome to TEO's APP! </h1>";
}
?>

