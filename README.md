Copy code
# IP Geolocation Lookup in Go

This is a simple Go program that performs a reverse DNS lookup and provides the geolocation of an IP address. It uses the [IPStack](https://ipstack.com/) API to get the geolocation data.

## Requirements

- [Go](https://golang.org/dl/) programming language
- [IPStack](https://ipstack.com/) API access key

## Installation

First, clone the repository or download the source code. 

```bash
git clone https://github.com/yourusername/yourrepository.git
```
Navigate to the directory where you saved your main.go file:

```bash
cd path/to/directory
```
Then, compile the Go code:

```bash
go build
```
This command will create an executable file in the current directory with the same name as the Go file, but without the .go extension.

## Configuration

Before running the program, you need to set the IPSTACK_ACCESS_KEY environment variable with your IPStack access key.

In Linux or MacOS, you can do this in the terminal:

```bash
export IPSTACK_ACCESS_KEY=your_access_key
```
In Windows, you can set the environment variable in the command prompt:

```bash
set IPSTACK_ACCESS_KEY=your_access_key
```
Replace your_access_key with your actual IPStack access key.

## Running the program

Finally, you can run the program by executing the main file, followed by the IP address you want to look up:

```bash
./main 8.8.8.8
```
The program will perform a reverse DNS lookup on the specified IP address and print the location information to the standard output.

Replace 8.8.8.8 with the IP address you want to look up.

License

This project is licensed under the [MIT License](LICENSE).


Note: Replace `yourusername`, `yourrepository`, `path/to/directory`, and `your_access_key` with the appropriate values for your situation.


