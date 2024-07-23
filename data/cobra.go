

//Initialize a new cobra application : 

package main

import (
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.Execute()
}


// Add a version Command

package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

func main() {
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of the app",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("App v0.1")
		},
	}

	rootCmd.AddCommand(versionCmd)
	rootCmd.Execute()
}


// Add a help Command

package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

func main() {
	rootCmd.InitDefaultHelpCmd()
	rootCmd.Execute()
}



//Add custom command with arguments 

package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

func main() {
	var echoCmd = &cobra.Command{
		Use:   "echo [message]",
		Short: "Echo the provided message",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args[0])
		},
	}

	rootCmd.AddCommand(echoCmd)
	rootCmd.Execute()
}



// Add a persistent flag

package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	rootCmd   = &cobra.Command{Use: "app"}
	echoCmd   = &cobra.Command{Use: "echo [message]", Short: "Echo the provided message", Args: cobra.MinimumNArgs(1), Run: echoRun}
	rootName  string
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&rootName, "name", "n", "World", "name to greet")
	rootCmd.AddCommand(echoCmd)
}

func echoRun(cmd *cobra.Command, args []string) {
	fmt.Printf("Hello, %s!\n", rootName)
}

func main() {
	rootCmd.Execute()
}



// Add a local flag 

package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	rootCmd   = &cobra.Command{Use: "app"}
	echoCmd   = &cobra.Command{Use: "echo [message]", Short: "Echo the provided message", Args: cobra.MinimumNArgs(1), Run: echoRun}
	echoTimes int
)

func init() {
	echoCmd.Flags().IntVarP(&echoTimes, "times", "t", 1, "number of times to echo the message")
	rootCmd.AddCommand(echoCmd)
}

func echoRun(cmd *cobra.Command, args []string) {
	for i := 0; i < echoTimes; i++ {
		fmt.Println(args[0])
	}
}

func main() {
	rootCmd.Execute()
}



//Adds a command with subcommands 

package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{Use: "app"}
)

func init() {
	var addCmd = &cobra.Command{Use: "add", Short: "Add items"}
	var sumCmd = &cobra.Command{Use: "sum", Short: "Sum items"}
	var listCmd = &cobra.Command{Use: "list", Short: "List items"}

	addCmd.AddCommand(sumCmd, listCmd)
	rootCmd.AddCommand(addCmd)
}

func main() {
	rootCmd.Execute()
}



// Uses pre-run and post-run hooks 

package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{Use: "app"}
)

func main() {
	var echoCmd = &cobra.Command{
		Use:   "echo [message]",
		Short: "Echo the provided message",
		Args:  cobra.MinimumNArgs(1),
		PreRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("Preparing to echo...")
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args[0])
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("Echo complete.")
		},
	}

	rootCmd.AddCommand(echoCmd)
	rootCmd.Execute()
}


// Adds completion support  : 


package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

func init() {
	rootCmd.InitDefaultCompletionCmd()
}

func main() {
	rootCmd.Execute()
}



// Defines a custom usage template 

package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

func init() {
	rootCmd.SetUsageTemplate(`Usage:
  {{.UseLine}}
  
{{if .HasAvailableSubCommands}}{{.Commands}}{{end}}`)
}

func main() {
	rootCmd.Execute()
}



// Code to use annotations in commands 


package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

func init() {
	var echoCmd = &cobra.Command{
		Use:   "echo [message]",
		Short: "Echo the provided message",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args[0])
		},
		Annotations: map[string]string{"group": "utility"},
	}

	rootCmd.AddCommand(echoCmd)
}

func main() {
	rootCmd.Execute()
}




// Code to use commands aliases 

package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

func init() {
	var echoCmd = &cobra.Command{
		Use:     "echo [message]",
		Aliases: []string{"say", "repeat"},
		Short:   "Echo the provided message",
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args[0])
		},
	}

	rootCmd.AddCommand(echoCmd)
}

func main() {
	rootCmd.Execute()
}


// Code to disaply command usage : 


package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

func main() {
	var echoCmd = &cobra.Command{
		Use:   "echo [message]",
		Short: "Echo the provided message",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args[0])
		},
	}

	echoCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Println("This is a custom help message for the echo command.")
	})

	rootCmd.AddCommand(echoCmd)
	rootCmd.Execute()
}


// Code to define command groups : 

package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

func init() {
	var addCmd = &cobra.Command{Use: "add", Short: "Add items"}
	var sumCmd = &cobra.Command{Use: "sum", Short: "Sum items"}
	var listCmd = &cobra.Command{Use: "list", Short: "List items"}

	addCmd.AddCommand(sumCmd, listCmd)
	rootCmd.AddCommand(addCmd)
}

func main() {
	rootCmd.Execute()
}




// Code to mark a flag as required : 


package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	rootCmd   = &cobra.Command{Use: "app"}
	echoCmd   = &cobra.Command{Use: "echo [message]", Short: "Echo the provided message", Args: cobra.MinimumNArgs(1), Run: echoRun}
	echoTimes int
)

func init() {
	echoCmd.Flags().IntVarP(&echoTimes, "times", "t", 0, "number of times to echo the message")
	echoCmd.MarkFlagRequired("times")
	rootCmd.AddCommand(echoCmd)
}

func echoRun(cmd *cobra.Command, args []string) {
	for i := 0; i < echoTimes; i++ {
		fmt.Println(args[0])
	}
}

func main() {
	rootCmd.Execute()
}



// Code to use Cobra's Args package : 

package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

func main() {
	var echoCmd = &cobra.Command{
		Use:   "echo [message]",
		Short: "Echo the provided message",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args[0])
		},
	}

	rootCmd.AddCommand(echoCmd)
	rootCmd.Execute()
}


// Code to customize usage function 

package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

func init() {
	rootCmd.SetUsageFunc(func(cmd *cobra.Command) error {
		fmt.Println("This is a custom usage message.")
		return nil
	})
}

func main() {
	rootCmd.Execute()
}




// Code to use custom error handling : 

package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{Use: "app"}

func main() {
	rootCmd.Execute()
}

func init() {
	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		fmt.Println("Running persistent pre-run command")
		return nil
	}
	rootCmd.PersistentPostRunE = func(cmd *cobra.Command, args []string) error {
		fmt.Println("Running persistent post-run command")
		return nil
	}
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:   "help",
		Short: "Help for the app",
		RunE: func(cmd *cobra.Command, args []string) error {
			return fmt.Errorf("this is a custom error")
		},
	})
	rootCmd.SetFlagErrorFunc(func(cmd *cobra.Command, err error) error {
		fmt.Println("Custom flag error message")
		return err
	})
	rootCmd.SetArgs(os.Args[1:])
}



// COde to Add a persistent pre-run command

package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

func init() {
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		fmt.Println("This runs before any command.")
	}
}

func main() {
	rootCmd.Execute()
}



// Code to add persistent post-run command : 

package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

func init() {
	rootCmd.PersistentPostRun = func(cmd *cobra.Command, args []string) {
		fmt.Println("This runs after any command.")
	}
}

func main() {
	rootCmd.Execute()
}




// Code to Use variable argument functions

package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

func main() {
	var echoCmd = &cobra.Command{
		Use:   "echo [messages...]",
		Short: "Echo the provided messages",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for _, arg := range args {
				fmt.Println(arg)
			}
		},
	}

	rootCmd.AddCommand(echoCmd)
	rootCmd.Execute()
}


// Code to chain commands together : 


package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

func main() {
	var echoCmd = &cobra.Command{
		Use:   "echo [message]",
		Short: "Echo the provided message",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args[0])
		},
	}

	var upperCmd = &cobra.Command{
		Use:   "upper [message]",
		Short: "Echo the provided message in uppercase",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(strings.ToUpper(args[0]))
		},
	}

	rootCmd.AddCommand(echoCmd)
	echoCmd.AddCommand(upperCmd)
	rootCmd.Execute()
}




// Override default help command 

package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

func init() {
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:   "help",
		Short: "Show custom help for the app",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("This is a custom help message.")
		},
	})
}

func main() {
	rootCmd.Execute()
}


// Code to persist a configuration file : 

package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{Use: "app"}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
	}
}

func main() {
	rootCmd.Execute()
}


// code to print command line flags : 

package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

func main() {
	var echoCmd = &cobra.Command{
		Use:   "echo [message]",
		Short: "Echo the provided message",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Flags().VisitAll(func(f *cobra.Flag) {
				fmt.Printf("%s: %s\n", f.Name, f.Value)
			})
			fmt.Println(args[0])
		},
	}

	rootCmd.AddCommand(echoCmd)
	rootCmd.Execute()
}


// Code to run commands in sequence : 

package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

func main() {
	var echoCmd = &cobra.Command{
		Use:   "echo [message]",
		Short: "Echo the provided message",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Flags().VisitAll(func(f *cobra.Flag) {
				fmt.Printf("%s: %s\n", f.Name, f.Value)
			})
			fmt.Println(args[0])
		},
	}

	rootCmd.AddCommand(echoCmd)
	rootCmd.Execute()
}


// Go code to Chain flag values

package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

func main() {
	var times int
	var echoCmd = &cobra.Command{
		Use:   "echo [message]",
		Short: "Echo the provided message",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for i := 0; i < times; i++ {
				fmt.Println(args[0])
			}
		},
	}

	echoCmd.Flags().IntVarP(&times, "times", "t", 1, "number of times to echo the message")
	rootCmd.AddCommand(echoCmd)
	rootCmd.Execute()
}



// code to Get flag values directly

package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

func main() {
	var times int
	var echoCmd = &cobra.Command{
		Use:   "echo [message]",
		Short: "Echo the provided message",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			times, _ = cmd.Flags().GetInt("times")
			for i := 0; i < times; i++ {
				fmt.Println(args[0])
			}
		},
	}

	echoCmd.Flags().IntVarP(&times, "times", "t", 1, "number of times to echo the message")
	rootCmd.AddCommand(echoCmd)
	rootCmd.Execute()
}



// Go code to generate bash completions : 


package main

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{Use: "app"}

func main() {
	var completionCmd = &cobra.Command{
		Use:   "completion",
		Short: "Generate bash completion script",
		Run: func(cmd *cobra.Command, args []string) {
			rootCmd.GenBashCompletion(os.Stdout)
		},
	}

	rootCmd.AddCommand(completionCmd)
	rootCmd.Execute()
}




// code to Generate Zsh completions : 

package main

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{Use: "app"}

func main() {
	var completionCmd = &cobra.Command{
		Use:   "completion",
		Short: "Generate zsh completion script",
		Run: func(cmd *cobra.Command, args []string) {
			rootCmd.GenZshCompletion(os.Stdout)
		},
	}

	rootCmd.AddCommand(completionCmd)
	rootCmd.Execute()
}



// code to use command with persistent flags 

package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

func main() {
	var echoCmd = &cobra.Command{
		Use:   "echo [message]",
		Short: "Echo the provided message",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args[0])
		},
	}

	echoCmd.PersistentFlags().String("prefix", "", "prefix message")
	rootCmd.AddCommand(echoCmd)
	rootCmd.Execute()
}



// Code to pass arguments to subcommands 

package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

func main() {
	var echoCmd = &cobra.Command{
		Use:   "echo [message]",
		Short: "Echo the provided message",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args[0])
		},
	}

	var repeatCmd = &cobra.Command{
		Use:   "repeat [times] [message]",
		Short: "Repeat the provided message a number of times",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			times := args[0]
			message := args[1]
			fmt.Printf("%s repeated %s times\n", message, times)
		},
	}

	rootCmd.AddCommand(echoCmd)
	rootCmd.AddCommand(repeatCmd)
	rootCmd.Execute()
}




// Go code to use shorthand for commands 

package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

func main() {
	var echoCmd = &cobra.Command{
		Use:     "echo [message]",
		Short:   "Echo the provided message",
		Aliases: []string{"e"},
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args[0])
		},
	}

	rootCmd.AddCommand(echoCmd)
	rootCmd.Execute()
}


// code to add dynamic flags 

package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var rootCmd = &cobra.Command{Use: "app"}

func main() {
	var echoCmd = &cobra.Command{
		Use:   "echo [message]",
		Short: "Echo the provided message",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			times, _ := strconv.Atoi(args[0])
			message := args[1]
			for i := 0; i < times; i++ {
				fmt.Println(message)
			}
		},
	}

	rootCmd.AddCommand(echoCmd)
	rootCmd.Execute()
}


// Code to override global flags 

package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

func main() {
	var verbose bool
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")

	var echoCmd = &cobra.Command{
		Use:   "echo [message]",
		Short: "Echo the provided message",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if verbose {
				fmt.Println("Verbose mode enabled")
			}
			fmt.Println(args[0])
		},
	}

	rootCmd.AddCommand(echoCmd)
	rootCmd.Execute()
}


// Go code to use built-in root command :

package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

func main() {
	var echoCmd = &cobra.Command{
		Use:   "echo [message]",
		Short: "Echo the provided message",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args[0])
		},
	}

	rootCmd.AddCommand(echoCmd)
	rootCmd.Execute()
}




// Go code to use custom version command : 

package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

func init() {
	rootCmd.Version = "1.0.0"
}

func main() {
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of the app",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("App v1.0.0")
		},
	}

	rootCmd.AddCommand(versionCmd)
	rootCmd.Execute()
}



// GO code to use command groups

package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

func main() {
	var echoCmd = &cobra.Command{
		Use:   "echo [message]",
		Short: "Echo the provided message",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args[0])
		},
	}

	var repeatCmd = &cobra.Command{
		Use:   "repeat [message]",
		Short: "Repeat the provided message",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for i := 0; i < 2; i++ {
				fmt.Println(args[0])
			}
		},
	}

	rootCmd.AddCommand(echoCmd)
	rootCmd.AddCommand(repeatCmd)
	rootCmd.Execute()
}



// Code to use dynaic argument parsing : 

package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

func main() {
	var echoCmd = &cobra.Command{
		Use:   "echo [messages...]",
		Short: "Echo the provided messages",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for _, arg := range args {
				fmt.Println(arg)
			}
		},
	}

	rootCmd.AddCommand(echoCmd)
	rootCmd.Execute()
}



// code to use persistent flag values 

package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

func main() {
	var prefix string
	rootCmd.PersistentFlags().StringVarP(&prefix, "prefix", "p", "", "prefix message")

	var echoCmd = &cobra.Command{
		Use:   "echo [message]",
		Short: "Echo the provided message",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(prefix + args[0])
		},
	}

	rootCmd.AddCommand(echoCmd)
	rootCmd.Execute()
}



// go code to use context in commands : 

package main

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

func main() {
	var echoCmd = &cobra.Command{
		Use:   "echo [message]",
		Short: "Echo the provided message",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := cmd.Context()
			message := args[0]
			fmt.Println("Context:", ctx)
			fmt.Println("Message:", message)
		},
	}

	rootCmd.AddCommand(echoCmd)
	rootCmd.ExecuteContext(context.Background())
}




// go code to use silent errors : 

package main

import (
	"errors"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

func main() {
	var errorCmd = &cobra.Command{
		Use:   "error",
		Short: "Generate an error",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("this is a silent error")
		},
		SilenceErrors: true,
	}

	rootCmd.AddCommand(errorCmd)
	rootCmd.Execute()
}



// go code to use silent usage : 

package main

import (
	"errors"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

func main() {
	var errorCmd = &cobra.Command{
		Use:   "error",
		Short: "Generate an error",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("this is a silent error")
		},
		SilenceUsage: true,
	}

	rootCmd.AddCommand(errorCmd)
	rootCmd.Execute()
}



// go code to print the version of the application 

package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var version string

var rootCmd = &cobra.Command{Use: "app"}

func init() {
	rootCmd.Version = version
}

func main() {
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of the app",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("App version: %s\n", version)
		},
	}

	rootCmd.AddCommand(versionCmd)
	rootCmd.Execute()
}


// go code to use variable argument count 

package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

func main() {
	var sumCmd = &cobra.Command{
		Use:   "sum [numbers...]",
		Short: "Sum the provided numbers",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			sum := 0
			for _, arg := range args {
				var num int
				fmt.Sscanf(arg, "%d", &num)
				sum += num
			}
			fmt.Println("Sum:", sum)
		},
	}

	rootCmd.AddCommand(sumCmd)
	rootCmd.Execute()
}



// go code to use custom output 

package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{Use: "app"}

func main() {
	var echoCmd = &cobra.Command{
		Use:   "echo [message]",
		Short: "Echo the provided message",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			cmd.SetOut(os.Stdout)
			cmd.Println(args[0])
		},
	}

	rootCmd.AddCommand(echoCmd)
	rootCmd.Execute()
}



// go code to combine multiple flags : 

package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var rootCmd = &cobra.Command{Use: "app"}

func main() {
	var prefix, suffix string
	var echoCmd = &cobra.Command{
		Use:   "echo [message]",
		Short: "Echo the provided message",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			message := args[0]
			fmt.Println(prefix + message + suffix)
		},
	}

	echoCmd.Flags().StringVarP(&prefix, "prefix", "p", "", "prefix for the message")
	echoCmd.Flags().StringVarP(&suffix, "suffix", "s", "", "suffix for the message")
	rootCmd.AddCommand(echoCmd)
	rootCmd.Execute()
}




// code to handle subcommands errors : 

package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

func main() {
	var errorCmd = &cobra.Command{
		Use:   "error",
		Short: "Generate an error",
		RunE: func(cmd *cobra.Command, args []string) error {
			return fmt.Errorf("an error occurred")
		},
	}

	rootCmd.AddCommand(errorCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}



// code to use nested subcommands : 


package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

func main() {
	var nestedCmd = &cobra.Command{
		Use:   "nested",
		Short: "Nested command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Nested command executed")
		},
	}

	var subCmd = &cobra.Command{
		Use:   "sub",
		Short: "Sub command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Sub command executed")
		},
	}

	nestedCmd.AddCommand(subCmd)
	rootCmd.AddCommand(nestedCmd)
	rootCmd.Execute()
}




// code to use custom error messages 

package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "app"}

func main() {
	var errorCmd = &cobra.Command{
		Use:   "error",
		Short: "Generate an error",
		RunE: func(cmd *cobra.Command, args []string) error {
			return fmt.Errorf("custom error message")
		},
	}

	rootCmd.AddCommand(errorCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

