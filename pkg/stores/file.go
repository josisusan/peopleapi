package stores

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
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

// Update method to update the person record according to uid.
func (fs FileStore) Update(uid string, record map[string]string) error {
	f, err := ioutil.ReadFile(fs.Name)
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	lines := strings.Split(string(f), "\n")

	for i, line := range lines {
		if strings.Contains(line, uid) {
			lines[i] = fmt.Sprintf("%s,%s,%s,%s", uid, record["Name"], record["Age"], record["Gender"])
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(fs.Name, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}

func generateUUID() int64 {
	return time.Now().Unix()
}
