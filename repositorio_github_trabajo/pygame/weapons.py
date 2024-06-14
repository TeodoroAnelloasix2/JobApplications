#!/usr/bin/ython3

from typing import Any
import pygame, math

import random
import constantes

class Weapon():
    def __init__(self,image,imagen_bala):
        self.image_original=image
        self.imagen_bala=imagen_bala
        self.angulo=0
        self.imagen=pygame.transform.rotate(self.image_original,self.angulo) #Funcion que cambia angulo de las imagenes param: imagen, angulo
        self.forma=self.image_original.get_rect()
        self.dispara=False
        self.ultimo_disparo=pygame.time.get_ticks()

    
    def update(self,personaje):
        disparo_cooldown=constantes.COOLDOWN_BALAS

        bala=None
        self.forma.center=personaje.forma.center
        

        if personaje.flip==False:
            #Si el personaje no està girado
            self.forma.x +=(personaje.forma.width//2)
            self.rotar_arma(False)
            
            if pygame.mouse.get_pressed()[0] and self.dispara==False and ((pygame.time.get_ticks()-self.ultimo_disparo) >= disparo_cooldown): #click mouse[0]izquierdo,[1]ruedita,[2]derecho
                bala=Bullet(self.imagen_bala,self.forma.centerx,self.forma.centery,self.angulo)
                self.dispara=True
                self.ultimo_disparo=pygame.time.get_ticks()
            if pygame.mouse.get_pressed()[0]==False:
                self.dispara=False
            #return bala
        
        if personaje.flip==True:
            self.forma.x-=(personaje.forma.width//2)
            self.rotar_arma(True)

            #Mover pistola con mouse si el personaje estàa girando pero esta vez eje x negativo!
            
            # mouse_pos=pygame.mouse.get_pos()
            # distancia_x=-mouse_pos[0]-self.forma.centerx
            # distancia_y=mouse_pos[1]-self.forma.centery

            # self.angulo=math.degrees(math.atan2(distancia_y,distancia_x))
            if pygame.mouse.get_pressed()[0] and self.dispara==False and ((pygame.time.get_ticks()-self.ultimo_disparo) >= disparo_cooldown): #click mouse[0]izquierdo,[1]ruedita,[2]derecho
                bala=Bullet(self.imagen_bala,self.forma.centerx,self.forma.centery,self.angulo)
                self.dispara=True
                self.ultimo_disparo=pygame.time.get_ticks()
            if pygame.mouse.get_pressed()[0]==False:
                self.dispara=False
        
        mouse_pos=pygame.mouse.get_pos()
        distancia_x=mouse_pos[0]-self.forma.centerx
        distancia_y=-(mouse_pos[1]-self.forma.centery)
        self.angulo=math.degrees(math.atan2(distancia_y,distancia_x))
        return bala
            # print(self.angulo)


  


    def rotar_arma(self,rotar):
        if rotar:
            imagen_flip=pygame.transform.flip(self.image_original,True,False)
            self.imagen=pygame.transform.rotate(imagen_flip,self.angulo)

        else:
            imagen_flip=pygame.transform.flip(self.image_original,False,False)
            self.imagen=pygame.transform.rotate(imagen_flip,self.angulo)
    
      
    def dibujar(self,interfaz):
        self.imagen=pygame.transform.rotate(self.imagen,self.angulo)
        interfaz.blit(self.imagen,self.forma)
        pygame.draw.rect(interfaz,constantes.COLOR_ARMA,self.forma,1)


class Bullet(pygame.sprite.Sprite):
    def __init__(self,image,x,y,angle):
        pygame.sprite.Sprite.__init__(self)
        self.imagen_original=image
        self.angulo=angle
        self.image=pygame.transform.rotate(self.imagen_original,self.angulo)
        self.rect=self.image.get_rect()
        self.rect.center=(x,y)

        #Velocidad bala

        self.delta_x=math.cos(math.radians(self.angulo))*constantes.VELOCIDAD_BALA
        self.delta_y=-(math.sin(math.radians(self.angulo))*constantes.VELOCIDAD_BALA)

    def update(self,lista_enemigos):
        damage=0
        pos_damage=None
        self.rect.x+=self.delta_x
        self.rect.y+=self.delta_y

        if self.rect.top > constantes.ALTO or self.rect.bottom < 0  or self.rect.right <0  or self.rect.left > constantes.ANCHO:

            self.kill()


        #Restar vida a los enemigos por cada colision de cada bala    
        for enemigo in lista_enemigos:
            
            if enemigo.forma.colliderect(self.rect):
                damage=15+random.randint(-7,7)
                pos_damage=enemigo.forma
                enemigo.energia-=damage
                self.kill()
                break
        
        return damage,pos_damage



    def dibujar(self,interfaz):
        print("bala")
        interfaz.blit(self.image,(self.rect.centerx,self.rect.centery)) #Para ajustar de donde sale la bala hacer operaciones en centery o centerx self.rect.centery- self.image.get_height