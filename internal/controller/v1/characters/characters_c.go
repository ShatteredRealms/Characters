package characters

import (
    "github.com/ShatteredRealms/Accounts/pkg/service"
    "github.com/gin-gonic/gin"
)

type CharacterController interface {
    List(ctx *gin.Context)
    Get(ctx *gin.Context)
    Update(ctx *gin.Context)
    Delete(ctx *gin.Context)
    Create(ctx *gin.Context)
}

type characterController struct {
    jwtService service.JWTService
}

func NewAuthController(jwt service.JWTService) CharacterController {
    return characterController{
        jwtService: jwt,
    }
}

// List godoc
// @Summary Handles when a request to get all characters for a user.
// @Schemes object
// @Description Gets all of the characters for a given user. Allowed if requesting user has permission
// option.LIST_CHARACTERS and is accessing itself. Allowed if a user has permission option.LIST_CHARACTER_ANY.
// @Tags Characters
// @Accept json
// @Produce json
// @Success 200 {object} model.ResponseModel All characters for the user in the data field
// @Failure 401 {object} model.ResponseModel Error message
// @Router /v1/{user_id} [post]
func (c characterController) List(ctx *gin.Context) {

}

// Get godoc
// @Summary Gets information for a specific character
// @Schemes object
// @Description Gets general information about a specific character. Allowed if requesting user has permission
// option.GET_CHARACTER and is accessing itself. Allowed if a user has permission option.GET_CHARACTER_ANY.
// @Tags Characters
// @Accept json
// @Produce json
// @Success 200 {object} model.ResponseModel General information about the character
// @Failure 200 {object} model.ResponseModel Error message
// @Router /v1/{user_id}/{character_id} [get]
func (c characterController) Get(ctx *gin.Context) {

}

// Create godoc
// @Summary Creates a new character for a given user
// @Schemes object
// @Description Creates a new user for a given user. Allowed if requesting user has permission option.CREATE_CHARACTER
// and has enough space to create another.
// @Tags Characters
// @Accept json
// @Produce json
// @Success 200 {object} model.ResponseModel General information about the new character
// @Failure 200 {object} model.ResponseModel Error message
// @Router /v1/{user_id} [post]
func (c characterController) Create(ctx *gin.Context) {

}

// Update godoc
// @Summary Updates general information about the character
// @Schemes object
// @Description Updates general information about the character such as Name and Owner. Allowed if a user has
// @Tags Characters, Admin
// @Accept json
// @Produce json
// @Success 200 {object} model.Response General information about the updated character
// @Failure 200 {object} model.Response Error message
// @Router /v1/{user_id}/{character_id} [put]
func (c characterController) Update(ctx *gin.Context) {

}

// Delete godoc
// @Summary Deletes a character
// @Schemes object
// @Description Deletes a character
// @Tags Characters, Admin
// @Accept json
// @Produce json
// @Success 200 {object} model.Response Success message with no data
// @Failure 200 {object} model.Response Error message
// @Router /v1/{user_id}/{character_id} [delete]
func (c characterController) Delete(ctx *gin.Context) {

}

// godoc
// @Summary
// @Schemes
// @Description
// @Tags
// @Accept
// @Produce
// @Success
// @Failure
// @Router
