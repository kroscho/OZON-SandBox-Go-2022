package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Time struct {
	hh int
	mm int
	ss int
}

func (t *Time) getDatetime() time.Time {
	return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), t.hh, t.mm, t.ss, 0, time.UTC)
}

type SegmentTime struct {
	Start Time
	End   Time
}

func (s *SegmentTime) isEndAfterStart() bool {
	if s.Start == s.End {
		return true
	}
	t1 := s.Start.getDatetime()
	t2 := s.End.getDatetime()

	return t2.After(t1)
}

func inputInt(in *bufio.Reader, text string) int {
	var x int
	fmt.Print(text)
	fmt.Fscan(in, &x)

	return x
}

func inputStr(in *bufio.Reader, text string) string {
	var x string
	fmt.Print(text)
	fmt.Fscan(in, &x)

	return x
}

func isValidFormatTime(text string) bool {
	matched, _ := regexp.MatchString(`\b(0\d|1\d|2[0-3]):(0\d|[1-5]\d):(0\d|[1-5]\d)\b`, text)

	return matched
}

func splitTimes(timesStr string) (string, string) {
	times := strings.Split(timesStr, "-")
	time1 := times[0]
	time2 := times[1]

	return time1, time2
}

func getTimeStruct(timeStr string) Time {
	splitTime := strings.Split(timeStr, ":")

	hh, _ := strconv.Atoi(splitTime[0])
	mm, _ := strconv.Atoi(splitTime[1])
	ss, _ := strconv.Atoi(splitTime[2])

	return Time{
		hh: hh,
		mm: mm,
		ss: ss,
	}
}

func isIntersectionTimes(s1 SegmentTime, s2 SegmentTime) bool {
	t1 := s1.Start.getDatetime()
	t2 := s1.End.getDatetime()
	t3 := s2.Start.getDatetime()
	t4 := s2.End.getDatetime()

	if t3.After(t2) || t1.After(t4) {
		return false
	}

	return true
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	testsCount := inputInt(in, "Введите количество тестов: ")

	for i := 0; i < testsCount; i++ {
		timesCount := inputInt(in, "Введите кол-во отрезков: ")
		isError := false
		segmentsTime := []SegmentTime{}

		for j := 0; j < timesCount; j++ {
			timeStr := inputStr(in, "Ввредите отрезок: ")
			if isError {
				continue
			}
			timeStr1, timeStr2 := splitTimes(timeStr)

			if !isValidFormatTime(timeStr1) || !isValidFormatTime(timeStr2) {
				fmt.Println("Error format")
				isError = true
				continue
			}

			time1 := getTimeStruct(timeStr1)
			time2 := getTimeStruct(timeStr2)

			curSegment := SegmentTime{
				Start: time1,
				End:   time2,
			}

			if !curSegment.isEndAfterStart() {
				fmt.Println("Error time2 before")
				isError = true
				continue
			}

			for _, segment := range segmentsTime {
				if isIntersectionTimes(curSegment, segment) {
					fmt.Println("Error times intersection")
					isError = true
					break
				}
			}
			if isError {
				continue
			}

			segmentsTime = append(segmentsTime, curSegment)
		}

		if isError {
			fmt.Fprintln(out, "NO")
		} else {
			fmt.Fprintln(out, "YES")
		}
	}
}
