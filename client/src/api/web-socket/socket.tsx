const ws = new WebSocket("ws://localhost:8080/api/ws");
export default ws;
import React, { useState, useEffect } from "react";

const User = () => {
    const [connectToken, setConnectToken] = useState("");
    const [messages, setMessages] = useState([]);
    const [message, setMessage] = useState("");
    const [userID, setUserID] = useState(0);
    const [chatId, setChatId] = useState(0);
    const [msgSent, setMessageSent] = useState(false);
    const [socket, setSocket] = useState(null);
    const [error, setError] = useState("");

    // Connect to the WebSocket on component mount
    useEffect(() => {
        console.log("use effect of connect token");

        if (connectToken) {
            const newSocket = new WebSocket("ws://localhost:8080/api/ws");
            newSocket.onopen = () => {
                console.log("Successfully Connected");

                // sending token
                console.log("!!!!! sending token if you see socket closed then check token is expired !!!!!");
                newSocket.send(JSON.stringify({
                    token: connectToken,
                }))

            };
            newSocket.onclose = event => {
                console.log("Socket Closed Connection: ", event);
            };

            newSocket.onerror = error => {
                console.log("Socket Error: ", error);
            };

            // set socket   
            setSocket(newSocket);

            return () => {
                newSocket.close();
            };
        }
    }, [connectToken]);

    // Handle incoming messages
    

    const handleClick = () => {
        if (connectToken) {
            setError(""); // Reset error state
        } else {
            setError("Please enter a valid access token."); // Show error if access token is empty
        }
    };

    const sendMessage = async () => {
        try {
            if (!connectToken || !chatId || !userID || !message) {
                setError("Please fill in all the required fields.");
                return;
            }

            setError(""); // Reset error state
            const data = {
                content: message,
                receiver_id: parseInt(userID),
            };

            const jsonData = JSON.stringify(data);
            setMessages((prevMessages) => [...prevMessages, "yours:      " + message]);

            //   sendMsg(jsonData);
            socket.send(jsonData);
            setMessageSent(!msgSent);

            const response = await fetch(
                `http://localhost:8080/api/chats/${chatId}/messages`,
                {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                        Authorization: "Bearer " + connectToken,
                    },
                    body: jsonData,
                }
            );

            if (!response.ok) {
                setError("Failed to send the message.");
                return;
            }

            const parsedResponse = await response.json();
            console.log(parsedResponse);
        } catch (error) {
            setError("An error occurred while sending the message.");
            console.log(error);
        }
    };

    return (
        
            <div
              style={{
                width:"100%",
                display: "flex",
                flexDirection: "",
                justifyContent: "space-evenly",
                alignItems: "center",
                marginTop: "2rem",
              }}
            >
              {/* Left Side - Access Token */}
              <div
                style={{
                  padding: "1rem",
                  border: "1px solid black",
                  width: "90%",
                  maxWidth: "400px",
                  marginBottom: "1rem",
                }}
              >
                <h2>Access Token</h2>
                <input
                  type="text"
                  name="token"
                  onInput={(e) => setConnectToken(e.target.value)}
                  value={connectToken}
                  placeholder="Access token"
                />
                <button style={{ marginTop: "1rem" }} onClick={handleClick} type="button">
                  Connect
                </button>
                {error && <p style={{ color: "red" }}>{error}</p>}
              </div>
        
              {/* Send Message */}
              <div
                style={{
                  padding: "1rem",
                  border: "1px solid black",
                  width: "90%",
                  maxWidth: "400px",
                  marginBottom: "1rem",
                }}
              >
                <h2>Send Message</h2>
                <div style={{ padding: "1rem" }}>
                  <h5>Chat Id</h5>
                  <input
                    type="number"
                    onInput={(e) => setChatId(e.target.value)}
                    value={chatId}
                    placeholder="Chat Id"
                  />
                </div>
                <div style={{ padding: "1rem" }}>
                  <h5>Receiver User ID</h5>
                  <input
                    type="number"
                    onInput={(e) => setUserID(e.target.value)}
                    value={userID}
                    placeholder="Receiver User ID"
                  />
                </div>
                <div style={{ padding: "1rem" }}>
                  <h5>Message Content</h5>
                  <input
                    type="text"
                    onInput={(e) => setMessage(e.target.value)}
                    value={message}
                    placeholder="Message"
                  />
                </div>
                <button style={{ marginLeft: "1rem" }} onClick={sendMessage} type="button">
                  Submit
                </button>
                {error && <p style={{ color: "red" }}>{error}</p>}
              </div>
        
              {/* Right Side - Messages */}
              <div
                style={{
                  padding: "1rem",
                  border: "1px solid black",
                  width: "90%",
                  maxWidth: "800px",
                }}
              >
                {/* Message Display */}
                <div>
                  <h1>WebSocket Messages</h1>
                  <ul>
                    {messages.map((msg, index) => {
                      return <ol key={index}>{msg}</ol>;
                    })}
                  </ul>
                </div>
              </div>
            </div>
    );
};

export default User;