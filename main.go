package main

import (

	hangman "hangman/GAME"
)

 func main() {
    var h hangman.HANGMAN
	h.Flag()
	h.DrawDisplay()
	h.Init()
	h.Start()
	h.Run()
	h.Stop()
	
 }