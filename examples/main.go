package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/william-hood/boolog-go"
)

func main() {
	userHome, _ := os.UserHomeDir()
	rootLogPath := filepath.Join(userHome, "Documents", "Test Results", "Boolog Example.html")
	log := boolog.NewBoolog("Boolog Example", rootLogPath, boolog.THEME_DARK_FLAT)

	log.Info("Boolog is an HTML-based rich logging system capable of visual renditions of HTTP requests & responses, errors, and any other struct or type. One Boolog instance can even embed another as a log subsection.")
	log.Info("When used for debugging control flow, HTTP requests & responses, activity logging, or any other purpose, output from Boolog will easier to read and work with than ordinary console output (though it does provide counterpart output to the console in addition to its HTML log file).")
	log.Info("All of the above messages represent \"normal\" log output with the .Info() function.")
	log.Info("When debugging a program, you might need a single line of information to stand out.")
	log.Info("If you use the .Debug() function instead of .Info() the message will be highlighted in yellow like this...")
	log.Debug("Boolog is the spritual successor to a similar Golang log system I created at work years ago!")
	log.Info("Similar to that is the .Error() function. The only difference is an icon in the log identifying the line as an error...")
	log.Error("Uh-oh... That wasn't supposed to happen!")
	log.SkipLine()
	log.Info("Why would you want to log directly to HTML?")
	log.Info("Because it's very hard, using ordinary plain-text logging, to visualize the workings of a cloud service, test suite, or other back-end process.")
	log.Info("Let's suppose you need to check on the state of a data structure at a certain point in the program.")
	log.Info("Look at the class \"TestStruct\" at the bottom of this source code file. Let's render one!")

	myStruct := getTestStruct()
	log.ShowAsJson(myStruct, "myStruct")
	log.SkipLine()

	log.Info("Boolog can be very useful for testing HTTP Requests. Let's use Golang's standard HTTP client to send a request and get a response.")

	req, _ := http.NewRequest("GET", "https://httpbin.org/get?param1=latida&param2=tweedledee&param3=whatever", nil)
	log.ShowHttpTransaction(*req, nil)

	log.Info("Boolog also has a .ShowError() function. Here it is with a simple, standard Golang error...")
	_, err := os.Open("nonexistent_file.txt")
	log.ShowError(err, "err")
	log.Info("Search the internet and you may be able to find implementations of errors with attached stack traces.")
	log.SkipLine()

	log.ShowBoolog(sublog())
	log.SkipLine()

	log.Debug("One caveat: If you .Conclude() a Boolog, it's done. That function closes any output streams and makes it read-only.")
	log.Info("A Boolog also gets concluded if you embed it in another Boolog with the .ShowBoolog() function.")
	log.SkipLine()
	log.Info("Well, that's the demo. Go forth and do great things!")
	log.Conclude()
}

func sublog() boolog.Boolog {
	log := boolog.NewBoolog("Click this to see one of Boolog's biggest tricks!", "", "")
	log.Info("The truth is that all of the stuff above could've been put into it's own little click-to-expand subsection.")
	log.Info("A Boolog can embed another Boolog. Time stamps, icons, and all!")
	log.SkipLine()
	log.Info("Let's show another of those TestStruct things...")
	myStruct := getFullTestStruct()
	log.ShowAsJson(myStruct, "myStruct")
	log.SkipLine()

	log.Info("Let's repeat some of the things we did in the old log, just for show...")
	log.Debug("Yet another debug line!")
	log.Error("Uh-oh... That wasn't supposed to happen!")
	log.SkipLine()

	log.Info("Boolog can be very useful for testing HTTP Requests. Let's use Golang's standard HTTP client to send a request and get a response.")

	body := strings.NewReader("This is the request body")
	req, _ := http.NewRequest("POST", "https://httpbin.org/get?param1=latida&param2=tweedledee&param3=whatever", body)
	log.ShowHttpTransaction(*req, nil)
	log.SkipLine()

	subLog := boolog.NewBoolog("Keep on embedding Boologs within Boologs within Boologs", "", "")
	subLog.Info("The other day")
	subLog.Info("Upon the stair")
	subLog.Info("I saw a man")
	subLog.Info("Who wasn't there")
	subLog.SkipLine()
	subLog.Info("He wasn't there")
	subLog.Info("Again today...")
	subLog.SkipLine()
	subLog.Error("Gee, I wish he'd go away!")
	log.ShowBoologDetailed(subLog, boolog.EMOJI_BOOLOG, "neutral", 0)

	return log
}

type TestStruct struct {
	Name       string
	Value      int
	OtherValue float32
	Child      *TestStruct
	troll      string
	Rogue      map[string]string
}

func getTestStruct() TestStruct {
	result := new(TestStruct)
	result.Name = "Hi"
	result.Value = 7
	result.OtherValue = 42.9
	result.Child = nil
	result.troll = "Unexported (lowercase) fields are not visible to Golang reflection"
	result.Rogue = map[string]string{
		"LOTR":      "Sauron",
		"Star Wars": "Darth Vader",
		"It":        "Pennywise",
	}

	return *result
}

func getChildStruct() *TestStruct {
	result := new(TestStruct)
	result.Name = "Child"
	result.Value = 9
	result.OtherValue = 3.14159
	result.Child = nil
	result.troll = "arrrrrrgh!"
	result.Rogue = map[string]string{
		"Coffee": "Kona",
		"Tea":    "Earl Grey",
		"Soda":   "Cola",
		"Water":  "Distilled",
		"Spirit": "Wine",
	}

	return result
}

func getFullTestStruct() TestStruct {
	child := getChildStruct()
	another := getTestStruct()
	child.Child = &another
	result := getTestStruct()
	result.Child = child
	return result
}
