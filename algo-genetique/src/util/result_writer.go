package util

import "io/ioutil"

// écrit la string content dans le fichier fileName
func WriteResult(filename string, content string) {
	byteArr := []byte(content)
	err := ioutil.WriteFile(filename, byteArr, 0644)

	// gestion de l'erreur
	if err != nil {
		panic(err)
	}
}