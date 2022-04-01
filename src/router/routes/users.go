package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:        "/users",
		Method:     http.MethodPost,
		Function:   controllers.CreateUser,
		AuthNeeded: false,
	},
	{
		URI:        "/users",
		Method:     http.MethodGet,
		Function:   controllers.SearchUsers,
		AuthNeeded: true,
	},
	{
		URI:        "/users/{userId}",
		Method:     http.MethodGet,
		Function:   controllers.SearchUser,
		AuthNeeded: true,
	},
	{
		URI:        "/users/{userId}",
		Method:     http.MethodPut,
		Function:   controllers.UpdateUser,
		AuthNeeded: true,
	},
	{
		URI:        "/users/{userId}",
		Method:     http.MethodDelete,
		Function:   controllers.DeleteUser,
		AuthNeeded: true,
	},
	{
		URI:        "/users/{userId}/follow",
		Method:     http.MethodPost,
		Function:   controllers.FollowUser,
		AuthNeeded: true,
	},
	{
		URI:        "/users/{userId}/unfollow",
		Method:     http.MethodPost,
		Function:   controllers.UnfollowUser,
		AuthNeeded: true,
	},
	{
		URI:        "/users/{userId}/followers",
		Method:     http.MethodGet,
		Function:   controllers.SearchFollowers,
		AuthNeeded: true,
	},
	{
		URI:        "/users/{userId}/following",
		Method:     http.MethodGet,
		Function:   controllers.SearchFollowing,
		AuthNeeded: true,
	},
	{
		URI:        "/users/{userId}/update-password",
		Method:     http.MethodPost,
		Function:   controllers.UpdatePassword,
		AuthNeeded: true,
	},
}
