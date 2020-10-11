package secret

import "os"

// WriteToFile - writes config to file
func WriteToFile(data []byte, fileName string) {
	f, err := os.Create(fileName)
	handleErr(err)

	defer f.Close()

	_, err = f.Write(data)
	handleErr(err)
}

func handleErr(e error) {
	if e != nil {
		panic(e)
	}
}
