[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-24ddc0f5d75046c5622901739e7c5dd533143b0c8e959d652212380cedb1ea36.svg)](https://classroom.github.com/a/Bp585G7b)

## Description

- This Github repo contains code to spin up servers and client app in CLI.
- The code is written in Golang.
- User can spin up multiple servers. 
- With one instance of client it will randomly hit on of the server to return a value
- User will have to give a number in the client cli app. And it will return a value which is a double number.

## Setting Up The Project

- Go version go1.21.0 [Install GO](https://go.dev/doc/install)

## How to Use the CLI App
#### Provided in the Makefile are the shortcut to run each comment.

1) In order to create servers run the following command:

```bash
    make server <port>
```
_By default if you don't specify the port it is taken as 5000_

Example:

```bash
make server PORT=5001
```

You can run multiple server by running this command with different port.

```bash
make server
make server PORT=5001
make server PORT=5002``
```

2) In order to create client instance run the following command:

```bash
    make client
```

_This will spin up the client instance._

Once The Client instance is running it will ask user for the following details:

a) List the no of server

b) The server address and server port.

**Example:**

```bash
List the no of server: 3

List the server address #1: localhost
List the server port: 5000

List the server address #2: localhost
List the server port: 5001
  
List the server address #3: localhost
List the server port: 5002
```

After setting up the connection between the **CLIENT** and **SERVER**

The cli will ask to input a number. If provided correct number client will send the input to on of the server (randomized)
and returns backs the double of that number.
