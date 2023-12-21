# Let's check for any Go source files, particularly server.go, in the extracted directory and its subdirectories
# go_files = []

# for root, dirs, files in os.walk(server_go_directory):
#     for file in files:
#         if file.endswith(".go"):
#             go_files.append(os.path.join(root, file))

# go_files


# Reading the content of server.go file
server_go_file_path = '/mnt/data/server_go_extracted/server-go/cmd/server/server.go'

with open(server_go_file_path, 'r') as file:
    server_go_content = file.read()

server_go_content[:1000]  # Displaying the first 1000 characters to get an idea of the file content
