## Project Structure

```bash
├── client/       # React frontend
└── server/       # Go backend API
```

## Features

### Backend (Go)
- RESTful API built with `fiber` package.
- Parses raw JSON data at startup using a custom `Parse()` function to extract only relevant information for later usage.
- Modular routing for different sports: `/cricket` and `/volleyball` endpoints.
- Dynamic odds fetching for different sports (volleyball, cricket).
- Predication evaluation logic implemented via `/evaluate-*` endpoints to assess user predictions.
- ⚠️ Due to JSON structure inconsistencies between cricket and volleyball, each sport is handled separately with tailored route handlers and logic.

### Frontend (React)
- Responsive UI for browsing fixtures and betting options.
- Dynamic rendering of sports and bet types.
- Uses `react-router-dom` for navigation.
- Clean component structure for reusability (`BetGroup`, `BetOption`, etc.).


## Getting Started

### Backend (Go)

#### Prerequisites
- Go 1.20+
- `go mod` initialized

#### Run the Server
```bash
cd server
go run main.go
```
The server runs on http://localhost:7777

### Frontend(React)

#### Prerequisites
- Node.js (v18+)
- `npm`

### Install and Run
```bash
cd client
npm install
npm run dev
```

## API Endpoints
### Volleyball

| Method | Endpoint                                  | Description                    |
|--------|-------------------------------------------|--------------------------------|
| GET    | `/volleyball/fixtures`                    | Get list of volleyball fixtures|
| GET    | `/volleyball/1x2/:id`                     | Get 1X2 betting odds           |
| GET    | `/volleyball/over-under/:id`              | Get over/under odds            |
| GET    | `/volleyball/correct-score/:id`           | Get correct score bets         |
| POST   | `/volleyball/evaluate-1x2`                | Evaluate 1X2 user prediction   |
| POST   | `/volleyball/evaluate-under-above`        | Evaluate over/under prediction |
| POST   | `/volleyball/evaluate-correct-score`      | Evaluate correct score guess   |

### Cricket

| Method | Endpoint                                  | Description                      |
|--------|-------------------------------------------|----------------------------------|
| GET    | `/cricket/fixtures`                       | Get list of cricket fixtures     |
| GET    | `/cricket/1x2/:id`                        | Get 1X2 betting odds             |
| GET    | `/cricket/over-under/:id`                 | Get over/under odds              |
| POST   | `/cricket/evaluate-1x2`                   | Evaluate 1X2 user prediction     |
| POST   | `/cricket/evaluate-under-above`           | Evaluate over/under prediction   |

> ⚠️ **Note on Evaluation Endpoints**
>
> The `/evaluate-*` endpoints for both volleyball and cricket are fully implemented on the **backend**, allowing score evaluation based on user predictions.
> However, **frontend integration for these endpoints is not yet implemented**. You can still test them using tools like Postman or curl.

