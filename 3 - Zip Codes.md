# Zip Codes

In order to get weather data for a zip code, we'll need to translate that zip code into coordinates.

## Create a SQLite database with a table for the zip codes

Download SQLite for Windows (https://sqlite.org/download.html, the `sqlite-tools-win32` one) and put `sqlite3.exe` in your app's folder.

Create a file in your app's folder called `.gitignore` (note the leading period) and put `sqlite3.exe` in it. This makes it so that you won't commit the exe file to git.

Now you can run `sqlite3.exe weather.db` and you'll be in the SQLite shell. SQLite is pretty similar to databases you've likely used before like MySQL or PostgreSQL. The big difference is that it just opens the file rather than connecting to a server (which makes it easy to set up).

## Create a table to hold the zip code information

It's going to need the zip code and the lat/lng. One tip: trying to wedge a zip code into a numeric field tends to bite people (leading zeros get dropped, people don't add them back when displaying information, etc).

Let me know if you want more information here, but this feels up your alley. Don't forget to index what you're going to query on (and it's good practice to put the SQL commands into a file so it's in git and others can see what happened).

## Download zip code CSV

https://gist.github.com/erichurst/7882666

## Insert the data into the zip code table

I'm going to leave a lot of this up to you with some pointers.

* You can read a file using `os.Open("my_csv.csv")`
* You can use the `encoding/csv` package to read the file as a CSV `csv.NewReader(f)`
* You can use the `github.com/jmoiron/sqlx` package to read/write to the database

Finally, since you can only have one `main()` function, you'll need to differentiate between "run the server" and "insert the CSV data". Programs can take flags/arguments. In your main function, you can do:

```go
var importCsv bool
flag.BoolVar(&importCsv, "import-csv", false, "Whether it should import the csv instead of running the server")
flag.Parse()

if (importCsv) {
    //do the import
} else {
    //run the server
}
```