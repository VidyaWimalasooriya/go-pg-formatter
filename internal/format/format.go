package format

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func Format(path string, args ...string) error {
	// Find files need to be formatted
	filePaths, err := findFiles(path)
	if err != nil {
		return err
	}

	errChan := make(chan error)
	for _, filePath := range filePaths {
		go formatFile(filePath, errChan, args...)
	}

	var errs []string
	for i := 0; i < len(filePaths); i++ {
		if err := <-errChan; err != nil {
			errs = append(errs, err.Error())
		}
	}

	close(errChan)

	if len(errs) > 0 {
		var sb strings.Builder
		sb.WriteString("encountered the following errors:\n")
		for _, e := range errs {
			sb.WriteString("- ")
			sb.WriteString(e)
			sb.WriteString("\n")
		}
		return errors.New(sb.String())
	}

	return nil
}

func FormatSQLString(filePath string, args ...string) (str string, err error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}

	defer file.Close()

	// Read the entire file content
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	formattedData, err := formatContent(string(content), args...)
	if err != nil {
		return "", err
	}

	return formattedData, nil
}

func formatFile(filePath string, ch chan error, args ...string) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		ch <- err
		return
	}

	defer file.Close()

	// Read the entire file content
	content, err := io.ReadAll(file)
	if err != nil {
		ch <- err
		return
	}

	formattedData, err := formatContent(string(content), args...)
	if err != nil {
		ch <- err
		return
	}

	if err := updateFileWithFormattedData(filePath, formattedData); err != nil {
		ch <- err
		return
	}

	fmt.Printf("Format completed\n")
	ch <- nil
}

func formatContent(content string, args ...string) (string, error) {
	_, f, _, _ := runtime.Caller(0)
	toolDirectory := filepath.Join(filepath.Dir(f), "tools/pg_format")
	cmd := exec.Command("perl", append([]string{toolDirectory}, args...)...)

	// Set up input and output buffers
	var outBuffer, errBuffer bytes.Buffer
	cmd.Stdin = strings.NewReader(content)
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer

	// Run the command
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error running command: %v\nStderr: %s", err, errBuffer.String())
	}

	return outBuffer.String(), nil
}

func updateFileWithFormattedData(filePath string, formattedData string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the formatted data to the file
	_, err = file.WriteString(formattedData)
	if err != nil {
		return err
	}

	return nil
}
