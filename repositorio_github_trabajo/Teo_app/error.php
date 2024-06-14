<?php
session_start();
function gestio_errors($err){

    switch($err){
        case 1:
            return "PASSWORD Y/O USUARIO NO VALIDOS! El login ha fallado!";       
        case 2:
            return "Assegurate de llenar todos los campos!";
        case 3:
            return "Email NO valido! ";     
        case 4: 
            return "Nombre de usuario No valido!";
        case 5:
            return "Nombre o Apllido NO validos!";
        case 6:
            return "Id de usuario No valido en el formulario de registro!";
        case 11:
            return "Elige valores positivos para las filas y las columnas!";
        case 12:
            return "Campo mail sin datos!";
        case 13:
            return "Campo password sin datos!";
        case 15:
            return "Campo usuario sin datos!";

        case 25:
            return "Campo Nombre cliente sin datos!";
        case 26:
            return "Campo Apellido cliente sin datos!";
        
    }
        
}

?>