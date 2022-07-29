package july2022

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Day17() {
	// use the ReadFile from ioutil package to read the entire file in memory
	data, err := ioutil.ReadFile("/Users/rosie/Documents/Job/Wiley_GoLang/learning/july2022/flatland01.txt")
	if err != nil {
		// feedback message in case of error
		fmt.Println(err)
	}

	// convert the file contents to a string and display it
	fmt.Print(strings.ToUpper(string(data)))
	fmt.Println("")
	//You can edit it with strings. Text documents can act like a string.

	//os panic, forces the program to stop.
	// _, err = os.Create("/erronous_path/file.txt")
	// if err != nil {
	// 	panic(err)
	// }
	//Read Bytes from a File
	f, err := os.Open("C:/Users/rosie/Documents/Job/Wiley_GoLang/learning/july2022/flatland01.txt")
	if err != nil {
		fmt.Println(err)
	}

	//  create a slice of bytes
	b1 := make([]byte, 5)
	data1, err1 := f.Read(b1)

	// feedback message in case of error; otherwise nil
	if err != nil {
		fmt.Println(err1) // print error due to reading from the file. nil is no errors
	}

	// display the slice
	fmt.Printf(string(b1[:data1]))

	// close the file after completing the operations
	f.Close()
	// We use the os library, rather than the io/ioutils The os library includes file management tools.
	// We retrieve the first five bytes of data from the file. Because this is a text file, those bytes are characters in the file.
	// We save the retrieved bytes to a slice and we display the contents of the slice.
	// We close the file after completing the operations.
	f, err = os.Open("/Users/rosie/Documents/Job/Wiley_GoLang/learning/july2022/flatland01.txt")
	if err != nil {
		fmt.Println(err)
	}

	// skip the first 100 bytes
	s, err := f.Seek(100, 0)
	if err != nil {
		fmt.Println(err)
	}

	// display the offset
	fmt.Println(s)

	// read 5 bytes starting from the offset
	data = make([]byte, 20)
	n, err := f.Read(data)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Bytes read", n)
	fmt.Println("Reading starting from byte", s, ":", string(data[:n]))
	//reads from the end
	s, err = f.Seek(100, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	n, err = f.Read(data)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Bytes read", n)
	fmt.Println("Reading starting from byte", s, ":", string(data[n:]))
	// close the file
	f.Close()
	//The following example uses the bufio.NewReader function to create a buffered reader. We then use the Peek function to see the first five characters in the data file.
	f, err = os.Open("/Users/rosie/Documents/Job/Wiley_GoLang/learning/july2022/flatland01.txt")
	br := bufio.NewReader(f) // create a buffered reader
	data, err = br.Peek(5)   // read 5 bytes
	if err != nil {
		fmt.Println(err)
	}
	// display the peeked data
	fmt.Printf(string(data))
	/*
		// create a slice of bytes (utf-8 code) from an input string
		data := []byte("Hello, world!")

		fmt.Println(data) // display the slice

		err := ioutil.WriteFile("/Users/username/Documents/golang/fileio/new_file.txt", data, 0644)
		if err != nil {
			fmt.Println(err)
		}
		new_file, err := ioutil.ReadFile("/Users/username/Documents/golang/fileio/new_file.txt")

		// feedback message in case of error
		if err != nil {
			fmt.Println(err)
		}
		// convert the file contents to a string and display it
		fmt.Print(string(new_file))
		//The the slice data includes only UTF-8 byte values to represent the characters in the original string.
		//The WriteFile function creates a file with the specified path if the file does not already exist. It will overwrite an existing file.
			//In the function, we specify data as the source, so Go will write the values from data into the new file.
			//We also include 0644 as the permission setting on the file. It will grant read/write access to the current user, but read-only access for any other user.
		//We then read the new file. Note that the UTF-8 bytes are reconverted to string characters when we read the file.
	*/
	//In the last example, we used the WriteFile function to both create a new file and write data to that file in the same operation. In some cases, we simply want to create a new file that we can add data to later. We can use the os.Create function for this purpose.
	/*
		data = "Hello, world!"
		fmt.Println("data string:", data)

		// create a new file
		f, err = os.Create("/Users/username/Documents/golang/fileio/another_file.txt")

		// display the new file
		fmt.Println(err)
		fmt.Println("new file:", f)


		// write the string to the new file
		n, err := f.WriteString(data)
		fmt.Println(err)
		fmt.Println("characters in file:", n)

		// close the file
		f.Close()
	*/
	/*
			data :="Hello, world!!!"
		   fmt.Println("original string:", data)

		   // create a new file
		   f, err := os.Create("/Users/username/Documents/golang/fileio/another_file.txt")
		   fmt.Println(err)

		   // create a buffered writer that we can use to write data to the new file
		   bw := bufio.NewWriter(f)

		   // write the data to the buffer writer
		   n, err := bw.WriteString(data)

		   // display the number of bytes written
		   fmt.Println("bytes written:", n)

		   // flush flushes/submits the data to the underlying io.Writer
		   bw.Flush()

		   newFile, err := ioutil.ReadFile("/Users/username/Documents/golang/fileio/another_file.txt")
		   // feedback message in case of error
		    fmt.Println(err)
		   // convert the file contents to a string and display it
		   fmt.Print("file contents:", string(newFile))

		   f.Close()

		    //We create a string.
		    //We create a new file.
		    //We create a buffered writer that will write data to the new file.
		    //We display the number of characters written to the buffer.
		    //We use flush to transfer the data from the buffer to the new file.
		    //We display the updated contents of the new file.

		//Note in particular that no data is added to the new file until we execute the Flush function.
	*/
}
func IoAssign() {
	/*
		1)Create Text file(Documants containing summary of context like movie review ,either in possitive or negetive )

	*/
	data := `With no apologies to Joss Whedon, Zack Snyder, or anyone else, it took bringing in a bunch of animal characters, the writers of "The LEGO Batman Movie," and shifting to animation to give us the best theatrical Justice League movie we've had to date.`
	f, err := os.Create("/Users/rosie/Documents/Job/Wiley_GoLang/learning/july2022/movieReview.txt")

	// display the new file
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("new file:", f)

	// write the string to the new file
	n, err := f.WriteString(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("characters in file:", n)

	// close the file
	f.Close()
	reading, err := ioutil.ReadFile("/Users/rosie/Documents/Job/Wiley_GoLang/learning/july2022/movieReview.txt")
	if err != nil {
		fmt.Println(err)
	}
	pFile, err := os.Open("/Users/rosie/Documents/Job/Wiley_GoLang/learning/july2022/positive-words.txt")
	if err != nil {
		panic("File not found")
	}
	scanner := bufio.NewScanner(pFile)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	pFile.Close()
	positive := 0
	negative := 0
	//removes puncuation
	replacer := strings.NewReplacer(",", "", ".", "", ";", "")
	reading = []byte(replacer.Replace(string(reading)))
	words := strings.Fields(strings.ToLower(string(reading)))
	for _, word := range words {
		for _, line := range lines {
			if line == word {
				positive++
			}
		}
	}
	fmt.Println(positive)
	nFile, err := os.Open("/Users/rosie/Documents/Job/Wiley_GoLang/learning/july2022/negative-words.txt")
	if err != nil {
		panic("File not found")
	}
	scanner = bufio.NewScanner(nFile)
	scanner.Split(bufio.ScanLines)
	lines = nil
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	nFile.Close()
	for _, word := range words {
		for _, line := range lines {
			if line == word {
				negative++
			}
		}
	}
	goodReview := float32(positive) / float32(len(words))
	badReview := float32(negative) / float32(len(words))
	if goodReview > badReview {
		goodReview *= 100
		fmt.Println("This is a positive review. ", goodReview, "% of the review had positive words")
	} else if badReview > goodReview {
		badReview *= 100
		fmt.Println("This is a negative review. ", badReview, "% of the review had positive words")
	}
	/*
		2)Read the text file and depending upon the words available in the Document give a conclusion about review (Possitive/negetive)
		in terms of % of possitive/negetive
		3)Give high priority to Awesome than Good , and then batter and so on
	*/
}
