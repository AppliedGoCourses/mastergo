// Time To Practice: Maps

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/pkg/errors"
)

func count(s string, m map[string]int) {

	// TODO:
	// * split s into words (hint: use the strings package)
	// * iterate over the words
	// * for each word, add it to the map and/or increase its counter by one
	//
	// Tip: trim each word using strings.Trim(<wordvariable>, " \t\n\"'.,:;?!()-")
	// Tip: turn each word into lowercase before counting - use the strings package for this.

}

func printCount(m map[string]int) {

	// TODO:
	// * iterate over the map variable m
	// * for each word in m, if its count is greater than 1,
	//   print the word and its count

}

func main() {
	// If the first argument is a valid file name/path,
	// open that file and read it line by line.
	// Else read the test text at the end of this file.
	in := bufio.NewReader(strings.NewReader(testText))
	if len(os.Args) > 1 {
		f, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Println("Error opening file", os.Args[1], "\n", errors.WithStack(err))
			os.Exit(1)
		}
		in = bufio.NewReader(f)
	}

	wordList := map[string]int{} // remember this is not a nil map but rather an existing map with zero elements. No need to check for nil before using it

	for more := true; more; {
		line, err := in.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Failed reading next line\n", errors.WithStack(err))
		}
		count(line, wordList)
	}
	printCount(wordList)

}

var testText = `So they began solemnly dancing round and round Alice, every now and
then treading on her toes when they passed too close, and waving their
forepaws to mark the time, while the Mock Turtle sang this, very slowly
and sadly:--

 '"Will you walk a little faster?" said a whiting to a snail.
 "There’s a porpoise close behind us, and he’s treading on my tail.

 See how eagerly the lobsters and the turtles all advance!
 They are waiting on the shingle--will you come and join the dance?

 Will you, won’t you, will you, won’t you, will you join the dance?
 Will you, won’t you, will you, won’t you, won’t you join the dance?

 "You can really have no notion how delightful it will be
 When they take us up and throw us, with the lobsters, out to sea!"
  But the snail replied "Too far, too far!" and gave a look askance--
 Said he thanked the whiting kindly, but he would not join the dance.

 Would not, could not, would not, could not, would not join the dance.
 Would not, could not, would not, could not, could not join the dance.

 '"What matters it how far we go?" his scaly friend replied.
 "There is another shore, you know, upon the other side.
 The further off from England the nearer is to France--
 Then turn not pale, beloved snail, but come and join the dance.

 Will you, won’t you, will you, won’t you, will you join the dance?
 Will you, won’t you, will you, won’t you, won’t you join the dance?"'`
