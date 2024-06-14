#!/usr/bin/python3

from typing import Any
import sys
import pygame 
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
        self.image = pygame.Surface((200, 200))  # Corrected typo

       
        self.image.fill("purple")

        self.rect=self.image.get_rect()
        self.rect.center=(ANCHO//2,ALTO//2) 

    def update(self):


        ###############################
        # self.rect.x -=10  va de derecha a izquierda
        # if self.rect.right < 0:
        #     self.rect.left= ANCHO
        ##############################################
        # self.rect.x +=10  Va de izquierda hacia derecha
        # if self.rect.left > ANCHO:
        #     self.rect.right= 0
        ##################################### va de arriba hacia abajo.
        self.rect.y -= 10   
        if self.rect.bottom < 0:
            self.rect.top=ALTO
        ######################################
        # self.rect.y += 10  va de arriba hacia abajo. 
        # if self.rect.top > ALTO:
        #     self.rect.bottom=0

clock = pygame.time.Clock()
pygame.init()
sprite1=pygame.sprite.Group()
jugador=Jugador()
sprite1.add(jugador)
while True:
    clock.tick(FPS)
    for event in pygame.event.get(): #Bucle for que recoge  pone en cola todos los eventos de input del usuario que juega
        if event.type == pygame.QUIT:
            sys.exit()
    
    sprite1.update()    #Actualizamos el estado despues de trabajar con los eventos recogidos en el bucle for 
    
    pantalla.fill((0, 0, 0))  # Filling the screen with black color
    sprite1.draw(pantalla)  # Drawing the sprite
    pygame.display.flip()

#pygame.quit()
