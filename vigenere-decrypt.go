package main

import "os"
import "strings"
import (
   "fmt"
   "io/ioutil"
   "log"
   "regexp"
)


func main(){
  arguments := os.Args;
  encipherment_key := arguments[1];
  plaintext_file_name := arguments[2];

  isAlpha := regexp.MustCompile(`^[A-Z]`).MatchString

  if(!isAlpha(encipherment_key)){
    fmt.Println("The Key should contain all uppercase letters.");
    return;
  }

  //Declaring the Alphabet Array
  alphabetArray := []string{"A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z"}; 

  //Putting the File Content in a variable
  file_content, err := ioutil.ReadFile(plaintext_file_name);
  if err != nil {
   fmt.Print(err);
  }

   file, err := os.Open(plaintext_file_name)

   stat, err := file.Stat()
   if err != nil {
       return
   }

   if(stat.Size()>100){
      fmt.Println("The File Size is greater than 100Kb");
//      return;
   }
  
  //Converting content of file to string
  str := string(file_content);
  // Make a Regex to say we only want
  reg, err := regexp.Compile("[^a-zA-Z]+")
    if err != nil {
        log.Fatal(err)
  }
  processedString := reg.ReplaceAllString(str, "")
  
  processedString = strings.ToUpper(processedString);

  //Putting the string into a character array
  contentArray := strings.Split(processedString, "");

  //Putting the Key string to array
  keyArray := strings.Split(encipherment_key, "");


  //Creating Key Stream Array
  length := len(contentArray);

  streamArray := make([]string,length);
  finalArray := make([]string,length);
  j := 0;
  for i :=0 ; i<len(contentArray) ; i++ {
       streamArray[i] = keyArray[j];
       j++;
       if j == len(keyArray) {
           j = 0;
       }
  } 
  

  //Iterating Over the Content Array and Stream Array and Fetching the new Encrypted Alphabets
  for i:=0 ; i< len(contentArray) ; i++{
    alphabetMap := map[string]int{"A":1,"B":2,"C":3,"D":4,"E":5,"F":6,"G":7,"H":8,"I":9,"J":10,"K":11,"L":12,"M":13,"N":14,"O":15,"P":16,"Q":17,"R":18,"S":19,"T":20,"U":21,"V":22,"W":23,"X":24,"Y":25,"Z":26};
    positionContentAlphabet := 0;
    positionStreamAlphabet :=0;
    newPosition :=0;
    
    characterContent := contentArray[i];
    characterStream := streamArray[i];

    positionContentAlphabet = alphabetMap[characterContent];
    positionStreamAlphabet = alphabetMap[characterStream];

    difference := (positionContentAlphabet) - (positionStreamAlphabet);
   
    if difference<26 && difference>=0 {
      newPosition = difference ;
    
    }else {
      newPosition = (26 + difference)  ;
    }
    
      finalArray[i] = alphabetArray[newPosition];
    
  }
  encryptedString := strings.Join(finalArray,"");
  fmt.Println(encryptedString);
}
