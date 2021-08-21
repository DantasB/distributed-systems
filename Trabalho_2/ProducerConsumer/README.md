# Producer Consumer

![demonstration](https://cdn.discordapp.com/attachments/878687554470289438/878690124861411398/teste_ratton-2021-08-21_14.02.23_online-video-cutter.com.gif)

## Table of Contents

<!--ts-->
   * [About](#about)
   * [Requirements](#requirements)
   * [How to use](#how-to-use)
      * [Setting up Program](#program-setup)
   * [Technologies](#technologies)
<!--te-->

## About

This repository is a study about Distributed Systems, where i and my classmate were developing a project for an UFRJ class. The idea of this class assignment was to study about Semaphores where we developt it using golang.

## Requirements

To run this repository by yourself you will need to install golang in your machine and clone this repository to run the code bellow.

## How to use

### Program Setup

```bash
# Clone this repository
$ git clone <https://github.com/DantasB/Distributed-Systems>

# Access the project page on your terminal
$ cd Distributed-Systems/Trabalho_2/ProducerConsumer/results/

# Set the bash script permission to executable
$ chmod +x runner

# Execute the main program
$ ./runner.sh

# The code will start and them you will generate a bunch of files:
# An image containing a graph about the run;
# A folder containing informations about all runs for each parameter;
# A csv containing all data used to generate the graph;
```
![demonstration](https://cdn.discordapp.com/attachments/878687554470289438/878687564129787914/unknown.png)


## Technologies

* Golang
* Go modules
* Weighted Semaphores
