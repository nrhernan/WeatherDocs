# Up and Running

Create an API that can be queried for the weather forecast by zip code and deploy it

## Install Go & Visual Studio code

Download and install Go from https://go.dev and Visual Studio Code from https://code.visualstudio.com

## Create the Project

Create a folder called `Dev` and inside that create a folder called `weather`.

Open the Windows Start menu and start typing `PowerShell` and open PowerShell. Navigate to the `weather` folder you just created. For example `cd Dev/weather` if you created `Dev` in your home folder.

Still in PowerShell, run `go mod init` and then `go get github.com/labstack/echo/v4`

## Starting the Project

Open Visual Studio Code and File -> Open the weather folder. Create a file called `server.go`. VSCode will ask if you want to install the recommended extensions and click install. It will say that gopls isn't installed and you can click "Install All". You'll see some text in the output pane of VSCode as it downloads and installs the Go Language Server.

When it's done, it'll say "All tools successfully installed. You are ready to Go. :)".

You can close the tabs "Go for VS Code" and "Extension: Go" and return to the server.go file. You can also close the drawer showing the output.

Inside server.go, paste:

```go
package main

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
```

Go to the run menu and choose "Run without debugging". It'll say that the dlv command is not available. Click install. The output drawer will come up again and you'll wait for it to finish. Wait for "All tools successfully installed. You are ready to Go. :)".

Now Run -> Run without debugging again. Windows Defender might tell you that you need to allow firewall access. You can just allow access on home networks.

You'll see the drawer show some output and say "http server started on [::]:1323". Open your web browser and naviate to localhost:1323 and you'll see it say "Hello, World!"

You now have it working and can start doing real things.

## Explaining the Code

Feel free to skip this if you don't care.

`package main`

Every file in Go belongs to a package. The `main` package is a special one telling Go where to start. The `main` method inside the `main` package is the entry point of your program. You have to have some way of saying where to start and `main` is a pretty common convention.

`import`

Programs use code from other people and you want a way of referencing that.

`e := echo.New()`

You're creating a new variable `e` and assigning the result of calling `echo.New()` to it.

In Go, you need to declare variables before using or assigning anything to them. The `:=` syntax is a shortcut that declars and assigns. In other languages, you'd usually do something like `var e = echo.New()` or in languages like Python you don't need to declare the variable before using it.

In Go, all functions from other packages will be prefixed by the package name. In Python, you can do `from json import loads` and then do `loads('{"Hello": "World"}')` In Go, you'll import the package rather than individual functions.

`echo.New()` returns a new Echo server instance. If you hover over `New` VSCode will give you a bit of documentation. If you Ctrl-Click on `New` you can see the code of the function. Basically, it creates a instance of the `Echo` struct and fills it in with some default values and then returns it.

A struct, object, record, or other name that a language wnats to use is a place to store data, potentially with some methods that act on that data (and methods are just functions on the struct's data).

If you click on the `Echo` name, you can see that it's just a bunch of fields like `Debug bool`.

`e.GET("/", func(c echo.Context) error {...`

This sets a route for Echo. If a request comes in looking for `/`, it'll get routed to the function provided.

The function, in this case, is a lambda function which just means that it's a function we haven't given a name to.

`e.GET` is a method on the `Echo` instance created on the previous line. It simply stores that when `/` is visited, call that function.

`e.Logger.Fatal(e.Start(":1323"))`

`e.Start(":1323")` tells Echo to start resonding to requests on port 1323. Ports are just a construct that lets the operating system know what program should be sent data it receives on that port. Echo uses 1323, but you could use another port (some are quasi-reserved, especially below 1,000).

`e.Logger.Fatal(...)` just makes it so that if the server can't start, you'll see the message telling you why. Echo has a built in logger and loggers usually have a bunch of log levels like debug, info, warn, error, and fatal.

## Your First API

We'll now create an endpoint that will tell us the weather by zip code (with fake data at first).

First, we need to design the response. We'll create a struct that conveys what we want to return.

```go
type WeatherResponse struct {
	ZipCode     string
	Temperature int
}
```

That's how you declare a struct in Go. VSCode will auto-format things when you save so you don't have to worry about the spacing.

Fields must be capitalized if you want them to be accessible. Lower-case fields are private and won't be accessable to other packages or sent to users. In Java and many languages, there are explicit `public` and `private` modifiers on fields and methods. Go decided to use capitalization.

We now want to make a function that returns a `WeatherResponse`:

```go
func Weather(c echo.Context) error {
	zip := c.Param("zip")
	response := WeatherResponse{
		ZipCode:     zip,
		Temperature: 0,
	}
	return c.JSON(http.StatusOK, response)
}
```

The `echo.Context` contains the state of the request (the data the user has sent to the server) and the response (what you're sending back to the user). We get the zip code from this context with `c.Param`. We create a new `WeatherContext` and fill in the zip code and set the temperature to 0. We then use the context to send JSON back to the user with an `OK` status.

If you Ctrl-click on `StatusOK` you can see that it's just a constant for the integer 200.

If anything goes wrong in sending that JSON to the user, `c.JSON()` will return an error (which you're returning from the function) and that can be handled elsewhere (Echo can log it or you could have more advanced error handling like sending things to a metrics dashboard at a big company).

Now we just need to create a route for it. Back in `main` we can add a new route (before the server is started in that method): `e.GET("/weather/:zip", Weather)`

Like before, we call `e.GET`. This time, we've provided a larger route and it has a parameter `:zip`. The colon there tells Echo that portion of the path is going to be a variable that should map to the name `zip`. Other frameworks might do `/weather/{zip}` or `/weather/<zip>`.

In this case, instead of giving `e.GET` a lambda function, we've given it a named function (`Weather`). Notice how we're not calling the function, but just passing the function (it will get called by Echo when someone visits the route).

Now if you start (or restart) the server, you can go to `localhost:1323/weather/90210` and get a weather report...sort of.

