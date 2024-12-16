# Go Conference Booking App

This is a simple **Conference Booking Application** written in Go. It allows users to book tickets for a conference, validates user input, and sends ticket details asynchronously via email simulation. The program demonstrates core Go concepts such as goroutines, slices, structs, and concurrency.

## Features

- User-friendly console-based interface.
- Real-time validation of user input (name, email, and ticket count).
- Concurrency to simulate sending tickets via email without blocking the main thread.
- Maintains a list of bookings with user details.

## Requirements

- Go installed on your system (version 1.16 or higher).

## How It Works

### Key Components

1. **Structs**:
   - `UserData`: Stores user details like first name, last name, email, and number of tickets booked.
2. **Functions**:
   - `greetUser`: Displays a welcome message and details about the conference.
   - `getUserInput`: Collects user input for booking tickets.
   - `bookTicket`: Processes the booking, updates the remaining tickets, and stores user data.
   - `printFirstName`: Extracts and prints the first names of users who booked tickets.
   - `sendTicket`: Simulates sending tickets via email using a goroutine.
3. **Validation**:
   - Input validation is performed using the `helper.ValidateUserInput` function from an external package `helper`.

### Flow

1. The program welcomes the user and displays available tickets.
2. Users input their details (name, email, and ticket count).
3. Input is validated; if valid, the booking is processed and confirmation is displayed.
4. Tickets are sent asynchronously using a goroutine, simulating email delivery.
5. If all tickets are booked, the program notifies users that the conference is sold out.

## Running the Program

1. Clone the repository or copy the source code.
2. Ensure that the `helper` package is correctly implemented and available in your Go workspace.
3. Run the program:

```bash
go run main.go
```

4. Follow the instructions in the console to book tickets.

## Example Usage

```
Welcome to our Go Conference booking App
We have total of 50 tickets and 50 tickets are still available
Get your ticket here to attend

Enter your firstName: John
Enter your lastName: Doe
Enter your email: john.doe@example.com
Enter the number of tickets: 2

Thank you John Doe for booking 2 tickets. You will receive a confirmation email at john.doe@example.com
48 tickets remaining for Go conference

The first name of the booking are: [John]
```

## Project Structure

```plaintext
.
├── main.go           # Main program file
├── helper            # Package containing validation logic (not included in this example)
```

## Concepts Demonstrated

- **Concurrency**: Use of `sync.WaitGroup` and goroutines to handle asynchronous tasks.
- **Data Structures**: Use of slices and structs to manage user data.
- **Validation**: Basic input validation for better user experience.

## Improvements and Next Steps

- Add file persistence to save bookings.
- Enhance email simulation with more realistic features.
- Build a web interface for the application.

## License

This project is open-source and available under the [MIT License](https://opensource.org/licenses/MIT).

---

Feel free to modify and expand this program as you continue learning Go!
