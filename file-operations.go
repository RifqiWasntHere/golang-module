package file_operations

import (
	"bufio"
	"io"
	"os"
)

func CreateNewFile(name string, content string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	defer file.Close()
	file.WriteString(content)

	return nil
}

// Buatkan fungsi untuk read file, WElllllll
func ReadFile(filename string) (string, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}

	defer file.Close()

	var content string
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		content += string(line)
	}
	return content, err
}

// Bangg kalo mau nambahin konten ke existing file gimana bang ? siyaaaaaap
func AppendToFile(filename string, content string) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	defer file.Close()

	file.WriteString(content)

	return nil
}
