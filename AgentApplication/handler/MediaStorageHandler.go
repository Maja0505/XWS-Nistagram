package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type MediaStorageHandler struct {}

func (handler *MediaStorageHandler) GetMediaImage(c *gin.Context) {
	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Content-Type", "image/jpeg")

	imagepath := c.Query("id")
	if imagepath == "" {
		c.JSON(http.StatusUnprocessableEntity, "image path not provided")
		return
	}

	file,err := ioutil.ReadFile("product-images/" + imagepath)
	if err != nil{
		c.JSON(http.StatusBadRequest,"")
		fmt.Println(err)
		return
	}
	img, _ :=c.Writer.Write(file)
	c.JSON(http.StatusOK,img)


}

func (handler *MediaStorageHandler) UploadMediaImage(c *gin.Context){
	imagePath := c.Query("id")
	c.Request.ParseMultipartForm(10 << 20)
	file,_, err := c.Request.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()

	dst, err := os.Create("product-images/" + imagePath +".jpg")
	defer dst.Close()
	if err != nil {
		c.JSON(http.StatusBadRequest,err)
		return
	}

	if _, err := io.Copy(dst, file); err != nil {
		c.JSON(http.StatusBadRequest,err)
		return
	}

	c.JSON(http.StatusOK,"Successfully uploaded image")


}

