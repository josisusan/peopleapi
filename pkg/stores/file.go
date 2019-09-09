package stores

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type FileStore struct {
	Name string
}

func (fs FileStore) Write(content string) (string, error) {
	fmt.Println(content)
	f, err := os.OpenFile(fs.Name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	if err != nil {
		fmt.Printf(err.Error())
		log.Fatal(err)
		return "", err
	}

	uid := strconv.FormatInt(generateUUID(), 10)
	newContent := uid + "," + content
	if _, err := f.Write([]byte(newContent)); err != nil {
		log.Fatal(err)
		return "", err
	}

	return newContent, nil
}

func (fs FileStore) Read() ([][]string, error) {
	f, err := os.OpenFile(fs.Name, os.O_CREATE|os.O_RDONLY, 0644)
	defer f.Close()

	if err != nil {
		fmt.Printf(err.Error())
		log.Fatal(err)
		return [][]string{}, err
	}

	data, err := csv.NewReader(f).ReadAll()
	if err != nil {
		fmt.Printf(err.Error())
		log.Fatal(err)
		return [][]string{}, err
	}

	return data, nil
}

func generateUUID() int64 {
	return time.Now().Unix()
}
