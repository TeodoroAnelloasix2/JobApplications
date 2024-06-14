#!/usr/bin/python3

import pygame

pygame.init()
screen= pygame.display.set_mode((1580,920))
clock = pygame.time.Clock()
running=True
pygame.display.set_caption("Naruto")
dt= 1

player_pos=pygame.Vector2(screen.get_width()/2, screen.get_height()/2)

while running:

    for event in pygame.event.get():
        if event.type == pygame.QUIT:
            running= False

    screen.fill("purple") #fill background with purple color

    bg_image= pygame.image.load("./fondos/villa_oculta_hoja.jpg").convert()
    screen.blit(bg_image,(0,0))
    #pygame.display.update()
    #pygame.draw.circle(screen,"red",player_pos,40)
    pain=pygame.image.load("./personajes/traced-pain1.png")#.convert() #Personaje
    pain=pygame.transform.scale(pain,(220,220)) #Set size personaje

    keys= pygame.key.get_pressed()
    if keys[pygame.K_UP]:
        player_pos.y -=10 *dt
    if keys[pygame.K_DOWN]:
        player_pos.y +=10*dt
    if keys[pygame.K_RIGHT]:
        player_pos.x +=10 * dt
    if keys[pygame.K_LEFT]:
        player_pos.x -=10 *dt

    screen.blit(pain,player_pos)
    pygame.display.flip()

    clock.tick(60)/1000

pygame.quit()