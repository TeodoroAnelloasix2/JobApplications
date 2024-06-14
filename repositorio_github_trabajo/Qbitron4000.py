#!/usr/bin/python3
from selenium import webdriver
from selenium.webdriver.common.keys import Keys
import sh
import time
from selenium.webdriver.common.by import By
from selenium.webdriver.common.action_chains import ActionChains
import pyautogui
from selenium.webdriver.support.ui import Select
import datetime
from datetime import datetime
from datetime import timedelta
from datetime import date
import pywhatkit

try:
   url="https://www.empresaiformacio.org/sBid/"

   user="tu usuario"
   passwd="tu password"

   driver= webdriver.Firefox()
   driver.get(url)
   driver.maximize_window()

   x,y=100,200
   time.sleep(3)

   pyautogui.click(x,y)

   ActionChains(driver).key_down(Keys.TAB).key_up(Keys.TAB).perform()
   time.sleep(1)
   ActionChains(driver).key_down(Keys.TAB).key_up(Keys.TAB).perform()
   time.sleep(1)
   ActionChains(driver).key_down(Keys.TAB).key_up(Keys.TAB).perform()
   time.sleep(1)
   ActionChains(driver).key_down(Keys.TAB).key_up(Keys.TAB).perform()
   time.sleep(1)
   ActionChains(driver).key_down(Keys.ENTER).key_up(Keys.ENTER).perform()
   time.sleep(1)

   driver.switch_to.frame('jspContainer')

   insert_usuario= driver.find_element(By.ID,'username')

   if insert_usuario:
      insert_usuario.click()
      time.sleep(1)
      insert_usuario.send_keys(user)

   insert_password=driver.find_element(By.ID,'password')

   if insert_password:
      insert_password.click()
      time.sleep(1)
      insert_password.send_keys(passwd)


   click_login=driver.find_element(By.XPATH, "//span[@class='glyphicon glyphicon-play']")

   if click_login:
      time.sleep(1)
      click_login.click()
      time.sleep(1)



   driver.switch_to.frame('contentmain')

   tareas_pendientes=driver.find_element(By.ID, 'collapse1')
   #Stop temporal

   #Crea una lista que contiene los enlaces a los dias pendientes para insertar horas
   enlaces_a_tareas=driver.find_elements(By.XPATH,"//a[@class='list-group-item list-group-item-action flex-column align-items-start']")

   time.sleep(1)
   #Sacamos la longitud de la lista. Si la longitud es igual a 0 no hay tareas pendientes
   numero_tareas_pendientes=len(enlaces_a_tareas)
   print(f" El numero de tareas es: {numero_tareas_pendientes}")

   # for a in enlaces_a_tareas:
   #    print("enalce 1")
   #    print(a)
   
   #time.sleep(15)
   sms_no_tareas="No hay tareas pendientes"
   if numero_tareas_pendientes <=1:
      time.sleep(0.5)
      driver.close()
      pywhatkit.sendwhatmsg_instantly("+tu numero",sms_no_tareas,20,True)
   else:   
      enlace_a_tarea=driver.find_element(By.XPATH,"//a[@class='list-group-item list-group-item-action flex-column align-items-start']")
      if enlace_a_tarea:
         time.sleep(0.5)
         enlace_a_tarea.click()
      

   

      """
      pyautogui.drag(0,distance,duration=0.5) #moverse hacia abajo
      pyautogui.drag(distance,0,duration=0.5) #moverse hacia derecha
      pyautogui.drag(0,-distance,duration=0.5) #moverse hacia arriba
      pyautogui.drag(-distance,0,duration=0.5) #moverse hacia izquierda
      """
      distance=200

      while distance >0:
         pyautogui.drag(distance,0,duration=0.5) #moverse a la derecha
         distance -=20

      distance=100

      time.sleep(1)

      while distance >0:
         pyautogui.drag(0,distance,duration=0.5) #moverse hacia abajo
         distance-=15

      pyautogui.click()
      time.sleep(2)


      #################################################################

      def rellenar():
         act1_web=driver.find_element(By.ID,'activitat8671')
         if act1_web:
            print("ok")
            time.sleep(1)
         escribir_web=driver.find_element(By.ID,'inp_8671')
         if escribir_web:
            escribir_web.click()
            time.sleep(1)
            for i in range(2):
               escribir_web.send_keys("1")
               time.sleep(0.25)

         ActionChains(driver).key_down(Keys.TAB).key_up(Keys.TAB).perform()

         escribir_resolucio_incidencies=driver.find_element(By.ID,'inp_8676')
         if escribir_resolucio_incidencies:
            escribir_resolucio_incidencies.click()
            time.sleep(1)
      
            escribir_resolucio_incidencies.send_keys("2")
         
         ActionChains(driver).key_down(Keys.TAB).key_up(Keys.TAB).perform()


         insertar_depuracion=driver.find_element(By.ID,'inp_8677')
         if insertar_depuracion:
            insertar_depuracion.click()
            time.sleep(1)
            insertar_depuracion.send_keys("2")
         ActionChains(driver).key_down(Keys.TAB).key_up(Keys.TAB).perform()
         ActionChains(driver).key_down(Keys.TAB).key_up(Keys.TAB).perform()

         text="Creación de informes, limpieza de imágenes, programación de aplicaciones SIA, cierre de tickets"

         comentario=driver.find_element(By.ID,'observacionsInutilsAlumneMaiOmplira')
         if comentario:
            comentario.click()
            time.sleep(1)
            comentario.send_keys(text)

         time.sleep(0.5)
      
         #Almacenar las horas insertadas 

         almacenar=driver.find_element(By.XPATH,'/html/body/form/div[1]/div[3]/span[1]')
         if almacenar:
            time.sleep(0.5)
            almacenar.click()
            time.sleep(0.5)
      
      ###############################fin def rellenar

      #  Para saber si hay que seguir rellenando podemos extraer la fecha desde el enlace al dia siguiente, si es menor o igual a la de hoy pues seguimos rellenando horas!
      #Funcion sacar fecha dia rellenando y fecha actual
      def fechas():
         fecha_evaluar=driver.find_element(By.XPATH,'/html/body/form/div[1]/div[2]/div[7]/div[1]/label/a[2]')
         if fecha_evaluar:
            print("Encontrado")
            valor=fecha_evaluar.get_attribute('href')
            print(valor)
            start_index = valor.find('data=')


            end_index = valor.find('&', start_index)


            data_value = valor[start_index + 5:end_index]
            print(data_value)


            fecha=data_value[0:4]+str("-")+data_value[4:6]+str("-")+data_value[6:]
            print(fecha)

            fecha2=datetime.strptime(fecha,'%Y-%m-%d')
            print(fecha2)
            print(type(fecha2))

            fecha2=fecha2.date()
            print(type(fecha2))

            fecha_actual=date.today()
            return fecha_actual,fecha2
            #FIN def fechas
      #/html/body/form/div[1]/div[2]/div[7]/div[1]/label/a[2]/img
      ##############################################################
      #Function para ir al dia siguiente
      def dia_siguiente():
         ir_a_dia_siguiente=driver.find_element(By.XPATH,'/html/body/form/div[1]/div[2]/div[7]/div[1]/label/a[2]/img')
         if ir_a_dia_siguiente:
            #print("Encontrado boton dia siguiente")
            time.sleep(0.5)
            ir_a_dia_siguiente.click()
            time.sleep(1)
      ##############################fin def fechas
      #Fecha actual y fecha a rellenar seràn iguales a lo que devuelve def fechas!
      fecha_actual,fecha2=fechas()


      #Si llegamos aqui es que hay tareas pendientes.
      #Rellenamos la actual y seguimos !
      rellenar()
      while fecha2 <= fecha_actual:
         dia_siguiente()
         rellenar()
         fecha_actual,fecha2=fechas()
         
         # fecha_actual=date.today()
         # fecha_evaluar=driver.find_element(By.XPATH,'/html/body/form/div[1]/div[2]/div[7]/div[1]/label/a[2]')
         # if fecha_evaluar:
         #    print("Encontrado")
         #    valor=fecha_evaluar.get_attribute('href')
         #    print(valor)
         #    start_index = valor.find('data=')


         #    end_index = valor.find('&', start_index)


         #    data_value = valor[start_index + 5:end_index]
         #    print(data_value)


         #    fecha=data_value[0:4]+str("-")+data_value[4:6]+str("-")+data_value[6:]
         #    print(fecha)

         #    fecha2=datetime.strptime(fecha,'%Y-%m-%d')
         #    print(fecha2)
         #    print(type(fecha2))

         #    fecha2=fecha2.date()
         #    print(type(fecha2))

         #    fecha_actual=date.today()
      if fecha2==fecha_actual:
         
         sms1=f"Las horas se han rellenado hasta el dia {fecha_actual} "
         
      else:
         sms1=f"Las horas se han rellenado hasta el dia {fecha2}"
      pywhatkit.sendwhatmsg_instantly('+tu numero',sms1,20,True)
      driver.close()
      """
      Escribir con pyautogui!
      pyautogui.write("2")
      """

except:
   print("Errore durante il processo, avvisare!")
   sms="È avvenuto un errore durante il processo del qbid! Revisare!"
   pywhatkit.sendwhatmsg_instantly('+tu numero',sms,20,True)#Envia whatsapp Alerta
   driver.close()
