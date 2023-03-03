package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/ozbekburak/estimation-assistant/chatgpt"
)

const promptsDir = "prompts"

func main() {
	var (
		prompt string
		err    error
	)

	// Define flags
	promptFlag := flag.String("prompt", "", "if you have a pre written prompt, you can use this flag to use it")
	// Parse command-line arguments
	flag.Parse()

	if len(*promptFlag) > 0 {
		prompt, err = readFromPromptFile(*promptFlag)
		if err != nil {
			log.Printf("Failed to read prompt file error: %v", err)
		}
	} else {
		prompt = completePrompt()
		if _, err := savePrompt(prompt); err != nil {
			log.Printf("Failed to save prompt to file error: %v", err)
		}
	}
	estimation, err := chatgpt.AskChatGPT(prompt)
	if err != nil {
		log.Printf("Failed to get answer from chatgpt error: %v", err)
		return
	}

	fmt.Println("Here is your estimation to plan your sprint!")
	fmt.Println(estimation[0])
}

// savePrompts saves the prompt to a file to be used for -p flag
func savePrompt(prompt string) (*os.File, error) {
	// fileName: transcript-english-20210101120000.txt
	timestamp := time.Now().Format("20060102150405")
	fileName := filepath.Join(promptsDir, fmt.Sprintf("prompt-%s.txt", timestamp))

	// Open a file for writing, create it if it doesn't exist, and append to it
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Printf("Failed to open file error: %v", err)
		return nil, err
	}
	defer file.Close()

	// Create a bufio.Writer for the file
	writer := bufio.NewWriter(file)

	words := strings.Split(prompt, " ")
	lineLength := 150
	line := ""
	for _, word := range words {
		if len(line)+len(word)+1 > lineLength {
			// Write the current line to the file and start a new one
			fmt.Fprintln(writer, line)
			line = word
		} else {
			// Add the word to the current line
			if line == "" {
				line = word
			} else {
				line += " " + word
			}
		}
	}
	// Write the last line to the file
	fmt.Fprintln(writer, line)

	// Flush the buffer to ensure that all data is written to the file
	writer.Flush()

	return file, nil
}

func readFromPromptFile(path string) (string, error) {
	textData, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Failed to read keywords file error: %v", err)
		return "", err
	}

	return string(textData), nil
}

func completePrompt() string {
	basePrompt := `I want you to estimate how much time it takes to complete a task. Please consider the duration of one workday is 6 hours. Please tell me your estimated duration using days. Please consider there will be unforeseen challenges and there will be other work to do. 
	So, adding 1 or 2 days as a buffer to your estimation would be nice. No explanations. Just list the estimated time breakdown for each task. Please add explanations to breakdowns that contain coding jobs. Do not add "In summary" line. Just total days to complete. 
	Please suggest a scrum story point (using Fibonacci) for the total days to complete the task at the end of the estimation. `

	// Create a new scanner to read user input
	scanner := bufio.NewScanner(os.Stdin)

	// Ask the user for their task
	fmt.Print("Task description: ")
	scanner.Scan()
	task := scanner.Text()
	basePrompt += "Task description: " + task

	// Ask the user for their task complexity
	fmt.Print("Task complexity: ")
	scanner.Scan()
	taskComplexity := scanner.Text()
	basePrompt += " Task complexity: " + taskComplexity

	// Ask the user for their task
	fmt.Print("Skill level: ")
	scanner.Scan()
	skillLevel := scanner.Text()
	basePrompt += " Skill level: " + skillLevel

	// Ask the user for their task
	fmt.Print("Resources: ")
	scanner.Scan()
	resources := scanner.Text()
	basePrompt += " Resources: " + resources

	// Ask the user for their task
	fmt.Print("Timeline: ")
	scanner.Scan()
	timeline := scanner.Text()
	basePrompt += " Timeline: " + timeline

	// Ask the user for their task
	fmt.Print("Risks and Issues: ")
	scanner.Scan()
	risks := scanner.Text()
	basePrompt += " Risks: " + risks

	s := regexp.MustCompile(`\s+`).ReplaceAllString(basePrompt, " ")
	basePrompt = strings.TrimSuffix(s, "\n")

	return basePrompt
}
