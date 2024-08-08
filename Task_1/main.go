package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func calculateAverage(total float32, numberOfSubjects int) float32 {
	ave := float32(total) / float32(numberOfSubjects)
	return ave
}

func display(name string, subjectToGrade map[string]float32, Average float32) {
	fmt.Println("Student Report Card")
	fmt.Println("---------------------")
	fmt.Println("Name:", name)
	fmt.Println("---------------------")

	fmt.Println("Subjects and Grades:")
	fmt.Println("---------------------")
	for key, value := range subjectToGrade {
		fmt.Printf("%-15s: %3.2f\n", key, value)
	}
	fmt.Println("---------------------")

	fmt.Printf("Average of All subjects: %.2f\n", Average)
	fmt.Println("---------------------")
}

func validateName(name string) bool {

	//names should always be letters
	if len(name) == 0 {
		return false
	}
	for _, c := range name {
		if !unicode.IsLetter(c) && !unicode.IsSpace(c) {
			return false
		}
	}
	return true
}

func validateSubjectName(subjectName string) bool {
	//Subject name might include numbers. Example CS101
	return len(subjectName) != 0
}

func validateNumberOfSubjects(num int) bool {
	//no upper limit because it depends on the country/places
	return num >= 0
}

func validateGrade(grade float32) bool {
	if grade < 0 || grade > 100 {
		return false
	}
	return true
}

func main() {
	var name string
	var numberOfSubjects int
	var subject string
	var grade float32
	var totalGrade float32
	subjectToGrade := make(map[string]float32)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Your name: ")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)

	if !validateName(name) {
		fmt.Println("Invalid name. Please Enter a name that contains letters and spaces.")
		return
	}

	fmt.Print("How many subjects did you take: ")
	fmt.Scanln(&numberOfSubjects)

	if !validateNumberOfSubjects(numberOfSubjects) {
		fmt.Println("Invalid input.Enter a positive number")
		return
	}

	for i := 0; i < numberOfSubjects; i++ {
		fmt.Print("Subject Name: ")
		subject, _ = reader.ReadString('\n')
		subject = strings.TrimSpace(subject)

		if !validateSubjectName(subject) {
			fmt.Println("Invalid subject name.")
			i-- //decrement i to ask the user again
			continue
		}

		fmt.Printf("%s grade: ", subject)
		fmt.Scanln(&grade)

		if !validateGrade(float32(grade)) {
			fmt.Println("Invalid grade. Please enter a grade between 0 and 100.")
			i--
			continue
		}

		subjectToGrade[subject] = grade
		totalGrade += grade
	}

	average := calculateAverage(totalGrade, numberOfSubjects)

	display(name, subjectToGrade, average)

}
