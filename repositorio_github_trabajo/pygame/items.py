#!/usr/bin/python3
import pygame.sprite


#La clase sprite necesitan las coordenadas, tipo de item

class Item(pygame.sprite.Sprite):
    def __init__(self,x,y,item_type,animacion_list):
        pygame.sprite.Sprite.__init__(self)
        self.item_type= item_type     #0 =monedas 1= pocion
        self.animacion_list=animacion_list
        self.frame_index=0

        self.update_time=pygame.time.get_ticks()
        self.image=self.animacion_list[self.frame_index]
        self.rect=self.image.get_rect()
        self.rect=(x,y)
    
    def update(self):
        cooldown_animacion= 50 #Tiempo en milisegundos, 1000 = 1 segundo
        self.image=self.animacion_list[self.frame_index]
        
        if pygame.time.get_ticks() - self.update_time > cooldown_animacion:
            self.frame_index+=1
            self.update_time=pygame.time.get_ticks()
        
        if self.frame_index >= len(self.animacion_list):
            self.frame_index=0