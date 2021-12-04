package main

type Command struct {
	direction string
	amount    int
}

func Challenge2aFile(s string) (int,int) {
	var commands = StringToCommand(ReadMeasurements(s))
	return Challenge2a(commands)
}

func Challenge2a(commands []Command) (int,int) {
	var horizontal = 0
	var depth = 0

	for i := 0; i < len(commands); i++ {
		if commands[i].direction == "forward" {
			horizontal += commands[i].amount
		} else if commands[i].direction == "up" {
			depth -= commands[i].amount
		} else  if commands[i].direction == "down" {
			depth += commands[i].amount
		}
	}
	return horizontal, depth
}


