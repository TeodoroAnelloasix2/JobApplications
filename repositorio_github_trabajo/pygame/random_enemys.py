#!/usr/bin/python3


import sys
import pygame 
import random


ANCHO=800
ALTO=600
FPS=30

pantalla= pygame.display.set_mode((ANCHO,ALTO))
# 
# running=True
pygame.display.set_caption("Vamonosss")
# dt= 1

class Jugador(pygame.sprite.Sprite):  #Classe de pygame para los sprite

    def __init__(self):
        super().__init__()

        # self.image = pygame.Surface((200, 200))  # Corrected typo

        self.image= pygame.image.load("/home/teodoro_root/python/pygame/data/images/entities/player.png")
        #self.image.fill("purple")
        #pain=pygame.transform.scale(pain,(220,220))

        #self.image.set_colorkey(color)Hace desaparecer el color especificado

        self.image=pygame.transform.scale(self.image,(150,150)).convert() #Cargamos imagen y con .convert() mejoramos el rendimiento optimizando la imagen
        self.rect=self.image.get_rect()#obtener posicion
        
        
        self.rect.center=(ANCHO//2,ALTO) #Modificar para cambiar posicion inicial jugador(ANCHO,ALTO)

        #Velocidad inicial x / y ejes
        self.velocity_x=0
        self.velocity_y=0

    def update(self):

        #Para el personaje si no pasa nada 
        self.velocity_x=0
        self.velocity_y=0
        #Mantiene las teclas pulsadas

        teclas=pygame.key.get_pressed()
        #Movemos el pesonaje
        if teclas[pygame.K_LEFT]:
            self.velocity_x= -10

        if teclas[pygame.K_RIGHT]:
            self.velocity_x= 10
        
        if teclas[pygame.K_UP]:
            self.velocity_y= -10
        if teclas[pygame.K_DOWN]:
            self.velocity_y= 10
        
        # self.rect.y -= 10   
        # if self.rect.bottom < 0:
        #     self.rect.top=ALTO
        
        #Actualizamos posicion del jugador sumando a la posicion de su rectangulo sobre los ejes la variabel velocidad
        self.rect.x += self.velocity_x
        self.rect.y +=self.velocity_y

        #Controlamos que no supere los limites de la ventana
        #limitar izquierda, derecha
        if self.rect.left < 0:
            self.rect.left=0
        if self.rect.right > ANCHO:
            self.rect.right=ANCHO
        #limitar Abajo, Arriba
        if self.rect.bottom > ALTO:
            self.rect.bottom=ALTO
        if self.rect.top < 0:
            self.rect.top=0


#Clase para el enemigo
class Enemigos(pygame.sprite.Sprite):
    def __init__(self):
        super().__init__()
        #/home/teodoro_root/python/pygame/data/images/entities/enemy/run
        self.image=pygame.image.load("/home/teodoro_root/python/pygame/data/images/entities/enemy/run/5.png").convert()
        self.image=pygame.transform.scale(self.image,(150,150))
        self.rect=self.image.get_rect()
        self.image.set_colorkey((0,0,0))
        #Posicion aleatoria en eje x,y (restar el ancho y el alto del enemigo para que no sobresalga)
        self.rect.x=random.randrange(ANCHO -self.rect.width)
        self.rect.y=random.randrange(ALTO - self.rect.height)

clock = pygame.time.Clock()
pygame.init()
sprite1=pygame.sprite.Group()



enemigo=Enemigos()
sprite1.add(enemigo)

#Por cada vez que repetimos estas lineas aparecen 1 enemigo.
#Podemos poden en bucle

# enemigo2=Enemigos()
# sprite1.add(enemigo2)

for i in range(random.randrange(5) +1):
    enemigo2=Enemigos()
    sprite1.add(enemigo2)
    

jugador=Jugador()
sprite1.add(jugador)
while True:
    clock.tick(FPS)
    for event in pygame.event.get(): #Bucle for que recoge  pone en cola todos los eventos de input del usuario que juega
        if event.type == pygame.QUIT:
            sys.exit()
    
    sprite1.update()    #Actualizamos el estado despues de trabajar con los eventos recogidos en el bucle for 
    
    pantalla.fill((0, 0, 0))  # Filling the screen with black color
    sprite1.draw(pantalla)# Drawing the sprite
    pygame.display.flip()

#pygame.quit()