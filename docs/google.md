# Google Package

The `google` package provides functions for interacting with Google services, specifically the Google Maps API.

## Functions

### `NewMapClient`

```go
func NewMapClient(apiKey string) (*MapClient, error)
```

Creates a new `MapClient` using the provided API key.

**Parameters:**

- `apiKey`: The API key for accessing Google Maps services.

**Returns:**

- `*MapClient`: A new instance of `MapClient`.
- `error`: An error if there is an issue creating the client.

### `FindCoordinates`

```go
func (c *MapClient) FindCoordinates(context context.Context, address string) (lat float64, lng float64, err error)
```

Finds the coordinates (latitude and longitude) for a given address.

**Parameters:**

- `context`: The context for the request.
- `address`: The address to find the coordinates for.

**Returns:**

- `lat`: The latitude of the address.
- `lng`: The longitude of the address.
- `error`: An error if there is an issue with the request.

**Usage Example:**

```go
ctx := context.Background()
client, err := maps.NewMapClient("YOUR_API_KEY")
if err != nil {
    log.Fatal(err)
}

lat, lng, err := client.FindCoordinates(ctx, "1600 Amphitheatre Parkway, Mountain View, CA")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Latitude: %f, Longitude: %f\n", lat, lng)
```

## Types

### `MapClient`

```go
type MapClient struct {
    client *maps.Client
}
```

A client for interacting with the Google Maps API.
