# Discord Imposter Go

A Discord bot to secretly split your voice channel into teams.

[Invite the bot to your own server!](https://discord.com/oauth2/authorize?client_id=1529308930856194048)

... Or build and host it yourself

## Building

Install go and run `go build` from the repositories root folder.

## Testing

Run benchmarks

``` bash
go test ./formatting -bench=GetDisplayNames -benchmem
```

Run regular tests

``` bash
go test ./formatting
```

## Running

It is recommended to use a `.env` file to protect your application secrets by keeping them out of your terminal's command history, however both `.env` files and cli flags are supported.

### .env File

Create a file called `.env` in the repository's root and add a a value `BOT_TOKEN` with your bot token from the Discord developers portal.

### CLI Flags

Run your built executable with the flag `-token <your token>` with your bot token from the Discord developers portal.
