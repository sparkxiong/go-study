package gin

import "net/http"
import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	//GET
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Geektutu")
	})
	r.GET("/user/:name/*test", func(c *gin.Context) {
		c.String(http.StatusOK, "hello, %s", c.Param("name"))
	})
	r.GET("/user", func(c *gin.Context) {
		name := c.Query("name")
		age := c.Query("age")
		c.String(http.StatusOK, "hello, %s, %s", name, age)
	})

	//POST
	r.POST("/form", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.DefaultPostForm("password", "000000") // 可设置默认值

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	r.POST("/post/dddd", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		c.JSON(http.StatusOK, gin.H{
			"ids":   ids,
			"names": names,
		})
	})

	// GET 和 POST 混合
	r.POST("/posts", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		username := c.PostForm("username")
		password := c.DefaultPostForm("username", "000000") // 可设置默认值

		c.JSON(http.StatusOK, gin.H{
			"id":       id,
			"page":     page,
			"username": username,
			"password": password,
		})
	})

	r.Run("0.0.0.0:8888") // listen and serve on 0.0.0.0:8080
}
