# Go Client for Nasa Open APIs

## Description

The Go Client for Nasa Open APIs is a Go Client for the following Nasa Open APIs:

- [x] **APOD**: Astronomy Picture of the Day
- [ ] **Asteroids NeoWs**: Near Earth Object Web Service
- [ ] **DONKI**: Space Weather Database of Notifications, Knowledge, Information
- [ ] **Earth**: Unlock the significant public investment in earth observation data
- [ ] **EONET**: The Earth Observatory Natural Event Tracker
- [ ] **EPIC**: Earth Polychromatic Imaging Camera
- [ ] **Exoplanet**: Programmatic access to NASA's Exoplanet Archive database
- [ ] **GeneLab**: Programmatic interface for GeneLab's public data repository website
- [ ] **Insight**: Mars Weather Service API
- [ ] **Mars Rover Photos**: Image data gathered by NASA's Curiosity, Opportunity, and Spirit rovers on Mars
- [ ] **NASA Image and Video Library**: API to access the NASA Image and Video Library site at images.nasa.gov
- [ ] **TechTransfer**: Patents, Software, and Tech Transfer Reports
- [ ] **Satallite Situation Center**: System to cast geocentric spacecraft location information into a framework of (empirical) geophysical regions
- [ ] **SSD/CNEOS**: Solar System Dynamics and Center for Near-Earth Object Studies
- [ ] **Techport**: API to make NASA technology project data available in a machine-readable format
- [ ] **TLE API**: Two line element data for earth-orbiting objects at a given point in time
- [ ] **Vesta/Moon/Mars Trek WMTS**: A Web Map Tile Service for the Vesta, Moon, and Mars Trek imagery projects

More information on Nasa Open APIs can be found here https://api.nasa.gov/

## Nasa API Key

Retrieve your personal API Key from https://api.nasa.gov/. If no API Key is given `"DEMO_KEY"` will be used as default.

## Requirements

- Go 1.17 or later

## Installation
```cmd
go get github.com/ernestohdz/go-nasa
```

## Usage

Ways of initializing Nasa Client

```go
package main

import (
    "net/http"
    "github.com/ernestohdz/go-nasa"
)

func main() {
    // Default Client with DEMO_KEY
    client := nasa.NewClient()

    // Create Client with your personal API Key
    client := nasa.NewClient(nasa.WithKey("API_KEY"))

    // Create Client with personal API Key and configured HTTP Client
    c := &http.Client{
		Timeout: 2 * time.Second,
	}
    client := nasa.NewClient(nasa.WithKey("API_KEY"), nasa.WithClient(c))
}

```