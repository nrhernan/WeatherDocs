# Deploying

We're going to use fly.io to deploy. They have a free tier that lets you run a few apps so we'll use that (it does require a credit card to prevent abuse).

Go to https://fly.io and copy/paste the "Install flyctl on Windows" into PowerShell (`iwr https://fly.io/install.ps1 -useb | iex`). You can sign up for an account at https://fly.io/app/sign-up.

Once you have an account, you can `flyctl auth login` in PowerShell.

We'll need to make some small adjustments to the app to get it ready for deployment on Fly.io.

In the `main()` method, we'll need to get the port the app should listen on in Fly's service.

```go
port := os.Getenv("PORT")
if port == "" {
    port = "1323"
}
e.Logger.Fatal(e.Start(":" + port))
```

VS Code should automatically add an import for "os", but if it doesn't, just add it to the imports.

Many PaaS (platform as a service) offerings will route things to your program based on the port. As noted in the previous section, ports are just a construct that lets the operating system know what program should be sent data it receives on that port. Fly.io will set an environment variable for your program called `PORT` and then your program can read the port and listen on that instead of Echo's default 1323. This allows Fly to run many programs on the same machine and it delivers traffic based on which port the person is listening on.

Now if you use `flyctl launch` in PowerShell, it should package up your application into a container and give you the option to deploy it. In the future, you can deploy changes using `flyctl deploy`.

If you run `flyctl status` you can get the information on the running service. You'll see a `Hostname` which will show you the URL it's deployed at. You can open up your browser and use that URL and you'll see your application. It's now deployed!

You'll notice that it'll add a file `fly.toml` to your directory. This is the build and deploy configuration for your app. Each service is a tiny bit different, but they're all basically the same: they list how something is going to get built, what services it's running, some health checks to make sure that the service is up, etc.

## If you're curious

Fly.io uses "buildpacks". They're a semi-standard way of building and packaging apps for many different languages. Fly.io also supports Docker so that you can build a Docker container of any program and use it with their service. Docker isn't as hard as a lot of people think, but it would require a few extra steps and (realistically) if you're getting a job at a reasonably sized company, they'll already have a build/deploy system.