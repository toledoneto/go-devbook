package routes

import (
	"api/src/controllers"
	"net/http"
)

var postRoutes = []Route{
	{
		URI:        "/posts",
		Method:     http.MethodPost,
		Function:   controllers.CreatePost,
		AuthNeeded: true,
	},
	{
		URI:        "/posts",
		Method:     http.MethodGet,
		Function:   controllers.SearchPosts,
		AuthNeeded: true,
	},
	{
		URI:        "/posts/{postId}",
		Method:     http.MethodGet,
		Function:   controllers.SearchPost,
		AuthNeeded: true,
	},
	{
		URI:        "/posts/{postId}",
		Method:     http.MethodPut,
		Function:   controllers.UpdatePost,
		AuthNeeded: true,
	},
	{
		URI:        "/posts/{postId}",
		Method:     http.MethodDelete,
		Function:   controllers.DeletePost,
		AuthNeeded: true,
	},
	{
		URI:        "/users/{userId}/posts",
		Method:     http.MethodGet,
		Function:   controllers.SearchPostsByUser,
		AuthNeeded: true,
	},
	{
		URI:        "/posts/{postId}/like",
		Method:     http.MethodPost,
		Function:   controllers.LikePost,
		AuthNeeded: true,
	},
	{
		URI:        "/posts/{postId}/unlike",
		Method:     http.MethodPost,
		Function:   controllers.UnlikePost,
		AuthNeeded: true,
	},
}
