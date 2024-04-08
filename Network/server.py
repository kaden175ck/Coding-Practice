import socket

def start_server(host='127.0.0.1', port=65432):
    # Create a socket object
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        # Bind the socket to the address and port
        s.bind((host, port))
        # Enable the server to accept connections
        s.listen()
        print(f"Server started, listening on {host}:{port}")
        # Wait for a client to connect
        conn, addr = s.accept()
        with conn:
            print(f"Connected by {addr}")
            while True:
                # Receive data from the client
                data = conn.recv(1024)
                if not data:
                    break
                print(f"Received from client: {data.decode()}")
                # Send data back as a response
                conn.sendall(data)

if __name__ == "__main__":
    start_server()
