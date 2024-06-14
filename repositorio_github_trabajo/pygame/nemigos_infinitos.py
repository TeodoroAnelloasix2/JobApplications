#!/usr/bin/python3


import sys
import pygame 
import random




ANCHO=1280
ALTO=800
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
        self.image=pygame.transform.scale(pygame.image.load("/home/teodoro_root/python/pygame/data/images/particles/leaf/01.png").convert(),(30,30))
        self.image.set_colorkey((0,0,0))
        self.rect=self.image.get_rect()
        self.rect.bottom=y
        self.rect.centerx= x
    def update(self):
        self.rect.y -=25
        if self.rect.bottom < 0:
            self.kill()


class Meteoritos(pygame.sprite.Sprite):
    def __init__(self):
        super().__init__()

        #generar numero aleatorio
        self.img_aleatoria=random.randrange(3)

        if self.img_aleatoria==0:
            self.image=pygame.transform.scale(pygame.image.load("/home/teodoro_root/python/pygame/data/images/entities/enemy/duende1.png").convert(),(100,100))
            self.radius=50#Implementar collisiones,declarar radius mitad medida imagen

        if self.img_aleatoria==1:
            self.image=pygame.transform.scale(pygame.image.load("/home/teodoro_root/python/pygame/data/images/entities/enemy/duende1.png").convert(),(50,50))
            self.radius=20
        if self.img_aleatoria==2:
            self.image=pygame.transform.scale(pygame.image.load("/home/teodoro_root/python/pygame/data/images/entities/enemy/duende1.png").convert(),(25,25))
            self.radius=11
        
        self.image.set_colorkey((0,0,0))
        self.rect=self.image.get_rect()#Obtener rectangulo entidad
        self.rect.x=random.randrange(ANCHO-self.rect.width)#Coordenada x posiciona random
        self.rect.y =-self.rect.width#Restar ancho para evitar que se genere dentro de la pantalla
        self.velocidad_y=random.randrange(1,7)

    def update(self):
        self.rect.y += self.velocidad_y #Se mueve en y eje
        if self.rect.top > ALTO:
            self.rect.x=random.randrange(ANCHO - self.rect.width)
            self.rect.y= -self.rect.width
            self.velocidad_y= random.randrange(1,10)


clock = pygame.time.Clock()


pygame.init()

sprite1=pygame.sprite.Group() 
enemy2=pygame.sprite.Group()
balas=pygame.sprite.Group()
# for i in range(random.randrange(5) +1):
#     enemigo=Enemigos()
   
#     sprite1.add(enemigo)

meteoritos=pygame.sprite.Group()  #crear grupo de sprite 
#Crear varios a la vez
contador_meteoritos=10
for x in range(contador_meteoritos):
# for x in range(contador_meteoritos):
    meteorito=Meteoritos()#1 entidad
    meteoritos.add(meteorito)#Por cada entidad, insertar entidad al grupo meteoritos



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
    meteoritos.update()
    
    colision_p=pygame.sprite.spritecollide(jugador,meteoritos,False,pygame.sprite.collide_circle) #Mirar bien spritecollide Basicamente implementa colisiones 
    
    disparo=pygame.sprite.groupcollide(meteoritos,balas,False,True)  
    

    if disparo:
        meteorito.velocidad_y+=100
        contador_meteoritos-=1
        if meteorito.rect.top > ALTO:
            meteorito.kill()
            
    if colision_p:
        jugador.kill() #Muere Jugador
    # if colision_p or colision:
    #     enemigo.image=pygame.image.load("/home/teodoro_root/python/pygame/data/images/clouds/cloud_1.png").convert()
    #     enemigo.velocity_y+=10
    # elif enemigo.rect.top > ALTO:
    #     enemigo.kill() 

    bg_image= pygame.image.load("villa_oculta_hoja.jpg").convert()
    pantalla.fill((255, 255, 255))  # Filling the screen with black color
    pantalla.blit(bg_image,(0,0))
    enemy2.draw(pantalla) #despues de actualizar hay que pintar en pantalla
    sprite1.draw(pantalla)# Drawing the sprite
    balas.draw(pantalla)
    meteoritos.draw(pantalla)
    pygame.display.flip()

#pygame.quit()