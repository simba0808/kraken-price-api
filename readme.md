- Kraken Price API

Prerequisites
  Go 1.22.5
  Docker and Docker Compose: For containerized deployment.
  Internet Connection: Required for API calls to Kraken.

How to run the application
  Locally
    1. Clone the repository:
    git clone https://github.com/simba0808/kraken-price-api.git
    cd kraken-price-api

    2. Build and run the application:
    go build -o main ./cmd/server/main.go
    ./main


  Using Docker
    1. Clone the repository:
      git clone https://github.com/yourusername/bitcoin-ltp-service.git
      cd bitcoin-ltp-service  
    2. Build and run using Docker Compose:
      docker-compose up --build
  
The service will be available at http://localhost:8080.

API Usage
  GET /api/v1/ltp: Retrieves the LTP for Bitcoin in various currency pairs.
  Response: {"ltp":[{"pair":"BTC/USD","price":86263.4},{"pair":"BTC/CHF","price":77447.6},{"pair":"BTC/EUR","price":82335.3}]}