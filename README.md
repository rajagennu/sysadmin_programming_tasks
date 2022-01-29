Any Sys Administrator / DevOps Engineer would like to learn new language, can follow below to practice.

# Programming Tasks

[ ] Execute `uptime` command and fetch load of the given system
[ ] Connect to a remote host/server and execute the same uptime command `uptime` and fetch the output.
[ ] Check whether a website is up/down, ex input www.google.com -> should return up if site up else down.
[ ] parse a json file, append a new a json attribute. You can follow below example
```json
[
 { 
  "name" : "Python",
  "type" : "Interpreter",
 },
 {
  "name" : "C",
  "type: : "Compiler"
 }
]
```
Store above json object in a .json file, read it as input from your program using command line args, print the object.
[ ] For above json Object, append below object
```json
{
  "name" : "GoLang",
  "type" : "Compiler"
}
```
so output should be 
```json
[
 { 
  "name" : "Python",
  "type" : "Interpreter",
 },
 {
  "name" : "C",
  "type: : "Compiler"
 },
 {
  "name" : "GoLang",
  "type" : "Compiler"
 }
]
```

[ ] same follow for yaml format

```yaml
---
 - name: "python"
   type : "Interpreter"
```
- store this yaml in a .yaml file and read it as input by your program, print the yaml data

[ ] append below yaml data to above file and print the final content and update the file with new content
```yaml
- name: "C"
  type : "Compiler"
```

so final content from the file should be like below
```yaml
---
 - name: "python"
   type : "Interpreter"
 - name: "C"
   type : "Compiler"

```

[ ] check which service is running on port 22 or any port you think is in use ?
[ ] clone a git repository
[ ] write a program to print all users in your system in below format
```
username  default_shell expiry_date
```
[ ] Generate a JSON object using above user information, like below and store that in a file.
```json
[
  {
    "username" : 
    "default_shell" :
    "expiry_date" :
  }
]
```

[ ]  Generate a yaml object for same and store it in a yaml format
```yaml
---
  - username:
    default_shell
    expiry_date
```

[ ] extend the same programs JSON, YAML to execute on a remote host and this time fetch the file from remote system to local system.
[ ] Manual: Install mySQL DB in your host/remote host
[ ] Manual: Configure user authentication
[ ] Manual: Create a database and simple table with name as column
  [ ] Connect to a MySQL DB
    [ ] execute a query and print the output
    [ ] execute a query and store the output in csv format
    [ ] execute a query to write a new row to MySQL/Maria DB
    [ ] execute a query to delete a row ( user WHERE clause to delete a specific row)
    [ ] execute a query to update a row
[ ] Write a program to start/stop/restart a specific service.
[ ] Start an apache server, which runs on port 80, now using your program update config file of apache to run the server on 8080 and restart the server.
[ ] Install a package by using your program
[ ] now do all above tasks with a single program: apache installation, change default port to 8080, restart apache server
[ ] Build a program, which restarts a remote server, monitor till server comes online and once server online execute uptime command. 
  [ ] User INFO, ERROR, WARNING logging and let the execute know whats going on with logging.
  [ ] If you want to get into a bit advanced use color coding green, red and yellow accoridingly. 
[ ] Write a program to read the data from URL https://api.github.com/users/rajagennu in JSON format and store it in a file.


