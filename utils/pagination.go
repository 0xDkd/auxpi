package utils

import (
	"math"
	"strconv"
)

func GetPage(page, limit int) (int, int) {
	num := (page - 1) * limit
	return num, limit
}

func GetStringPage(page, limit string) (int, int) {
	iPage, _ := strconv.Atoi(page)
	iLimit, _ := strconv.Atoi(limit)

	if iLimit == 0 {
		iLimit = 20
	}

	num := (iPage - 1) * iLimit
	return num, iLimit
}

func PageHtml(nowPage int, all, pageSize int) string {
	if nowPage <= 0 {
		nowPage = 1
	}
	firstActive := ""
	firstDisable := ""
	url := "?page="
	pageCount := math.Ceil(float64(all) / float64(pageSize))
	IntPageCount := int(pageCount)
	previous := strconv.Itoa(nowPage - 1)
	next := strconv.Itoa(nowPage + 1)
	firstStatus := ""
	if nowPage == 1 {
		firstActive = "active"
		firstDisable = "disabled"
		firstStatus = "true"
	}

	//第一页
	pageTpl := `<nav aria-label="...">
        <ul class="pagination">
			<li class="page-item ` + firstDisable + `">
      			<a class="page-link" href="` + url + previous + `" tabindex="-1" aria-disabled="` + firstStatus + `">Previous</a>
    		</li>
            <li class="page-item ` + firstActive + `"><a class="page-link" href="?page=1"+ >1</a></li>`

	//省略号
	if nowPage > 5 {
		pageTpl += `<li class="page-item disabled">
      <a class="page-link" href="#" tabindex="-1" aria-disabled="true">...</a>
    </li>`
	}
	for i := nowPage - 3; i < nowPage+3; i++ {
		active := ""
		if i < 2 || i > IntPageCount {
			continue
		}

		if nowPage == i {
			active = "active"
		} else {
			active = ""
		}

		if i == IntPageCount {
			break
		}

		s := strconv.Itoa(i)
		pageTpl += `<li class="page-item ` + active + `"><a class="page-link" href="` + url + s + `">` + s + `</a></li>`

	}

	if nowPage < IntPageCount-3 {
		pageTpl += `<li class="page-item disabled">
      <a class="page-link" href="#" tabindex="-1" aria-disabled="true">...</a>
    </li>`
	}

	lastActive := ""
	lastDisable := "false"
	if IntPageCount == nowPage {
		lastActive = "active"
		lastDisable = "disabled"
	}
	sCount := strconv.Itoa(IntPageCount)
	pageTpl += `<li class="page-item ` + lastActive + `"><a class="page-link" href="` + url + sCount + `">` + sCount + `</a></li>
			<li class="page-item ` + lastDisable + `">
      			<a class="page-link" href="` + url + next + `" tabindex="-1" aria-disabled="` + lastDisable + `">Next</a>
    		</li>
        </ul>
    </nav>`
	if IntPageCount <= 1 {
		pageTpl = ""
	}

	return pageTpl
}
