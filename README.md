# How to run server and client

1. Download the latest version of the server and client.
To run the server:
2. Open terminal/cmd. Move to the folder of the server. Run the .jar file:
   java -jar <filename>.jar
3. Choose the map you want to run your codes on.
To run clients:
3. Choose Edit Configurations... from Run tab on GoLand IDE.
4. Set name of the configuration to :
    `go build main.go network.go controller.go ai.go`
5. Run your code (Shift+F10) as the number of players in the game (4 times). Obviously, you can run different clients.
6. If you're not using GoLand IDE, you can run the preceding command in terminal/cmd (in src/client/ folder) and then run the executable file, which is built, four times.
7. You can also use following command in 4 different terminals/cmds instead of building the project :
   `go run .`
