package hangman





func (hangman *HANGMAN) Start() {
	if hangman.wordIsGood {
		hangman.wordlist()
	} else {
		hangman.Stop()
	}
	
}

func(hangman *HANGMAN) Stop() {
	// si la touche ctrl et q est pressé
	//on sauvegarde
	//on passe hangman.IsRunning = false ce qui arrête run()
	// save 
}

