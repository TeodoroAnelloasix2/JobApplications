#!/usr/bin/python3

import pygame,sys
from textos import DamageText
import constantes #Importar un archivo
import personajes # Para importar todo el archivo
from weapons import Weapon
from personajes import Personaje #importar solo una clase from archivo import class
import os

from items import Item
pygame.init()

#Contar elementos de una carpeta

def contar_elementos(directorio):

    return len(os.listdir(directorio))

#print(contar_elementos("/home/teodoro_root/python/pygame/2_curso_copia/recursos/Anime_warrior/Warrior/enemigo"))

#Devuelve una lista con elementos del directorio
def nombres_carpetas(directory):
    return os.listdir(directory)



def escalar_img(image,scale):
    w= image.get_width()
    h= image.get_height()

    nueva_imagen=pygame.transform.scale(image,(w*scale,h*scale))

    return nueva_imagen

ventana= pygame.display.set_mode((constantes.ANCHO,constantes.ALTO)) #(ancho,alto)(width,height)

pygame.display.set_caption("Prueba1")

#Energia

life_bar1=pygame.image.load("/home/teodoro_root/Scrivania/Sprites para juegos/life_bar/lifebar_1.png").convert_alpha()
life_bar1=escalar_img(life_bar1,constantes.ESCALA_VIDA)
life_bar2=pygame.image.load("/home/teodoro_root/Scrivania/Sprites para juegos/life_bar/lifebar_2.png").convert_alpha()
life_bar2=escalar_img(life_bar2,constantes.ESCALA_VIDA)
life_bar3=pygame.image.load("/home/teodoro_root/Scrivania/Sprites para juegos/life_bar/lifebar_3.png").convert_alpha()
life_bar3=escalar_img(life_bar3,constantes.ESCALA_VIDA)
life_bar4=pygame.image.load("/home/teodoro_root/Scrivania/Sprites para juegos/life_bar/lifebar_4.png").convert_alpha()
life_bar4=escalar_img(life_bar4,constantes.ESCALA_VIDA)
life_bar5=pygame.image.load("/home/teodoro_root/Scrivania/Sprites para juegos/life_bar/lifebar_5.png").convert_alpha()
life_bar5=escalar_img(life_bar5,constantes.ESCALA_VIDA)
life_bar6=pygame.image.load("/home/teodoro_root/Scrivania/Sprites para juegos/life_bar/lifebar_6.png").convert_alpha()
life_bar6=escalar_img(life_bar6,constantes.ESCALA_VIDA)



#Fuentes
font=pygame.font.Font("/home/teodoro_root/python/pygame/2_curso_copia/recursos/fonts/InflateptxBase-ax3da.ttf",25)

animaciones=[]
for i in "1234567":
    img=pygame.image.load(f"/home/teodoro_root/python/pygame/2_curso/recursos/Knight_player/andar_hacia_derecha{i}.png")
    img=escalar_img(img,constantes.ESCALA_PERSONAJE)
    animaciones.append(img)



#Vida con baremos de valores

def vida_jugador():

    energia=jugador.energia
    
    if energia >=100 and energia <=120:
        ventana.blit(life_bar6,(20, 5))
    elif energia >=80 and energia < 100:
        ventana.blit(life_bar5,(20, 5))
    elif energia >=60 and energia < 80:
        ventana.blit(life_bar4,(20,5))
    elif energia >=40 and energia < 60:
        ventana.blit(life_bar3,(20,5))
    elif energia >=20 and energia < 40:
        ventana.blit(life_bar2,(20,5))
    elif energia >=5 and energia < 20:
        ventana.blit(life_bar2,(20,5))
  
    
"""
def vida2():
implementar vida con corazones llenos medios y vacios. Tanto corazones como rango(numero)-->range(4) 4 corazones. Dibujar medios y vacios segun vida que queda.
dividir 100 entre numero del range ejemplo--> range(5) 100//5=20 --> if jugador.energia >=((i+1)*20<---(aqui resultado)):Esto es si tenemos medidores como corazones llenos, medios y vacios 

for i in range(4):
    dibujo_vida_mitad=False
    if jugador.energia >=((i+1)*25):
        ventana.blit(corazon_lleno,(5+1*50, 5))
    
    elif jugador.energia % 25 >=0 and dibujo_vida_mitad ==False:
        ventana.blit(corazon_lleno,(5+1*50, 5))
        dibujo_vida_mitad = True
    
        else:
        ventana.blit(corazon_vacio,((5+i*50,5)))



"""

#Dibuar ajedrezz fondo pantalla
def bibujar_grid():

    for x in range(30):
        pygame.draw.line(ventana,constantes.NEGRO,(x*constantes.TILE_SIZE,0),(x*constantes.TILE_SIZE,constantes.ALTO))
        pygame.draw.line(ventana,constantes.NEGRO,(0,x*constantes.TILE_SIZE),(constantes.ANCHO,x*constantes.TILE_SIZE))



jugador=Personaje((constantes.ANCHO//2),(constantes.ALTO*0.6),animaciones,120)


#Enemigos

directorio_enemigos="/home/teodoro_root/python/pygame/2_curso_copia/recursos/Anime_warrior/Warrior/enemigo"
tipo_enemigos=nombres_carpetas(directorio_enemigos)
#print(tipo_enemigos)
animaciones_enemigos=[]

for eni in tipo_enemigos:
    lista_temp=[]
    ruta_temp=f"/home/teodoro_root/python/pygame/2_curso_copia/recursos/Anime_warrior/Warrior/enemigo/{eni}"
    num_animaciones=contar_elementos(ruta_temp)
    for i in range(num_animaciones):
        img_enemigo=pygame.image.load(f"{ruta_temp}//{eni}_{i+1}.png").convert_alpha()
        img_enemigo=escalar_img(img_enemigo,constantes.ESCALA_ENEMIGO)
        lista_temp.append(img_enemigo)
    
    animaciones_enemigos.append(lista_temp)

#print(animaciones_enemigos)

#Crear enemigos con classe personaje
    
goblin=Personaje(200,200,animaciones_enemigos[0],100)

goblin2=Personaje(500,500,animaciones_enemigos[0],100)


lista_enemigos=[]

lista_enemigos.append(goblin)
lista_enemigos.append(goblin2)





#print(lista_enemigos)
#Armas

imagen_pistola=pygame.image.load("/home/teodoro_root/python/pygame/2_curso/recursos/Knight_player/mitralladora_simple2.gif").convert()
#imagen_pistola.set_colorkey((255,255,255))

imagen_pistola=escalar_img(imagen_pistola,0.25)
imagen_balas=pygame.image.load("/home/teodoro_root/python/pygame/2_curso_copia/recursos/bala1_fuego1.png").convert_alpha()



coin_images=[]
ruta_img="/home/teodoro_root/python/pygame/2_curso_copia/items/Super Pixel Objects Sample/outline_dark/coin_large"

num_coin_images=contar_elementos(ruta_img)
for i in range(num_coin_images):
    img=pygame.image.load(f"{ruta_img}/coin_large_{i+1}.png")
    img=escalar_img(img,0.8)
    coin_images.append(img)

#print(f"numero monedas: {num_coin_images}")
#imagen_balas.set_colorkey((255,255,255))
imagen_balas=escalar_img(imagen_balas,0.3)
pistola=Weapon(imagen_pistola,imagen_balas)


pocion_roja=pygame.image.load("/home/teodoro_root/python/pygame/2_curso_copia/items/Super Pixel Objects Sample/outline_dark/potion_small/potion_C_red_full.png").convert_alpha()
pocion_roja=escalar_img(pocion_roja,1.5)
#instanciar texto
grupo_damage_text=pygame.sprite.Group()




#Instanciar grupo texto
# damage_text=DamageText(100,240,"3",font,constantes.COLOR_TEXTO)
# grupo_damage_text.add(damage_text)

#Puede haber màas balas a la vez en pantalla y para eso creamos un grupo
grupo_balas=pygame.sprite.Group()
#print(grupo_balas)
# player_image=pygame.image.load("/home/teodoro_root/python/pygame/2_curso/recursos/Knight_player/andar_hacia_derecha1.png")
# player_image=escalar_img(player_image,constantes.ESCALA_PERSONAJE)
#jugador=Personaje(50,50,player_image)(x,y)

grupo_items=pygame.sprite.Group()

coin=Item(50,50,0,coin_images)

potion=Item(300,20,1,[pocion_roja])

grupo_items.add(coin)
grupo_items.add(potion)


#Variables movimiento jugador
mover_arriba=False
mover_abajo= False
mover_izquierda=False
mover_derecha=False
reloj=pygame.time.Clock()
run= True

while run:
    reloj.tick(constantes.FPS)
    ventana.fill(constantes.COLOR_BG) #A cada ronda rellenamos ventana para borrar lo que habia en antiguo frame
    bibujar_grid()
    #Calcular movimiento jugador
    delta_x=0
    delta_y=0

    if mover_izquierda==True:
        delta_x=-constantes.VELOCIDAD_PERSONAJE
    if mover_derecha==True:
        delta_x=constantes.VELOCIDAD_PERSONAJE
    if mover_abajo==True:
        delta_y=constantes.VELOCIDAD_PERSONAJE
    if mover_arriba==True:
        delta_y=-constantes.VELOCIDAD_PERSONAJE
    
    jugador.movimiento(delta_x,delta_y) 
    #Actualizar estado jugador
    jugador.update()

    #Actualizar estado enemigos
    for ene in lista_enemigos:
        ene.update()
        print(ene.energia)
        


    bala=pistola.update(jugador)
    if bala:
        grupo_balas.add(bala)
    for bala in grupo_balas:
        damage,pos_damage=bala.update(lista_enemigos)

        if damage >0:
            damage_text=DamageText(pos_damage.centerx,pos_damage.centery,"-"+" "+str(damage),font,constantes.COLOR_TEXTO)
            grupo_damage_text.add(damage_text)

    grupo_damage_text.update()

    grupo_items.update()
    #Dibujar items

    jugador.dibujar(ventana)
    pistola.dibujar(ventana)
    #jugador.energia-=1
    vida_jugador()

    #Dibujar enemigos

    for ene in lista_enemigos:
        ene.dibujar(ventana)
    #dibujar balas

    for bala in grupo_balas:
        
        bala.dibujar(ventana)

    grupo_damage_text.draw(ventana)
    grupo_items.draw(ventana)
    for event in pygame.event.get():
        if event.type== pygame.QUIT:
            run = False
            #sys.exit()
        
        if event.type==pygame.KEYDOWN: #Reconoce si se pressiona una tecla
            if event.key==pygame.K_UP:            #Que tecla presionamos?
                mover_arriba=True
                #print("Arriba")
            if event.key==pygame.K_DOWN:            
                mover_abajo=True
                #print("Abajo")
            if event.key==pygame.K_LEFT:           
                mover_izquierda=True
                #print("Izquierda")
            if event.key==pygame.K_RIGHT:            
                mover_derecha=True
                #print("derecha")
            #Mover jugador
        if event.type==pygame.KEYUP:
            if event.key==pygame.K_UP:
                mover_arriba=False
               # print("Arriba False")
            if event.key==pygame.K_DOWN:
                mover_abajo=False
                #print("Abajo False")
            if event.key==pygame.K_RIGHT:
                mover_derecha=False
                #print("derecha false")
            if event.key==pygame.K_LEFT:
                mover_izquierda=False
                #print("izquierda False")
        
          
    pygame.display.update()
pygame.quit()