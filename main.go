package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	features, digits := digits()
	for {
		now := time.Now()
		millis := now.Nanosecond() / 100_000_000
		hour, min, second := now.Hour(), now.Minute(), now.Second()

		clock := []placeholder{
			digits[hour/10], digits[hour%10],
			features[0],
			digits[min/10], digits[min%10],
			features[0],
			digits[second/10], digits[second%10], features[2], digits[millis],
		}

		for line := range clock[0] {
			for index, digit := range clock {
				next := clock[index][line]
				if digit == features[0] && second%2 == 0 {
					next = "   "
				}
				if index == len(clock)-2 || index == len(clock)-1 {
					fmt.Print(next)
				} else {
					fmt.Print("  ", next)
				}

			}
			fmt.Println()
		}

		fmt.Println()
		
		for _, a := range features[1] {
			if second%10 == 0 {
				fmt.Println(" ", a)
			}
		}

		fmt.Println()
		time.Sleep(time.Second)
		fmt.Println()
		clearScreen()
	}

}

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
