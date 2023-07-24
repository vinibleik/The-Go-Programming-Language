package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var mainTemplate = `package main

import "fmt"

func main() {
	fmt.Println("Chapter {{.Chapter}} Exercise {{.Ex}}")
}
`

type Exercise struct {
	Chapter, Ex string
}

var exercise Exercise
var goMain = template.Must(template.New("goMain").Parse(mainTemplate))

// Returns the path to the correct directory received by command-line
func getDirPath(dir string) (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}

	dirEx := filepath.Dir(ex)
	dirPath := filepath.Join(dirEx, "..", dir)

	return dirPath, nil
}

// Checks if the exercises directory exists in the path
func checkDirExercises(dir string) (bool, error) {
	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		return false, err
	}

	for _, de := range dirEntries {
		if de.IsDir() && de.Name() == "exercises" {
			return true, nil
		}
	}

	return false, nil
}

// Create the exercises directory in the dir path
func createExerciseDir(dir string) error {
	return os.MkdirAll(filepath.Join(dir, "exercises"), 0750)
}

// Returns the chapter from the command-line argument
func getChapter(s string) (string, bool) {
	return strings.CutPrefix(os.Args[1], "ch")
}

// Makes the new exercise on the dir path
func makeNewDirExercise(dir string, chapter string) (string, error) {
	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		return "", err
	}

	nextExercise := fmt.Sprintf("ex%s.%d", chapter, len(dirEntries)+1)
	nextExercise = filepath.Join(dir, nextExercise)

	if err := os.MkdirAll(nextExercise, 0750); err != nil {
		return "", err
	}
	exercise.Chapter = chapter
	exercise.Ex = fmt.Sprint(len(dirEntries) + 1)
	return nextExercise, nil
}

// Makes the main.go
func makeMain(dir string) error {
	if err := os.Chdir(dir); err != nil {
		return err
	}

	file, err := os.Create("main.go")
	if err != nil {
		return err
	}
	defer file.Close()

	return goMain.Execute(file, exercise)
}

func main() {

	if len(os.Args) < 2 {
		log.Fatal("new_exercise: Too few command-line arguments. Please provide the directory to construct the template!")
	}

	chapter, isChapter := getChapter(os.Args[1])
	if !isChapter {
		log.Fatal("new_exercise: The command-line argument is not a valid directory...(ch**)!")
	}

	chapter = strings.TrimPrefix(chapter, "0")

	dirPath, err := getDirPath(os.Args[1])
	if err != nil {
		log.Fatalf("new_exercise: %v\n", err)
	}

	existsDir, err := checkDirExercises(dirPath)
	if err != nil {
		log.Fatalf("new_exercise: %v\n", err)
	}

	if !existsDir {
		if err := createExerciseDir(dirPath); err != nil {
			log.Fatalf("new_exercise: %v\n", err)
		}
	}

	dirPath, err = makeNewDirExercise(filepath.Join(dirPath, "exercises"), chapter)
	if err != nil {
		log.Fatalf("new_exercise: %v\n", err)
	}

	if err := makeMain(dirPath); err != nil {
		log.Fatalf("new_exercise: %v\n", err)
	}

	fmt.Println("Successfully created the new exercise!")
}
