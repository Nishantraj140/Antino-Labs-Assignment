package blog

import (
	"Antino-Labs-Assignment/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateArticle(g *gin.Context) {
	a := &Article{}

	err := g.ShouldBindJSON(a)
	if err != nil {
		logger.ErrorLogger.Printf("error in CreateArticle ShouldBindJSON, err:%v", err)
		g.JSON(http.StatusUnprocessableEntity, gin.H{
			"data":     nil,
			"debugMsg": "Invalid json provided",
		})
		return
	}
	logger.InfoLogger.Printf("CreateArticle, req:%v", a)

	err = a.Create()
	if err != nil {
		logger.ErrorLogger.Printf("error in CreatArticle Create, err:%v", err)
		g.JSON(http.StatusInternalServerError, gin.H{
			"data":     nil,
			"debugMsg": "Error in creating article",
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"data":     a,
		"debugMsg": "",
	})
	return
}

func GetArticleByID(g *gin.Context) {
	articleID := g.Param("id")

	articleIDInt, err := strconv.Atoi(articleID)
	if err != nil {
		logger.ErrorLogger.Printf("Error converting article id to int, err:%v", err)
		g.JSON(http.StatusUnprocessableEntity, gin.H{
			"data":     nil,
			"debugMsg": "",
		})
		return
	}

	a := &Article{}
	a.ID = articleIDInt

	err = a.GetArticleByID()
	if err != nil {
		logger.ErrorLogger.Printf("Error while getting article details from DB, err:%v", err)
		g.JSON(http.StatusInternalServerError, gin.H{
			"data":     nil,
			"debugMsg": "",
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"data":     a,
		"debugMsg": "",
	})
	return
}

func GetAllArticles(g *gin.Context) {
	a := &Article{}
	allArticles, err := a.GetAllArticles()
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"data":     nil,
			"debugMsg": "Error in getting all articles get method",
		})
		return
	}
	logger.InfoLogger.Println("articles list:%v", a)
	g.JSON(http.StatusOK, gin.H{
		"data":     allArticles,
		"debugMsg": "",
	})
}

func UpdateArticle(g *gin.Context) {
	articleID := g.Param("id")

	articleIDInt, err := strconv.Atoi(articleID)
	if err != nil {
		logger.ErrorLogger.Printf("Error converting article id to int, err:%v", err)
		g.JSON(http.StatusUnprocessableEntity, gin.H{
			"data":     nil,
			"debugMsg": "",
		})
		return
	}

	a := &Article{}
	a.ID = articleIDInt

	err = a.GetArticleByID()
	if err != nil {
		logger.ErrorLogger.Printf("Error while getting article details from DB, err:%v", err)
		g.JSON(http.StatusInternalServerError, gin.H{
			"data":     nil,
			"debugMsg": "Please provide a valid article id whose details needs to updated",
		})
		return
	}

	err = g.ShouldBindJSON(a)
	if err != nil {
		logger.ErrorLogger.Printf("error in UpdateArticle ShouldBindJSON, err:%v", err)
		g.JSON(http.StatusUnprocessableEntity, gin.H{
			"data":     nil,
			"debugMsg": "Invalid json provided",
		})
		return
	}
	logger.InfoLogger.Printf("UpdateArticle, req:%v", a)

	err = a.Update()

	if err != nil {
		logger.ErrorLogger.Printf("error in updating article, Id:%d, err:%v", articleID, err)
		g.JSON(http.StatusInternalServerError, gin.H{
			"data":     nil,
			"debugMsg": "Error in updating article",
		})
		return
	}
	g.JSON(http.StatusOK, gin.H{
		"data":     fmt.Sprintf("the %v article updated successfully ", a.ID),
		"debugMsg": "",
	})
}

func DeleteArticle(g *gin.Context) {
	articleID := g.Param("id")

	articleIDInt, err := strconv.Atoi(articleID)
	if err != nil {
		logger.ErrorLogger.Printf("Error converting article id to int, err:%v", err)
		g.JSON(http.StatusUnprocessableEntity, gin.H{
			"data":     nil,
			"debugMsg": "",
		})
		return
	}

	a := &Article{}
	a.ID = articleIDInt

	err = a.GetArticleByID()
	if err != nil {
		logger.ErrorLogger.Printf("Error while getting article details from DB, err:%v", err)
		g.JSON(http.StatusInternalServerError, gin.H{
			"data":     nil,
			"debugMsg": "Please provide a valid article id which needs to be deleted",
		})
		return
	}

	err = a.Delete()

	if err != nil {
		logger.ErrorLogger.Printf("error in deleting article, Id:%d, err:%v", articleID, err)
		g.JSON(http.StatusInternalServerError, gin.H{
			"data":     nil,
			"debugMsg": "Error in deleting article",
		})
		return
	}
	g.JSON(http.StatusOK, gin.H{
		"data":     fmt.Sprintf("the %v article deleted successfully ", a.ID),
		"debugMsg": "",
	})

}
