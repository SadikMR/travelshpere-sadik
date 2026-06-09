# TravelSphere

TravelSphere is a Go web application built with Beego. It helps users explore countries, search nearby attractions, and manage a personal travel wishlist with a simple web interface and JSON APIs.

## Features

- Country search and detail pages
- Attraction lookup by latitude/longitude
- User wishlist creation, update, and deletion
- Dashboard summary for authenticated users
- Server-side rendering with Beego templates

## Technology

- Go 1.26
- Beego v2
- In-memory session and wishlist storage
- REST API and server-rendered HTML pages

## Prerequisites

- Go 1.26 installed
- Docker and Docker Compose (optional)
- Internet access for external country and attraction APIs

## Quick setup

```bash
git clone https://github.com/SadikMR/travelshpere-sadik.git
cd travelshpere-sadik
cp conf/app.conf.example conf/app.conf
```

Then update `conf/app.conf`:

```ini
opentripApi = YOUR_OPENTRIPMAP_API_KEY
restcountriesBaseURL = https://restcountries.com/v3.1
opentripBaseURL = https://api.opentripmap.com/0.1/en/places
```

If `conf/app.conf` does not exist yet, create it with:

```bash
touch conf/app.conf
```

## Run locally

```bash
go mod download
go run main.go
```

Open the application at:

```text
http://localhost:8080
```

## Run with Docker

```bash
docker compose up --build
```

The app will be available on `http://localhost:8080`.

## Application routes

### User-facing pages

- `/` — home page
- `/login` — login page
- `/logout` — logout action
- `/countries` — browse countries
- `/countries/:slug` — country detail page
- `/wishlist` — authenticated wishlist page
- `/wishlist/rows` — wishlist row fragment
- `/dashboard` — authenticated dashboard

### Rendering and AJAX

The app uses Beego server-side templates for SSR pages. Country, wishlist, and dashboard pages are rendered on the server, with client-side JavaScript calling JSON API endpoints for dynamic data updates.

- `GET /countries` and `GET /countries/:slug` render server-side pages
- Search and attraction lookups use API calls from the page
- Wishlist actions are performed through AJAX to `/api/wishlist`

### API endpoints

#### Country API

- `GET /api/countries`
- `GET /api/countries/search?q=<term>`
- `GET /api/countries/:slug`

Examples:

```bash
curl http://localhost:8080/api/countries
curl "http://localhost:8080/api/countries/search?q=bang"
curl http://localhost:8080/api/countries/bangladesh
```

#### Attraction API

- `GET /api/attractions?lat=<lat>&lon=<lon>`

Example:

```bash
curl "http://localhost:8080/api/attractions?lat=23.8&lon=90.4"
```

#### Wishlist API

These endpoints require an authenticated session.

- `GET /api/wishlist`
- `POST /api/wishlist`
- `PUT /api/wishlist/:id`
- `DELETE /api/wishlist/:id`

Example create request:

```bash
curl -X POST http://localhost:8080/api/wishlist \
  -H "Content-Type: application/json" \
  -d '{"country_name":"Bangladesh","note":"Visit next year"}'
```

## Configuration

Use `conf/app.conf` to configure:

- `httpport` — server port
- `sessionon`, `sessionname`, `sessionprovider` — session settings
- `opentripApi` — OpenTripMap API key
- `restcountriesBaseURL` — country data API base URL
- `opentripBaseURL` — attraction API base URL

## Testing

Run all tests with:

```bash
go test ./...
```

Run coverage reporting with:

```bash
go test ./... -coverprofile=coverage.out
go tool cover -func=coverage.out
```

Current total coverage: **89.1%**

## Notes

- This project uses in-memory wishlist storage, so data is not persisted across restarts.
- Keep `opentripApi` set before running attraction-related features.
