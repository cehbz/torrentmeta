# torrentmeta

The `torrentmeta` package provides utilities for working with torrent metadata using the `github.com/anacrolix/torrent/metainfo` package. This package enables you to parse torrent files and extract useful information such as the torrent name and tracker domain.

## Installation

To use the `torrentmeta` package, you must first install the `anacrolix/torrent` library:

```bash
go get github.com/anacrolix/torrent
```

## Usage

### Importing the Package

To use the `torrentmeta` package, import it in your Go code:

```go
import "github.com/cehbz/torrentmeta"
```

### Parsing Torrent Metadata

To parse torrent metadata, use the `NewTorrentMetadata` function. This function takes a byte slice containing the torrent data and returns a `Metadata` instance:

```go
torrentData := []byte{...} // Your torrent file content
metadata, err := torrentmeta.NewTorrentMetadata(torrentData)
if err != nil {
    log.Fatalf("Failed to parse torrent data: %v", err)
}
```

### Accessing Torrent Metadata Fields

Once you have a `Metadata` instance, you can access various fields of the torrent:

```go
// Get the name of the torrent
name := metadata.Name()
fmt.Println("Torrent Name:", name)

// Get the tracker domain (e.g., "example.com")
tracker, err := metadata.Tracker()
if err != nil {
    log.Fatalf("Failed to extract tracker domain: %v", err)
}
fmt.Println("Tracker Domain:", tracker)
```

## API Reference

### `type Metadata struct`

The `Metadata` struct holds the torrent metadata and provides methods to access specific fields:

- **`Name() string`**: Returns the name of the torrent.
- **`Tracker() (string, error)`**: Returns the tracker domain (second-level and top-level domain) extracted from the announce URL.

### `func NewTorrentMetadata(torrentData []byte) (*Metadata, error)`

Parses the torrent data and returns a `Metadata` instance.

**Parameters:**

- `torrentData []byte`: The raw content of the torrent file.

**Returns:**

- `*Metadata`: A pointer to the `Metadata` instance.
- `error`: An error if the torrent data could not be parsed.

### Example

```go
package main

import (
    "fmt"
    "log"

    "github.com/yourusername/torrentmeta"
)

func main() {
    torrentData := []byte{...} // Load your torrent file data here

    metadata, err := torrentmeta.NewTorrentMetadata(torrentData)
    if err != nil {
        log.Fatalf("Failed to parse torrent metadata: %v", err)
    }

    fmt.Println("Torrent Name:", metadata.Name())

    tracker, err := metadata.Tracker()
    if err != nil {
        log.Fatalf("Failed to extract tracker domain: %v", err)
    }
    fmt.Println("Tracker Domain:", tracker)
}
```

## License

This package is released under the MIT License. See the `LICENSE` file for more information.
