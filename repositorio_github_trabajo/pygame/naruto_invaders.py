#!/usr/bin/python3

import pygame, sys, random,os  

ANCHO=1800
ALTO=900
FPS=30

pantalla=pygame.display.set_mode((ANCHO,ALTO))

name_game="NARUTO'S INVADERS"
pygame.display.set_caption(name_game)

pygame.mixer.init()
musica_juego=pygame.mixer.Sound("/home/teodoro_root/python/pygame/soundtracks/bluebird.mpga")
musica_juego.play(-1)
#pygame.mixer.music.set_volume(0.0)


muerte_enemigo=pygame.mixer.Sound("/home/teodoro_root/python/pygame/soundtracks/explosion.wav")
shinra_tensei=pygame.mixer.Sound("/home/teodoro_root/python/pygame/soundtracks/shinratensei1.mpga")
bijudama=pygame.mixer.Sound("/home/teodoro_root/python/pygame/soundtracks/bijudama.mpga")
class Jugador(pygame.sprite.Sprite):
    def __init__(self):
        super().__init__()
        self.image=pygame.image.load("/home/teodoro_root/python/pygame/data/images/entities/traced-pain1.png").convert()
        self.image.set_colorkey((0,0,0))
        self.image=pygame.transform.scale(self.image,(50,90)).convert()
        self.rect=self.image.get_rect()
        self.rect.center=(ANCHO//2,ALTO)
        self.pos_x=0
        self.pos_y=0
        self.cadencia=550  #tiempo en milisegundos
        self.ultimo_disparo=pygame.time.get_ticks()
        self.radius=50
        #pygame.draw.circle(self.image,"green",self.rect.center,self.radius)

    def update(self):
        self.pos_x=0
        self.pos_y=0
        teclas=pygame.key.get_pressed()

        if teclas[pygame.K_LEFT]:
            self.pos_x= -10

        if teclas[pygame.K_RIGHT]:
            self.pos_x= 10
        
        if teclas[pygame.K_UP]:
            self.pos_y= -10
        
        if teclas[pygame.K_DOWN]:
            self.pos_y=10


        if teclas[pygame.K_SPACE]:
            ahora=pygame.time.get_ticks()
            if ahora -self.ultimo_disparo > self.cadencia:
                bijudama.play()
                jugador.disparo()
                self.ultimo_disparo=ahora
            #shinra_tensei.play()

        self.rect.x+= self.pos_x
        self.rect.y+=self.pos_y

        if self.rect.left < 0:
            self.rect.left=0
        
        if self.rect.right > ANCHO:
            self.rect.right=ANCHO
        
        if self.rect.bottom > ALTO:
            self.rect.bottom=ALTO
        
        if self.rect.top < 400:
            self.rect.top=400
    def disparo(self):
        bala=Disparos(self.rect.center,self.rect.top)
        balas.add(bala)


class Disparos(pygame.sprite.Sprite):
    def __init__(self,x,y):
        super().__init__()
        self.image=pygame.transform.scale(pygame.image.load("/home/teodoro_root/python/pygame/balas_teo/bala_de_fuego.jpeg").convert(),(10,20))
        self.image.set_colorkey((255,255,255))
        self.rect=self.image.get_rect()
        self.rect.bottom=y
        self.rect.center=x
    def update(self):
        self.rect.y-=25
        if self.rect.bottom < 0:
            self.kill()

class Enemigos(pygame.sprite.Sprite):
    def __init__(self):
        super().__init__()
        self.image=pygame.image.load("/home/teodoro_root/python/pygame/data/images/entities/enemy/naruto1.jpg").convert()
        self.image.set_colorkey((255,255,255))
        self.image=pygame.transform.scale(self.image,(100,100))
        self.radius=50
        self.rect=self.image.get_rect()
        pygame.draw.circle(self.image,"red",self.rect.center,self.radius,1)
        self.rect.x=random.randrange(ANCHO-self.rect.width)
        self.rect.y=random.randrange(ALTO//2)
        self.velocidad_y=random.randrange(4,8)
        self.velocidad_x=random.randrange(3,9)
    
    
    def update(self):
        self.rect.y+=self.velocidad_y
        
        if self.rect.top > ALTO:
            self.rect.x=random.randrange(ANCHO-self.rect.width)
            self.rect.y=random.randrange(ALTO//2)
            self.velocidad_y=random.randrange(4,8)   

clock=pygame.time.Clock()
pygame.init()

running=True
#######################################
#Instanciar grupo, instancia singolo entidad, insertar entidad en grupo
sprites=pygame.sprite.Group()
jugador=Jugador()
sprites.add(jugador)
######################################
balas=pygame.sprite.Group()


enemys=pygame.sprite.Group()

for i in range(10):
    enemigo=Enemigos()
    enemigo.add(enemys)
#eventos=pygame.event.get()
puntos=0
consolas=pygame.font.match_font('consolas')
times=pygame.font.match_font('times')
arial=pygame.font.match_font('arial')
courier=pygame.font.match_font('courier') 

def muestra_texto(pantalla,fuente,texto,color,dimensiones,x,y):
    tipo_letra=pygame.font.Font(fuente,dimensiones,)
    superficie=tipo_letra.render(texto,True,color)
    rectangulo=superficie.get_rect()
    rectangulo.center=(x,y)
    pantalla.blit(superficie,rectangulo)

while running:

   

    clock.tick(FPS)
    for event in pygame.event.get():
        if event.type==pygame.QUIT:
            running=False
            sys.exit()
    bg_image= pygame.image.load("/home/teodoro_root/python/pygame/villa_oculta_hoja.jpg").convert()
    pantalla.fill((255,255,255))
    pantalla.blit(bg_image,(0,0))
    enemys.update()    
    sprites.update()
    balas.update()


    contacto_con_enemigo=pygame.sprite.spritecollide(enemigo,sprites,True,pygame.sprite.collide_circle)

    disparo=pygame.sprite.groupcollide(enemys,balas,True,True)
    # if disparo:
    #     enemigo.kill()

    for accierto in disparo.keys():
        muerte_enemigo.play()
        enemigo=Enemigos()
        enemys.add(enemigo)
        puntos+=10
    
    if contacto_con_enemigo:
        jugador.kill()
    
    enemys.draw(pantalla)
    balas.draw(pantalla)
    sprites.draw(pantalla)

    puntuacion=muestra_texto(pantalla,courier,str(puntos).zfill(6),(255,0,0),40,900,50)      
    pygame.display.flip()
    
