package main

import (
        "net/http"
        "strconv"

        "github.com/labstack/echo/v4"
        "github.com/labstack/echo/v4/middleware"
)

type (
      user struct {
              ID int `json:"id"`
              Name string `json:"name"`
      }
)

var (
      users = map[int]*user{}
      seq = 1
)

//---------
//Handlers
//---------

func createUser(c echo.Context) error {
        u := &user{
              ID: seq,
        }
        if err := c.Bind(u); err != nil {
                return err 
        }
        users[u.ID] = u
        seq++
        return c.JSON(http.StatusCreated, u)
}

func getUser(c echo.Context) error {
        id, _ := strconv.Atoi(c.Param("id"))
        return c.JSON(http.StatusOK, users[id])
}

func updatUser(c echo.Context) error{
        u := new(user)
        if err := c.Bind(u); err != nil {
                return err
        }
        id, _ := strconv.Atoi(c.Param("id"))
        users[id].Name = u.Name
        return c.JSON(http.StatusOK, users[id])
}

func deleteUser(c echo.Context) error {
        id, _:= strconv.Atoi(c.Param("id"))
        delete(users, id)
        return c.NoContent(http.StatusNoContent)
}

func main() {
        e := echo.New()

        // middleware
        e.Use(middleware.Logger())
        e.Use(middleware.Recover())

        // Routes
        e.POST("/users", createUser)
        e.GET("/users/:id", getUser)
        e.PUT("/users/:id", updatUser)
        e.DELETE("/users/:id", deleteUser)
        

        // Start the Server
        e.Logger.Fatal(e.Start(":1323"))
        
        // helloHandler := func (w http.ResponseWriter, req *http.Request)  {
        //         io.WriteString(w, "Hello, world!\n")
        // }
        //
      // http.HandleFunc("/hello", helloHandler)
      // log.Println("Listening for requests at http://localhost:8000/hello")
      //     log.Fatal(http.ListenAndServe(":8000", nil))
      //
      }
