For this project, I implemented the following commands:
    alias
    echo
    history
    pwd
    time
Each command does what you'd generally expect it to do except alias which requires a specific format for its arguement. If you want to add the alias "c" for "cd," you'd type "alias (cd = c)" in the shell. Typing "c" will now grant the same functionality as "cd." This can be stacked for the commands as well, giving a command many aliases if desired. Removing aliases must be done in the "aliases.txt" file in the "Project2" directory.

The "history" command will print historical commands (not aliases) by creating a temporary document "history.txt" in the "Project2" directory. The document is created when the first command is successfully entered and deleted upon gracefully exiting the program. Interrupting the program will cause the document to survive until the next run of the shell with no known negative repercussions other than the command printing more than one session's worth of histroical input. 

"echo" simply repeats the arguement entered.

"pwd" prints the working directory.

"time" outputs the current time in this timezone.

"main.go" will need to be copied completely, not just the switch and GitHub path. I installed the linting package mentioned and corrected my code and repository until it stopped providing recommended changes. I will provide a .png in the "Project2" directory as proof. No automated test coverage will be included in the submission. I had too much trouble trying to figure out how to fulfill that requirement and I've already sunk a rediculous amount of time over the last week learning golang. I apologize for this. Nothing further should be needed in order to build the shell.
