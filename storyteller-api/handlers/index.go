package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const fancyIndex = `
<style>
    div {
    	color: white;
    	background: linear-gradient(to right, #ff6b6b, #6b47ff);
    	width: 83.5vh;
    	height: 80vh;
    	text-align: center;
    }
  </style>
  <div>
    <h1>
    ⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠟⠁⣠⣤⣤⣤⣤⣤⣀⣀⠉⠻⣿⣿⣿⣿⣿⣿⣿
    ⣿⣿⣿⣿⣿⣿⣿⣿⣿⠏⠄⣾⣬⣽⣿⣿⣿⣿⡿⢿⣿⣆⠈⢻⣿⣿⣿⣿⣿
    ⣿⣿⣿⣿⣿⣿⣿⣿⡿⢀⣞⡉⢩⣙⣿⡿⠉⠄⣠⣤⠤⠉⠄⠄⢿⣿⣿⣿⣿
    ⣿⣿⣿⣿⣿⣿⣿⣿⠁⣼⣿⣿⣯⣿⣿⠁⢰⣾⣦⡤⠄⢀⣶⡀⠸⣿⣿⣿⣿
    ⣿⣿⣿⣿⣿⣿⣿⡏⢀⣿⣿⣿⣿⣿⠟⠁⠄⠈⢿⣿⣿⣿⣿⡇⠄⣿⣿⣿⣿
    ⣿⣿⣿⣿⣿⣿⣿⠇⢸⣿⣿⡟⠛⠃⡠⠄⠄⠄⠈⣿⣿⣿⣿⡇⠄⣿⣿⣿⣿
    ⣿⣿⣿⣿⣿⣿⣿⠄⣿⣿⣿⣶⣾⣿⣿⣿⣤⣤⣄⣘⣿⣿⠁⡀⠄⢻⣿⣿⣿
    ⣿⣿⣿⣿⣿⣿⡇⢰⣿⣿⣿⣿⣿⣏⣉⣽⣿⣿⣿⣿⣿⣿⣿⣿⠄⢸⣿⣿⣿
    ⣿⣿⣿⣿⣿⣿⠄⣼⣿⣁⣸⣿⣿⣿⣿⣿⡿⠟⠉⠙⠋⠹⠟⠁⠄⢸⣿⣿⣿
    ⣿⣿⣿⣿⣿⡏⢀⣿⣿⣿⣿⠋⢠⣤⣤⣤⣤⠈⢿⣿⣷⣦⣄⠄⠄⢸⣿⣿⣿
    ⣿⠋⣀⣤⣄⣠⣼⣿⣿⣿⣿⡀⢹⣿⣿⣿⣿⠄⢸⣿⣿⣿⣿⣿⠄⢸⣿⣿⣿
    ⣧⠄⢿⣿⣿⣿⣿⣿⣿⣿⡿⠃⢸⠿⠛⠉⣁⣠⣿⣿⣿⣿⣿⣿⠄⣼⣿⣿⣿
    ⣿⣷⣄⣉⠉⠉⢉⣉⣉⣁⣤⣾⡏⠄⣾⣿⣿⣿⣿⣿⣿⣿⣿⡟⠄⣿⣿⣿⣿
    ⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣤⣈⠙⠛⠛⠟⠛⠛⢉⣁⣤⣾⣿⣿⣿⣿
    </h1>
    <h2>
      Welcome to Index!
    </h2>
  </div>
`

func HandleGetIndex(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	t := template.New("Index")
	t, err := t.Parse(fancyIndex)
	if err != nil {
		fmt.Fprint(w, err)
	}
	t.Execute(w, nil)
}
