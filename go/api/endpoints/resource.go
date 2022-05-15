package endpoints

import (
	"api/data"
	"api/db"
	"api/utils"
	"fmt"
	"lib/logger"
	"net/http"

	"github.com/gorilla/mux"
)

//////////////////////////////////////////////////
// Structures and Variables

type CreateResourceStruct struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Capacity    *int    `json:"capacity,omitempty"`
	Picture     *string `json:"picture,omitempty"`
}

/////////////////////////////////////////////
// Endpoints

//ResourceHandlers
func ResourceHandlers(router *mux.Router) error {
	router.HandleFunc("/create", TempResourceHandler).Methods("POST")
	router.HandleFunc("/remove", TempResourceHandler).Methods("POST")
	router.HandleFunc("/information", TempResourceHandler).Methods("POST")

	router.HandleFunc("/room/create", CreateRoomHandler).Methods("POST")
	router.HandleFunc("/room/remove", DeleteRoomHandler).Methods("POST")
	router.HandleFunc("/room/information", InformationRoomsHandler).Methods("POST")

	router.HandleFunc("/room/association/create", TempResourceHandler).Methods("POST")
	router.HandleFunc("/room/association/remove", TempResourceHandler).Methods("POST")
	router.HandleFunc("/room/association/information", TempResourceHandler).Methods("POST")

	router.HandleFunc("/building/create", CreateBuildingHandler).Methods("POST")
	router.HandleFunc("/building/remove", DeleteBuildingHandler).Methods("POST")
	router.HandleFunc("/building/information", InformationBuildingsHandler).Methods("POST")

	return nil
}

/////////////////////////////////////////////
// Functions

////////////////
// Building

// CreateBuildingHandler creates or updates a building
func CreateBuildingHandler(writer http.ResponseWriter, request *http.Request) {
	var building data.Building

	err := utils.UnmarshalJSON(writer, request, &building)
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

	da := data.NewResourceDA(access)

	// TODO [KP]: Do more checks like if there exists a building already etc

	err = da.StoreBuildingResource(&building)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	logger.Access.Printf("%v created\n", building.Id)

	utils.Ok(writer, request)
}

// InformationBuildingsHandler gets Buildings
func InformationBuildingsHandler(writer http.ResponseWriter, request *http.Request) {
	var building data.Building

	err := utils.UnmarshalJSON(writer, request, &building)
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

	da := data.NewResourceDA(access)

	buildings, err := da.FindBuildingResource(&building)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	logger.Access.Printf("%v building information requested\n", building.Id)

	utils.JSONResponse(writer, request, buildings)
}

// DeleteBuildingHandler removes a Building
func DeleteBuildingHandler(writer http.ResponseWriter, request *http.Request) {
	var building data.Building

	err := utils.UnmarshalJSON(writer, request, &building)
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

	da := data.NewResourceDA(access)

	buildingRemoved, err := da.DeleteBuildingResource(&building)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	logger.Access.Printf("%v building removed\n", building.Id)

	utils.JSONResponse(writer, request, buildingRemoved)
}

////////////////
// Room

// CreateRoomHandler creates or updates a Room
func CreateRoomHandler(writer http.ResponseWriter, request *http.Request) {
	var room data.Room

	err := utils.UnmarshalJSON(writer, request, &room)
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

	da := data.NewResourceDA(access)

	// TODO [KP]: Do more checks like if there exists a Room already etc

	err = da.StoreRoomResource(&room)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	logger.Access.Printf("%v created\n", room.Id)

	utils.Ok(writer, request)
}

// InformationRoomsHandler gets Rooms
func InformationRoomsHandler(writer http.ResponseWriter, request *http.Request) {
	var room data.Room

	err := utils.UnmarshalJSON(writer, request, &room)
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

	da := data.NewResourceDA(access)

	rooms, err := da.FindRoomResource(&room)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	logger.Access.Printf("%v Room information requested\n", room.Id)

	utils.JSONResponse(writer, request, rooms)
}

// DeleteRoomHandler removes a Room
func DeleteRoomHandler(writer http.ResponseWriter, request *http.Request) {
	var room data.Room

	err := utils.UnmarshalJSON(writer, request, &room)
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

	da := data.NewResourceDA(access)

	roomRemoved, err := da.DeleteRoomResource(&room)
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	err = access.Commit()
	if err != nil {
		utils.InternalServerError(writer, request, err)
		return
	}

	logger.Access.Printf("%v Room removed\n", room.Id)

	utils.JSONResponse(writer, request, roomRemoved)
}

// DeleteBookingHandler rewrites fields for booking where applicable
func TempResourceHandler(writer http.ResponseWriter, request *http.Request) {
	utils.Ok(writer, request)
}
