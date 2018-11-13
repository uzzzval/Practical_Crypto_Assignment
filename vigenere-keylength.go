package main

import "os"
import "strings"
import (
   "fmt"
   "io/ioutil"
)


func main(){
  arguments := os.Args;
  plaintext_file_name := arguments[1];

  //Putting the File Content in a variable
  file_content, err := ioutil.ReadFile(plaintext_file_name);
  if err != nil {
   fmt.Print(err);
  }

  //Converting content of file to string
  str := string(file_content);

  //Putting the string into a character array
  contentArray := strings.Split(str, "");

  //Setting max length for key 
  maxKeyLength := 20;

  //Length of String
  length := len(contentArray);
  
  highestIOC := float32(0);
  highestIndex := 0;

  for j:=1; j<=maxKeyLength ; j++{
    totalioc := iterate(j, length, contentArray);
    averageIOC := float32(totalioc/float32(j));

  	
  	if(highestIOC == 0.0){
  		highestIOC = averageIOC;
  		highestIndex = j;
  	} else{
  		if(highestIOC < averageIOC){
  			highestIOC = averageIOC;
  			highestIndex = j;
  		}
  	}
    
  }

  	//fmt.Print("Length of the Key:");
  	fmt.Println(highestIndex);
}

func iterate(j int, length int, contentArray []string)(totalIoc float32){
	for l:=0; l<j; l++{
		sequence :="";
		for i:=0; i<length; i++ {
  			if((j*i)+l < length){
  			sequence = sequence + strings.ToUpper(contentArray[(j*i)+l]);
  		}
  	 }	
  	 iocOfSequence := indexOfIncidence(sequence);
     totalIoc = totalIoc + iocOfSequence;
	}  
	return totalIoc;  
}

 func indexOfIncidence (sequence string)(ioc float32){
   mapping := map[string]int {"A" :0, "B" :0, "C" :0, "D" :0, "E" :0, "F" :0, "G" :0, "H" :0, "I" :0, "J" :0, "K" :0, "L" :0, "M" :0, "N" :0, "O" :0, "P" :0, "Q" :0, "R" :0, "S" :0, "T" :0, "U" :0, "V" :0, "W" :0, "X" :0, "Y" :0, "Z" :0};
   sequenceArray := strings.Split(sequence, "");
   	 for i:=0; i<len(sequence)-1; i++{
   	 	if value, ok := mapping[sequenceArray[i]]; ok{
       	    value = mapping[sequenceArray[i]];
       	    mapping[sequenceArray[i]]=value +1;
       } else{
       	   mapping[sequenceArray[i]]=1;
       }
   	 } 
   	 
   	 numerator :=0;
   	 denominatorCount :=0;
   	 for k,v := range mapping{
   	 	k = k;
   	 	product :=0;
   	 	  if(v>1){
   	 	  	product = v*(v-1);
   	 	  }else{
   	 	  	product = v;
   	 	  }
          numerator = numerator + product;
          denominatorCount = denominatorCount+v;
   	 }
   	 denominator := float32(denominatorCount*(denominatorCount-1));
   	 numer := float32(numerator);
   	 ioc = numer/denominator;
   	
   	 return ioc;
}
