package main

import(
	"study/router"
	"study/model"
)

func init(){
	model.DB = model.InitDB(model.DB)
}

func main(){
	defer model.DB.Close()

	router.Router()

}
