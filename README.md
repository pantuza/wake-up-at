# Wake up at
Command line tool that helps you know what is the best time to go to sleep based on the time you want to wake up.

# How to use

```bash
#             hour | minute
$> wake-up-at -h 7 -m 15
To wake up at 7:15 AM, you should sleep at: 10:15 PM

Also at: 11:45 PM | 1:15 AM | 2:45 AM
```

# Command line options
        
**-h**
> Desired hour to wake up. Ex: -h 7

**-m**
> Desired minute to wake up. Ex: -m 15

**-p**
> Period of the day to wake up (am or pm). Default is am. Ex: -p pm

**--help**
> Prints wake-up-at help message

# Contributing

Constributions are very welcome. The project has a Makefile with the following target rules:

* build: Builds the source code. The binary file 'wake-up-at' is created;
* run: Run the program. Used for development purpose;
* test: Run the project tests;
* clean: Clean up the project by removing the binary file;

To contribute, just fork this project, add some code and create a Pull request. Don't forget to write and run tests.

### Author

Written by [Gustavo Pantuza](https://pantuza.com)

### Reporting Bugs

Report wake-up-at bugs by creating issues for this project or by email through gustavopantuza@gmail.com
