package handler

import(
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"log"
	"regexp"

	"study/model"
)

func Spider() {
	var book model.Book
	var books []model.Book

	model.DB.Find(&books)

	requestUrl := "https://book.douban.com/latest?icn=index-latestbook-all"
	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:60.0) Gecko/20100101 Firefox/60.0")
	response, _ := http.DefaultClient.Do(req)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer response.Body.Close()
	content := string(body)

	reg1 := regexp.MustCompile(`<li>(?s:(.*?))</li>`)
	reg2 := regexp.MustCompile(`<ahref="https://book.douban.com/subject/[0-9]+/">(?s:(.*?))</a>`)
	reg3 := regexp.MustCompile(`<pclass="color-gray">(?s:(.*?))</p>`)
	reg4 := regexp.MustCompile(`<p(class="detail")?>(?s:(.*?))</p>`)
	reg5 := regexp.MustCompile(`<imgsrc="(?s:(.*?))"/>`)

	result := reg1.FindAllStringSubmatch(content, -1) //所有书籍信息
	for _, aim := range result {
		bookInfo := strings.Replace(aim[1], " ", "", -1)
		bookInfo = strings.Replace(bookInfo, "\n", "", -1) //取消所有空格与回车,注意正则中也要取消

		//标题
		result2 := reg2.FindAllStringSubmatch(bookInfo, -1)
		book.Book_name = result2[0][1]
		//检测是否书籍已存在
		isHave := false
		for _, oldbook := range books {
			if oldbook.Book_name == book.Book_name {
				isHave = true
				break
			}
		}
		if isHave {
			continue
		}

		//作者
		result2 = reg3.FindAllStringSubmatch(bookInfo, -1)
		book.Book_auther = result2[0][1]

		//简介
		result2 = reg4.FindAllStringSubmatch(bookInfo, -1)
		book.Book_information = result2[0][2]

		//图片
		result2 = reg5.FindAllStringSubmatch(bookInfo, -1)
		book.Book_picture = result2[0][1]

		//类别、点击量
		book.Class_id = 0
		book.Click_sum = 0

		fmt.Println(book)
		model.DB.Create(&book)
	}
}
