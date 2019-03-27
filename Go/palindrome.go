package main

import "os"
import "fmt"

func main() {
    args := os.Args[1:]
    inputString := args[0]
    fmt.Println("Input string : ", inputString)
    reverseString := reverse(inputString)
    fmt.Println("Reverse string : ", reverseString)
    isPalindrome := checkPalindrome(inputString, reverseString)
    fmt.Println("Input string is a palindrome", isPalindrome)
}

func reverse(s string) string {
  runes := []rune(s)
  for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
    runes[i], runes[j] = runes[j], runes[i]
  }
  return string(runes)
}

func checkPalindrome(inputString string, reverseString string) bool {
  if inputString == reverseString {
    return true
  } else {
    return false
  }
}
