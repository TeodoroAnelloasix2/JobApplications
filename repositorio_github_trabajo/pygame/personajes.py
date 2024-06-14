#!/usr/bin/python3

import pygame
import constantes
class Personaje():
    def __init__(self,x,y,animaciones,energia):
        self.energia=energia
        self.vivo=True
        self.flip=False
        self.animaciones=animaciones
        #Imagene de la animacion que se esta mostrando
        self.frame_index=0
        self.update_time=pygame.time.get_ticks() #Almacena en milisegundos el tiempo transcurrido
        self.image=animaciones[self.frame_index]
        # self.forma = pygame.Rect(0,0,constantes.ALTO_PERSONAJE,constantes.ANCHO_PERSONAJE)  #pygame.rect(coordenadasx, y ,pixeles lado1, pixeles lado2)
        self.forma=self.image.get_rect()
        self.forma.center= (x,y) #Mueve el personaje en las coorrdenadas entregadas(x,y)



    def movimiento(self,delta_x,delta_y):

        if delta_x < 0:
            self.flip= True
        if delta_x > 0:
            self.flip=False

        #Limites para el jugador. Evitar que se salga de la ventana de juego
        if self.forma.top < 0:
            self.forma.top=0
        
        if self.forma.bottom > constantes.ALTO:
            self.forma.bottom=constantes.ALTO
        
        if self.forma.right > constantes.ANCHO:
            self.forma.right=constantes.ANCHO
        
        if self.forma.left < 0:
            self.forma.left=0
   


        self.forma.x=self.forma.x+ delta_x  #self.forma.x eje x posicion actual, + delta_x que define cuantos pixel se tiene que deplazar
        self.forma.y=self.forma.y+ delta_y  #La cordenada y va a ser  igual a y actual mas y nueva
        ejex=self.forma.x=self.forma.x+ delta_x
        #print(f"x: {ejex}")
        ejey=self.forma.y=self.forma.y+ delta_y 
        #print(f"y: {ejey}")

    def update(self):

        #Comprobar si el pèersonaje sigue vivo
        if self.energia <=0:
            self.energia=0
            self.vivo=False
            


        cooldown_animacion=100
        self.image=self.animaciones[self.frame_index]

        if pygame.time.get_ticks() - self.update_time >= cooldown_animacion:
            self.frame_index+=1
            self.update_time=pygame.time.get_ticks()
        if self.frame_index >=len(self.animaciones):
            self.frame_index=0
    def dibujar(self,interfaz):
        imagen_flip=pygame.transform.flip(self.image,self.flip,False)#Invierte imagenes(imagen,True|Flase (eje x) , True|False eje y) parametros imagen a invertir, True(invierte en eje x), True (invierte en eje y)
        interfaz.blit(imagen_flip,self.forma)
        pygame.draw.rect(interfaz,constantes.COLOR_PERSONAJE,self.forma,1)
