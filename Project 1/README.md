Author: Alexander Seneca
Date: 10/15/2023
Course: CSCE4600.002
Project #1: Process Scheduling

This program should not need anything special in order to compile and run. Please execute program with command-syntax "go run main.go example_processes.csv" where "example_processes.csv" can be changed to any relevant testing data. The written code is derivative of prepared code that was provided and is therefore very similar in much of the body. Credit to Professor Jacob Hochstetler for the prepared code.

Note: The code I wrote seems to work fine for me, though the "Round-robin" portion looks to be a rearanged version of the SJF table. The round-robin portion doesn't actually use a time-quantum as one would in practice, but instead works on iterations of 0.01 units per process until the total burst-duration of that process is met. For example, if a process has a burst-duration of 4, the RR process will have to round back to that process 400 times in order to mark the process as fully completed. Also, as you would expect the RR process scheduler does actually move between processes that are incomplete until all processes are resolved. All other schedulers work as intended.
