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

        self.image= pygame.image.load("/home/teodoro_root/python/pygame/data/images/entities/traced-pain1.png").convert()
        self.image.set_colorkey((0,0,0))
        #self.image.fill("purple")
        #pain=pygame.transform.scale(pain,(220,220))

        #self.image.set_colorkey(color)Hace desaparecer el color especificado

        self.image=pygame.transform.scale(self.image,(150,150)).convert() #Cargamos imagen y con .convert() mejoramos el rendimiento optimizando la imagen
        self.rect=self.image.get_rect()#obtener posicion
        
        self.radius=30
        #pygame.draw.circle(self.image,"green",self.rect.center,self.radius)
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
        
        if teclas[pygame.K_SPACE]:
            jugador.disparo()
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
    def disparo(self):
        bala= Disparos(self.rect.centerx,self.rect.top + 55)
        balas.add(bala)

#Clase para el enemigo
class Enemigos(pygame.sprite.Sprite):
    def __init__(self):
        super().__init__()
        #/home/teodoro_root/python/pygame/data/images/entities/enemy/run
        self.image=pygame.image.load("/home/teodoro_root/python/pygame/data/images/entities/enemy/run/1.png").convert()
        self.image=pygame.transform.scale(self.image,(100,100))

        self.rect=self.image.get_rect()
        self.image.set_colorkey((0,0,0))
        self.radius=10
        pygame.draw.circle(self.image,"red",self.rect.center,self.radius)
        
        #Posicion aleatoria en eje x,y (restar el ancho y el alto del enemigo para que no sobresalga)
        self.rect.x=random.randrange(ANCHO -self.rect.width)
        self.rect.y=random.randrange(ALTO - self.rect.height)

        #Insert velocidad inicial a los enemigos para los dos ejes
        self.velocity_x=2
        self.velocity_y=2
    def update(self):

        #Determina movimiento
        self.rect.x += self.velocity_x
        self.rect.y +=self.velocity_y
        #Controlamos que no supere los limites de la ventana
        #limitar izquierda, derecha
        if self.rect.left < 0:
            self.velocity_x+=1
        if self.rect.right > ANCHO:
            self.velocity_x-=1
        #limitar Abajo, Arriba
        if self.rect.bottom > ALTO: 
            self.velocity_y-=1
        if self.rect.top < 0:
            self.velocity_y+=1

class Disparos(pygame.sprite.Sprite):
    def __init__(self,x,y):
        super().__init__()
        self.image=pygame.transform.scale(pygame.image.load("/home/teodoro_root/python/pygame/data/images/particles/leaf/01.png").convert(),(10,10))
        self.image.set_colorkey((0,0,0))
        self.rect=self.image.get_rect()
        self.rect.bottom=y
        self.rect.centerx= x
    def update(self):
        self.rect.y -=25
        if self.rect.bottom < 0:
            self.kill()
        
clock = pygame.time.Clock()


pygame.init()

sprite1=pygame.sprite.Group()
enemy2=pygame.sprite.Group()
balas=pygame.sprite.Group()
# for i in range(random.randrange(5) +1):
#     enemigo=Enemigos()
   
#     sprite1.add(enemigo)


enemigo=Enemigos()

sprite1.add(enemigo)
enemy2.add(enemigo)#Insertamos en enemigo que esta instanciado a Enemigos()

#Por cada vez que repetimos estas lineas aparecen 1 enemigo.
#Podemos poden en bucle

# 
# sprite1.add(enemigo2)
# for i in range(random.randrange(5) +1):
#     enemigo2=Enemigos()
#     sprite1.add(enemigo2)
#     enemy2.add(enemigo)
    

jugador=Jugador()
sprite1.add(jugador)

while True:
    clock.tick(FPS)
    for event in pygame.event.get(): #Bucle for que recoge  pone en cola todos los eventos de input del usuario que juega
        if event.type == pygame.QUIT:
            sys.exit()
    
    sprite1.update()    #Actualizamos el estado despues de trabajar con los eventos recogidos en el bucle for 
    enemy2.update()  #Actualizamos lo que pasa en update de enemy2
    balas.update()

    colision_p=pygame.sprite.spritecollide(jugador,enemy2,False,pygame.sprite.collide_circle) #Mirar bien spritecollide Basicamente implementa colisiones 

    colision=pygame.sprite.groupcollide(enemy2,balas,False,True)  
    if colision_p or colision:
        enemigo.image=pygame.image.load("/home/teodoro_root/python/pygame/data/images/clouds/cloud_1.png").convert()
        enemigo.velocity_y+=10
    elif enemigo.rect.top > ALTO:
        enemigo.kill()   

    pantalla.fill((255, 255, 255))  # Filling the screen with black color

    enemy2.draw(pantalla) #despues de actualizar hay que pintar en pantalla
    sprite1.draw(pantalla)# Drawing the sprite
    balas.draw(pantalla)
    pygame.display.flip()

#pygame.quit()