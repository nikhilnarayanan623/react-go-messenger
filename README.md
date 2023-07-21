# React-Go Messenger

## Description

React-Go Messenger is a real-time messaging application built using React, Go, TypeScript, and Socket.IO. It enables users to communicate with each other instantly and effortlessly.

The application provides a sleek and intuitive user interface, allowing users to send and receive messages in real-time. It leverages the power of React to deliver a seamless and responsive user experience, while the Go backend handles the communication and data processing using Socket.IO.

## Features

- Real-time messaging: Users can send and receive messages in real-time, enabling instant communication.
- User authentication: Secure user authentication system ensures that only authorized users can access the application.
- Typing indicators: Shows when other users are currently typing a message.
- Online presence: Users can see the online status of other users.
- Message history: The application stores and displays the message history for easy reference.

## Technologies Used

- React: A popular JavaScript library for building user interfaces.
- Go: A programming language known for its efficiency and concurrency features.
- TypeScript: A statically-typed superset of JavaScript that enhances the development experience.
- Socket.IO: A library that enables real-time, bidirectional communication between the server and the client.

## Prerequisites

Make sure you have the following software installed on your system:

- Node.js (v12 or above)
- Go (v1.16 or above)

## Configuration

The application can be configured using the .env.example file in both client and server folder
Make sure to set the environment variables before starting the application.

## Getting Started

To run the React-Go Messenger application on your local machine, follow these steps:

1. Clone the repository:

   ```
   git clone https://github.com/your-username/react-go-messenger.git](https://github.com/nikhilnarayanan623/react-go-messenger
   ```

2. Change to the project directory:

   ```
   cd react-go-messenger
   ```
3. Change to the client directory
   ```
   npm install
   ```
3. Install the dependencies:

   ```
   npm install
   ```

5. Start the React frontend:

   ```
   npm run dev
   ```

   This will launch the application in your default browser.

6. Change to the Go backend directory:

   ```
   cd server
   ```
3. Install the dependencies:

   ```
   make deps || go mod tidy
   ```

7. Run the Go backend:

   ```
   make run || go run ./cmd/api/main.go
   ```

8. Open the application in your browser and start messaging!

## Contributing

Contributions to React-Go Messenger are welcome! If you find any bugs or want to suggest improvements, please open an issue or submit a pull request on the GitHub repository.

## License

This project is licensed under the [MIT License](LICENSE). Feel free to use, modify, and distribute the code as per the license terms.

## Acknowledgments

- React: https://reactjs.org/
- Go: https://golang.org/
- TypeScript: https://www.typescriptlang.org/
- Socket.IO: https://socket.io/

## Contact

If you have any questions or need further assistance, please feel free to contact the project maintainers:

- Nikhil N - nikhilnarayanan623@gmail.com
- Abin T H - abinth250@gmail.com

We hope you enjoy using React-Go Messenger! Happy messaging!
