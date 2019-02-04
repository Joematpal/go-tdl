// Please provide your answer after the `?:`.
// What is your first name?:
// What is your last name?:
// What are the last four digits of your phone number?:
// What is your age?:
// What month were you born in?:
// What day were you born on; monday, tuesday, etc?:
// What date were you born on; 1, 2, 3, etc?:
// What year were you born on; 2000, 1985, 2019, etc?:
// What time where you born on, military (13:00)?:
// Do you have any siblings, yes or no?:
// Place 0 if you don't have any. If you have siblings, how many do you have?:
// How tall are you in inches, whole number?:

package tdl

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	// Learner ...
	Learner             = Student{}
	fnameRegexp         = regexp.MustCompile(`.*first\sname\?:`)
	lnameRegexp         = regexp.MustCompile(`.*last\sname\?:`)
	ageRegexp           = regexp.MustCompile(`.*age\?:`)
	phoneNumberRegexp   = regexp.MustCompile(`.*phone\snumber\?:`)
	monthRegexp         = regexp.MustCompile(`.*month.*?:`)
	dayRegexp           = regexp.MustCompile(`.*day.*?:`)
	dateRegexp          = regexp.MustCompile(`.*date.*?:`)
	yearRegexp          = regexp.MustCompile(`.*year.*?:`)
	timeRegexp          = regexp.MustCompile(`.*time.*?:`)
	siblingsRegexp      = regexp.MustCompile(`.*siblings, yes.*?:`)
	siblingsCountRegexp = regexp.MustCompile(`.*siblings, how many.*?:`)
	heightRegexp        = regexp.MustCompile(`.*tall.*?:`)
)

// Student ...
type Student struct {
	FirstName         string
	LastName          string
	PhoneNumber       int
	Age               int
	BirthMonth        string
	BirthDay          string
	BirthDate         int
	BirthYear         int
	BirthMilitaryTime string
	Siblings          bool
	SiblingsCount     int
	Height            int
}

// GetBirthMilitaryTimeFloat ...
func (student *Student) GetBirthMilitaryTimeFloat() float64 {
	return float64(student.GetBirthHour()) + (float64(student.GetBirthMin()) / float64(60))
}

// GetBirthHour ...
func (student *Student) GetBirthHour() int {
	times := strings.Split(student.BirthMilitaryTime, ":")
	hour, _ := strconv.Atoi(times[0])
	return hour
}

// GetBirthMin ...
func (student *Student) GetBirthMin() int {
	times := strings.Split(student.BirthMilitaryTime, ":")
	minutes, _ := strconv.Atoi(times[1])
	return minutes
}

// ParseQuestions reads comments for `?:` to find answers to questions
func ParseQuestions() Student {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pwd)

	file, err := os.Open("./questions.go")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 2 {
			continue
		}
		sub := line[0:2]
		if sub != "//" {
			continue
		}
		switch true {
		case fnameRegexp.Match([]byte(line)):
			Learner.FirstName = getStringAnswer(line)

		case lnameRegexp.Match([]byte(line)):
			Learner.LastName = getStringAnswer(line)

		case ageRegexp.Match([]byte(line)):
			Learner.Age = getIntAnswer(line)

		case phoneNumberRegexp.Match([]byte(line)):
			Learner.PhoneNumber = getIntAnswer(line)

		case monthRegexp.Match([]byte(line)):
			Learner.BirthMonth = getStringAnswer(line)

		case dayRegexp.Match([]byte(line)):
			Learner.BirthDay = getStringAnswer(line)

		case dateRegexp.Match([]byte(line)):
			Learner.BirthDate = getIntAnswer(line)

		case yearRegexp.Match([]byte(line)):
			Learner.BirthYear = getIntAnswer(line)

		case timeRegexp.Match([]byte(line)):
			Learner.BirthMilitaryTime = getStringAnswer(line)

		case siblingsRegexp.Match([]byte(line)):
			Learner.Siblings = getBoolAnswer(line)

		case siblingsCountRegexp.Match([]byte(line)):
			Learner.SiblingsCount = getIntAnswer(line)

		case heightRegexp.Match([]byte(line)):
			Learner.Height = getIntAnswer(line)

		default:
			// fmt.Println("Error...", line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return Learner
}

func isStringEmpty(s string) bool {
	if len(s) < 1 {
		fmt.Println("Please answer the question.")
		return true
	}
	return false
}

func getStringAnswer(line string) string {
	questionLine := regexp.MustCompile(`\?:`).Split(line, 2)
	if isStringEmpty(questionLine[1]) {
		return ""
	}
	return strings.Trim(questionLine[1], " ")
}

func getIntAnswer(line string) int {
	i, err := strconv.Atoi(getStringAnswer(line))
	if err != nil {
		log.Fatal("getIntAnswer()")
		return 0
	}
	return i
}

func getFloat64Answer(line string) float64 {
	f, err := strconv.ParseFloat(line, 64)
	if err != nil {
		log.Fatal("getFloat64Answer()")
		return 0.0
	}
	return f
}

func getBoolAnswer(line string) bool {
	answer := strings.ToLower(getStringAnswer(line))
	if answer == "yes" {
		return true
	}
	if answer == "no" {
		return false
	}
	fmt.Println(answer)

	b, err := strconv.ParseBool(answer)
	if err != nil {
		log.Fatal("getBoolAnswer()")
	}
	return b
}
