package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	
)

func main() {

	// fmt.Println(todoes)
	router := gin.Default()
	router.GET("/todoes", getTodoes)
	router.GET("/todoes/:id", getTodo)
	router.POST("/addTodo",addTodo)
	router.DELETE("delet/:id",deletTodo)
	router.Run("localhost:9090")
}

type Todo struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Descripiton string `json:"descripiton"`
	Compleat    bool   `json:"compleat"`
}

var todoes = []Todo{
	{Id: 1, Title: "go to work", Descripiton: "", Compleat: false},
	{Id: 2, Title: "play game", Descripiton: "plan with your frind  unio", Compleat: false},
	{Id: 3, Title: "eate your fotor", Descripiton: "ater adan eat your fotor", Compleat: false},
}

func getTodoes(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todoes)
}
func addTodo(context *gin.Context){
	var newTodo Todo
	if err:=context.BindJSON(&newTodo);err!=nil{
		fmt.Println(err)
		return
	}
	todoes=append(todoes, newTodo)
	context.IndentedJSON(http.StatusCreated,todoes)

}
func gettodoById(id int )(*Todo ,error){
for i,t := range todoes{
	fmt.Print(i,t)
	if t.Id==id{

		return &todoes[i],nil
	}
}
return nil,errors.New("msg:not found")

}
func getTodo(context *gin.Context){
strID:=context.Param("id")
id ,err:=strconv.Atoi(strID)
if err!=nil{
	context.JSON(400,"Invaled Id!")
	return
}
mytodo,err :=gettodoById(id)
if err!=nil{
	context.JSON(203,"not found task")
	return
}
context.JSON(200,mytodo)
}


func deletTodoById(id int)(string,error){
	for i,t:=range todoes{
		
		if t.Id==id{	
			todoes = append(todoes[:i], todoes[i+1:]...)
			return "doen",nil
		}

	}
	return "status:ther is no todo with this id",errors.New("not Found")
}

func deletTodo(context *gin.Context)  {
	strID:=context.Param("id")
	id,err:=idChecker(strID)
	fmt.Println(err)
	if err!=nil{
		 context.IndentedJSON(400,gin.H{"msg":err.Error()})
		 return
	}
	msg,err:=deletTodoById(id)
	if err!=nil{
		context.IndentedJSON(400,gin.H{"msg":err.Error()})
		return
	}
	context.IndentedJSON(200,gin.H{"msg":msg})

	
}

func idChecker(strID string )(int,error) {
		id,err:=strconv.Atoi(strID)
		if err!=nil{
			
			return 0,errors.New("ID not Valid")
		}
		return id,nil
	}

