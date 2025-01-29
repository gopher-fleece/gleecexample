package controllers

import (
	"github.com/google/uuid"
	"github.com/gopher-fleece/gleece/external"
)

// UsersController
// @Tag(Users) Users
// @Route(/users)
// @Description The Users API
type UsersController struct {
	external.GleeceController // Embedding the GleeceController to inherit its methods
}

// @Description User's domicile
type Domicile struct {
	// @Description The address
	Address string `json:"address" validate:"required"`
	// @Description The number of the house (must be at least 1)
	HouseNumber int `json:"houseNumber" validate:"gte=1"`
}

// @Description Create a new user
// @Method(POST) This text is not part of the OpenAPI spec
// @Route(/user/{user_name}/{user_id}/{serial}) Same here
// @Query(email, { validate: "required,email" }) The user's email
// @Path(id, { name: "user_id", validate:"gt=1" }) The user's ID
// @Path(serial, { validate:"gte=10" }) The user's serial number
// @Path(name, { name: "user_name" }) The user's name
// @Body(domicile) The user's domicile
// @Header(origin, { name: "x-origin", validate: "validate_starts_with_letter" }) The request origin
// @Header(option, { name: "x-option" }) The request option header
// @Header(trace) The trace info
// @Response(200) The ID of the newly created user
// @ErrorResponse(500) The error when process failed
// @Security(securitySchemaName, { scopes: ["read:users", "write:users"] })
func (ec *UsersController) CreateNewUser(id int, serial int, email string, name string, origin string, option *string, trace string, domicile Domicile) (string, error) {
	ec.SetHeader("x-powered-by", "Gleece")
	userId := uuid.New()
	return userId.String(), nil
}

// @Description Get user's domicile
// @Method(GET)
// @Route(/domicile/{id})
// @Path(id, { name: "id" }) The id of the user to get the domicile to
// @Response(200) The user's domicile
// @ErrorResponse(404) The user not found
// @ErrorResponse(500) The error when process failed
// @Security(securitySchemaName, { scopes: ["read:users" ] }) Consumer should pass this security schema
// @Security(securitySchemaName, { scopes: ["read:all_data"] }) -OR- that one
func (ec *UsersController) GetUserDomicile(id string) (Domicile, error) {
	return Domicile{
		Address:     "221B Baker Street, London",
		HouseNumber: 221,
	}, nil
}
