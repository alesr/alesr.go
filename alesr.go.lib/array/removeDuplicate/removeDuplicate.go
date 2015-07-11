package main

import "fmt"

func main() {

   words := []string{"A", "B", "C", "C", "D", "E", "E", "F", "F", "F", "F", "F", "A", "B", "F", "G", "H", "H", "I"}
   arr := words[:1]
   Loop:
      for i := 1; i < len(words); {
         for j := 0; j < len(arr); {
            if words[i] != arr[j] {
               j++
            } else {
               i++
               continue Loop
            }
         }
         arr = append(arr, words[i])
         i++
      }
   fmt.Println(arr)
}
