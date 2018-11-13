	package main

	import (
		"fmt"
		"io/ioutil"
		"math"
		"os"
		"regexp"
		"sort"
		"strings"
		"strconv"
	)


	func main() {
        
        //Reading File
		data, err := ioutil.ReadFile(os.Args[1])
		if err != nil {	
		}

		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Println(err);
		}

		defer file.Close()

		var sequence []string
        //Converting the string to uppercase
		message := strings.ToUpper(strippingCharacters(string(data)))
		keylength,errConv := strconv.Atoi(os.Args[2]);
		if errConv != nil{
			fmt.Println(errConv);
		}

		sequence = sequences(string(message), keylength)
		chiSquareMap := make(map[float64]string)

		for i:= 0; i < keylength; i++ {
			var chiSqList []float64
			for k := 0; k < len(alphabet); k++ {
				tempChisquare := calcChisquare(sequence[i], string(alphabet[k]))
				chiSquareMap[tempChisquare] = string(alphabet[k])
				chiSqList = append(chiSqList, tempChisquare)
			}
			sort.Float64s(chiSqList)
			fmt.Print(chiSquareMap[chiSqList[0]]);

		}
	    fmt.Println("");
		var keys[]float64
		for k := range chiSquareMap {
			keys = append(keys, k)
		}
		sort.Float64s(keys)
	}

	func sequenceGenerator (encryptedtext string)([]string){
		max_keylength := 25;
		lengthEncrypted := len(encryptedtext);
		var sequences[] string

		for k :=2; k < max_keylength; k++ {
			for j:= 0; j < k; j++ {
				var tempseq string = "";
				for i := j; i < lengthEncrypted; i +=k {
					tempseq = tempseq + string(encryptedtext[i])
				}
				sequences = append(sequences, tempseq)
				tempseq = ""
			}

		}
		return sequences
	}

    //Calculating the number of occurance of each alphabet 
	func alphabetFrequqncy (sequence string, character string) float64 {
		var count float64 = 0;
		for i:=0 ; i < len(sequence); i++ {
			if string(sequence[i]) == character {
				count ++;
			}
		}
		return count;
	}

    //Getting all the seqences
	func sequences ( encryptedtext string, assumed_key int) []string {
		var theSequences []string
		theSequences = sequenceGenerator(encryptedtext)
		startpos := 0
		var pulled_sequence []string
		for j:= 2; j < assumed_key; j++ {
			startpos = startpos + j
		}
		for i:= startpos; i < startpos + assumed_key; i ++ {
			pulled_sequence = append (pulled_sequence, theSequences[i])
		}
		return pulled_sequence
	}

    //Getting the character positions
	func characterToPosition(character string) int {
		var position int;
		for i := 0; i < len(alphabet); i++ {
			if character == string(alphabet[i]) {
				position = i;
			}
		}
		return position;
	}

    
	func characterize(p int) string {
		var character string;
		for i := 0; i < len(alphabet); i++ {
			if p == i {
				character = character + string(alphabet[i]);
			}
		}
		return character;
	}

    //Deciphering all the sequqnces
	func decrypt( sequence string, k string) string {
		var plain string;
		for i, j := 0, 0; i <len(sequence) && j <len(sequence); i, j = i+1, j+1 {
			var temp int;
			calc := characterToPosition(string(sequence[i])) - characterToPosition(string(k[j])) % 26;
			if math.Signbit(float64(calc)) == false {
				temp = calc;
			}else {
				temp = calc + 26;
			}
			plain = plain + characterize(temp);
		}
		return plain;
	}

     //Generating the Key
	func keyGenerator (keystring string, messagelength int) string {
		var newkey string;
		j := 0;
		if messagelength > len(keystring) {
			for i := 0; i < messagelength; i++ {
				if j == len(keystring){
					j = 0;
				}
				newkey =  newkey + string(keystring[j]);
				j++;
			}
		}
		return newkey;
	}

    //Calculating chi square
	func calcChisquare (sequqnce string, trychar string) float64 {
		newkey := keyGenerator(trychar, len(sequqnce));
		tempdecryp := decrypt(sequqnce, newkey);
		total_len := float64(len(tempdecryp));
		var Chisquare float64;

		for c:= 0; c < len(alphabet); c++ {
			count_trychar := alphabetFrequqncy(tempdecryp, string(alphabet[c]));
			ec := (englishFrequency(string(alphabet[c])) / 100) * total_len;
			Chisquare = Chisquare + math.Pow(count_trychar-ec, 2) / ec;
		}
	   return Chisquare;
	}

    //Calculating Frequqncy based on English frequency distribution
	func englishFrequency (character string) float64 {
		character = strings.ToUpper(character);
		freqMapping := map[string]float64{"A" :8.17, "B" :1.49, "C" :2.78, "D" :4.25, "E" :12.70, "F" : 2.23, "G" :2.02, "H" :6.09, "I" :7.00, "J" :0.15, "K" :0.77, "L" :4.03, "M" :2.41, "N" :6.75, "O" :7.51, "P" :1.93, "Q" :0.10, "R" :5.99, "S" :6.33, "T" :9.06, "U" :2.76, "V" :0.98, "W" :2.36, "X" :0.15, "Y" :1.97, "Z" :0.07};
		return freqMapping[character];
	}

    //Stripping all the characters from the input file content
	func strippingCharacters (content string) string {
		regex := regexp.MustCompile("[a-z]*[A-Z]*");
		check := regex.FindAllString(content,-1);
		content = strings.Join(check,"");
		content = strings.Replace(content, " ", "", -1);
		return content;
	}
    
	const alphabet string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
