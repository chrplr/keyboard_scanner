#! /usr/bin/env python
# Time-stamp: <2024-12-07 11:42:55 christophe@pallier.org>

""" basic HID (keyboard, mouse, ...) event monitor """ 

#    Copyright (C) 2024 Christophe Pallier
#
#    This program is free software: you can redistribute it and/or modify
#    it under the terms of the GNU General Public License as published by
#    the Free Software Foundation, either version 3 of the License, or
#    (at your option) any later version.
#
#    This program is distributed in the hope that it will be useful,
#    but WITHOUT ANY WARRANTY; without even the implied warranty of
#    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
#    GNU General Public License for more details.
#
#    You should have received a copy of the GNU General Public License
#    along with this program.  If not, see <https://www.gnu.org/licenses/>.

import os
os.environ['PYGAME_HIDE_SUPPORT_PROMPT'] = "hide"
import pygame

event_names = { pygame.KEYUP: "keyup", 
                pygame.KEYDOWN: "keydown", 
                pygame.MOUSEBUTTONUP: "buttonup", 
                pygame.MOUSEBUTTONDOWN: "buttondown",
            }


class App(object):
    def __init__(self, screen_width=1280, screen_height=800, font_size=40):
        pygame.init()
        self.screen = pygame.display.set_mode((screen_width, screen_height))
        pygame.display.set_caption('Monitor keyboard & mouse buttons')
        self.screen_width = screen_width
        self.screen_height = screen_height
        self.font_size = font_size
        self.font = pygame.font.Font(None, font_size)
        self.ytext = 10
        

    def display(self, txt):
        text = self.font.render(txt.replace("\t", "  "), True, pygame.Color('white'))
        self.screen.blit(text, (100, self.ytext))

        self.ytext = self.ytext + self.font_size
        if self.ytext + self.font_size > self.screen_height:
            self.screen.fill(pygame.Color('black'))
            self.ytext = 10
            
        pygame.display.flip()

    def main_loop(self):
        txt = "Time\tDelta\tKeyCode\tChar\tEvent"
        print(txt)
        self.display(txt)
        
        t0 = pygame.time.get_ticks()
        last_event_time = 0
        running = True
        while running:
            for event in pygame.event.get():
                txt = ""
                now = pygame.time.get_ticks() - t0
                delta = now - last_event_time
        
                if event.type == pygame.QUIT: 
                    running = False
                if event.type in [pygame.KEYDOWN, pygame.KEYUP]:
                    txt = "\t".join(map(str, [now, delta, event.key, pygame.key.name(event.key), event_names[event.type]]))
                if event.type in [pygame.MOUSEBUTTONDOWN, pygame.MOUSEBUTTONUP]:
                    txt = "\t".join(map(str, [now, delta, event.button, event.button, event_names[event.type]]))

                if txt != "":
                    print(txt)
                    self.display(txt)
                    last_event_time = now
                    

        pygame.quit()


app = App()
app.main_loop()
