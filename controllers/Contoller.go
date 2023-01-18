package controllers

import (
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/module_page/models"
	"github.com/module_page/services"
)

type PageController struct {
	PageService services.PageService
}

func SortByPriority_Pages(mpp map[string]int) []string {
	keys := make([]string, 0, len(mpp))

	for key := range mpp {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return mpp[keys[i]] > mpp[keys[j]]
	})
	return keys
}

func New(pageservice services.PageService) PageController {
	return PageController{
		PageService: pageservice,
	}
}

func (pgc *PageController) CreateNewPage(ctx *gin.Context) {
	var page models.Pages
	if err := ctx.ShouldBindJSON(&page); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	temp := strings.Split(page.Key, " ")
	//checking total number of keys
	if len(temp) <= 10 {

		err := pgc.PageService.AddPage(&page)

		if err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "success"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "key length out of limit"})
	}

}

func (pgc *PageController) GetByQuery(ctx *gin.Context) {
	temp := []string{}
	var query string = ctx.Param("query")
	queries := strings.Split(query, " ")

	user, err := pgc.PageService.GetAllPages()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	for _, t := range user {
		temp = append(temp, t.Key)
	}
	val := []int{}
	for i := range temp {
		var sum int = 0
		temp2 := strings.Split(temp[i], " ")
		for k := 0; k < len(temp2); k++ {
			for j := range queries {
				if strings.EqualFold(temp2[k], queries[j]) {
					sum += (10 - k) * (10 - j)
				}
			}
		}
		val = append(val, sum)
	}
	mpp := map[string]int{}
	for l, varr := range val {
		if varr != 0 {
			str := strconv.Itoa(l + 1)
			var pageNo string = "P" + str
			mpp[pageNo] = varr
		}
	}

	ctx.JSON(http.StatusOK, SortByPriority_Pages(mpp))

}

func (pgc *PageController) GetAllPage(ctx *gin.Context) {
	users, err := pgc.PageService.GetAllPages()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}
func (pgc *PageController) active(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "api is active"})
}
func (pgc *PageController) Routes(route *gin.RouterGroup) {
	route.POST("/addpage", pgc.CreateNewPage)
	route.GET("/pages", pgc.GetAllPage)
	route.GET("/:query", pgc.GetByQuery)
	route.GET("/", pgc.active)
}
