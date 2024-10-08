package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)




func (hangman *HANGMAN) Run() {
	hangman.wordIsGood = false
	 for hangman.IsRunning {
		hangman.hangman()
		hangman.readletter()
		hangman.testword()
		
	 }			
}

func (hangman *HANGMAN) readletter() { 
	fmt.Println(hangman.motIconnu) //affiche dans la console la valeur de motIconnu
	var Reader = bufio.NewReader(os.Stdin) //créé un reader
	String,_:= Reader.ReadString('\n') //lit la ligne et la stocke dans la variable string
	String = strings.TrimSpace(String) //supprime les espaces ou les retours a la ligne au début et à la fin de la chaine
	if len(String) > 1 {
		if String == hangman.Mot { //si le mot écrit est le bon,
			hangman.win()          //tu as gagné
		} else {
			
			hangman.erreur += 2 //sinon erreur +2
		}
		
	} else if len(String) == 1 {
			hangman.lettre = String
			hangman.UsedLetter = append(hangman.UsedLetter, hangman.lettre) //ajoute dans le tableau les lettres déjà utilisées
			hangman.testLetter()
	} 
		
}

func (hangman *HANGMAN) testLetter() {
	hangman.lettreIsGood = false
	for l := range hangman.MotAdeviner {
		
		if hangman.lettre == hangman.MotAdeviner[l] {
			hangman. lettreIsGood = true
			hangman.motIconnu[l] = hangman.lettre
			fmt.Println(hangman.motIconnu)
		}
	}
	if !hangman.lettreIsGood {
		hangman.erreur++
		
	}
	if hangman.erreur >= 9 {
		hangman.gameOver()
	}
	
}
func (hangman *HANGMAN) testword() {
	hangman.wordIsGood = true
	for l := range hangman.motIconnu {
		if !(hangman.motIconnu[l] == hangman.MotAdeviner[l]) {
			hangman.wordIsGood = false
			break
		}
	}
	if hangman.wordIsGood {
		hangman.win()
	}
	if hangman.erreur >= 9 {
		hangman.gameOver()
	}
}


func (hangman *HANGMAN) win() {
		fmt.Println(`
 		__     __          __          ___       _ 
 		\ \   / /          \ \        / (_)     | |
	 	 \ \_/ /__  _   _   \ \  /\  / / _ _ __ | |
 	 	  \   / _ \| | | |   \ \/  \/ / | | '_ \| |
  		   | | (_) | |_| |    \  /\  /  | | | | |_|
  	 	   |_|\___/ \__,_|     \/  \/   |_|_| |_(_)
                                           
                                           

		`)
		fmt.Println("fin du jeu vous avez gagner le mot était bien", hangman.Mot)
	hangman.IsRunning = false
	
}

func (hangman *HANGMAN) gameOver() {
	fmt.Println(` 
		   ____                          ___                 
		  / ___| __ _ _ __ ___   ___    / _ \__   _____ _ __ 
		 | |  _ / _' | '_ ' _ \ / _ \  | | | \ \ / / _ \ '__|
		 | |_| | (_| | | | | | |  __/  | |_| |\ V /  __/ |
	    	  \____|\__,_|_| |_| |_|\___|   \___/  \_/ \___|_|
                                                                                              

	`)

	fmt.Println("fin du jeu le mot était : ", hangman.Mot)
	hangman.IsRunning = false
	
}