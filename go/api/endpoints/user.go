package endpoints

import (
	"api/data"
	"api/db"
	"api/utils"
	"fmt"
	"lib/logger"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
)

//////////////////////////////////////////////////
// Structures and Variables

var emailRegex = regexp.MustCompile(`^(?:[^@\t\n ])+@(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z0-9][a-zA-Z0-9-]{0,61}[a-zA-Z0-9]*$`)

type RegisterUserStruct struct {
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Email     string  `json:"email"`
	Picture   *string `json:"picture,omitempty"`
	Password  *string `json:"password"`
}

/////////////////////////////////////////////
// Endpoints

//UserHandlers registers the user
func UserHandlers(router *mux.Router) error {
	router.HandleFunc("/register", RegisterUserHandler).Methods("POST")
	router.HandleFunc("/information", InformationUserHandler).Methods("POST")
	router.HandleFunc("/update", UpdateUserHandler).Methods("POST")
	router.HandleFunc("/remove", RemoveUserHandler).Methods("POST")
	return nil
}

/////////////////////////////////////////////
// Functions

func TempUserHandlerfunc(writer http.ResponseWriter, request *http.Request) {
	utils.Ok(writer, request)
}

// RegisterUserHandler registers a new user
func RegisterUserHandler(writer http.ResponseWriter, request *http.Request) {
	// TODO [KP]: Add default permissions to users once they register
	var registerUserStruct RegisterUserStruct

	err := utils.UnmarshalJSON(writer, request, &registerUserStruct)
	if err != nil {
		utils.BadRequest(writer, request, "invalid_request")
		return
	}

	if !emailRegex.MatchString(registerUserStruct.Email) {
		utils.BadRequest(writer, request, "invalid_email")
		return
	}

	access, err := db.Open()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	defer access.Close()

	da := data.NewUserDA(access)

	users, err := da.FindIdentifier(&data.User{Email: &registerUserStruct.Email})
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	user := users.FindHead()

	if user != nil {
		utils.BadRequest(writer, request, "user_already_exists")
		return
	}

	user = &data.User{
		Identifier: &registerUserStruct.Email,
		Email:      &registerUserStruct.Email,
		FirstName:  registerUserStruct.FirstName,
		LastName:   registerUserStruct.LastName,
		Picture:    registerUserStruct.Picture,
	}

	err = da.StoreIdentifier(user)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	clientID := "local." + *user.Identifier
	err = da.StoreCredential(clientID, registerUserStruct.Password, *user.Identifier)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	logger.Access.Printf("%v registered\n", user.Identifier)

	utils.Ok(writer, request)
}

func InformationUserHandler(writer http.ResponseWriter, request *http.Request) {
	var user data.User

	err := utils.UnmarshalJSON(writer, request, &user)
	if err != nil {
		fmt.Println(err)
		utils.BadRequest(writer, request, "invalid_request")
		return
	}

	access, err := db.Open()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}
	defer access.Close()

	da := data.NewUserDA(access)

	users, err := da.FindIdentifier(&user)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	logger.Access.Printf("%v user information requested\n", user.Id)

	utils.JSONResponse(writer, request, users)
}

func UpdateUserHandler(writer http.ResponseWriter, request *http.Request) {
	logger.Info.Println("user update requested")
	utils.Ok(writer, request)
}

func RemoveUserHandler(writer http.ResponseWriter, request *http.Request) {
	logger.Info.Println("user remove requested")
	utils.Ok(writer, request)
}
