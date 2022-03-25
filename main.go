package main

func main() {
	commandOutput, _, _ := runCommand("git dsd")
	log(1, commandOutput)
}
