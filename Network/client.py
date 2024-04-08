import socket

def start_client(server_host='127.0.0.1', server_port=65432):
    # Create a socket object
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        # Connect to the server
        s.connect((server_host, server_port))
        # Send data to the server
        s.sendall(b'Hello, server')
        # Receive the response from the server
        data = s.recv(1024)
        print(f"Received from server: {data.decode()}")

if __name__ == "__main__":
    start_client()
