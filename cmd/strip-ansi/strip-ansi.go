package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"

	"golang.org/x/text/unicode/norm"
)

func backspaceHelper(text []byte, p int) []byte {
	if p > 0 {
		if text[p] == '' {
			if len(text) > p {
				text = append(text[0:p-1], text[p+1:]...)
			} else {
				text = text[0 : p-1]
			}
		}
		text = backspaceHelper(text, p-1)
	} else {
		if text[p] == '' {
			if len(text) == 1 {
				text = []byte{}
			} else {
				text = text[1:]
			}
		}
	}
	return text
}

func backspace(text []byte) []byte {
	return backspaceHelper(text, len(text)-1)
}
func main() {
	nonAscii := regexp.MustCompile("[[:^ascii:]]")
	notEscape := regexp.MustCompile("[[1@.]")
	text, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	// wc := norm.NFC.Writer(w)
	// defer wc.Close()
	// // write as before...
	// If you have a small string and want to do a quick conversion, you can use this simpler form:

	// norm.NFC.Bytes(b)
	text = norm.NFC.Bytes(text)

	text = backspace(text)
	text = bytes.Replace(text, []byte{''}, []byte{}, -1)

	if false {
		x := string(norm.NFC.Bytes(text))
		fmt.Println(x)
	}
	// fmt.Println(string(norm.NFC.Bytes(text)))
	strip := nonAscii.ReplaceAllLiteralString(string(text), "")
	strip = notEscape.ReplaceAllLiteralString(strip, "")
	fmt.Println(strip)
	// text, err = ansi.Strip(text)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

/*
string outputString;
bool inEscape = false;
bool justEnteredEscape = false;
foreach(c; inputString) {
   if(justEnteredEscape) {
       justEnteredEscape = false;
       if(c == '[')
            inEscape = true;
       else {
            // NOTE: this is actually likely wrong but prolly good enough
            outputString ~= c;
       }
   else if(inEscape) {
       if(c >= 'A')
           inEscape = false;
       // otherwise we want to skip this character, since it is part of e.g. a color sequence
   } else if(c == '\033') {
       justEnteredEscape = true;
   } else {
       if(c == 8) {
   if i > 0 {
      i--
}
// continue; // backspace in string
}
       if(c == 0) continue; // and so on for whatever else you don't want....

      text = text[0:i]
      outputString ~= c;
   }
}

// outputString should be ok now


*/
