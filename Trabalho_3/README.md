# Third Exercise

![demonstration](https://cdn.discordapp.com/attachments/878687554470289438/896991575937269820/unknown.png)

## Table of Contents

<!--ts-->
   * [About](#about)
   * [Requirements](#requirements)
   * [How to use](#how-to-use)
      * [Setting up Program](#program-setup)
   * [Technologies](#technologies)
<!--te-->

## About

This repository is a study about Distributed Systems, where i and my classmate were developing a project for an UFRJ class. The idea of this class assignment was to study about IPC (Inter-process communication) where we developt it using golang.

## Requirements

To run this repository by yourself you will need to install golang in your machine and clone this repository to run the code bellow.

## How to use

### Program Setup

```bash
# Clone this repository
$ git clone <https://github.com/DantasB/Distributed-Systems>

# Access the project page on your terminal
$ cd Distributed-Systems/Trabalho_3/

# Access the coordinator proccess
$ cd Coordinator/

# Start the coordinator code

$ go run .

# In another terminal run the runner.sh code
# You will need to pass 3 parameters
# 1°: number of processes (n)
# 2°: number of repetitions (r)
# 3°: number of seconds to be awaited (k)
$ ./runner.sh 10 5 1

# The code will run and you will get a log containing a check successful message if the logs are ok.
```
![demonstration](https://cdn.discordapp.com/attachments/878687554470289438/896989790019416084/unknown.png)


## Technologies

* Golang
* Go modules