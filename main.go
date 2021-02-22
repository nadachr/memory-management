package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	process []string
	table   []string
	size    []int
	start   []int
)

func initialized() {
	process = append(process, "free")
	start = append(start, 0)
	size = append(size, 1000)
}

func showProcess() {
	fmt.Printf("\n-----------\n")
	fmt.Printf("Process | Start | Size \n")
	for i := range process {
		fmt.Printf("  %s  |  %d  |  %d  \n", process[i], start[i], size[i])
	}
	fmt.Printf("\n")
	fmt.Printf("\nCommand > ")
}

func getCommand() string {
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString('\n')
	data = strings.Trim(data, "\n")
	return data
}

func command_create(p string, s int) {
	for i := range process {
		if process[i] == "free" { //check size
			if size[i] < s {
				fmt.Printf("\nSize error!!!\n")
				continue
			}
		}
		if i == 0 { // check if process is the first process in system (index = 0)
			if process[i] == "free" {
				process = append(process, " ")
				copy(process[i+1:], process[i:])
				process[i] = p
				size = append(size, 0)
				copy(size[i+1:], size[i:])
				size[i] = s
				size[i+1] = size[i+1] - s
				start = append(start, 0)
				copy(start[i+1:], start[i:])
				start[i] = start[i+1]
				start[i+1] = s
				break
			}
		} else { // or it's another process in system (another = 1, 2, 3, ..., n)
			if process[i] == "free" {
				if size[i] > 0 { // if size of free process is more than 0 then...
					process = append(process, " ")
					copy(process[i+1:], process[i:])
					process[i] = p
					size = append(size, 0)
					copy(size[i+1:], size[i:])
					size[i] = s
					size[i+1] = size[i+1] - s
					start = append(start, 0)
					copy(start[i+1:], start[i:])
					start[i] = start[i+1]
					start[i+1] = size[i] + start[i]
					if size[i+1] == 0 { // if size of free process after create process is 0 then...
						if i+1 == len(process)-1 { // if free is the last process in system then...
							process = process[:len(process)-1]
							start = start[:len(start)-1]
							size = size[:len(size)-1]
						} else { // if free is not the last process in system then...
							process = append(process[:i+1], process[i+2:]...)
							size = append(size[:i+1], size[i+2:]...)
							start = append(start[:i+1], start[i+2:]...)
						}
					}
					break
				} else {
					fmt.Println("error...")
				}
			}
		}
	}
}

func command_terminate(p string) {
	if len(process) > 0 { //check if process is not empty
		for i := range process {
			if process[i] == p { //check if process name
				if i-1 >= 0 { //check if this process is not the first process in system
					if i < len(process)-1 { //check if this process is not the last process in system
						if process[i+1] == "free" && process[i-1] != "free" { //check if there is free process after this process
							size[i+1] = size[i] + size[i+1]
							start[i+1] = start[i]
							process = append(process[:i], process[i+1:]...)
							size = append(size[:i], size[i+1:]...)
							start = append(start[:i], start[i+1:]...)
							break
						} else if process[i-1] == "free" && process[i+1] != "free" { //check if there is free process before this process
							size[i-1] = size[i] + size[i-1]
							start[i-1] = start[i-1]
							process = append(process[:i], process[i+1:]...)
							size = append(size[:i], size[i+1:]...)
							start = append(start[:i], start[i+1:]...)
							break
						} else if process[i-1] == "free" && process[i+1] == "free" { //check if this process is in between 2 free processes
							size[i-1] = size[i] + size[i-1] + size[i+1]
							start[i-1] = start[i-1]
							process = append(process[:i], process[i+1:]...)
							size = append(size[:i], size[i+1:]...)
							start = append(start[:i], start[i+1:]...)
							process = process[:len(process)-1]
							break
						} else { //check if there is no free process near this process
							process[i] = "free"
						}
					} else { //check if this process is the last process in system
						if process[i-1] == "free" { //check if there is free process before this process
							size[i-1] = size[i] + size[i-1]
							start[i-1] = start[i-1]
							process = append(process[:i], process[i+1:]...)
							size = append(size[:i], size[i+1:]...)
							start = append(start[:i], start[i+1:]...)
							break
						} else { //check if there is no free process near this process
							process[i] = "free"
						}
					}
				} else { //check if this process is the first process in system
					if process[i+1] == "free" { //check if there is free process after this process
						size[i+1] = size[i] + size[i+1]
						start[i+1] = start[i]
						process = append(process[:i], process[i+1:]...)
						size = append(size[:i], size[i+1:]...)
						start = append(start[:i], start[i+1:]...)
						fmt.Println("loop8")
						break
					} else { //check if there is no free process near this process
						process[i] = "free"
						fmt.Println("loop9")
					}
				}
			}
		}
	}
}

func main() {
	initialized()
	for {
		showProcess()
		command := getCommand()
		commandx := strings.Split(command, " ")
		switch commandx[0] {
		case "exit":
			return
		case "create":
			pr := strings.Split(commandx[1], "-")
			sz, _ := strconv.Atoi(pr[1])
			if sz > 1000 {
				fmt.Printf("\n Size error!! Please create process's size less than 1000.")
			} else {
				command_create(pr[0], sz)
			}
		case "terminate":
			command_terminate(commandx[1])
		default:
			fmt.Printf("\n Sorry, Command error!")
		}
	}
}
