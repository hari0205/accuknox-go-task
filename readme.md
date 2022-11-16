# Program to find top three foods in a restaurant

This is a program to find top 3 foods in the restaurant.


## Approach

- The restaurant stores a log file "{somename}.txt" 
- Edit the textfile provided to see different scenarios.
- The format for storing logs are customer_id : item_id1,item_id2,.....,item_idn
- Reading the file as buffer using `bufio` 
- Then splitting the strings and formatting them in required format.
- Create a function to check if the customer_id has a duplicate item. If yes, throw an error
- Map the number of times a particular item is consumed.
- Make a map and sort the key-value pairs so that items that are most consumed are at the top
- Display the top 3 item ids


## Steps
- Clone this repository
- Run `go run main.go`
- For testing `go test`