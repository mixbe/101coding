package go_interview

import (
	"fmt"
	"sync"
	"testing"
)

/*
# 1: Print numbers and letters alternately
## Problem description

Use two goroutines to print sequences alternately, one goroutine prints numbers, and the other goroutine prints letters. The final effect is as follows:

12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728

## Solution

The problem is very simple. Use channels to control the progress of printing.
Use two channels to control the printing sequence of numbers and letters respectively.
After the number is printed, the channel notifies the letter to print, and after the letter is printed,
the number is notified to print, and then the work is repeated.
*/

func TestPrintNumbersAndLetters(t *testing.T) {
	printNumbersAndLetters()
}

func printNumbersAndLetters() {
	letter, number := make(chan bool), make(chan bool)
	wg := sync.WaitGroup{}

	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
			}
		}
	}()
	wg.Add(1)
	go func(wait *sync.WaitGroup) {
		i := 'A'
		for {
			select {
			case <-letter:
				if i >= 'Z' {
					wait.Done()
					return
				}

				fmt.Print(string(i))
				i++
				fmt.Print(string(i))
				i++
				number <- true
			}

		}
	}(&wg)
	number <- true
	wg.Wait()
	fmt.Println()
}
