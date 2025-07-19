# Pokedex CLI

A command-line Pokédex application built in Go that allows you to explore Pokémon locations and catch Pokémon using the [PokéAPI](https://pokeapi.co/).

## Features

- **Interactive REPL**: Command-line interface with persistent session
- **Location Exploration**: Browse Pokémon location areas with pagination
- **Pokémon Discovery**: Explore specific areas to find Pokémon
- **Pokémon Catching**: Attempt to catch Pokémon with success based on base experience
- **Caching System**: Built-in cache to reduce API calls and improve performance

## Commands

- `help` - Display available commands
- `map` - Show next 20 location areas
- `mapb` - Show previous 20 location areas
- `explore <area-name>` - List Pokémon in a specific area
- `catch <pokemon-name>` - Attempt to catch a Pokémon
- `exit` - Exit the application

## Usage

```bash
go run .
```

Then use the interactive commands:
```
pokedex > help
pokedex > map
pokedex > explore canalave-city-area
pokedex > catch pikachu
```

## Architecture

### REPL

The pokedex REPL is waits for user input using `bufio.NewScanner` which blocks and waits for input, once the user types something and presses enter, the code continues and the input is available in the returned `bufio.Scanner`.

The `for` loop runs infinitely and only breaks when a certain command is entered. I am using the scanner's `.Scan()` and `.Text()` methods to get the user's input as a string.

The `getCommand` function returns a `map` of supported commands.

```go
map[string]cliCommand{
 "exit": {
  name:        "exit",
  description: "Exit the Pokedex",
  callback:    exitCommand,
 },
 "help": {
  name:        "help",
  description: "Displays a help message",
  callback:    helpCommand,
 },
 "map": {
  name: "map",
  description: "Displays names of 20 location areas",
  callback: mapCommand,
 },
}
```

This gives me a nice abstraction for managing all the many commands for the REPL. In the REPL I used the function to check for the existence of a command. If the command exists, its callback is called. If the command does not exist, an `"Unknown command"` message is printed and the user is prompted to put in another command.

I created a `struct` type that describes a command:

```go
type cliCommand struct {
 name string
 description string
 callback func(input []string, c *config) error
}
```

### Pokecache

Pokecache is a package responsible for all my caching logic.

I used a `Cache` struct to hold an `entry`, a `mutex`, and an `interval` field.

```go
type Cache struct {
 entry map[string]cacheEntry
 mu *sync.Mutex
 interval time.Duration
}
```

- The `entry` is a map of type `map[string]cacheEntry` that stores the cached data, where the URL is the key and the `cacheEntry` holds the raw data I am caching (`[]byte`) along with its creation timestamp (`time.Time`).
- Because maps are not thread-safe, I used a `mutex` to protect it across multiple goroutines.

### Functions

I exposed a `NewCache()` function that creates a new cache with a configurable interval (`time.Duration`). The interval is what determines how long an entry should remain in the cache.

### Methods

The `cache.Add()` methods adds a new entry to the cache. It takes a key (`string`) and a val (`[]byte`) as parameters. On adding an entry, I used the mutex to ensure thread safety.

The `cache.Get()` method gets an entry from the cache. It takes a key (`string`) and returns a `[]byte` and a `bool`. If an entry was found, it returns true, otherwise it returns false.

The `cache.reapLoop()` method is called each time a cache is created and runs in a separate goroutine. Each time an `interval` passes (the `time.Duration` passed to `NewCache`), it removes entries older than the interval. This makes sure that the cache doesn't grow too large over time. For example, if the interval is 10 seconds and an entry was added 15 seconds ago, that entry would be removed.
