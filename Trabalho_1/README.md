# First Exercise

![demonstration](https://cdn.discordapp.com/attachments/539836343094870016/871496358144638986/unknown.png)

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
$ git clone <https://github.com/DantasB/distributed-systems>

# Access the project page on your terminal
$ cd distributed-systems/Trabalho_1/


# Execute the main program
$ go run main.go

# The code will start and them you will need to choose which program built do you want to run. 

# There are three implementations:
# IPC using Pipes (you just need to run the pipe code)
# IPC using Signals (you will need to run the main program twice, one for signal_rec and other for signal_sen)
# IPC using Sockets (you will need to run the main program twice, one for socket_server and other for socket_client)
```
![demonstration](https://cdn.discordapp.com/attachments/539836343094870016/871498055273299998/unknown.png)


## Technologies

* Golang
* Go modules