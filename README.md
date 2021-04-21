# Api-go
gorilla/mux: go get github.com/gorilla/mux

go run .\main.go

ROUTES

ALL GET : http://localhost:3000/ticket

SPECIFIC GET : http://localhost:3000/ticket/{ID}

DELETE : http://localhost:3000/ticket/{ID}

POST : http://localhost:3000/ticket
  
  body: {
  	"User":    "BlackJeldo",
	  "Fecha_creacion": "20/10/2021",
	  "Fecha_actualizacion": "25/10/2021",
	  "Status": true
  }
  
PUT : http://localhost:3000/ticket/{ID}
  
  body: {
      "User":    "Jeldo05",
      "Fecha_creacion": "20/10/2021",
      "Fecha_actualizacion": "25/10/2021",
      "Status": true
    }

