// This code logs a formatted message indicating that the application "Example" is starting and specifies the version of Go used to build the binary.
package main

import (
	"log"
	"runtime"
)

const info = `
Application %s starting.
The binary was build by GO: %s`

func main() {
	log.Printf(info, "Example", runtime.Version())
}

// This code prints all command line arguments, the name of the binary, and each subsequent argument with its index.
package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args

	// This call will print
	// all command line arguments.
	fmt.Println(args)

	// The first argument, zero item from array,
	// is the name of the called binary.
	programName := args[0]
	fmt.Printf("The binary name is: %s \n", programName)

	// The rest of the arguments could be naturally obtained
	// by omitting the first argument.
	otherArgs := args[1:]
	fmt.Println(otherArgs)

	for idx, arg := range otherArgs {
		fmt.Printf("Arg %d = %s \n", idx, arg)
	}

}


// This code defines and parses command line flags for retry count, log prefix, and an array, then logs retry attempts and the array using the specified log prefix.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// Custom type needs to implement
// flag.Value interface to be able to
// use it in flag.Var function.
type ArrayValue []string

func (s *ArrayValue) String() string {
	return fmt.Sprintf("%v", *s)
}

func (a *ArrayValue) Set(s string) error {
	*a = strings.Split(s, ",")
	return nil
}

func main() {

	// Extracting flag values with methods returning pointers
	retry := flag.Int("retry", -1, "Defines max retry count")

	// Read the flag using the XXXVar function.
	// In this case the variable must be defined
	// prior to the flag.
	var logPrefix string
	flag.StringVar(&logPrefix, "prefix", "", "Logger prefix")

	var arr ArrayValue
	flag.Var(&arr, "array", "Input array to iterate through.")

	// Execute the flag.Parse function, to
	// read the flags to defined variables.
	// Without this call the flag
	// variables remain empty.
	flag.Parse()

	// Sample logic not related to flags
	logger := log.New(os.Stdout, logPrefix, log.Ldate)

	retryCount := 0
	for retryCount < *retry {
		logger.Println("Retrying connection")
		logger.Printf("Sending array %v\n", arr)
		retryCount++
	}
}



// This code sets an environment variable, retrieves its value (or a default if not set), logs the values, and then unsets the environment variable.
package main

import (
	"log"
	"os"
)

func main() {

	key := "DB_CONN"
	// Set the environmental variable.
	os.Setenv(key, "postgres://as:as@example.com/pg?sslmode=verify-full")

	val := GetEnvDefault(key, "postgres://as:as@localhost/pg?sslmode=verify-full")

	log.Println("The value is :" + val)

	// Unset the environmental variable.
	os.Unsetenv(key)

	val = GetEnvDefault(key, "postgres://as:as@127.0.0.1/pg?sslmode=verify-full")

	log.Println("The default value is :" + val)

}

// GetEnvDefault retrieves the value of the environment variable
// specified by key, or returns defVal if the variable is not set.
func GetEnvDefault(key, defVal string) string {
	val, ex := os.LookupEnv(key)
	if !ex {
		return defVal
	}
	return val
}

// This code looks up an environment variable and logs a message if it is not set, then prints the value of the environment variable.
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	key := "DB_CONN"
	connStr, ex := os.LookupEnv(key)
	if !ex {
		log.Printf("The env variable %s is not set.\n", key)
	}
	fmt.Println(connStr)
}

// This code retrieves the value of the "DB_CONN" environment variable and logs it.
package main

import (
	"log"
	"os"
)

func main() {
	connStr := os.Getenv("DB_CONN")
	log.Printf("Connection string: %s\n", connStr)
}


// This code prints the path to the current executable, resolves and prints its directory, and evaluates and prints any symbolic links in the directory path.
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	// Path to executable file
	fmt.Println(ex)

	// Resolve the directory
	// of the executable
	exPath := filepath.Dir(ex)
	fmt.Println("Executable path :" + exPath)

	// Use EvalSymlinks to get
	// the real path.
	realPath, err := filepath.EvalSymlinks(exPath)
	if err != nil {
		panic(err)
	}
	fmt.Println("Symlink evaluated:" + realPath)
}

// This code retrieves the current process ID, runs the "ps" command to display process information, and prints the output.
package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func main() {

	// Get the current process ID.
	pid := os.Getpid()
	fmt.Printf("Process PID: %d \n", pid)

	// Execute the "ps" command to display process information for the current process.
	prc := exec.Command("ps", "-p", strconv.Itoa(pid), "-v")
	out, err := prc.Output()
	if err != nil {
		panic(err)
	}

	// Print the output of the "ps" command.
	fmt.Println(string(out))

}


// This code sets up signal handling to catch specific termination signals, prints corresponding messages, and exits with an appropriate status code.
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Create the channel where the received
	// signal will be sent. The Notify
	// will not block when the signal
	// is sent and the channel is not ready.
	// So it is better to
	// create a buffered channel.
	sChan := make(chan os.Signal, 1)

	// Notify will catch the
	// given signals and send
	// the os.Signal value
	// through the sChan.
	signal.Notify(sChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGKILL)

	// Create a channel to wait until the
	// signal is handled.
	exitChan := make(chan int)
	go func() {
		signal := <-sChan
		switch signal {
		case syscall.SIGHUP:
			fmt.Println("The calling terminal has been closed")
			exitChan <- 0

		case syscall.SIGINT:
			fmt.Println("The process has been interrupted by CTRL+C")
			exitChan <- 1

		case syscall.SIGTERM:
			fmt.Println("kill SIGTERM was executed for process")
			exitChan <- 1

		case syscall.SIGKILL:
			fmt.Println("SIGKILL handler")
			exitChan <- 1

		case syscall.SIGQUIT:
			fmt.Println("kill SIGQUIT was executed for process")
			exitChan <- 1
		}
	}()

	code := <-exitChan
	os.Exit(code)
}

// This code executes the "ls -a" command to list all files, captures the output, and prints it if the command succeeds.
package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {

	// Create a command to execute "ls -a"
	prc := exec.Command("ls", "-a")

	// Create a buffer to capture the command's output
	out := bytes.NewBuffer([]byte{})
	prc.Stdout = out

	// Start the command
	err := prc.Start()
	if err != nil {
		fmt.Println(err)
	}

	// Wait for the command to complete
	prc.Wait()

	// Check if the command executed successfully
	if prc.ProcessState.Success() {
		fmt.Println("Process run successfully with output:\n")
		fmt.Println(out.String())
	}
}

// This Go program executes the 'ls -a' command to list all files and directories in the current directory,
// captures the command's output into a buffer, and prints the output if the command executes successfully.
package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	prc := exec.Command("ls", "-a")
	out := bytes.NewBuffer([]byte{})
	prc.Stdout = out

	err := prc.Run()
	if err != nil {
		fmt.Println(err)
	}

	if prc.ProcessState.Success() {
		fmt.Println("Process run successfully with output:\n")
		fmt.Println(out.String())
	}
}

// This program executes a platform-specific command ('timeout' on Windows, 'sleep' on other systems) for 1 second, then prints process information including PID, execution time in milliseconds, and success status.

package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"
)

func main() {

	var cmd string
	if runtime.GOOS == "windows" {
		cmd = "timeout"
	} else {
		cmd = "sleep"
	}

	proc := exec.Command(cmd, "1")
	proc.Start()

	// Wait function will
	// wait till the process ends.
	proc.Wait()

	// After the process terminates
	// the *os.ProcessState contains
	// simple information
	// about the process run
	fmt.Printf("PID: %d\n", proc.ProcessState.Pid())
	fmt.Printf("Process took: %dms\n", proc.ProcessState.SystemTime()/time.Microsecond)
	fmt.Printf("Exited sucessfuly : %t\n", proc.ProcessState.Success())
}


// This program starts a platform-specific command ('timeout' on Windows, 'sleep' on other systems) for 1 second, then prints the process state and PID of the running process.
package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {

	var cmd string
	if runtime.GOOS == "windows" {
		cmd = "timeout"
	} else {
		cmd = "sleep"
	}
	proc := exec.Command(cmd, "1")
	proc.Start()

	// No process state is returned
	// till the process finish.
	fmt.Printf("Process state for running process: %v\n", proc.ProcessState)

	// The PID could be obtain
	// event for the running process
	fmt.Printf("PID of running process: %d\n\n", proc.Process.Pid)
}


// Executes a Go program "sample.go" using "go run" command, sends input via standard input and prints output from stdout until killed after 2 seconds.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"time"
)

func main() {
	cmd := []string{"go", "run", "sample.go"}

	// The command line tool
	// "ping" is executed for
	// 2 seconds
	proc := exec.Command(cmd[0], cmd[1], cmd[2])

	// The process input is obtained
	// in form of io.WriteCloser. The underlying
	// implementation use the os.Pipe
	stdin, _ := proc.StdinPipe()
	defer stdin.Close()

	// For debugging purposes we watch the
	// output of the executed process
	stdout, _ := proc.StdoutPipe()
	defer stdout.Close()

	go func() {
		s := bufio.NewScanner(stdout)
		for s.Scan() {
			fmt.Println("Program says:" + s.Text())
		}
	}()

	// Start the process
	proc.Start()

	// Now the the following lines
	// are written to child
	// process standard input
	fmt.Println("Writing input")
	io.WriteString(stdin, "Hello\n")
	io.WriteString(stdin, "Golang\n")
	io.WriteString(stdin, "is awesome\n")

	time.Sleep(time.Second * 2)

	proc.Process.Kill()

}


// This program demonstrates interprocess communication by executing "go run sample.go" command,
// sending multiple lines of input to the child process, and printing its responses until killed after 2 seconds.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"time"
)

func main() {
	cmd := []string{"go", "run", "sample.go"}

	// The command line tool "go run sample.go" is executed for 2 seconds
	proc := exec.Command(cmd[0], cmd[1], cmd[2])

	// The process input is obtained in the form of io.WriteCloser using os.Pipe
	stdin, _ := proc.StdinPipe()
	defer stdin.Close()

	// For capturing and printing output from the executed process
	stdout, _ := proc.StdoutPipe()
	defer stdout.Close()

	go func() {
		s := bufio.NewScanner(stdout)
		for s.Scan() {
			fmt.Println("Program says:" + s.Text())
		}
	}()

	// Start the process
	proc.Start()

	// Writing input lines to the child process standard input
	fmt.Println("Writing input")
	io.WriteString(stdin, "Hello\n")
	io.WriteString(stdin, "Golang\n")
	io.WriteString(stdin, "is awesome\n")

	// Allow 2 seconds for the process to run
	time.Sleep(time.Second * 2)

	// Kill the process
	proc.Process.Kill()
}


// This program demonstrates logging to a dynamically named file, manages resource cleanup with signals,
// and gracefully shuts down a goroutine writing logs until termination signals are received.
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var writer *os.File

func main() {

	// The file is opened as
	// a log file to write into.
	// This way we represent the resources
	// allocation.
	var err error
	writer, err = os.OpenFile(fmt.Sprintf("test_%d.log", time.Now().Unix()), os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// The code is running in a goroutine
	// independently. So in case the program is
	// terminated from outside, we need to
	// let the goroutine know via the closeChan
	closeChan := make(chan bool)
	go func() {
		for {
			time.Sleep(time.Second)
			select {
			case <-closeChan:
				log.Println("Goroutine closing")
				return
			default:
				log.Println("Writing to log")
				io.WriteString(writer, fmt.Sprintf("Logging access %s\n", time.Now().String()))
			}

		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGINT)

	// This is blocking read from
	// sigChan where the Notify function sends
	// the signal.
	<-sigChan

	// After the signal is received
	// all the code behind the read from channel could be
	// considered as a cleanup
	close(closeChan)
	releaseAllResources()
	fmt.Println("The application shut down gracefully")
}

func releaseAllResources() {
	io.WriteString(writer, "Application releasing all resources\n")
	writer.Close()
}


// This program demonstrates the use of functional options to configure a Client struct,
// allowing configuration from a JSON file and environmental variables, and prints the resulting configuration.
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Client struct {
	consulIP   string
	connString string
}

func (c *Client) String() string {
	return fmt.Sprintf("ConsulIP: %s , Connection String: %s",
		c.consulIP, c.connString)
}

var defaultClient = Client{
	consulIP:   "localhost:9000",
	connString: "postgres://localhost:5432",
}

// ConfigFunc works as a type to be used
// in functional options
type ConfigFunc func(opt *Client)

// FromFile func returns the ConfigFunc
// type. So this way it could read the configuration
// from the json.
func FromFile(path string) ConfigFunc {
	return func(opt *Client) {
		f, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		decoder := json.NewDecoder(f)

		fop := struct {
			ConsulIP   string `json:"consul_ip"`
			ConnString string `json:"conn_string"`
		}{}
		err = decoder.Decode(&fop)
		if err != nil {
			panic(err)
		}
		opt.consulIP = fop.ConsulIP
		opt.connString = fop.ConnString
	}
}

// FromEnv reads the configuration
// from the environmental variables
// and combines them with existing ones.
func FromEnv() ConfigFunc {
	return func(opt *Client) {
		connStr, exist := os.LookupEnv("CONN_DB")
		if exist {
			opt.connString = connStr
		}
	}
}

func NewClient(opts ...ConfigFunc) *Client {
	client := defaultClient
	for _, val := range opts {
		val(&client)
	}
	return &client
}

func main() {
	client := NewClient(FromFile("config.json"), FromEnv())
	fmt.Println(client.String())
}


// This program demonstrates the usage of strings.Contains, strings.HasPrefix, and strings.HasSuffix
// functions to check if a reference string contains a substring, starts with a prefix, or ends with a suffix.
package main

import (
	"fmt"
	"strings"
)

const refString = "Mary had a little lamb"

func main() {

	lookFor := "lamb"
	contain := strings.Contains(refString, lookFor)
	fmt.Printf("The \"%s\" contains \"%s\": %t \n", refString, lookFor, contain)

	lookFor = "wolf"
	contain = strings.Contains(refString, lookFor)
	fmt.Printf("The \"%s\" contains \"%s\": %t \n", refString, lookFor, contain)

	startsWith := "Mary"
	starts := strings.HasPrefix(refString, startsWith)
	fmt.Printf("The \"%s\" starts with \"%s\": %t \n", refString, startsWith, starts)

	endWith := "lamb"
	ends := strings.HasSuffix(refString, endWith)
	fmt.Printf("The \"%s\" ends with \"%s\": %t \n", refString, endWith, ends)

}


// This program splits the constant string refString using the underscore character "_"
// as a delimiter and prints each split word along with its index.
package main

import (
	"fmt"
	"strings"
)

const refString = "Mary_had a little_lamb"

func main() {
	words := strings.Split(refString, "_")
	for idx, word := range words {
		fmt.Printf("Word %d is: %s\n", idx, word)
	}
}


// This program splits the constant string refString using the characters '*', ',', '%', and '_' as delimiters
// and prints each split word along with its index.
package main

import (
	"fmt"
	"regexp"
)

const refString = "Mary*had,a%little_lamb"

func main() {
	words := regexp.MustCompile("[*,%_]{1}").Split(refString, -1)
	for idx, word := range words {
		fmt.Printf("Word %d is: %s\n", idx, word)
	}
}



// This program splits the constant string refString using a custom split function defined by splitFunc,
// which checks if each rune in the string is '*', ',', '%', or '_',
// and prints each split word along with its index.
package main

import (
	"fmt"
	"strings"
)

const refString = "Mary*had,a%little_lamb"

func main() {
	// The splitFunc is called for each
	// rune in a string. If the rune
	// equals any of the characters '*', ',', '%', '_',
	// the refString is split.
	splitFunc := func(r rune) bool {
		return strings.ContainsRune("*%,_", r)
	}

	words := strings.FieldsFunc(refString, splitFunc)
	for idx, word := range words {
		fmt.Printf("Word %d is: %s\n", idx, word)
	}
}

// This program splits the constant string refString into words using whitespace characters as delimiters
// and prints each word along with its index.
package main

import (
	"fmt"
	"strings"
)

const refString = "Mary had	a little lamb"

func main() {
	words := strings.Fields(refString)
	for idx, word := range words {
		fmt.Printf("Word %d is: %s\n", idx, word)
	}
}


// This program constructs a SQL SELECT statement template with placeholders,
// joins an array of conditions refStringSlice using "AND" as a delimiter,
// and prints the formatted SELECT statement.
package main

import (
	"fmt"
	"strings"
)

const selectBase = "SELECT * FROM user WHERE %s "

var refStringSlice = []string{
	" FIRST_NAME = 'Jack' ",
	" INSURANCE_NO = 333444555 ",
	" EFFECTIVE_FROM = SYSDATE ",
}

func main() {
	sentence := strings.Join(refStringSlice, " AND ")
	fmt.Printf(selectBase+"\n", sentence)
}


// This program constructs a SQL WHERE clause by joining an array of conditions refStringSlice
// using a custom join function jF, which determines whether to use "AND" or "OR" based on the content of each condition.
package main

import (
	"fmt"
	"strings"
)

const selectBase = "SELECT * FROM user WHERE "

var refStringSlice = []string{
	" FIRST_NAME = 'Jack' ",
	" INSURANCE_NO = 333444555 ",
	" EFFECTIVE_FROM = SYSDATE ",
}

type JoinFunc func(piece string) string

func main() {
	jF := func(p string) string {
		if strings.Contains(p, "INSURANCE") {
			return "OR"
		}
		return "AND"
	}
	result := JoinWithFunc(refStringSlice, jF)
	fmt.Println(selectBase + result)
}

func JoinWithFunc(refStringSlice []string, joinFunc JoinFunc) string {
	concatenate := refStringSlice[0]
	for _, val := range refStringSlice[1:] {
		concatenate = concatenate + joinFunc(val) + val
	}
	return concatenate
}


// This program demonstrates efficient string concatenation by copying each string
// from the slice `strings` into a byte slice `bs` using the `copy` function,
// and then converting the byte slice back to a string for output.
package main

import (
	"fmt"
)

func main() {
	strings := []string{"This ", "is ", "even ", "more ", "performant "}

	bs := make([]byte, 100)
	bl := 0

	for _, val := range strings {
		bl += copy(bs[bl:], []byte(val))
	}

	fmt.Println(string(bs[:bl]))
}



// This program efficiently concatenates multiple strings from the slice `strings`
// into a single string using a bytes.Buffer to accumulate the result,
// demonstrating a more performant approach compared to direct string concatenation.
package main

import (
	"bytes"
	"fmt"
)

func main() {
	strings := []string{"This ", "is ", "even ", "more ", "performant "}
	buffer := bytes.Buffer{}
	for _, val := range strings {
		buffer.WriteString(val)
	}

	fmt.Println(buffer.String())
}

// This program demonstrates the usage of the tabwriter package
// to format and align columns of text, particularly for tabular data output.
package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {
	w := tabwriter.NewWriter(os.Stdout, 15, 0, 1, ' ', tabwriter.AlignRight)
	fmt.Fprintln(w, "username\tfirstname\tlastname\t")
	fmt.Fprintln(w, "sohlich\tRadomir\tSohlich\t")
	fmt.Fprintln(w, "novak\tJohn\tSmith\t")
	w.Flush()
}


// This program demonstrates the usage of regular expressions
// to replace all substrings matching the pattern "l[a-z]+" in the constant refString
// with the string "replacement".
package main

import (
	"fmt"
	"regexp"
)

const refString = "Mary had a little lamb"

func main() {
	regex := regexp.MustCompile("l[a-z]+")
	out := regex.ReplaceAllString(refString, "replacement")
	fmt.Println(out)
}


// This program demonstrates the usage of strings.Replace function
// to replace all occurrences of "lamb" with "wolf" in the constant refString,
// and replace only the first two occurrences of "lamb" with "wolf" in the constant refStringTwo.
package main

import (
	"fmt"
	"strings"
)

const refString = "Mary had a little lamb"
const refStringTwo = "lamb lamb lamb lamb"

func main() {
	out := strings.Replace(refString, "lamb", "wolf", -1)
	fmt.Println(out)

	out = strings.Replace(refStringTwo, "lamb", "wolf", 2)
	fmt.Println(out)
}


// This program demonstrates the usage of strings.NewReplacer to create a custom string replacer,
// replacing occurrences of "lamb" with "wolf" and "Mary" with "Jack" in the constant refString.
package main

import (
	"fmt"
	"strings"
)

const refString = "Mary had a little lamb"

func main() {
	replacer := strings.NewReplacer("lamb", "wolf", "Mary", "Jack")
	out := replacer.Replace(refString)
	fmt.Println(out)
}



// This program demonstrates the usage of regular expressions to extract email addresses
// from the constant refString, using a simplified pattern matching approach.
package main

import (
	"fmt"
	"regexp"
)

const refString = `[{ "email": "email@example.com" "phone": 555467890},
{ "email": "other@domain.com" "phone": 555467890}]`

func main() {
	// This pattern is simplified for brevity
	emailRegexp := regexp.MustCompile("[a-zA-Z0-9]{1,}@[a-zA-Z0-9]{1,}\\.[a-z]{1,}")
	first := emailRegexp.FindString(refString)
	fmt.Println("First: ")
	fmt.Println(first)

	all := emailRegexp.FindAllString(refString, -1)
	fmt.Println("All: ")
	for _, val := range all {
		fmt.Println(val)
	}
}

// This program demonstrates reading and decoding a file encoded in Windows-1250 charset.
// It reads the file "win1250.txt", displays its content in its raw form,
// and then decodes it from Windows-1250 to Unicode using golang.org/x/text/encoding/charmap package.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

func main() {
	// Open windows-1250 encoded file.
	f, err := os.Open("win1250.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read all content in raw form.
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	content := string(b)

	fmt.Println("Without decode: " + content)

	// Decode to Unicode.
	decoder := charmap.Windows1250.NewDecoder()
	reader := decoder.Reader(strings.NewReader(content))
	b, err = ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decoded: " + string(b))
}


// This program demonstrates writing Unicode text "Gdańsk" to a file "out.txt"
// encoded in Windows-1250 charset using golang.org/x/text/encoding/charmap package.
package main

import (
	"io"
	"os"

	"golang.org/x/text/encoding/charmap"
)

func main() {
	f, err := os.OpenFile("out.txt", os.O_CREATE|os.O_RDWR, os.ModePerm|os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Encode Unicode to Windows-1250.
	encoder := charmap.Windows1250.NewEncoder()
	writer := encoder.Writer(f)
	io.WriteString(writer, "Gdańsk")
}

// This program demonstrates various string manipulations such as case conversion,
// matching case-insensitive strings, and converting snake_case to camelCase.
package main

import (
	"fmt"
	"strings"
	"unicode"
)

const (
	email     = "ExamPle@domain.com"
	name      = "isaac newton"
	upc       = "upc"
	i         = "i"
	snakeCase = "first_name"
)

func main() {
	// Compare email case insensitively.
	input := "Example@domain.com"
	input = strings.ToLower(input)
	emailToCompare := strings.ToLower(email)
	matches := input == emailToCompare
	fmt.Printf("Email matches: %t\n", matches)

	// Convert to upper case.
	upcCode := strings.ToUpper(upc)
	fmt.Println("UPPER case: " + upcCode)

	// Convert to upper case and title case.
	str := "ǳ"
	fmt.Printf("%s in upper: %s and title: %s \n",
		str,
		strings.ToUpper(str),
		strings.ToTitle(str))

	// Compare ToTitle and ToTitleSpecial functions.
	title := strings.ToTitle(i)
	titleTurk := strings.ToTitleSpecial(unicode.TurkishCase, i)
	if title != titleTurk {
		fmt.Printf("ToTitle is different: %#U vs. %#U \n",
			title[0],
			[]rune(titleTurk)[0])
	}

	// Correct the case of a name.
	correctNameCase := strings.Title(name)
	fmt.Println("Corrected name: " + correctNameCase)

	// Convert snake_case to camelCase.
	firstNameCamel := toCamelCase(snakeCase)
	fmt.Println("Camel case: " + firstNameCamel)
}

func toCamelCase(input string) string {
	titleSpace := strings.Title(strings.Replace(input, "_", " ", -1))
	camel := strings.ReplaceAll(titleSpace, " ", "")
	return strings.ToLower(camel[:1]) + camel[1:]
}


// This program reads a CSV file, ignores lines starting with '#', and ensures each record has exactly 3 fields.
package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// Open the CSV file.
	file, err := os.Open("data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a CSV reader.
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 3 // Ensure each record has exactly 3 fields.
	reader.Comment = '#'       // Ignore lines starting with '#'.

	// Read and print each record.
	for {
		record, err := reader.Read()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(record)
	}
}



// This program reads a CSV file where the field delimiter is ';' instead of the default ','.
package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// Open the CSV file.
	file, err := os.Open("data_uncommon.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a CSV reader with ';' as the delimiter.
	reader := csv.NewReader(file)
	reader.Comma = ';'

	// Read and print each record.
	for {
		record, err := reader.Read()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(record)
	}
}



// This program demonstrates various string manipulation operations such as trimming whitespace,
// replacing multiple spaces with a single space, and padding strings with spaces based on alignment.
package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Example of trimming leading and trailing whitespace.
	stringToTrim := "\t\t\n   Go \tis\t Awesome \t\t"
	trimResult := strings.TrimSpace(stringToTrim)
	fmt.Println("Trimmed:", trimResult)

	// Example of replacing multiple spaces with a single space.
	stringWithSpaces := "\t\t\n   Go \tis\n Awesome \t\t"
	r := regexp.MustCompile("\\s+")
	replace := r.ReplaceAllString(stringWithSpaces, " ")
	fmt.Println("Spaces replaced:", replace)

	// Examples of padding strings with spaces based on alignment.
	needSpace := "need space"
	fmt.Println("Center padded:", pad(needSpace, 14, "CENTER"))
	fmt.Println("Left padded:", pad(needSpace, 14, "LEFT"))
}

// pad function pads the input string with spaces to achieve the desired length and alignment.
func pad(input string, padLen int, align string) string {
	inputLen := len(input)

	if inputLen >= padLen {
		return input
	}

	repeat := padLen - inputLen
	var output string
	switch align {
	case "RIGHT":
		output = fmt.Sprintf("% "+strconv.Itoa(-padLen)+"s", input)
	case "LEFT":
		output = fmt.Sprintf("% "+strconv.Itoa(padLen)+"s", input)
	case "CENTER":
		bothRepeat := float64(repeat) / float64(2)
		left := int(math.Floor(bothRepeat)) + inputLen
		right := int(math.Ceil(bothRepeat))
		output = fmt.Sprintf("% "+strconv.Itoa(left)+"s% "+strconv.Itoa(right)+"s", input, "")
	}
	return output
}



// This program demonstrates functions for indenting and unindenting strings based on spaces and runes.
package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	// Example of indenting the text by prefixing with spaces.
	text := "Hi! Go is awesome."
	text = Indent(text, 6)
	fmt.Println("Indented:", text)

	// Example of unindenting the text by removing a specified number of leading spaces.
	text = Unindent(text, 3)
	fmt.Println("Unindented:", text)

	// Trying to unindent more than the current indent level won't affect the string.
	text = Unindent(text, 10)
	fmt.Println("Unindented beyond limit:", text)

	// Example of indenting the text by prefixing with a specified rune.
	text = IndentByRune(text, 10, '.')
	fmt.Println("Indented by rune:", text)
}

// Indent adds spaces to the beginning of the input string to achieve the desired indentation level.
func Indent(input string, indent int) string {
	padding := indent + len(input)
	return fmt.Sprintf("% "+strconv.Itoa(padding)+"s", input)
}

// Unindent removes a specified number of leading spaces from the input string.
// If the input is indented by fewer spaces than the specified indent, it removes all leading spaces.
func Unindent(input string, indent int) string {
	count := 0
	for _, val := range input {
		if unicode.IsSpace(val) {
			count++
		}
		if count == indent || !unicode.IsSpace(val) {
			break
		}
	}
	return input[count:]
}

// IndentByRune adds a specified rune at the beginning of the input string to achieve the desired indentation level.
func IndentByRune(input string, indent int, r rune) string {
	return strings.Repeat(string(r), indent) + input
}







// This code demonstrates how to parse string representations of different numeric types (decimal, hexadecimal, binary, and floating-point) into their respective numeric values using the strconv package.

package main

import (
	"fmt"
	"strconv"
)

const bin = "00001"
const hex = "2f"
const intString = "12"
const floatString = "12.3"

func main() {

	// Decimals
	res, err := strconv.Atoi(intString)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed integer: %d\n", res)

	// Parsing hexadecimals
	res64, err := strconv.ParseInt(hex, 16, 32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed hexadecimal: %d\n", res64)

	// Parsing binary values
	resBin, err := strconv.ParseInt(bin, 2, 32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed binary: %d\n", resBin)

	// Parsing floating points
	resFloat, err := strconv.ParseFloat(floatString, 32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed float: %.5f\n", resFloat)

}



// This code demonstrates the comparison of floating-point numbers using both float64 and big.Float types. It highlights the precision control with big.Float and its impact on comparison results.

package main

import (
	"fmt"
	"math/big"
)

var da float64 = 0.299999992
var db float64 = 0.299999991

var prec uint = 32
var prec2 uint = 16

func main() {

	fmt.Printf("Comparing float64 with '==' equals: %v\n", da == db)

	daB := big.NewFloat(da).SetPrec(prec)
	dbB := big.NewFloat(db).SetPrec(prec)

	fmt.Printf("A: %v \n", daB)
	fmt.Printf("B: %v \n", dbB)
	fmt.Printf("Comparing big.Float with precision: %d : %v\n", prec, daB.Cmp(dbB) == 0)

	daB = big.NewFloat(da).SetPrec(prec2)
	dbB = big.NewFloat(db).SetPrec(prec2)

	fmt.Printf("A: %v \n", daB)
	fmt.Printf("B: %v \n", dbB)
	fmt.Printf("Comparing big.Float with precision: %d : %v\n", prec2, daB.Cmp(dbB) == 0)

}



// This code demonstrates different methods for comparing floating-point numbers, including string formatting and numerical tolerance to handle precision limitations.

package main

import (
	"fmt"
	"math"
)

const da = 0.29999999999999998889776975374843459576368331909180
const db = 0.3

func main() {

	daStr := fmt.Sprintf("%.10f", da)
	dbStr := fmt.Sprintf("%.10f", db)

	// While formatting the number to string
	// it is rounded to 3.
	fmt.Printf("Strings %s = %s equals: %v \n", daStr, dbStr, dbStr == daStr)

	// Numbers are not equal
	fmt.Printf("Number equals: %v \n", db == da)

	// As the precision of float representation
	// is limited. For the float comparison it is
	// better to use comparison with some tolerance.
	fmt.Printf("Number equals with TOLERANCE: %v \n", Equals(da, db))

}

const TOLERANCE = 1e-8

// Equals compares the floating point numbers
// with tolerance 1e-8
func Equals(numA, numB float64) bool {
	delta := math.Abs(numA - numB)
	return delta < TOLERANCE
}



// This code demonstrates the difference between truncating a floating-point number to an integer and rounding it properly using a custom function.

package main

import (
	"fmt"
	"math"
)

var valA float64 = 3.55554444

func main() {

	// Bad assumption on rounding
	// the number by casting it to
	// integer.
	intVal := int(valA)
	fmt.Printf("Bad rounding by casting to int: %v\n", intVal)

	fRound := Round(valA)
	fmt.Printf("Rounding by custom function: %v\n", fRound)

}

// Round returns the nearest integer.
func Round(x float64) float64 {
	t := math.Trunc(x)
	if math.Abs(x-t) >= 0.5 {
		return t + math.Copysign(1, x)
	}
	return t
}



// This code demonstrates high-precision arithmetic using the math/big package to perform calculations on large floating-point numbers, including the calculation of a circle's circumference and basic arithmetic operations.

package main

import (
	"fmt"
	"math/big"
)

const PI = `3.14159265358979323846264338327950288419716939937510582097494459230781640628620899862803482534211706798214808651328230664709384460955058223172535940812848111745028410270193852110555964462294895493038196`
const diameter = 3.0
const precision = 400

func main() {

	pi, _ := new(big.Float).SetPrec(precision).SetString(PI)
	d := new(big.Float).SetPrec(precision).SetFloat64(diameter)

	circumference := new(big.Float).Mul(pi, d)

	pi64, _ := pi.Float64()
	fmt.Printf("Circumference big.Float = %.100f\n", circumference)
	fmt.Printf("Circumference float64   = %.100f\n", pi64*diameter)

	sum := new(big.Float).Add(pi, pi)
	fmt.Printf("Sum = %.100f\n", sum)

	diff := new(big.Float).Sub(pi, pi)
	fmt.Printf("Diff = %.100f\n", diff)

	quo := new(big.Float).Quo(pi, pi)
	fmt.Printf("Quotient = %.100f\n", quo)

}



// This code demonstrates various formatting options using fmt.Printf to print integers and floating-point numbers in different ways, including different bases, padding, and scientific notation.

package main

import (
	"fmt"
)

var integer int64 = 32500
var floatNum float64 = 22000.456

func main() {

	// Common way to print the decimal number
	fmt.Printf("%d \n", integer)

	// Always show the sign
	fmt.Printf("%+d \n", integer)

	// Print in other bases: x - 16, o - 8, b - 2, d - 10
	fmt.Printf("%x \n", integer)
	fmt.Printf("%#x \n", integer)

	// Padding with leading zeros
	fmt.Printf("%010d \n", integer)

	// Left padding with spaces
	fmt.Printf("% 10d \n", integer)

	// Right padding
	fmt.Printf("% -10d \n", integer)

	// Print floating point number
	fmt.Printf("%f \n", floatNum)

	// Floating point number with limited precision = 5
	fmt.Printf("%.5f \n", floatNum)

	// Floating point number in scientific notation
	fmt.Printf("%e \n", floatNum)

	// Floating point number in %e for large exponents or %f otherwise
	fmt.Printf("%g \n", floatNum)

}



// This code demonstrates formatting numbers according to different locales using the golang.org/x/text/message package.

package main

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

const num = 100000.5678

func main() {
	p := message.NewPrinter(language.English)
	p.Printf("%.2f \n", num)

	p = message.NewPrinter(language.German)
	p.Printf("%.2f \n", num)
}



// This code demonstrates converting integer values between different bases (binary, hexadecimal, octal, and decimal) using custom functions.

package main

import (
	"fmt"
	"strconv"
)

const bin = "10111"
const hex = "1A"
const oct = "12"
const dec = "10"

func main() {

	// Converts binary value into hex
	v, _ := ConvertInt(bin, 2, 16)
	fmt.Printf("Binary value %s converted to hex: %s\n", bin, v)

	// Converts hex value into dec
	v, _ = ConvertInt(hex, 16, 10)
	fmt.Printf("Hex value %s converted to dec: %s\n", hex, v)

	// Converts oct value into hex
	v, _ = ConvertInt(oct, 8, 16)
	fmt.Printf("Oct value %s converted to hex: %s\n", oct, v)

	// Converts dec value into oct
	v, _ = ConvertInt(dec, 10, 8)
	fmt.Printf("Dec value %s converted to oct: %s\n", dec, v)

	// Analogically, any other conversion could be done.

}

// ConvertInt converts the given string value of base to defined toBase.
func ConvertInt(val string, base, toBase int) (string, error) {
	i, err := strconv.ParseInt(val, base, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, toBase), nil
}



// This code demonstrates handling pluralization and localization of messages using the golang.org/x/text/message package, customizing messages based on variable values and locale settings.

package main

import (
	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {

	// Set pluralization rules for the message "%d items to do" in English
	message.Set(language.English, "%d items to do",
		plural.Selectf(1, "%d",
			"=0", "no items to do",
			plural.One, "one item to do",
			"<100", "%[1]d items to do",
			plural.Other, "lot of items to do",
		))

	// Set pluralization rules for the message "The average is %.2f" in English
	message.Set(language.English, "The average is %.2f",
		plural.Selectf(1, "%.2f",
			"<1", "The average is zero",
			"=1", "The average is one",
			plural.Other, "The average is %[1]f ",
		))

	// Create a new printer for the English language
	prt := message.NewPrinter(language.English)

	// Print messages based on the set pluralization rules
	prt.Printf("%d items to do", 0)
	prt.Println()
	prt.Printf("%d items to do", 1)
	prt.Println()
	prt.Printf("%d items to do", 10)
	prt.Println()
	prt.Printf("%d items to do", 1000)
	prt.Println()

	prt.Printf("The average is %.2f", 0.8)
	prt.Println()
	prt.Printf("The average is %.2f", 1.0)
	prt.Println()
	prt.Printf("The average is %.2f", 10.0)
	prt.Println()

}



// This code demonstrates the generation of random numbers using both the math/rand package and crypto/rand package.
// It compares sequences generated by math/rand and ensures cryptographic random numbers from crypto/rand are distinct.

package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
)

func main() {

	// Using math/rand package
	sec1 := rand.New(rand.NewSource(10))
	sec2 := rand.New(rand.NewSource(10))
	for i := 0; i < 5; i++ {
		rnd1 := sec1.Int()
		rnd2 := sec2.Int()
		if rnd1 != rnd2 {
			fmt.Println("Rand generated non-equal sequence")
			break
		} else {
			fmt.Printf("Math/Rand1: %d , Math/Rand2: %d\n", rnd1, rnd2)
		}
	}

	// Using crypto/rand package
	for i := 0; i < 5; i++ {
		safeNum := NewCryptoRand()
		safeNum2 := NewCryptoRand()
		if safeNum == safeNum2 {
			fmt.Println("Crypto generated equal numbers")
			break
		} else {
			fmt.Printf("Crypto/Rand1: %d , Crypto/Rand2: %d\n", safeNum, safeNum2)
		}
	}
}

// NewCryptoRand generates a random number using crypto/rand package.
func NewCryptoRand() int64 {
	safeNum, err := rand.Int(rand.Reader, big.NewInt(100234))
	if err != nil {
		panic(err)
	}
	return safeNum.Int64()
}



// This code demonstrates basic operations and functions related to complex numbers in Go.

package main

import (
	"fmt"
	"math/cmplx"
)

func main() {

	// Complex numbers are defined with real and imaginary parts as float64.
	a := complex(2, 3)

	fmt.Printf("Real part: %f \n", real(a))
	fmt.Printf("Imaginary part: %f \n", imag(a))

	b := complex(6, 4)

	// Basic arithmetic operations on complex numbers
	c := a - b
	fmt.Printf("Difference : %v\n", c)
	c = a + b
	fmt.Printf("Sum : %v\n", c)
	c = a * b
	fmt.Printf("Product : %v\n", c)
	c = a / b
	fmt.Printf("Quotient : %v\n", c)

	// Calculating conjugate of a complex number
	conjugate := cmplx.Conj(a)
	fmt.Println("Complex number a's conjugate : ", conjugate)

	// Calculating cosine of a complex number
	cos := cmplx.Cos(b)
	fmt.Println("Cosine of b : ", cos)

}



// This program demonstrates conversion between radians and degrees using both standalone functions and type methods in Go.

package main

import (
	"fmt"
	"math"
)

type Radian float64

func (rad Radian) ToDegrees() Degree {
	return Degree(float64(rad) * (180.0 / math.Pi))
}

func (rad Radian) Float64() float64 {
	return float64(rad)
}

type Degree float64

func (deg Degree) ToRadians() Radian {
	return Radian(float64(deg) * (math.Pi / 180.0))
}

func (deg Degree) Float64() float64 {
	return float64(deg)
}

func main() {

	// Using standalone functions for conversion
	val := radiansToDegrees(1)
	fmt.Printf("One radian is : %.4f degrees\n", val)

	val2 := degreesToRadians(val)
	fmt.Printf("%.4f degrees is %.4f rad\n", val, val2)

	// Using type methods for conversion
	val = Radian(1).ToDegrees().Float64()
	fmt.Printf("Degrees: %.4f degrees\n", val)

	val = Degree(val).ToRadians().Float64()
	fmt.Printf("Rad: %.4f radians\n", val)
}

// Function to convert degrees to radians
func degreesToRadians(deg float64) float64 {
	return deg * (math.Pi / 180.0)
}

// Function to convert radians to degrees
func radiansToDegrees(rad float64) float64 {
	return rad * (180.0 / math.Pi)
}



// This program demonstrates logarithmic functions in Go:
// Ln (natural logarithm), Log10 (base-10 logarithm), Log2 (base-2 logarithm),
// and a custom logarithm function (Log) with a specified base.
package main

import (
	"fmt"
	"math"
)

func main() {

	ln := math.Log(math.E)
	fmt.Printf("Ln(E) = %.4f\n", ln)

	log10 := math.Log10(-100)
	fmt.Printf("Log10(10) = %.4f\n", log10)

	log2 := math.Log2(2)
	fmt.Printf("Log2(2) = %.4f\n", log2)

	log_3_6 := Log(3, 6)
	fmt.Printf("Log3(6) = %.4f\n", log_3_6)

}

// Log computes the logarithm of base > 1 and x greater 0
func Log(base, x float64) float64 {
	return math.Log(x) / math.Log(base)
}



// This program calculates MD5 checksums for a string and a file.

package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

var content = "This is content to check"

func main() {

	checksum := MD5(content)
	checksum2 := FileMD5("content.dat")

	fmt.Printf("Checksum 1: %s\n", checksum)
	fmt.Printf("Checksum 2: %s\n", checksum2)
	if checksum == checksum2 {
		fmt.Println("Content matches!!!")
	}

}

// MD5 calculates the MD5 hash of the given data and returns it as a hex-encoded string.
func MD5(data string) string {
	h := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", h)
}

// FileMD5 calculates the MD5 hash of a file's content and returns it as a hex-encoded string.
func FileMD5(path string) string {
	h := md5.New()
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = io.Copy(h, f)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}



// This program demonstrates how to create a new SHA-1 hash instance using the crypto package in Go.

package main

import (
	"crypto"
)

func main() {
	crypto.SHA1.New()
}



// This program demonstrates how to retrieve and print the current date and time using the time package in Go.

package main

import (
	"fmt"
	"time"
)

func main() {

	today := time.Now()
	fmt.Println(today)

}



// This program demonstrates various ways to format a specific time value using the time package in Go.

package main

import (
	"fmt"
	"time"
)

func main() {
	tTime := time.Date(2017, time.March, 5, 8, 5, 2, 0, time.Local)

	// Formatting with a custom layout
	fmt.Printf("tTime is: %s\n", tTime.Format("2006/1/2"))

	// Formatting hours and minutes
	fmt.Printf("The time is: %s\n", tTime.Format("15:04"))

	// Using predefined RFC1123 format
	fmt.Printf("The time is: %s\n", tTime.Format(time.RFC1123))

	// Space padding for days (Go 1.9.2+)
	fmt.Printf("tTime is: %s\n", tTime.Format("2006/1/_2"))

	// Zero-padding for days, months, and hours
	fmt.Printf("tTime is: %s\n", tTime.Format("2006/01/02"))

	// Fractional seconds with leading zeros
	fmt.Printf("tTime is: %s\n", tTime.Format("15:04:05.00"))

	// Fractional seconds without leading zeros
	fmt.Printf("tTime is: %s\n", tTime.Format("15:04:05.999"))

	// AppendFormat example
	fmt.Println(string(tTime.AppendFormat([]byte("The time is up: "), "03:04PM")))
}



// This program demonstrates parsing of date and time strings using the time package in Go,
// handling different time zone scenarios with Parse and ParseInLocation functions.

package main

import (
	"fmt"
	"time"
)

func main() {

	// If timezone is not defined, Parse function returns the time in UTC timezone.
	t, err := time.Parse("2/1/2006", "31/7/2015")
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

	// If timezone is given, it is parsed in the specified timezone.
	t, err = time.Parse("2/1/2006  3:04 PM MST", "31/7/2015  1:25 AM DST")
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

	// ParseInLocation parses the time in the given location if the string does not contain time zone definition.
	t, err = time.ParseInLocation("2/1/2006  3:04 PM ", "31/7/2015  1:25 AM ", time.Local)
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

}



// This program demonstrates how to work with epoch time (Unix time) using the time package in Go.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Set the epoch time from int64
	t := time.Unix(0, 0)
	fmt.Println(t)

	// Get the epoch time from a Time instance
	epoch := t.Unix()
	fmt.Println(epoch)

	// Current epoch time in seconds
	epochNow := time.Now().Unix()
	fmt.Printf("Epoch time in seconds: %d\n", epochNow)

	// Current epoch time in nanoseconds
	epochNano := time.Now().UnixNano()
	fmt.Printf("Epoch time in nano-seconds: %d\n", epochNano)

}



// This program demonstrates extracting date and time units from a specific time instance using the time package in Go.

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2017, 11, 29, 21, 0, 0, 0, time.Local)
	fmt.Printf("Extracting units from: %v\n", t)

	dOfMonth := t.Day()
	weekDay := t.Weekday()
	month := t.Month()

	fmt.Printf("The %dth day of %v is %v\n", dOfMonth, month, weekDay)
}



// This program demonstrates manipulating dates and times using the time package in Go,
// including adding and subtracting durations and using a more convenient API for adding years, months, and days.

package main

import (
	"fmt"
	"time"
)

func main() {

	l, err := time.LoadLocation("Europe/Vienna")
	if err != nil {
		panic(err)
	}
	t := time.Date(2017, 11, 30, 11, 10, 20, 0, l)
	fmt.Printf("Default date is: %v\n", t)

	// Add 3 days
	r1 := t.Add(72 * time.Hour)
	fmt.Printf("Default date +3 days is: %v\n", r1)

	// Subtract 3 days
	r1 = t.Add(-72 * time.Hour)
	fmt.Printf("Default date -3 days is: %v\n", r1)

	// Using AddDate to add years, months, and days
	r1 = t.AddDate(1, 3, 2)
	fmt.Printf("Default date +1 year +3 months +2 days is: %v\n", r1)
}



// This program demonstrates calculating durations between dates using the time package in Go,
// including calculating durations between specific dates, from a date to the present, and from the present to a date.

package main

import (
	"fmt"
	"time"
)

func main() {

	l, err := time.LoadLocation("Europe/Vienna")
	if err != nil {
		panic(err)
	}

	t := time.Date(2000, 1, 1, 0, 0, 0, 0, l)
	t2 := time.Date(2000, 1, 3, 0, 0, 0, 0, l)
	fmt.Printf("First Default date is %v\n", t)
	fmt.Printf("Second Default date is %v\n", t2)

	dur := t2.Sub(t)
	fmt.Printf("The duration between t and t2 is %v\n", dur)

	dur = time.Since(t)
	fmt.Printf("The duration between now and t is %v\n", dur)

	dur = time.Until(t)
	fmt.Printf("The duration between t and now is %v\n", dur)

}



// This program demonstrates how to convert a time from one timezone (Europe/Vienna) to another (America/Phoenix) using Go's time package and the In() method.

package main

import (
	"fmt"
	"time"
)

func main() {
	eur, err := time.LoadLocation("Europe/Vienna")
	if err != nil {
		panic(err)
	}

	t := time.Date(2000, 1, 1, 0, 0, 0, 0, eur)
	fmt.Printf("Original Time: %v\n", t)

	phx, err := time.LoadLocation("America/Phoenix")
	if err != nil {
		panic(err)
	}

	t2 := t.In(phx)
	fmt.Printf("Converted Time: %v\n", t2)
}



// This program demonstrates how to handle OS signals to gracefully stop a goroutine that uses a ticker.

package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {

	c := make(chan os.Signal, 1)
	signal.Notify(c)

	ticker := time.NewTicker(time.Second)
	stop := make(chan bool)

	go func() {
		defer func() { stop <- true }()
		for {
			select {
			case <-ticker.C:
				fmt.Println("Tick")
			case <-stop:
				fmt.Println("Goroutine closing")
				return
			}
		}
	}()

	// Block until the signal is received
	<-c
	ticker.Stop()

	// Stop the goroutine
	stop <- true
	// Wait until the goroutine stops
	<-stop
	fmt.Println("Application stopped")
}



// This program demonstrates different ways to wait for a duration using time.Timer, time.AfterFunc, and time.After.

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	t := time.NewTimer(3 * time.Second)
	fmt.Printf("Start waiting at %v\n", time.Now().Format(time.UnixDate))
	<-t.C
	fmt.Printf("Code executed at %v\n", time.Now().Format(time.UnixDate))

	wg := &sync.WaitGroup{}
	wg.Add(1)
	fmt.Printf("Start waiting for AfterFunc at %v\n", time.Now().Format(time.UnixDate))
	time.AfterFunc(3*time.Second, func() {
		fmt.Printf("Code executed for AfterFunc at %v\n", time.Now().Format(time.UnixDate))
		wg.Done()
	})

	wg.Wait()

	fmt.Printf("Waiting on time.After at %v\n", time.Now().Format(time.UnixDate))
	<-time.After(3 * time.Second)
	fmt.Printf("Code resumed at %v\n", time.Now().Format(time.UnixDate))

}



// This program demonstrates how to insert items into a list until a timeout using time.After.

package main

import (
	"fmt"
	"time"
)

func main() {

	to := time.After(3 * time.Second)
	list := make([]string, 0)
	done := make(chan bool, 1)

	fmt.Println("Starting to insert items")
	go func() {
		defer fmt.Println("Exiting goroutine")
		for {
			select {
			case <-to:
				fmt.Println("The time is up")
				done <- true
				return
			default:
				list = append(list, time.Now().String())
			}
		}
	}()

	<-done
	fmt.Printf("Managed to insert %d items\n", len(list))

}



// This program demonstrates how to serialize and deserialize time.Time values using JSON encoding.

package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {

	// Load the Europe/Vienna time zone location
	eur, err := time.LoadLocation("Europe/Vienna")
	if err != nil {
		panic(err)
	}

	// Create a time instance in the Europe/Vienna time zone
	t := time.Date(2017, 11, 20, 11, 20, 10, 0, eur)

	// Serialize as RFC 3339
	b, err := t.MarshalJSON()
	if err != nil {
		panic(err)
	}
	fmt.Println("Serialized as RFC 3339:", string(b))

	// Deserialize from RFC 3339
	t2 := time.Time{}
	t2.UnmarshalJSON(b)
	fmt.Println("Deserialized from RFC 3339:", t2)

	// Serialize as epoch
	epoch := t.Unix()
	fmt.Println("Serialized as Epoch:", epoch)

	// Deserialize from epoch
	jsonStr := fmt.Sprintf("{ \"created\":%d }", epoch)
	data := struct {
		Created int64 `json:"created"`
	}{}
	json.Unmarshal([]byte(jsonStr), &data)
	deserialized := time.Unix(data.Created, 0)
	fmt.Println("Deserialized from Epoch:", deserialized)
}










// The code  prompts for and reads the user's name and age, then prints a greeting with that information.
package main

import (
	"fmt"
)

func main() {

	var name string
	fmt.Println("What is your name?")
	fmt.Scanf("%s\n", &name)

	var age int
	fmt.Println("What is your age?")
	fmt.Scanf("%d\n", &age)

	fmt.Printf("Hello %s, your age is %d\n", name, age)

}

// The code reads input from the user in 8-byte chunks, then prints the hexadecimal and string representation of each chunk.
package main

import (
	"fmt"
	"os"
)

func main() {

	for {
		data := make([]byte, 8)
		n, err := os.Stdin.Read(data)
		if err == nil && n > 0 {
			process(data)
		} else {
			break
		}
	}

}

func process(data []byte) {
	fmt.Printf("Received: %X 	%s\n", data, string(data))
}

// This code reads lines of input from the user and echoes each line back to the console prefixed with "Echo:".
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// The Scanner is able to
	// scan input by lines
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		txt := sc.Text()
		fmt.Printf("Echo: %s\n", txt)
	}

}

// This code writes strings to standard output and error, then writes a byte buffer to standard output repeatedly, followed by a newline.
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// Simply write string
	io.WriteString(os.Stdout,
		"This is string to standard output.\n")

	io.WriteString(os.Stderr,
		"This is string to standard error output.\n")

	// Stdout/err implements
	// writer interface
	buf := []byte{0xAF, 0xFF, 0xFE}
	for i := 0; i < 200; i++ {
		if _, e := os.Stdout.Write(buf); e != nil {
			panic(e)
		}
	}

	// The fmt package
	// could be used too
	fmt.Fprintln(os.Stdout, "\n")
}

// This code reads and prints a file's content, then creates or opens another file to write "Test string" into it.
//file.txt content: This is file content.
//test.txt content: Test string

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	f, err := os.Open("temp/file.txt")
	if err != nil {
		panic(err)
	}

	c, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("### File content ###\n%s\n", string(c))
	f.Close()

	f, err = os.OpenFile("temp/test.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	io.WriteString(f, "Test string")
	f.Close()

}


// Read "temp/file.txt" using bufio.Scanner line by line into bytes.Buffer,
// and print accumulated content. Then read entire file using ioutil.ReadFile,
// convert to string, and print directly, achieving the same file reading goal.
/*file content: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris id pretium eros. Aliquam imperdiet mi ut elit faucibus porta.
Donec facilisis nunc at risus dapibus elementum.
*/
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("### Read as reader ###")
	f, err := os.Open("temp/file.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read the
	// file with reader
	wr := bytes.Buffer{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		wr.WriteString(sc.Text())
	}
	fmt.Println(wr.String())

	fmt.Println("### ReadFile ###")
	// for smaller files
	fContent, err := ioutil.ReadFile("temp/file.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(fContent))

}

// This code encodes a string in Windows-1252 encoding, writes it to "example.txt", then reads and decodes it back to UTF-8, printing the decoded content.
//example.txt content: This is sample text with runes Š

package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/text/encoding/charmap"
)

func main() {

	// Write the string
	// encoded to Windows-1252
	encoder := charmap.Windows1252.NewEncoder()
	s, e := encoder.String("This is sample text with runes Š")
	if e != nil {
		panic(e)
	}
	ioutil.WriteFile("example.txt", []byte(s), os.ModePerm)

	// Decode to UTF-8
	f, e := os.Open("example.txt")
	if e != nil {
		panic(e)
	}
	defer f.Close()
	decoder := charmap.Windows1252.NewDecoder()
	reader := decoder.Reader(f)
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}



//This Go program manipulates a fixed-format flat file (flatfile.txt) with each line //of 25 characters.
//It reads specific lines, writes data into columns ("id", "first", "last"), and //manages file operations and data errors.
//It showcases structured data manipulation in a flat file, emphasizing basic file //handling and error management.
/*
content of flatfile.txt: 123.Jun.......Wong......
12..Novak.....Jurgen....
10..Radomir...Sohlich...                                                                                                                                                                                    Andrew....
*/
package main

import (
	"errors"
	"fmt"
	"os"
)

const lineLegth = 25

func main() {

	f, e := os.OpenFile("flatfile.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if e != nil {
		panic(e)
	}
	defer f.Close()

	fmt.Println(readRecords(2, "last", f))
	if err := writeRecord(2, "first", "Radomir", f); err != nil {
		panic(err)
	}
	fmt.Println(readRecords(2, "first", f))
	if err := writeRecord(10, "first", "Andrew", f); err != nil {
		panic(err)
	}
	fmt.Println(readRecords(10, "first", f))
	fmt.Println(readLine(2, f))
}

func readLine(line int, f *os.File) (string, error) {
	lineBuffer := make([]byte, 24)
	f.Seek(int64(line*lineLegth), 0)
	_, err := f.Read(lineBuffer)
	return string(lineBuffer), err
}

func writeRecord(line int, column, dataStr string, f *os.File) error {
	definedLen := 10
	position := int64(line * lineLegth)
	switch column {
	case "id":
		definedLen = 4
	case "first":
		position += 4
	case "last":
		position += 14
	default:
		return errors.New("Column not defined")
	}

	if len([]byte(dataStr)) > definedLen {
		return fmt.Errorf("Maximum length for '%s' is %d", column, definedLen)
	}

	data := make([]byte, definedLen)
	for i := range data {
		data[i] = '.'
	}
	copy(data, []byte(dataStr))
	_, err := f.WriteAt(data, position)
	return err
}

func readRecords(line int, column string, f *os.File) (string, error) {
	lineBuffer := make([]byte, 24)
	f.ReadAt(lineBuffer, int64(line*lineLegth))
	var retVal string
	switch column {
	case "id":
		return string(lineBuffer[:3]), nil
	case "first":
		return string(lineBuffer[4:13]), nil
	case "last":
		return string(lineBuffer[14:23]), nil
	}

	return retVal, errors.New("Column not defined")
}

// Writing binary values: writes a float64 (1.004) and a string ("Hello") to a buffer using binary encoding.
// Reading the written values: reads a float64 and a string from the buffer and prints them formatted.
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {

	// Writing binary values
	buf := bytes.NewBuffer([]byte{})
	if err := binary.Write(buf, binary.BigEndian, 1.004); err != nil {
		panic(err)
	}
	if err := binary.Write(buf, binary.BigEndian, []byte("Hello")); err != nil {
		panic(err)
	}

	// Reading the written values
	var num float64
	if err := binary.Read(buf, binary.BigEndian, &num); err != nil {
		panic(err)
	}
	fmt.Printf("float64: %.3f\n", num)
	greeting := make([]byte, 5)
	if err := binary.Read(buf, binary.BigEndian, &greeting); err != nil {
		panic(err)
	}
	fmt.Printf("string: %s\n", string(greeting))
}

// Creates a buffer and a file "sample.txt", writes a string into both using MultiWriter,
// then prints the contents of the buffer.
//sample.txt content: Hello, Go is awesome!
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {

	buf := bytes.NewBuffer([]byte{})
	f, err := os.OpenFile("sample.txt", os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	wr := io.MultiWriter(buf, f)
	_, err = io.WriteString(wr, "Hello, Go is awesome!")
	if err != nil {
		panic(err)
	}

	fmt.Println("Content of buffer: " + buf.String())
}

// Uses a pipe to capture output from executing "echo Hello Go!\nThis is example",
// then prints the output to the console using io.Copy and goroutines.
package main

import (
	"io"
	"log"
	"os"
	"os/exec"
)

func main() {
	pReader, pWriter := io.Pipe()

	cmd := exec.Command("echo", "Hello Go!\nThis is example")
	cmd.Stdout = pWriter

	go func() {
		defer pReader.Close()
		if _, err := io.Copy(os.Stdout, pReader); err != nil {
			log.Fatal(err)
		}
	}()

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

}

// Encodes a User struct into a byte buffer using gob encoding,
// then decodes it back into a User struct and prints its string representation.
// Also tries to decode into a SimpleUser struct but fails due to data format mismatch.
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type User struct {
	FirstName string
	LastName  string
	Age       int
	Active    bool
}

func (u User) String() string {
	return fmt.Sprintf(`{"FirstName":%s,"LastName":%s,"Age":%d,"Active":%v }`,
		u.FirstName, u.LastName, u.Age, u.Active)
}

type SimpleUser struct {
	FirstName string
	LastName  string
}

func (u SimpleUser) String() string {
	return fmt.Sprintf(`{"FirstName":%s,"LastName":%s}`,
		u.FirstName, u.LastName)
}

func main() {

	var buff bytes.Buffer

	// Encode value
	enc := gob.NewEncoder(&buff)
	user := User{
		"Radomir",
		"Sohlich",
		30,
		true,
	}
	enc.Encode(user)
	fmt.Printf("%X\n", buff.Bytes())

	// Decode value
	out := User{}
	dec := gob.NewDecoder(&buff)
	dec.Decode(&out)
	fmt.Println(out.String())

	enc.Encode(user)
	out2 := SimpleUser{}
	dec.Decode(&out2)
	fmt.Println(out2.String())

}

// Compresses "This is my file content" into a ZIP file "data.zip" and writes it to disk.
// Then decompresses "data.zip", reads "newfile.txt", and prints its content to stdout.
package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	var buff bytes.Buffer

	// Compress content
	zipW := zip.NewWriter(&buff)
	f, err := zipW.Create("newfile.txt")
	if err != nil {
		panic(err)
	}
	_, err = f.Write([]byte("This is my file content"))
	if err != nil {
		panic(err)
	}
	err = zipW.Close()
	if err != nil {
		panic(err)
	}

	//Write output to file
	err = ioutil.WriteFile("data.zip", buff.Bytes(), os.ModePerm)
	if err != nil {
		panic(err)
	}

	// Decompress the content
	zipR, err := zip.OpenReader("data.zip")
	if err != nil {
		panic(err)
	}

	for _, file := range zipR.File {
		fmt.Println("File " + file.Name + " contains:")
		r, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(os.Stdout, r)
		if err != nil {
			panic(err)
		}
		err = r.Close()
		if err != nil {
			panic(err)
		}
		fmt.Println()
	}

}

// Reads and parses XML data from "data.xml" into a slice of Book structs,
// using xml.Decoder to decode each <book> element into a Book struct and prints them.
package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Book struct {
	Title  string `xml:"title"`
	Author string `xml:"author"`
}

func main() {

	f, err := os.Open("data.xml")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	decoder := xml.NewDecoder(f)

	// Read the book one by one
	books := make([]Book, 0)
	for {
		tok, err := decoder.Token()
		if err != nil {
			panic(err)
		}
		if tok == nil {
			break
		}
		switch tp := tok.(type) {
		case xml.StartElement:
			if tp.Name.Local == "book" {
				// Decode the element to struct
				var b Book
				decoder.DecodeElement(&b, &tp)
				books = append(books, b)
			}
		}
	}
	fmt.Println(books)
}

// Parses JSON data from a constant string js representing an array of User objects,
// using json.NewDecoder to decode each object into a User struct and prints the slice of users.
package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

const js = `
	[{
		"name":"Axel",
		"lastname":"Fooley"
	},
	{
		"name":"Tim",
		"lastname":"Burton"
	},
	{
		"name":"Tim",
		"lastname":"Burton"
`

type User struct {
	Name     string `json:"name"`
	LastName string `json:"lastname"`
}

func main() {

	userSlice := make([]User, 0)
	r := strings.NewReader(js)
	dec := json.NewDecoder(r)
	for {
		tok, err := dec.Token()
		if err != nil {
			break
		}
		if tok == nil {
			break
		}
		switch tp := tok.(type) {
		case json.Delim:
			str := tp.String()
			if str == "[" || str == "{" {
				for dec.More() {
					u := User{}
					err := dec.Decode(&u)
					if err == nil {
						userSlice = append(userSlice, u)
					} else {
						break
					}
				}
			}
		}
	}

	fmt.Println(userSlice)
}

// Opens and retrieves information about "test.file":
// prints its name, whether it's a directory, its size, and mode.
package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Open("test.file")
	if err != nil {
		panic(err)
	}
	fi, err := f.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Printf("File name: %v\n", fi.Name())
	fmt.Printf("Is Directory: %t\n", fi.IsDir())
	fmt.Printf("Size: %d\n", fi.Size())
	fmt.Printf("Mode: %v\n", fi.Mode())

}

// Creates a temporary file and directory using ioutil.TempFile and ioutil.TempDir respectively,
// printing their names. Deferred cleanup ensures removal after program execution.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	tFile, err := ioutil.TempFile("", "gostdcookbook")
	if err != nil {
		panic(err)
	}
	// The called is responsible for handling
	// the clean up.
	defer os.Remove(tFile.Name())

	fmt.Println(tFile.Name())

	// TempDir returns
	// the path in string.
	tDir, err := ioutil.TempDir("", "gostdcookbookdir")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tDir)
	fmt.Println(tDir)

}

// Creates a file "sample.file", writes "Go is awesome!" and "Yeah! Go is great." to it using os.Create and io.Copy,
// demonstrating file writing and copying in Go, with deferred file closure for cleanup.
package main

import (
	"io"
	"os"
	"strings"
)

func main() {

	f, err := os.Create("sample.file")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString("Go is awesome!\n")
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, strings.NewReader("Yeah! Go is great.\n"))
	if err != nil {
		panic(err)
	}
}

// Creates "sample.file" and writes concurrent greetings ("Hello!", "Ola!", "Ahoj!") using SyncWriter,
// ensuring thread safety with sync.Mutex and sync.WaitGroup.
package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

type SyncWriter struct {
	m      sync.Mutex
	Writer io.Writer
}

func (w *SyncWriter) Write(b []byte) (n int, err error) {
	w.m.Lock()
	defer w.m.Unlock()
	return w.Writer.Write(b)
}

var data = []string{
	"Hello!",
	"Ola!",
	"Ahoj!",
}

func main() {

	f, err := os.Create("sample.file")
	if err != nil {
		panic(err)
	}

	wr := &SyncWriter{sync.Mutex{}, f}
	wg := sync.WaitGroup{}
	for _, val := range data {
		wg.Add(1)
		go func(greetings string) {
			fmt.Fprintln(wr, greetings)
			wg.Done()
		}(val)
	}

	wg.Wait()
}

// Lists files and directories in the current directory using ioutil.ReadDir and filepath.Walk respectively.
// Walk lists recursively and skips directories, printing them in brackets.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {

	fmt.Println("List by ReadDir")
	listDirByReadDir(".")
	fmt.Println()
	fmt.Println("List by Walk")
	listDirByWalk(".")
}

func listDirByWalk(path string) {
	filepath.Walk(path, func(wPath string, info os.FileInfo, err error) error {

		// Walk the given dir
		// without printing out.
		if wPath == path {
			return nil
		}

		// If given path is folder
		// stop list recursively and print as folder.
		if info.IsDir() {
			fmt.Printf("[%s]\n", wPath)
			return filepath.SkipDir
		}

		// Print file name
		if wPath != path {
			fmt.Println(wPath)
		}
		return nil
	})
}

func listDirByReadDir(path string) {
	lst, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, val := range lst {
		if val.IsDir() {
			fmt.Printf("[%s]\n", val.Name())
		} else {
			fmt.Println(val.Name())
		}
	}
}

// Creates "test.file", retrieves its initial permissions using os.Create and f.Stat.
// Changes permissions to 0777 using f.Chmod and verifies the change with f.Stat, printing both sets of permissions.
package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Create("test.file")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Obtain current permissions
	fi, err := f.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Printf("File permissions %v\n", fi.Mode())

	// Change permissions
	err = f.Chmod(0777)
	if err != nil {
		panic(err)
	}
	fi, err = f.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Printf("File permissions %v\n", fi.Mode())

}

// Creates "created.file" using os.Create.
// Creates "created.byopen" with append mode using os.OpenFile.
// Creates directory "createdDir" using os.Mkdir.
// Creates nested directories "sampleDir/path1/path2" using os.MkdirAll.

package main

import (
	"os"
)

func main() {

	f, err := os.Create("created.file")
	if err != nil {
		panic(err)
	}
	f.Close()

	f, err = os.OpenFile("created.byopen", os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}
	f.Close()

	err = os.Mkdir("createdDir", 0777)
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll("sampleDir/path1/path2", 0777)
	if err != nil {
		panic(err)
	}

}

// Creates six files named "test.file1" to "test.file6" using os.Create.
// Glob retrieves files matching pattern "test.file[1-3]" using filepath.Glob.
// Removes all created files after execution as cleanup using os.Remove.

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	for i := 1; i <= 6; i++ {
		_, err := os.Create(fmt.Sprintf("./test.file%d", i))
		if err != nil {
			fmt.Println(err)
		}
	}

	m, err := filepath.Glob("test.file[1-3]")
	if err != nil {
		panic(err)
	}

	for _, val := range m {
		fmt.Println(val)
	}

	// Cleanup
	for i := 1; i <= 6; i++ {
		err := os.Remove(fmt.Sprintf("./test.file%d", i))
		if err != nil {
			fmt.Println(err)
		}
	}
}

// Creates three files with content and permissions defined in `data`.
// Compares files by checksum using MD5 and line by line using bufio.Scanner.
// Cleans up created files after comparisons using os.Remove.

package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

var data = []struct {
	name string
	cont string
	perm os.FileMode
}{
	{"test1.file", "Hello\nGolang is great", 0666},
	{"test2.file", "Hello\nGolang is great", 0666},
	{"test3.file", "Not matching\nGolang is great\nLast line", 0666},
}

func main() {

	files := []*os.File{}
	for _, fData := range data {
		f, err := os.Create(fData.name)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		_, err = io.WriteString(f, fData.cont)
		if err != nil {
			panic(err)
		}
		files = append(files, f)
	}

	// Compare by checksum
	checksums := []string{}
	for _, f := range files {
		f.Seek(0, 0) // reset to beginngin of file
		sum, err := getMD5SumString(f)
		if err != nil {
			panic(err)
		}
		checksums = append(checksums, sum)
	}

	fmt.Println("### Comparing by checksum ###")
	compareCheckSum(checksums[0], checksums[1])
	compareCheckSum(checksums[0], checksums[2])

	fmt.Println("### Comparing line by line ###")
	files[0].Seek(0, 0)
	files[2].Seek(0, 0)
	compareFileByLine(files[0], files[2])

	// Cleanup
	for _, val := range data {
		os.Remove(val.name)
	}

}

func getMD5SumString(f *os.File) (string, error) {
	file1Sum := md5.New()
	_, err := io.Copy(file1Sum, f)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%X", file1Sum.Sum(nil)), nil
}

func compareCheckSum(sum1, sum2 string) {
	match := "match"
	if sum1 != sum2 {
		match = " does not match"
	}
	fmt.Printf("Sum: %s and Sum: %s %s\n", sum1, sum2, match)
}

func compareLines(line1, line2 string) {
	sign := "o"
	if line1 != line2 {
		sign = "x"
	}
	fmt.Printf("%s | %s | %s \n", sign, line1, line2)
}

func compareFileByLine(f1, f2 *os.File) {
	sc1 := bufio.NewScanner(f1)
	sc2 := bufio.NewScanner(f2)
	for {
		sc1Bool := sc1.Scan()
		sc2Bool := sc2.Scan()
		if !sc1Bool && !sc2Bool {
			break
		}
		compareLines(sc1.Text(), sc2.Text())
	}
}

// Retrieves the current user's information using user.Current() and prints the home directory.

package main

import (
	"fmt"
	"log"
	"os/user"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The user home directory: " + usr.HomeDir)
}
























/*
Package main lists all network interfaces on the system and their associated IP addresses.

This program uses the `net` package to retrieve and display all network interfaces available
on the system along with their IP addresses. It handles any errors encountered during the
retrieval of interfaces or addresses by using `panic` to terminate the program with an error message.
*/
package main

import (
	"fmt"
	"net"
)

func main() {
	// Get all network interfaces on the system.
	// If there is an error in retrieving the interfaces, the program will terminate with a panic.
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	// Iterate through each network interface.
	for _, interf := range interfaces {
		// Resolve and display addresses for each interface.
		// If there is an error in retrieving the addresses, the program will terminate with a panic.
		addrs, err := interf.Addrs()
		if err != nil {
			panic(err)
		}

		// Print the name of the network interface.
		fmt.Println(interf.Name)

		// Iterate through each address associated with the interface.
		for _, add := range addrs {
			// Check if the address is of type *net.IPNet and print it.
			if ip, ok := add.(*net.IPNet); ok {
				fmt.Printf("\t%v\n", ip)
			}
		}
	}
}










/*
Package main demonstrates a simple HTTP server and client interaction in Go.

This package defines an HTTP server that responds with a static message and a client that connects to the server via plain TCP,
sends an HTTP GET request, and reads the response. The server is created using the `http` package, and the client uses the `net` package for TCP connection.
The server is gracefully shut down after handling the client request.
*/
package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

// StringServer represents a simple HTTP handler that serves a static string message.
type StringServer string

// ServeHTTP responds to HTTP requests with the static string message defined in StringServer.
func (s StringServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte(string(s)))
}

// createServer initializes and returns an HTTP server listening on the specified address.
// The server uses StringServer as its handler to serve a static message.
func createServer(addr string) http.Server {
	return http.Server{
		Addr:    addr,
		Handler: StringServer("HELLO GOPHER!\n"),
	}
}

const addr = "localhost:7070"

func main() {
	// Create and start the HTTP server in a separate goroutine.
	s := createServer(addr)
	go s.ListenAndServe()

	// Establish a plain TCP connection to the server.
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Send an HTTP GET request to the server.
	_, err = io.WriteString(conn, "GET / HTTP/1.1\r\nHost: localhost:7070\r\n\r\n")
	if err != nil {
		panic(err)
	}

	// Read and print the server's response.
	scanner := bufio.NewScanner(conn)
	conn.SetReadDeadline(time.Now().Add(time.Second)) // Set a read deadline for the TCP connection.
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// Gracefully shut down the HTTP server with a 5-second timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.Shutdown(ctx)
}










/*
Package main demonstrates DNS lookup functionalities for resolving IP addresses and hostnames.

This program resolves the hostnames associated with the loopback IP address (`127.0.0.1`) and the IP addresses for the hostname `localhost`.
It uses functions from the `net` package to perform these lookups and prints the results to the console.
*/
package main

import (
	"fmt"
	"net"
)

func main() {
	// Resolve the hostname associated with the loopback IP address (127.0.0.1).
	// If there is an error during the lookup, the program will terminate with a panic.
	addrs, err := net.LookupAddr("127.0.0.1")
	if err != nil {
		panic(err)
	}

	// Print each resolved hostname.
	for _, addr := range addrs {
		fmt.Println(addr)
	}

	// Resolve the IP addresses associated with the hostname "localhost".
	// If there is an error during the lookup, the program will terminate with a panic.
	ips, err := net.LookupIP("localhost")
	if err != nil {
		panic(err)
	}

	// Print each resolved IP address.
	for _, ip := range ips {
		fmt.Println(ip.String())
	}
}










/*
Package main demonstrates how to create a simple HTTP server in Go and interact with it using POST requests.

This package defines an HTTP server that responds with a static message and logs received form data. It also provides
two examples of sending POST requests to the server using the `http` package: one with `http.Post` and one with `http.NewRequest`.
The responses from the server are printed to the console.
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// StringServer represents a simple HTTP handler that serves a static string message.
type StringServer string

// ServeHTTP responds to HTTP requests with the static string message defined in StringServer.
// It also parses and prints any form data received in the request.
func (s StringServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	fmt.Printf("Received form data: %v\n", req.Form)
	rw.Write([]byte(string(s)))
}

// createServer initializes and returns an HTTP server listening on the specified address.
// The server uses StringServer as its handler to serve a static message.
func createServer(addr string) http.Server {
	return http.Server{
		Addr:    addr,
		Handler: StringServer("Hello world"),
	}
}

const addr = "localhost:7070"

func main() {
	// Create and start the HTTP server in a separate goroutine.
	s := createServer(addr)
	go s.ListenAndServe()

	// Send a POST request to the server using http.Post.
	simplePost()

	// Send a POST request to the server using a custom http.Request.
	useRequest()
}

// simplePost sends a POST request with form data to the HTTP server using http.Post.
// It prints the response from the server to the console.
func simplePost() {
	res, err := http.Post("http://localhost:7070",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=Radek&surname=Sohlich"))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Response from server:" + string(data))
}

// useRequest sends a POST request with form data to the HTTP server using a custom http.Request.
// It prints the response from the server to the console.
func useRequest() {
	hc := http.Client{}
	form := url.Values{}
	form.Add("name", "Radek")
	form.Add("surname", "Sohlich")

	req, err := http.NewRequest("POST",
		"http://localhost:7070",
		strings.NewReader(form.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := hc.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Response from server:" + string(data))
}










/*
Package main demonstrates the creation, parsing, and serialization of URLs in Go.

This program constructs a URL using the `url.URL` struct, prints it, and then parses the constructed URL back
into a URL object. It also serializes the parsed URL object into a JSON format for display.
*/
package main

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func main() {
	// Create and assemble a URL using the url.URL struct.
	u := &url.URL{}
	u.Scheme = "http"                     // Set the URL scheme (e.g., "http").
	u.Host = "localhost"                  // Set the host (e.g., "localhost").
	u.Path = "index.html"                 // Set the path (e.g., "index.html").
	u.RawQuery = "id=1&name=John"         // Set the raw query parameters (e.g., "id=1&name=John").
	u.User = url.UserPassword("admin", "1234") // Set the user credentials (e.g., username and password).

	// Print the assembled URL.
	fmt.Printf("Assembled URL:\n%v\n\n\n", u)

	// Parse the assembled URL string back into a URL object.
	parsedURL, err := url.Parse(u.String())
	if err != nil {
		panic(err)
	}

	// Serialize the parsed URL object to JSON.
	jsonURL, err := json.Marshal(parsedURL)
	if err != nil {
		panic(err)
	}

	// Print the serialized JSON representation of the parsed URL.
	fmt.Println("Parsed URL:")
	fmt.Println(string(jsonURL))
}










/*
Package main demonstrates creating an HTTP server in Go and sending a POST request with form data and headers.

This program sets up a simple HTTP server that handles incoming requests, prints received form data and headers,
and responds with a static message. It also creates a POST request to this server with specific form data and headers,
and then prints the server's response.
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// StringServer represents a simple HTTP handler that serves a static string message.
type StringServer string

// ServeHTTP responds to HTTP requests with the static string message defined in StringServer.
// It also parses and prints any form data and headers received in the request.
func (s StringServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	fmt.Printf("Received form data: %v\n", req.Form)
	fmt.Printf("Received header: %v\n", req.Header)
	rw.Write([]byte(string(s)))
}

// createServer initializes and returns an HTTP server listening on the specified address.
// The server uses StringServer as its handler to serve a static message.
func createServer(addr string) http.Server {
	return http.Server{
		Addr:    addr,
		Handler: StringServer("Hello world"),
	}
}

const addr = "localhost:7070"

func main() {
	// Create and start the HTTP server in a separate goroutine.
	s := createServer(addr)
	go s.ListenAndServe()

	// Prepare form data for the POST request.
	form := url.Values{}
	form.Set("id", "5")
	form.Set("name", "Wolfgang")

	// Create a new POST request with the form data and the appropriate content type.
	req, err := http.NewRequest(http.MethodPost,
		"http://localhost:7070",
		strings.NewReader(form.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send the POST request and handle the response.
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// Read and print the response from the server.
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Response from server:" + string(data))
}










/*
Package main demonstrates basic operations on HTTP headers using the `http.Header` type in Go.

This program performs various operations on HTTP headers, including setting, adding, retrieving, replacing,
and deleting header values. It prints the header state after each operation to illustrate the changes.
*/
package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Create a new HTTP header.
	header := http.Header{}

	// Set the header "Auth-X" to a single value "abcdef1234".
	header.Set("Auth-X", "abcdef1234")

	// Add another value "defghijkl" to the header "Auth-X".
	header.Add("Auth-X", "defghijkl")

	// Print the current state of the header.
	fmt.Println(header)

	// Retrieve all values associated with the header "Auth-X".
	resSlice := header["Auth-X"]
	fmt.Println(resSlice)

	// Get the first value associated with the header "Auth-X".
	resFirst := header.Get("Auth-X")
	fmt.Println(resFirst)

	// Replace all existing values of the header "Auth-X" with a new value "newvalue".
	header.Set("Auth-X", "newvalue")
	fmt.Println(header)

	// Remove the header "Auth-X".
	header.Del("Auth-X")
	fmt.Println(header)
}










/*
Package main demonstrates how to handle HTTP redirects and track redirection counts in Go.

This program sets up an HTTP server that responds with temporary redirects (/redirect1, /redirect2, etc.) based on the number of redirections.
It also creates an HTTP client that handles redirects and limits the maximum number of redirects to 2. It prints details of each redirect
and stops if the maximum redirect count is exceeded.
*/
package main

import (
	"fmt"
	"net/http"
)

// addr represents the address and port on which the HTTP server listens.
const addr = "localhost:7070"

// RedirecServer is a struct that implements the http.Handler interface.
// It handles incoming HTTP requests and performs redirects based on a redirect count.
type RedirecServer struct {
	redirectCount int
}

// ServeHTTP handles incoming HTTP requests and performs a temporary redirect to the next redirect path.
// It increments the redirect count and sets a custom header "Known-redirects" with the current count.
func (s *RedirecServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	s.redirectCount++
	fmt.Println("Received header: " + req.Header.Get("Known-redirects"))
	// Perform a temporary redirect to the next redirection path (/redirect1, /redirect2, etc.).
	http.Redirect(rw, req, fmt.Sprintf("/redirect%d", s.redirectCount), http.StatusTemporaryRedirect)
}

func main() {
	// Create an HTTP server with a RedirecServer handler.
	s := http.Server{
		Addr:    addr,
		Handler: &RedirecServer{0}, // Start with redirectCount set to 0.
	}
	go s.ListenAndServe() // Start the HTTP server in a separate goroutine.

	// Create an HTTP client.
	client := http.Client{}
	redirectCount := 0

	// Configure the client to handle redirects and limit the maximum number of redirects to 2.
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		fmt.Println("Redirected")
		// Limit the maximum number of redirects to 2.
		if redirectCount > 2 {
			return fmt.Errorf("Too many redirects")
		}
		// Set a custom header "Known-redirects" with the current redirect count.
		req.Header.Set("Known-redirects", fmt.Sprintf("%d", redirectCount))
		redirectCount++
		// Print details of each previous request in the redirection chain.
		for _, prReq := range via {
			fmt.Printf("Previous request: %v\n", prReq.URL)
		}
		return nil
	}

	// Perform a GET request to the HTTP server.
	_, err := client.Get("http://" + addr)
	if err != nil {
		panic(err)
	}
}










/*
Package main demonstrates basic CRUD operations (Create, Read) using an HTTP server and client in Go.

This program sets up an HTTP server that manages a list of cities. It allows clients to retrieve the list of cities
via a GET request, and add a new city via a POST request. It also provides a client interface to fetch the list of cities
and save a new city to the server.
*/
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// addr represents the address and port on which the HTTP server listens.
const addr = "localhost:7070"

// City represents a city with its ID, name, and location.
type City struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

// toJson converts a City struct to its JSON string representation.
func (c City) toJson() string {
	return fmt.Sprintf(`{"name":"%s","location":"%s"}`, c.Name, c.Location)
}

func main() {
	// Create and start the HTTP server in a separate goroutine.
	s := createServer(addr)
	go s.ListenAndServe()

	// Retrieve the list of cities from the server.
	cities, err := getCities()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Retrieved cities: %v\n", cities)

	// Save a new city "Paris" to the server.
	city, err := saveCity(City{"", "Paris", "France"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Saved city: %v\n", city)
}

// saveCity sends a POST request to save a new city to the server.
// It returns the saved City object and any error encountered.
func saveCity(city City) (City, error) {
	r, err := http.Post("http://"+addr+"/cities", "application/json", strings.NewReader(city.toJson()))
	if err != nil {
		return City{}, err
	}
	defer r.Body.Close()
	return decodeCity(r.Body)
}

// getCities sends a GET request to retrieve the list of cities from the server.
// It returns a slice of City objects and any error encountered.
func getCities() ([]City, error) {
	r, err := http.Get("http://" + addr + "/cities")
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	return decodeCities(r.Body)
}

// decodeCity decodes a JSON-encoded City object from the provided reader.
// It returns the decoded City object and any error encountered during decoding.
func decodeCity(r io.Reader) (City, error) {
	city := City{}
	dec := json.NewDecoder(r)
	err := dec.Decode(&city)
	return city, err
}

// decodeCities decodes a JSON-encoded slice of City objects from the provided reader.
// It returns the decoded slice of City objects and any error encountered during decoding.
func decodeCities(r io.Reader) ([]City, error) {
	cities := []City{}
	dec := json.NewDecoder(r)
	err := dec.Decode(&cities)
	return cities, err
}

// createServer creates and returns an HTTP server configured to handle city-related requests.
// It initializes with a predefined set of cities and uses a multiplexer (mux) to route requests.
func createServer(addr string) http.Server {
	// Predefined list of cities.
	cities := []City{
		{ID: "1", Name: "Prague", Location: "Czechia"},
		{ID: "2", Name: "Bratislava", Location: "Slovakia"},
	}

	// Create a new HTTP multiplexer (mux).
	mux := http.NewServeMux()

	// Define handler for "/cities" endpoint.
	mux.HandleFunc("/cities", func(w http.ResponseWriter, r *http.Request) {
		enc := json.NewEncoder(w)
		if r.Method == http.MethodGet {
			// Handle GET request: Return the list of cities.
			enc.Encode(cities)
		} else if r.Method == http.MethodPost {
			// Handle POST request: Add a new city to the list.
			data, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			r.Body.Close()

			// Decode the incoming JSON data into a City object.
			city := City{}
			if err := json.Unmarshal(data, &city); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Assign a new ID (incremental) to the new city.
			city.ID = strconv.Itoa(len(cities) + 1)

			// Append the new city to the list of cities.
			cities = append(cities, city)

			// Return the newly added city as a JSON response.
			enc.Encode(city)
		}
	})

	// Create and return an HTTP server configured with the multiplexer (mux).
	return http.Server{
		Addr:    addr,
		Handler: mux,
	}
}










/*
Package main demonstrates how to send an email using SMTP with authentication and TLS encryption in Go.

This program prompts the user for SMTP credentials (username and password), establishes a connection to Gmail's SMTP server,
and sends a simple email message using the provided credentials.
*/
package main

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
)

func main() {
	// Prompt the user to enter SMTP username (email address).
	var email string
	fmt.Println("Enter username for smtp: ")
	fmt.Scanln(&email)

	// Prompt the user to enter SMTP password.
	var pass string
	fmt.Println("Enter password for smtp: ")
	fmt.Scanln(&pass)

	// Authenticate using PlainAuth with the provided credentials.
	auth := smtp.PlainAuth("",
		email,
		pass,
		"smtp.gmail.com")

	// Connect to Gmail's SMTP server on port 587 (TLS encryption).
	c, err := smtp.Dial("smtp.gmail.com:587")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	// Enable TLS encryption on the SMTP connection.
	config := &tls.Config{ServerName: "smtp.gmail.com"}
	if err = c.StartTLS(config); err != nil {
		panic(err)
	}

	// Authenticate with the server using the provided credentials.
	if err = c.Auth(auth); err != nil {
		panic(err)
	}

	// Set the sender's email address.
	if err = c.Mail(email); err != nil {
		panic(err)
	}

	// Set the recipient's email address (same as sender in this example).
	if err = c.Rcpt(email); err != nil {
		panic(err)
	}

	// Open a data connection to send the email content.
	w, err := c.Data()
	if err != nil {
		panic(err)
	}

	// Define the email message content (in this case, a simple text message).
	msg := []byte("Hello, this is the email content")

	// Write the message content to the data connection.
	if _, err := w.Write(msg); err != nil {
		panic(err)
	}

	// Close the data connection.
	err = w.Close()
	if err != nil {
		panic(err)
	}

	// Quit the SMTP session.
	err = c.Quit()
	if err != nil {
		panic(err)
	}
}










/*
Package main demonstrates how to set up a simple RPC server and client using net/rpc and net/rpc/jsonrpc packages in Go.

This program defines a simple RPC server that exposes an Add method, which adds two integers received in Args struct
and returns the result in Result struct. The client connects to this server using JSON-RPC over TCP, sends an Add RPC
request, and prints the result received from the server.
*/
package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// Args represents the arguments for the Add method.
type Args struct {
	A, B int
}

// Result represents the result of the Add method.
type Result int

// RpcServer represents the RPC server type.
type RpcServer struct{}

// Add is an RPC method on RpcServer that adds two integers.
func (t RpcServer) Add(args *Args, result *Result) error {
	log.Printf("Adding %d to %d\n", args.A, args.B)
	*result = Result(args.A + args.B)
	return nil
}

// addr represents the address and port on which the RPC server listens.
const addr = ":7070"

func main() {
	// Start the RPC server in a separate goroutine.
	go createServer(addr)

	// Connect to the RPC server using JSON-RPC over TCP.
	client, err := jsonrpc.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// Prepare arguments for the RPC call.
	args := &Args{
		A: 2,
		B: 3,
	}
	var result Result

	// Call the Add method on the RPC server.
	err = client.Call("RpcServer.Add", args, &result)
	if err != nil {
		log.Fatalf("Error calling RpcServer.Add: %s", err)
	}

	// Print the result received from the RPC server.
	log.Printf("%d + %d = %d\n", args.A, args.B, result)
}

// createServer creates and starts an RPC server that listens for incoming connections on the specified address.
func createServer(addr string) {
	// Create a new RPC server instance.
	server := rpc.NewServer()

	// Register the RpcServer type to handle RPC requests.
	err := server.Register(RpcServer{})
	if err != nil {
		panic(err)
	}

	// Listen for incoming TCP connections on the specified address.
	l, e := net.Listen("tcp", addr)
	if e != nil {
		log.Fatalf("Couldn't start listening on %s: %s", addr, e)
	}

	// Accept incoming connections and serve RPC requests using JSON-RPC codec.
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}



















/*
Package main demonstrates how to connect to a PostgreSQL database using the "database/sql" package and the pq driver in Go.

This program connects to a local PostgreSQL database named "example" running on port 5432 with the username and password "postgres".
It disables SSL mode for simplicity in local development. It verifies the connection by pinging the database and prints "Ping OK"
if the connection is successful.
*/
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Import the PostgreSQL driver package anonymously
)

func main() {
	// Connection string for PostgreSQL database
	connStr := "postgres://postgres:postgres@localhost:5432/example?sslmode=disable"

	// Open a connection to the PostgreSQL database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close() // Ensure the database connection is closed when main function exits

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// Print message indicating successful ping
	fmt.Println("Ping OK")
}










/*
Package main demonstrates how to use PostgreSQL database connection and context handling with "database/sql" package and the pq driver in Go.

This program connects to a local PostgreSQL database named "example" running on port 5432 with the username and password "postgres".
It disables SSL mode for simplicity in local development. It verifies the connection by performing pings with and without a context.
It also demonstrates creating a connection using db.Conn() and verifying its ping using a context.

Note: The use of time.Nanosecond for context timeout is not practical and is only used here to demonstrate context handling.

Dependencies:
- github.com/lib/pq: PostgreSQL driver for Go

Usage:
- Update connStr with your PostgreSQL database connection details.
- Ensure PostgreSQL server is running on localhost:5432.

init.sql:
        DROP TABLE IF EXISTS post;
        CREATE TABLE post (
          ID serial,
          TITLE varchar(40),
          CONTENT varchar(255),
          CONSTRAINT pk_post PRIMARY KEY(ID)
        );
        SELECT * FROM post;
*/
package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq" // Import PostgreSQL driver package anonymously
)

func main() {
	// Connection string for PostgreSQL database
	connStr := "postgres://postgres:postgres@localhost:5432/example?sslmode=disable"

	// Open a connection to the PostgreSQL database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close() // Ensure the database connection is closed when main function exits

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Ping OK.")

	// Create a context with a timeout of 1 nanosecond (not practical, for demonstration purposes only)
	ctx, _ := context.WithTimeout(context.Background(), time.Nanosecond)

	// Ping the database with the context
	err = db.PingContext(ctx)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}

	// Create a connection using db.Conn()
	conn, err := db.Conn(context.Background())
	if err != nil {
		panic(err)
	}
	defer conn.Close() // Ensure the connection is closed when function exits

	// Ping the connection with a context
	err = conn.PingContext(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection Ping OK.")
}










/*
Package main demonstrates basic CRUD operations (Create, Read, Update, Delete) with PostgreSQL using Go's database/sql package and the pq driver.

This program connects to a local PostgreSQL database named "example" running on port 5432 with the username and password "postgres".
It disables SSL mode for simplicity in local development. It performs the following operations:
- Truncates the "post" table.
- Inserts predefined rows into the "post" table.
- Selects and counts the number of rows in the "post" table.

Dependencies:
- github.com/lib/pq: PostgreSQL driver for Go

Usage:
- Update connStr with your PostgreSQL database connection details.
- Ensure PostgreSQL server is running on localhost:5432 and the "post" table exists in the "example" database.

*/
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Import PostgreSQL driver package anonymously
)

// SQL statements for database operations
const (
	sel   = "SELECT * FROM post;"                                     // Select all rows from "post" table
	trunc = "TRUNCATE TABLE post;"                                     // Truncate (empty) "post" table
	ins   = "INSERT INTO post(ID,TITLE,CONTENT) VALUES (1,'Title 1','Content 1'), (2,'Title 2','Content 2');" // Insert rows into "post" table
)

func main() {
	// Create a database connection
	db := createConnection()
	defer db.Close() // Ensure the database connection is closed when main function exits

	// Truncate the "post" table
	_, err := db.Exec(trunc)
	if err != nil {
		panic(err)
	}
	fmt.Println("Table truncated.")

	// Insert rows into the "post" table
	r, err := db.Exec(ins)
	if err != nil {
		panic(err)
	}
	affected, err := r.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Inserted rows count: %d\n", affected)

	// Query and count rows from the "post" table
	rs, err := db.Query(sel)
	if err != nil {
		panic(err)
	}
	count := 0
	for rs.Next() {
		if rs.Err() != nil {
			fmt.Println(rs.Err())
			continue
		}
		count++
	}
	fmt.Printf("Total of %d rows selected.\n", count)
}

// createConnection establishes a connection to the PostgreSQL database and returns the *sql.DB object.
func createConnection() *sql.DB {
	// Connection string for PostgreSQL database
	connStr := "postgres://postgres:postgres@localhost:5432/example?sslmode=disable"

	// Open a connection to the PostgreSQL database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}










/*
Package main demonstrates how to perform batch insert operations into a PostgreSQL database table using prepared statements with Go's database/sql package and the pq driver.

This program connects to a local PostgreSQL database named "example" running on port 5432 with the username and password "postgres".
It disables SSL mode for simplicity in local development. It performs the following operations:
- Truncates the "post" table to clear existing data.
- Uses a prepared statement to insert multiple rows into the "post" table from a predefined slice of structs.
- Prints the number of rows successfully inserted.

Dependencies:
- github.com/lib/pq: PostgreSQL driver for Go

Usage:
- Update connStr with your PostgreSQL database connection details.
- Ensure PostgreSQL server is running on localhost:5432 and the "post" table exists in the "example" database.


init.sql:
DROP TABLE IF EXISTS post;
        CREATE TABLE post (
          ID serial,
          TITLE varchar(40),
          CONTENT varchar(255),
          CONSTRAINT pk_post PRIMARY KEY(ID)
        );
        SELECT * FROM post;
*/
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Import PostgreSQL driver package anonymously
)

// SQL statements for database operations
const (
	trunc = "TRUNCATE TABLE post;"                        // Truncate (empty) "post" table
	ins   = "INSERT INTO post(ID,TITLE,CONTENT) VALUES ($1,$2,$3)" // Insert statement with placeholders
)

// Struct for test data
var testTable = []struct {
	ID      int
	Title   string
	Content string
}{
	{1, "Title One", "Content of title one"},
	{2, "Title Two", "Content of title two"},
	{3, "Title Three", "Content of title three"},
}

func main() {
	// Create a database connection
	db := createConnection()
	defer db.Close() // Ensure the database connection is closed when main function exits

	// Truncate the "post" table
	_, err := db.Exec(trunc)
	if err != nil {
		panic(err)
	}
	fmt.Println("Table truncated.")

	// Prepare the insert statement
	stm, err := db.Prepare(ins)
	defer stm.Close() // Ensure the prepared statement is closed
	if err != nil {
		panic(err)
	}

	inserted := int64(0)
	// Iterate over testTable and insert rows using the prepared statement
	for _, val := range testTable {
		fmt.Printf("Inserting record ID: %d\n", val.ID)
		// Execute the prepared statement with values from the struct
		r, err := stm.Exec(val.ID, val.Title, val.Content)
		if err != nil {
			fmt.Printf("Cannot insert record ID : %d\n", val.ID)
		}
		// Retrieve the number of affected rows and accumulate the total
		if affected, err := r.RowsAffected(); err == nil {
			inserted += affected
		}
	}

	// Print the total number of rows successfully inserted
	fmt.Printf("Result: Inserted %d rows.\n", inserted)
}

// createConnection establishes a connection to the PostgreSQL database and returns the *sql.DB object.
func createConnection() *sql.DB {
	// Connection string for PostgreSQL database
	connStr := "postgres://postgres:postgres@localhost:5432/example?sslmode=disable"

	// Open a connection to the PostgreSQL database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}










/*
Package main demonstrates executing a large SQL query with context timeout using Go's database/sql package and the pq driver for PostgreSQL.

This program connects to a local PostgreSQL database named "example" running on port 5432 with the username and password "postgres".
It disables SSL mode for simplicity in local development. It performs the following operations:
- Creates a database connection.
- Uses a context with a timeout of 20 microseconds to limit the query execution time.
- Executes a SELECT query (`sel`) that performs a cross join with a large series to generate a significant number of rows.
- Cancels the query if it exceeds the context timeout.
- Prints the number of rows returned by the query.

Dependencies:
- github.com/lib/pq: PostgreSQL driver for Go

Usage:
- Update connStr with your PostgreSQL database connection details.
- Ensure PostgreSQL server is running on localhost:5432 and the "post" table exists in the "example" database.
- Adjust the context timeout (`20*time.Microsecond`) as needed for your query's expected execution time.


init.sql:
DROP TABLE IF EXISTS post;
        CREATE TABLE post (
          ID serial,
          TITLE varchar(40),
          CONTENT varchar(255),
          CONSTRAINT pk_post PRIMARY KEY(ID)
        );
        SELECT * FROM post;
        INSERT INTO post(ID,TITLE,CONTENT) VALUES
                        (1,'Title One','Content One'),
                        (2,'Title Two','Content Two');
*/
package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq" // Import PostgreSQL driver package anonymously
)

// SQL SELECT statement to generate a large number of rows
const sel = "SELECT * FROM post p CROSS JOIN (SELECT 1 FROM generate_series(1,1000000)) tbl"

func main() {
	// Create a database connection
	db := createConnection()
	defer db.Close() // Ensure the database connection is closed when main function exits

	// Create a context with a timeout of 20 microseconds
	ctx, canc := context.WithTimeout(context.Background(), 20*time.Microsecond)
	defer canc() // Ensure the cancellation function is called to cancel the query if it exceeds the timeout

	// Execute the query with context timeout
	rows, err := db.QueryContext(ctx, sel)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	// Count the number of rows returned by the query
	count := 0
	for rows.Next() {
		if rows.Err() != nil {
			fmt.Println(rows.Err())
			continue
		}
		count++
	}

	// Print the number of rows returned
	fmt.Printf("%d rows returned\n", count)
}

// createConnection establishes a connection to the PostgreSQL database and returns the *sql.DB object.
func createConnection() *sql.DB {
	// Connection string for PostgreSQL database
	connStr := "postgres://postgres:postgres@localhost:5432/example?sslmode=disable"

	// Open a connection to the PostgreSQL database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}










/*
Package main demonstrates querying PostgreSQL database for column information using Go's database/sql package and the pq driver.

This program connects to a local PostgreSQL database named "example" running on port 5432 with the username and password "postgres".
It disables SSL mode for simplicity in local development. It performs the following operations:
- Creates a database connection.
- Executes a SELECT query (`sel`) to fetch all columns from the "post" table.
- Retrieves and prints information about the selected columns, including their names, types, and other properties.

Dependencies:
- github.com/lib/pq: PostgreSQL driver for Go

Usage:
- Update connStr with your PostgreSQL database connection details.
- Ensure PostgreSQL server is running on localhost:5432 and the "post" table exists in the "example" database.


init.sql:
DROP TABLE IF EXISTS post;
        CREATE TABLE post (
          ID serial,
          TITLE varchar(40),
          CONTENT varchar(255),
          CONSTRAINT pk_post PRIMARY KEY(ID)
        );
        SELECT * FROM post;
        INSERT INTO post(ID,TITLE,CONTENT) VALUES
                        (1,'Title One','Content One'),
                        (2,'Title Two','Content Two');
*/
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Import PostgreSQL driver package anonymously
)

// SQL SELECT statement to fetch all columns from the "post" table
const sel = "SELECT * FROM post p"

func main() {
	// Create a database connection
	db := createConnection()
	defer db.Close() // Ensure the database connection is closed when main function exits

	// Execute the SELECT query to fetch all columns
	rs, err := db.Query(sel)
	if err != nil {
		panic(err)
	}
	defer rs.Close()

	// Retrieve and print the selected column names
	columns, err := rs.Columns()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected columns: %v\n", columns)

	// Retrieve and print information about each column type
	colTypes, err := rs.ColumnTypes()
	if err != nil {
		panic(err)
	}
	for _, col := range colTypes {
		fmt.Println()
		fmt.Printf("%+v\n", col)
	}
}

// createConnection establishes a connection to the PostgreSQL database and returns the *sql.DB object.
func createConnection() *sql.DB {
	// Connection string for PostgreSQL database
	connStr := "postgres://postgres:postgres@localhost:5432/example?sslmode=disable"

	// Open a connection to the PostgreSQL database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}










/*
Package main demonstrates querying PostgreSQL database for multiple result sets and single row using Go's database/sql package and the pq driver.

This program connects to a local PostgreSQL database named "example" running on port 5432 with the username and password "postgres".
It disables SSL mode for simplicity in local development. It performs the following operations:
- Creates a database connection.
- Executes multiple SELECT queries (`sel` and `selOne`) to fetch multiple result sets and a single row.
- Retrieves and prints the selected posts and number.
- Handles errors gracefully using panic and printing error messages.

Dependencies:
- github.com/lib/pq: PostgreSQL driver for Go

Usage:
- Update connStr with your PostgreSQL database connection details.
- Ensure PostgreSQL server is running on localhost:5432 and the "post" table exists in the "example" database.


init.sql:
DROP TABLE IF EXISTS post;
       CREATE TABLE post (
         ID serial,
         TITLE varchar(40),
         CONTENT varchar(255),
         CONSTRAINT pk_post PRIMARY KEY(ID)
       );
       SELECT * FROM post;
       INSERT INTO post(ID,TITLE,CONTENT) VALUES
                       (1,'Title One','Content One'),
                       (2,NULL,'Content Two');
*/

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Import PostgreSQL driver package anonymously
)

// SQL SELECT statement to fetch title and content from the "post" table and a constant number
const sel = `SELECT title,content FROM post;
			SELECT 1234 NUM;`

// SQL SELECT statement to fetch title and content from the "post" table based on ID parameter
const selOne = "SELECT title,content FROM post WHERE ID = $1;"

// Post struct represents a row from the "post" table
type Post struct {
	Name sql.NullString
	Text sql.NullString
}

func main() {
	// Create a database connection
	db := createConnection()
	defer db.Close() // Ensure the database connection is closed when main function exits

	// Execute the first SELECT query (`sel`) to fetch multiple result sets
	rs, err := db.Query(sel)
	if err != nil {
		panic(err)
	}
	defer rs.Close()

	// Retrieve and store posts from the first result set
	posts := []Post{}
	for rs.Next() {
		p := Post{}
		if err := rs.Scan(&p.Name, &p.Text); err != nil {
			panic(err)
		}
		posts = append(posts, p)
	}

	// Move to the next result set to retrieve the number
	var num int
	if rs.NextResultSet() {
		for rs.Next() {
			if err := rs.Scan(&num); err != nil {
				panic(err)
			}
		}
	}

	// Print retrieved posts and number
	fmt.Printf("Retrieved posts: %+v\n", posts)
	fmt.Printf("Retrieved number: %d\n", num)

	// Execute the second SELECT query (`selOne`) to fetch a single row based on ID parameter
	row := db.QueryRow(selOne, 100)
	or := Post{}
	if err := row.Scan(&or.Name, &or.Text); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	// Print retrieved single post
	fmt.Printf("Retrieved one post: %+v\n", or)
}

// createConnection establishes a connection to the PostgreSQL database and returns the *sql.DB object.
func createConnection() *sql.DB {
	// Connection string for PostgreSQL database
	connStr := "postgres://postgres:postgres@localhost:5432/example?sslmode=disable"

	// Open a connection to the PostgreSQL database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}










/*
Package main demonstrates querying a PostgreSQL database using Go's database/sql package and the pq driver. It retrieves a specific row from the "post" table based on the ID and demonstrates two methods to parse the result set into a map.

This program connects to a local PostgreSQL database named "example" running on port 5432 with the username and password "postgres".
It disables SSL mode for simplicity in local development. It performs the following operations:
- Queries a single row from the "post" table based on ID = 1.
- Parses the result set using both RawBytes and standard interface{} methods.

Dependencies:
- github.com/lib/pq: PostgreSQL driver for Go

Usage:
- Ensure PostgreSQL server is running on localhost:5432 and the "post" table exists in the "example" database.


init.sql:
DROP TABLE IF EXISTS post;
        CREATE TABLE post (
          ID serial,
          TITLE varchar(40),
          CONTENT varchar(255),
          CONSTRAINT pk_post PRIMARY KEY(ID)
        );
        SELECT * FROM post;
        INSERT INTO post(ID,TITLE,CONTENT) VALUES 
                        (1,NULL,'Content One'),
                        (2,'Title Two','Content Two');
*/
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// selOne is the SQL statement to select a specific row from the "post" table by ID.
const selOne = "SELECT id,title,content FROM post WHERE ID = $1;"

func main() {
	// Create a database connection
	db := createConnection()
	defer db.Close() // Ensure the database connection is closed when main function exits

	// Query the specific row from the "post" table based on ID = 1
	rows, err := db.Query(selOne, 1)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Get column names from the result set
	cols, _ := rows.Columns()

	// Iterate over the result set
	for rows.Next() {
		// Parse the row into a map using RawBytes method
		m := parseWithRawBytes(rows, cols)
		fmt.Println("Parsed with RawBytes:", m)

		// Parse the row into a map using standard interface{} method
		m = parseToMap(rows, cols)
		fmt.Println("Parsed with interface{}:", m)
	}
}

// parseWithRawBytes parses a row from sql.Rows into a map[string]interface{} using RawBytes method.
func parseWithRawBytes(rows *sql.Rows, cols []string) map[string]interface{} {
	vals := make([]sql.RawBytes, len(cols))
	scanArgs := make([]interface{}, len(vals))
	for i := range vals {
		scanArgs[i] = &vals[i]
	}
	if err := rows.Scan(scanArgs...); err != nil {
		panic(err)
	}
	m := make(map[string]interface{})
	for i, col := range vals {
		if col == nil {
			m[cols[i]] = nil
		} else {
			m[cols[i]] = string(col)
		}
	}
	return m
}

// parseToMap parses a row from sql.Rows into a map[string]interface{} using standard interface{} method.
func parseToMap(rows *sql.Rows, cols []string) map[string]interface{} {
	values := make([]interface{}, len(cols))
	pointers := make([]interface{}, len(cols))
	for i := range values {
		pointers[i] = &values[i]
	}

	if err := rows.Scan(pointers...); err != nil {
		panic(err)
	}

	m := make(map[string]interface{})
	for i, colName := range cols {
		if values[i] == nil {
			m[colName] = nil
		} else {
			m[colName] = values[i]
		}
	}
	return m
}

// createConnection establishes a connection to the PostgreSQL database and returns the connection object.
func createConnection() *sql.DB {
	connStr := "postgres://postgres:postgres@localhost:5432/example?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to PostgreSQL!")
	return db
}










/*
Package main demonstrates transaction management in PostgreSQL using Go's database/sql package and the pq driver. It showcases how to perform operations inside transactions, including querying and modifying data, rolling back transactions, and using contexts for transaction control.

This program connects to a local PostgreSQL database named "example" running on port 5432 with the username and password "postgres".
It disables SSL mode for simplicity in local development. It performs the following operations:
- Starts a transaction and inserts a new row into the "post" table.
- Queries the newly inserted row both outside and within the transaction.
- Rolls back the transaction to discard changes.
- Demonstrates a transaction with context, showing how to use contexts to control transaction lifespan and ensure proper cleanup.

Dependencies:
- github.com/lib/pq: PostgreSQL driver for Go

Usage:
- Ensure PostgreSQL server is running on localhost:5432 and the "post" table exists in the "example" database.


init.sql:
DROP TABLE IF EXISTS post;
        CREATE TABLE post (
          ID serial,
          TITLE varchar(40),
          CONTENT varchar(255),
          CONSTRAINT pk_post PRIMARY KEY(ID)
        );
        SELECT * FROM post;
        INSERT INTO post(ID,TITLE,CONTENT) VALUES
                        (1,'Title One','Content One'),
                        (2,NULL,'Content Two');
*/
package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// selOne is the SQL statement to select a specific row from the "post" table by ID.
const selOne = "SELECT id,title,content FROM post WHERE ID = $1;"

// insert is the SQL statement to insert a new row into the "post" table.
const insert = "INSERT INTO post(ID,TITLE,CONTENT) VALUES (4,'Transaction Title','Transaction Content');"

// Post represents the structure of a post entity.
type Post struct {
	ID      int
	Title   string
	Content string
}

func main() {
	// Create a database connection
	db := createConnection()
	defer db.Close()

	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// Insert a new row into the "post" table within the transaction
	_, err = tx.Exec(insert)
	if err != nil {
		panic(err)
	}

	// Query the newly inserted row from a separate session (outside transaction)
	p := Post{}
	if err := db.QueryRow(selOne, 4).Scan(&p.ID, &p.Title, &p.Content); err != nil {
		fmt.Println("Error querying outside transaction:", err)
	}
	fmt.Println("Query outside transaction:", p)

	// Query the newly inserted row from within the transaction
	if err := tx.QueryRow(selOne, 4).Scan(&p.ID, &p.Title, &p.Content); err != nil {
		fmt.Println("Error querying within transaction:", err)
	}
	fmt.Println("Query within transaction:", p)

	// Rollback the transaction to discard changes
	err = tx.Rollback()
	if err != nil {
		panic(err)
	}
	fmt.Println("Transaction rolled back successfully.")

	// Demonstrate transaction with context
	fmt.Println("\nTransaction with context")
	ctx, cancel := context.WithCancel(context.Background())
	tx, err = db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadUncommitted})
	if err != nil {
		panic(err)
	}

	// Insert a new row into the "post" table within the transaction
	_, err = tx.Exec(insert)
	if err != nil {
		panic(err)
	}

	// Cancel the context to simulate premature transaction termination
	cancel()

	// Commit the transaction (which should fail due to canceled context)
	err = tx.Commit()
	if err != nil {
		fmt.Println("Error committing transaction with canceled context:", err)
	}
}

// createConnection establishes a connection to the PostgreSQL database and returns the connection object.
func createConnection() *sql.DB {
	connStr := "postgres://postgres:postgres@localhost:5432/example?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to PostgreSQL!")
	return db
}










/*
Package main demonstrates calling stored procedures/functions in both PostgreSQL and MySQL databases using Go's database/sql package and respective drivers.

This program connects to a local PostgreSQL database named "example" running on port 5432 with the username and password "postgres".
It disables SSL mode for simplicity in local development. It performs the following operations:
- Calls a PostgreSQL function named format_name with three parameters and retrieves the result.

Dependencies:
- github.com/lib/pq: PostgreSQL driver for Go
- github.com/go-sql-driver/mysql: MySQL driver for Go

Usage:
- Ensure PostgreSQL server is running on localhost:5432 and the "example" database contains the format_name function.

init.sql:
CREATE OR REPLACE FUNCTION format_name
        (firstname Text,lastname Text,age INT) RETURNS 
        VARCHAR AS $$
        BEGIN
          RETURN trim(firstname) ||' '||trim(lastname) ||' ('||age||')';
        END;
        $$ LANGUAGE plpgsql;
*/

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	_ "github.com/lib/pq"              // PostgreSQL driver
)

// call is the SQL statement to call the format_name function in PostgreSQL.
const call = "select * from format_name($1,$2,$3)"

// callMySQL is the SQL statement to call the simpleproc stored procedure in MySQL.
const callMySQL = "CALL simpleproc(?)"

// Result represents the structure of the result from calling the stored procedure/function.
type Result struct {
	Name     string
	Category int
}

func main() {
	// Create a database connection
	db := createConnection()
	defer db.Close()

	// Initialize a Result struct to hold the returned values
	r := Result{}

	// Call the PostgreSQL function format_name with parameters and scan the result into Result struct
	if err := db.QueryRow(call, "John", "Doe", 32).Scan(&r.Name); err != nil {
		panic(err)
	}
	fmt.Printf("Result is: %+v\n", r)
}

// createConnection establishes a connection to the PostgreSQL database and returns the connection object.
func createConnection() *sql.DB {
	connStr := "postgres://postgres:postgres@localhost:5432/example?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to PostgreSQL!")
	return db
}




// This code implements a TCP server on port 8080 that accepts incoming connections,
// reads messages from clients, and responds with a confirmation message.
package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main() {

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	for {
		fmt.Println("Waiting for client...")
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}

		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			panic(err)
		}
		_, err = io.WriteString(conn, "Received: "+string(msg))
		if err != nil {
			fmt.Println(err)
		}
		conn.Close()
	}

}




// This code sets up a UDP server on port 7070 that listens for incoming packets,
// reads messages from clients, and echoes back a confirmation message to the sender.
package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	pc, err := net.ListenPacket("udp", ":7070")
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()

	buffer := make([]byte, 2048)
	fmt.Println("Waiting for client...")
	for {

		_, addr, err := pc.ReadFrom(buffer)
		if err == nil {
			rcvMsq := string(buffer)
			fmt.Println("Received: " + rcvMsq)
			if _, err := pc.WriteTo([]byte("Received: "+rcvMsq), addr); err != nil {
				fmt.Println("error on write: " + err.Error())
			}
		} else {
			fmt.Println("error: " + err.Error())
		}

	}

}



// This code establishes a TCP server on port 8080 that accepts incoming connections,
// assigns a unique client ID to each connection, sends a welcome message to clients,
// and echoes back received messages with a prefix indicating reception.
package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main() {

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	ID := 0
	for {
		fmt.Println("Waiting for client...")
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Client ID: %d connected.\n", ID)
		go func(c net.Conn, clientID int) {
			fmt.Fprintf(c, "Welcome client ID: %d \n", clientID)
			for {
				msg, err := bufio.NewReader(c).ReadString('\n')
				if err != nil {
					fmt.Println(err)
					break
				}
				_, err = io.WriteString(c, "Received: "+string(msg))
				if err != nil {
					fmt.Println(err)
					break
				}
			}
			fmt.Println("Closing connection")
			c.Close()
		}(conn, ID)
		ID++
	}

}


// This code starts an HTTP server on port 8080 that responds with "Hello world" to incoming requests.
package main

import (
	"fmt"
	"net/http"
)

type SimpleHTTP struct{}

func (s SimpleHTTP) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Hello world")
}

func main() {
	fmt.Println("Starting HTTP server on port 8080")
	s := &http.Server{Addr: ":8080", Handler: SimpleHTTP{}}
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}


// This code sets up an HTTP server on port 8080 with multiple routes:
// - "/user" responds differently based on GET and POST methods.
// - "/items/clothes" serves "Clothes" using a separate mux under "/items/".
// - "/admin/ports" serves "Ports" using a mux under "/admin/" with prefix stripping.
package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fmt.Fprintln(w, "User GET")
		}
		if r.Method == http.MethodPost {
			fmt.Fprintln(w, "User POST")
		}
	})

	// separate handler
	itemMux := http.NewServeMux()
	itemMux.HandleFunc("/items/clothes", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Clothes")
	})
	mux.Handle("/items/", itemMux)

	// Admin handlers
	adminMux := http.NewServeMux()
	adminMux.HandleFunc("/ports", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Ports")
	})

	mux.Handle("/admin/",
		http.StripPrefix("/admin", adminMux))

	// Default server
	http.ListenAndServe(":8080", mux)

}


// This code sets up an HTTP server on port 8080 with two protected endpoints:
// - "/api/users" returns a JSON array of users.
// - "/api/profile" requires authentication via "X-Auth" header and includes a user profile JSON response.
package main

import (
	"io"
	"log"
	"net/http"
)

type User string

func (u User) toString() string {
	return string(u)
}

type AuthHandler func(u User, w http.ResponseWriter, r *http.Request)

func main() {

	// Secured API
	mux := http.NewServeMux()
	mux.HandleFunc("/api/users", Secure(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w,
			`[{"id":"1","login":"ffghi"},{"id":"2","login":"ffghj"}]`)
	}))
	mux.HandleFunc("/api/profile", WithUser(func(u User, w http.ResponseWriter, r *http.Request) {
		log.Println(u.toString())
		io.WriteString(w, "{\"user\":\""+u.toString()+"\"}")
	}))

	http.ListenAndServe(":8080", mux)

}

func WithUser(h AuthHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Header.Get("X-User")
		if len(user) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		h(User(user), w, r)
	}
}

func Secure(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sec := r.Header.Get("X-Auth")
		if sec != "authenticated" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		h(w, r) // use the handler
	}
}


// This code sets up an HTTP server on port 8080:
// - "/welcome" serves the content of "welcome.txt" file.
// - "/html/" serves static files from the "html" directory, stripping "/html" prefix.
package main

import (
	"net/http"
)

func main() {

	fileSrv := http.FileServer(http.Dir("html"))
	fileSrv = http.StripPrefix("/html", fileSrv)

	http.HandleFunc("/welcome", serveWelcome)
	http.Handle("/html/", fileSrv)
	http.ListenAndServe(":8080", nil)
}

func serveWelcome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "welcome.txt")
}


// This code starts an HTTP server on port 8080 that renders a template file "template.tpl"
// and serves it when accessing the root ("/") endpoint.
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	fmt.Println("Server is starting...")
	tpl, err := template.ParseFiles("template.tpl")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tpl.Execute(w, "John Doe")
		if err != nil {
			panic(err)
		}
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}


// This code starts an HTTP server on port 8080 with three endpoints:
// - "/secured/handle" redirects using http.RedirectHandler to "/login".
// - "/secured/hadlefunc" redirects using http.Redirect to "/login".
// - "/login" responds with a message "Welcome user! Please login!".
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Server is starting...")

	http.Handle("/secured/handle", http.RedirectHandler("/login", http.StatusTemporaryRedirect))
	http.HandleFunc("/secured/hadlefunc", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome user! Please login!\n")
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}


// This code starts an HTTP server on port 8080 with three endpoints:
// - "/set" sets a cookie named "X-Cookie" with value "Go is awesome." and domain "localhost".
// - "/get" retrieves and displays the value of the "X-Cookie" cookie and lists all cookies sent with the request.
// - "/remove" removes the "X-Cookie" cookie by setting its MaxAge to -1.
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const cookieName = "X-Cookie"

func main() {
	log.Println("Server is starting...")

	http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		c := &http.Cookie{
			Name:    cookieName,
			Value:   "Go is awesome.",
			Expires: time.Now().Add(time.Hour),
			Domain:  "localhost",
		}
		http.SetCookie(w, c)
		fmt.Fprintln(w, "Cookie is set!")
	})
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		val, err := r.Cookie(cookieName)
		if err != nil {
			fmt.Fprintln(w, "Cookie err: "+err.Error())
			return
		}
		fmt.Fprintf(w, "Cookie is: %s \n", val.Value)
		fmt.Fprintf(w, "Other cookies:\n")
		for _, v := range r.Cookies() {
			fmt.Fprintf(w, "%s => %s \n", v.Name, v.Value)
		}
	})
	http.HandleFunc("/remove", func(w http.ResponseWriter, r *http.Request) {
		val, err := r.Cookie(cookieName)
		if err != nil {
			fmt.Fprintln(w, "Cookie err: "+err.Error())
			return
		}
		val.MaxAge = -1
		http.SetCookie(w, val)
		fmt.Fprintln(w, "Cookie is removed!")
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}


// This code starts an HTTP server on port 8080 that responds with "Hello world!" after a delay.
// It handles graceful shutdown using OS signals (SIGINT) to stop the server and waits up to 30 seconds for connections to close.
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Second)
		fmt.Fprintln(w, "Hello world!")
	})

	srv := &http.Server{Addr: ":8080", Handler: mux}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("Server error: %s\n", err)
		}
	}()

	log.Println("Server listening on : " + srv.Addr)

	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	<-stopChan // wait for SIGINT
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		30*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("Server gracefully stopped")
}



// This code starts an HTTPS server on port 8080 using TLS certificates "server.crt" and "server.key",
// with a handler that responds with "Hello world".
package main

import (
	"fmt"
	"net/http"
)

type SimpleHTTP struct{}

func (s SimpleHTTP) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Hello world")
}

func main() {
	fmt.Println("Starting HTTPS server on port 8080")
	s := &http.Server{Addr: ":8080", Handler: SimpleHTTP{}}
	if err := s.ListenAndServeTLS("server.crt", "server.key"); err != nil {
		panic(err)
	}
}



// This code defines an HTTP server that responds with "Hello world" and demonstrates handling form data:
// - It logs the request form data before and after calling req.ParseForm().
// - It prints the value of "param1" from the parsed form.
// - It serves "Hello world" as the response to any incoming request.
package main

import (
	"fmt"
	"net/http"
)

type StringServer string

func (s StringServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("Prior ParseForm: %v\n", req.Form)
	req.ParseForm()
	fmt.Printf("Post ParseForm: %v\n", req.Form)
	fmt.Println("Param1 is : " + req.Form.Get("param1"))
	fmt.Printf("PostForm : %v\n", req.PostForm)
	rw.Write([]byte(string(s)))
}

func createServer(addr string) http.Server {
	return http.Server{
		Addr:    addr,
		Handler: StringServer("Hello world"),
	}
}

func main() {
	s := createServer(":8080")
	fmt.Println("Server is starting...")
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}


// This code demonstrates the usage of a thread-safe synchronized list (SyncList) with mutex protection,
// allowing concurrent goroutines to safely append and retrieve values, ensuring data integrity.
package main

import (
	"fmt"
	"sync"
)

var names = []string{"Alan", "Joe", "Jack", "Ben",
	"Ellen", "Lisa", "Carl", "Steve", "Anton", "Yo"}

type SyncList struct {
	m     sync.Mutex
	slice []interface{}
}

func NewSyncList(cap int) *SyncList {
	return &SyncList{
		sync.Mutex{},
		make([]interface{}, cap),
	}
}

func (l *SyncList) Load(i int) interface{} {
	l.m.Lock()
	defer l.m.Unlock()
	return l.slice[i]
}

func (l *SyncList) Append(val interface{}) {
	l.m.Lock()
	defer l.m.Unlock()
	l.slice = append(l.slice, val)
}

func (l *SyncList) Store(i int, val interface{}) {
	l.m.Lock()
	defer l.m.Unlock()
	l.slice[i] = val
}

func main() {

	l := NewSyncList(0)
	wg := &sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(idx int) {
			l.Append(names[idx])
			wg.Done()
		}(i)
	}
	wg.Wait()

	for i := 0; i < 10; i++ {
		fmt.Printf("Val: %v stored at idx: %d\n", l.Load(i), i)
	}

}

// This code demonstrates the usage of sync.Map for concurrent-safe access to a map,
// storing and retrieving values with goroutines, and utilizing Load, LoadOrStore, and Range methods.
package main

import (
	"fmt"
	"sync"
)

var names = []string{"Alan", "Joe", "Jack", "Ben",
	"Ellen", "Lisa", "Carl", "Steve", "Anton", "Yo"}

func main() {

	m := sync.Map{}
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(idx int) {
			m.Store(fmt.Sprintf("%d", idx), names[idx])
			wg.Done()
		}(i)
	}
	wg.Wait()

	v, ok := m.Load("1")
	if ok {
		fmt.Printf("For Load key: 1 got %v\n", v)
	}

	v, ok = m.LoadOrStore("11", "Tim")
	if !ok {
		fmt.Printf("Key 11 missing stored val: %v\n", v)
	}

	m.Range(func(k, v interface{}) bool {
		key, _ := k.(string)
		t, _ := v.(string)
		fmt.Printf("For index %v got %v\n", key, t)
		return true
	})

}



// This code defines a Source type with a Pop method that ensures data loading occurs only once,
// simulating a delayed initialization of data with sync.Mutex and sync.Once synchronization mechanisms,
// and demonstrates concurrent access to the Pop method by multiple goroutines.
package main

import (
	"fmt"
	"sync"
	"time"
)

var names = []interface{}{"Alan", "Joe", "Jack", "Ben",
	"Ellen", "Lisa", "Carl", "Steve", "Anton", "Yo"}

type Source struct {
	m    *sync.Mutex
	o    *sync.Once
	data []interface{}
}

func (s *Source) Pop() (interface{}, error) {
	s.m.Lock()
	defer s.m.Unlock()
	s.o.Do(func() {
		time.Sleep(time.Second * 30) // Simulates data loading delay
		s.data = names
		fmt.Println("Data has been loaded.")
	})
	if len(s.data) > 0 {
		res := s.data[0]
		s.data = s.data[1:]
		return res, nil
	}
	return nil, fmt.Errorf("No data available")
}

func main() {

	s := &Source{&sync.Mutex{}, &sync.Once{}, nil}
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(idx int) {
			if val, err := s.Pop(); err == nil {
				fmt.Printf("Pop %d returned: %s\n", idx, val)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}




// This program demonstrates the usage of sync.Pool to manage a pool of Worker objects,
// allowing efficient reuse of objects across multiple goroutines with minimized memory allocations.
package main

import (
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	id string
}

func (w *Worker) String() string {
	return w.id
}

var globalCounter = 0

var pool = sync.Pool{
	New: func() interface{} {
		res := &Worker{fmt.Sprintf("%d", globalCounter)}
		globalCounter++
		return res
	},
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(idx int) {
			w := pool.Get().(*Worker) // Get a worker from the pool
			fmt.Println("Got worker ID: " + w.String())
			time.Sleep(time.Second) // Simulate work with the worker
			pool.Put(w)             // Put the worker back into the pool
			wg.Done()
		}(i)
	}
	wg.Wait()
}



// This code demonstrates the usage of sync.WaitGroup to synchronize and wait for a group of goroutines to complete,
// each printing an exit message with its index before signaling completion.
package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(idx int) {
			// Do some work
			defer wg.Done()
			fmt.Printf("Exiting %d\n", idx)
		}(i)
	}
	wg.Wait()
	fmt.Println("All done.")
}



// This program demonstrates concurrent searches from multiple sources (SearchSrc),
// using contexts to manage cancellation and merging results into a single channel.
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type SearchSrc struct {
	ID    string
	Delay int
}

func (s *SearchSrc) Search(ctx context.Context) <-chan string {
	out := make(chan string)
	go func() {
		time.Sleep(time.Duration(s.Delay) * time.Second)
		select {
		case out <- "Result " + s.ID:
		case <-ctx.Done():
			fmt.Println("Search received Done()")
		}
		close(out)
		fmt.Println("Search finished for ID: " + s.ID)
	}()
	return out
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	src1 := &SearchSrc{"1", 2}
	src2 := &SearchSrc{"2", 6}

	r1 := src1.Search(ctx)
	r2 := src2.Search(ctx)

	out := merge(ctx, r1, r2)

	for firstResult := range out {
		cancel() // Cancel context after receiving the first result
		fmt.Println("First result is: " + firstResult)
	}
}

func merge(ctx context.Context, results ...<-chan string) <-chan string {
	wg := sync.WaitGroup{}
	out := make(chan string)

	output := func(c <-chan string) {
		defer wg.Done()
		select {
		case <-ctx.Done():
			fmt.Println("Received ctx.Done()")
		case res := <-c:
			out <- res
		}
	}

	wg.Add(len(results))
	for _, c := range results {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}


// This program detects errors in each line of the provided data using goroutines managed by errgroup,
// reporting any lines containing the substring "error:" as errors.
package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"strings"

	"golang.org/x/sync/errgroup"
)

const data = `line one
line two with more words
error: This is erroneous line`

func main() {
	log.Printf("Application %s starting.", "Error Detection")
	scanner := bufio.NewScanner(strings.NewReader(data))
	scanner.Split(bufio.ScanLines)

	// Use errgroup to manage multiple goroutines and errors
	g, _ := errgroup.WithContext(context.Background())
	for scanner.Scan() {
		row := scanner.Text()
		g.Go(func() error {
			if strings.Contains(row, "error:") {
				return fmt.Errorf("Error detected: %s", row)
			}
			return nil
		})
	}

	// Wait for all goroutines to complete and check for any errors
	if err := g.Wait(); err != nil {
		fmt.Println("Error while waiting: " + err.Error())
	}
}



