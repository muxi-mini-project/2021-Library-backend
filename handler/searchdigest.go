package handler

import(
	"strings"

	"study/model"

	"github.com/gin-gonic/gin"
)

func SearchDigest(c *gin.Context){
	var summaries, summariesSelect []model.Summary
	var summaryRow model.SummaryRow
	var summaryRows []model.SummaryRow
	var search model.Search

	user_id := c.Param("user_id")
	model.DB.Where("user_id = ?", user_id).Find(&summaries)

	err := c.BindJSON(&search)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "搜索错误",
		})
	return
	}
	word := search.Word

	if len(summaries) <= 0 {
		c.JSON(200, gin.H{
			"message": "没有数据",
		})
		return
	}

	for i := 0; i < len(summaries); i++ {
		if strings.Contains(summaries[i].Title, word) {
			summariesSelect = append(summariesSelect, summaries[i])
		}
	}

	for i := 0; i < len(summaries); i++ {
		if strings.Contains(summaries[i].Summary_information + summaries[i].Thought, word) {
			isHave := false
			for j := 0; j < len(summariesSelect); j++ {
				if summaries[i].Id == summariesSelect[j].Id {
					isHave = true
					break
				}
			}

			if !(isHave) {
				summariesSelect = append(summariesSelect, summaries[i])
			}
		}
	}

	if len(summariesSelect) <= 0 {
		c.JSON(200, gin.H{
			"message": "没有数据",
		})
		return
	}

	for i := 0; i < len(summariesSelect); i++ {
		summaryRow.Id = summariesSelect[i].Id
		summaryRow.Title = summariesSelect[i].Title
		summaryRow.Date = summariesSelect[i].Date[0:10]
		summaryRow.Public = summariesSelect[i].Public

		summaryRows = append(summaryRows, summaryRow)
	}

	c.JSON(200, gin.H{
		"message": "获取成功",
		"data": summaryRows,
	})


}
