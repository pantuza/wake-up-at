# Wake up at
Command line tool that helps you know what is the best time to go to sleep based on the time you want to wake up.

# Description

Human sleep occurs in periods of approximately 90 minutes.
One of the cycles is known as *deep sleep*.
if you wake up in the middle of this cycle, probably you will feel a little desoriented and slow.
For that, this program calculates the best times to you to go to sleep based in the time you want to wake up trying to avoid the deep sleep cycle.

Imagine you are programming at 1 *am* and you just realized that you should go to sleep because you must wake up
7:15 at morning.
If you go to sleep right now, it is possible to you wake up unwell and slow if you wake during your deep sleep.
To solve that, you just have to type in your terminal `wake-up-at -h 7 -m 15` and the program calculates the best
times from now on that you should sleep.

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
