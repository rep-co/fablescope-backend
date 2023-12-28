# Fablescope Backend

## Getting started
1. Make sure that Go is installed on your machine
You can download Go from the official [website](https://go.dev/doc/install)
2. Clone the repository
```sh
git clone https://github.com/rep-co/fablescope-backend.git
```
3. Open cloned repository directory
4. Build server:
```sh
make build -C storyteller-api
```
5. Run server:
```sh
make run -C storyteller-api
```
## State of dev
Available endpoints:
- "/" index
- "/form/category" all categories with tags

