package main

func Challenge2bFile(s string) (int, int) {
	var commands = StringToCommand(ReadMeasurements(s))
	return Challenge2b(commands)
}

func Challenge2b(commands []Command) (int, int) {
	var horizontal = 0
	var depth = 0
	var aim = 0

	for i := 0; i < len(commands); i++ {
		if commands[i].direction == "forward" {
			horizontal += commands[i].amount
			depth += aim * commands[i].amount
		} else if commands[i].direction == "up" {
			aim -= commands[i].amount
		} else if commands[i].direction == "down" {
			aim += commands[i].amount
		}
	}
	return horizontal, depth
}
