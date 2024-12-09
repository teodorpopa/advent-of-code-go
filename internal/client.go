package internal

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/k3a/html2text"
	"html"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const MinAocYear int = 2015
const AocDayStart int = 1
const AocDayEnd int = 25
const December int = 12
const AocTzOffset int = -5 * 3600

//go:embed templates/*.go
var fs embed.FS

type AocClient struct {
	Cookie string
	Year   int
	Day    int
}

func NewAocClient(year, day int) *AocClient {
	credentials, err := NewCredentialsManager()

	if err != nil {
		panic("Error: you need to provide your AOC cookie using the ~/.aoc file")
	}

	validatedYear, validatedDay := validateYearAndDay(year, day)

	return &AocClient{
		credentials.Cookie,
		validatedYear,
		validatedDay,
	}
}

func validateYearAndDay(year int, day int) (int, int) {
	newYear := validateYear(year)
	newDay := validateDay(newYear, day)

	return newYear, newDay
}

func validateYear(year int) int {
	date := currentDate()

	if year >= MinAocYear && year <= date.Year() {
		return year
	}

	return lastAocYear(date)
}

func validateDay(year int, day int) int {
	date := currentDate()
	currentDay := date.Day()

	if year == date.Year() && int(date.Month()) == December {
		if day > 0 && day <= currentDay && day <= AocDayEnd {
			return day
		} else {
			return AocDayStart
		}
	} else if year < date.Year() {
		if day > 0 && day <= AocDayEnd {
			return day
		} else {
			return AocDayEnd
		}
	}

	return lastAocDay(year, date)
}

func lastAocYear(date time.Time) int {
	if int(date.Month()) < December {
		return date.Year() - 1
	}

	return date.Year()
}

func lastAocDay(year int, date time.Time) int {
	if year == date.Year() && int(date.Month()) == December {
		if date.Day() <= AocDayEnd {
			return date.Day()
		} else {
			return AocDayEnd
		}
	} else if year < date.Year() {
		return AocDayEnd
	}

	return AocDayStart
}

func currentDate() time.Time {
	loc := time.FixedZone("aoc", AocTzOffset)
	return time.Now().UTC().In(loc)
}

func (c *AocClient) Calendar() string {
	url := fmt.Sprintf("https://adventofcode.com/%s", strconv.Itoa(c.Year))
	resp := c.aocRequest(url, "text/html")

	loginRegex, _ := regexp.Compile("href=\"/[0-9]{4}/auth/login")
	isLoginContent := loginRegex.FindString(resp)
	if isLoginContent != "" {
		return "Cannot display calendar. You need to login to your AOC account."
	}

	r := regexp.MustCompile("(?i)(?s)<main>(.*)</main>")
	calendarContent := r.FindString(resp)
	calendarContent = c.cleanCalendar(calendarContent)

	return calendarContent
}

func (c *AocClient) cleanCalendar(calendarContent string) string {

	r := regexp.MustCompile("(?i)(?s)<pre[^>]*>(.*)</pre>")
	calendarContent = r.FindString(calendarContent)

	calendarRe := regexp.MustCompile(
		// 2015
		"(<div class=\"calendar-bkg\">[[:space:]]*(<div>[^<]*</div>[[:space:]]*)*</div>)" +
			//// 2017
			"|(<div class=\"calendar-printer\">(?s:.)*\\|O\\|</span></div>[[:space:]]*)" +
			//// 2018
			"|(<pre id=\"spacemug\"[^>]*>[^<]*</pre>)" +
			//// 2019
			"|(<span style=\"color[^>]*position:absolute[^>]*>\\.</span>)" +
			"|(<span class=\"sunbeam\"[^>]*><span style=\"animation-delay[^>]*>\\*</span></span>)" +
			//// 2023
			"|(<span class=\"(lava|gear|sand|isle|snow)fall\"[^>]*>(<span (style|class)=\"[^>]*>(.*)</span>)*</span>)",
	)

	calendarString := calendarRe.ReplaceAllString(calendarContent, "")

	preRegex := regexp.MustCompile(`</?pre[^>]*>`)
	calendarString = preRegex.ReplaceAllString(calendarString, "$2")

	classRegex := regexp.MustCompile(
		"<a [^>]*class=\"(?P<class>[^\"]*)\">(?P<content>.*)",
	)

	starRegex := regexp.MustCompile(
		"(?P<stars><span class=\"calendar-mark-complete\">\\*</span><span class=\"calendar-mark-verycomplete\">\\*</span>)",
	)

	dayRegex := regexp.MustCompile(
		"(?P<day><span class=\"calendar-day\">(.*)</span>)",
	)

	spanRegex := regexp.MustCompile(`</?span[^>]*>`)

	scanner := bufio.NewScanner(strings.NewReader(calendarString))

	calendar := "\n"
	for scanner.Scan() {
		content := scanner.Text()

		class := classRegex.ReplaceAllString(content, "$1")

		content = classRegex.ReplaceAllString(content, "$2")
		content = starRegex.ReplaceAllString(content, "")
		content = dayRegex.ReplaceAllString(content, "$2")
		content = strings.ReplaceAll(content, "</a>", "")
		content = strings.ReplaceAll(content, "<i>", "")
		content = strings.ReplaceAll(content, "</i>", "")
		content = spanRegex.ReplaceAllString(content, "$2")

		if strings.Contains(class, "calendar-verycomplete") || strings.Contains(class, "calendar-perfect") {
			content += "**"
		} else if strings.Contains(class, "calendar-complete") {
			content += "*"
		}

		calendar += html.UnescapeString(content) + "\n"
	}

	return calendar
}

func (c *AocClient) ViewPuzzle() string {
	url := fmt.Sprintf("https://adventofcode.com/%s/day/%s", strconv.Itoa(int(c.Year)), strconv.Itoa(int(c.Day)))
	resp := c.aocRequest(url, "text/html")

	r, _ := regexp.Compile("(?i)(?s)<main>(?P<main>.*)</main>")
	return html2text.HTML2Text(r.FindString(resp))
}

func (c *AocClient) aocRequest(url string, contentType string) string {
	httpClient := http.Client{Timeout: time.Duration(1) * time.Second}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Printf("error %s", err)
		return ""
	}

	req.Header.Add("Accept", contentType)

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: strings.Trim(c.Cookie, " \n"),
	})
	req.Header.Add("User-Agent", "Rust AoC 0.1.0")

	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
		return ""
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	return string(body)
}

func (c *AocClient) GenerateDay() {
	ts, err := template.ParseFS(fs, "templates/*.go")
	if err != nil {
		log.Fatalf("error parsing templates: %s", err)
	}

	currentDir, _ := os.Getwd()

	dir := filepath.Join(currentDir, fmt.Sprintf("y%d/day%02d/", c.Year, c.Day))
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Fatalf("error creating directory: %s", err)
	}

	writeTemplate(dir, "main.go", ts)
	writeTemplate(dir, "main_test.go", ts)
	writeFile(dir, "puzzle.md", c.ViewPuzzle())
	writeFile(dir, "input.txt", "")
	writeFile(dir, "input_test.txt", "")

	fmt.Printf("templates created in y%d/day%d\n", c.Year, c.Day)
}

func writeTemplate(dir, filename string, ts *template.Template) {
	filePath := dir + "/" + filename

	_, err := os.Stat(filePath)
	if err != nil {
		file, err := os.Create(filePath)
		if err != nil {
			log.Fatalf("error creating file: %s - %v", filePath, err)
		}
		ts.ExecuteTemplate(file, filename, nil)
	}
}

func writeFile(dir, filename string, content string) {
	filePath := dir + "/" + filename

	_, err := os.Stat(filePath)
	if err != nil {
		file, err := os.Create(filePath)
		if err != nil {
			log.Fatalf("error creating file: %s - %v", filePath, err)
		}
		file.WriteString(content)
	}
}
