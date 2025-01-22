package controllers

import "github.com/google/uuid"

// UsersController
// @Tag(Users) Users
// @Route(/users)
// @Description The Users API
type UsersController struct {
	// gleece.GleeceController // Embedding the GleeceController to inherit its methods
}

// @Description User's domicile
type Domicile struct {
	// @Description The address
	Address string `json:"address" validate:"required"`
	// @Description The number of the house (must be at least 1)
	HouseNumber int `json:"houseNumber" validate:"gte=1"`
}

// @Description() Create a new user
// @Method(POST) This text is not part of the OpenAPI spec
// @Route(/user/{user_name}) Same here
// @Query(email, { validate: "required,email" }) The user's email
// @Path(name, { name: "user_name" }) The user's name
// @Body(domicile, { validate: "required" }) The user's domicile info
// @Header(origin, { name: "x-origin" }) The request origin
// @Header(trace) The trace info
// @Response(200) The ID of the newly created user
// @ErrorResponse(500) The error when process failed
// @Security(ApiKeyAuth, { scopes: ["read:users", "write:users"] })
func (ec *UsersController) CreateNewUser(email string, name string, domicile Domicile, origin string, trace string) (string, error) {
	userId := uuid.New()
	return userId.String(), nil
}

// @Description() Get user's domicile
// @Method(GET)
// @Route(/{id}/domicile)
// @Path(id, { name: "id" })
// @Response(200) The user's domicile
// @ErrorResponse(404) The user not found
// @ErrorResponse(500) The error when process failed
// @Security(ApiKeyAuth, { scopes: ["read:users"] })
func (ec *UsersController) GetUserDomicile(id string) (Domicile, error) {
	return Domicile{
		Address:     "Jl. Jend. Sudirman",
		HouseNumber: 1,
	}, nil
}
