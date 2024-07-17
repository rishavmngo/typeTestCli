package utils

func CheckForEnd(currentWord, lenOfText int) bool {

	return currentWord == lenOfText

}

func CheckForTypo(input, str string) bool {

	if len(input) <= len(str) && input == str[:len(input)] {
		return false
	} else {
		return true
	}

}
