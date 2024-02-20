# Fablescope Backend

## Getting started
1. Make sure that Go is installed on your machine. 
You can download Go from the official [website](https://go.dev/doc/install)
2. Clone the repository:
```sh
git clone https://github.com/rep-co/fablescope-backend.git
```
3. Open cloned repository directory:
```sh
cd fablescope-backend/
```
4. Build and run server:
```sh
make run -C storyteller-api
```
### Additional:
You always can use this to just build the server:
```sh
make build -C storyteller-api
```
## State of dev
Available endpoints:
- GET "/" index
- GET "/form/category" get all categories with tags
- POST "/generate/story" get AI generated story
