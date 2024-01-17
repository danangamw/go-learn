package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"danangamw.com/note/note"
	"danangamw.com/note/todo"
)

type saver interface {
	Save() error
}

/* type displayer interface {
	Display()
} */

// Embedded Interface
type outputtable interface {
	saver
	Display()
}

/* type outputtable interface {
	Save() error
	Display()
} */

func main() {
	printSomething(1)
	printSomething(2.5)
	printSomething("Hello World!")
	title, content := getNoteData()
	todotext := getUserInput("Todo text:")
	todo, err := todo.New(todotext)
	printSomething(todo)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	err = outputData(userNote)

}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func printSomething(value interface{}) {
	intVal, ok := value.(int)
	if ok {
		fmt.Println("Integer: ", intVal)
		return
	}

	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("Integer: ", floatVal)
		return
	}

	stringVal, ok := value.(string)
	if ok {
		fmt.Println("String: ", stringVal)
		return
	}
	/* switch value.(type) {
	case int:
		fmt.Println("Integer:", value)
	case float64:
		fmt.Println("Float:", value)
	case string:
		fmt.Println(value)
	} */

}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving data failed.")
		return err
	}

	fmt.Println("Saving the note succeeded!")
	return nil
}

func getUserInput(prompt string) string {
	fmt.Printf("%v: ", prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}
