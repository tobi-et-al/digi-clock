package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {

	var digits = [...][5]string{
		[5]string{
			"██████",
			"█    █",
			"█    █",
			"█    █",
			"██████",
		},
		[5]string{
			"  █  ",
			"  ███",
			"  █  ",
			"  █  ",
			"██████",
		},
		[5]string{
			"██████",
			"     █",
			"██████",
			"█   ",
			"██████",
		},
		[5]string{
			"██████",
			"     █",
			"   ███",
			"     █",
			"██████",
		},
		[5]string{
			"    █ ",
			"  █ █ ",
			" █  █ ",
			"██████",
			"    █ ",
		},
		[5]string{
			"██████",
			"█     ",
			"██████",
			"     █",
			"██████ ",
		},
		[5]string{
			"██████",
			"█     ",
			"██████",
			"█    █",
			"██████ ",
		},
		[5]string{
			"██████",
			"    █ ",
			"   █  ",
			"  █   ",
			" █    ",
		},
		[5]string{
			"██████",
			"█    █",
			"██████",
			"█    █",
			"██████ ",
		},
		[5]string{
			"██████",
			"█    █",
			"██████",
			"     █",
			"██████ ",
		},
	}

	delimeter := [5]string{
		"     ",
		"  ▒  ",
		"     ",
		"  ▒  ",
		"     ",
	}

	for {
		timeNow := time.Now()
		var timeAsString = fmt.Sprintf("%02d", timeNow.Hour()) + fmt.Sprintf("%02d", timeNow.Minute()) + fmt.Sprintf("%02d", timeNow.Second())
		timeAsArray := strings.Split(timeAsString, "")
		fmt.Printf("\n")
		for i := 0; i < len(digits[0]); i++ {
			for j, val := range timeAsArray {
				val, _ := strconv.Atoi(val)
				fmt.Printf("%s	", digits[val][i])

				if 1 == j%2 && j < (len(timeAsArray)-1) {
					fmt.Printf("%s	", delimeter[i])
				}

			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n")

		time.Sleep(1000 * time.Millisecond)
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	return
}
