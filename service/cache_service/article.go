package cache_service

import (
	"gin-blog/pkg/e"
	"strconv"
	"strings"
)

type Article struct {
	ID int
	tagID int
	State int

	PageNum int
	PageSize int
}

func (a *Article) GetArticleKey() string{
	return e.CACHE_ARTICLE + "_" + strconv.Itoa(a.ID)
}

func (a *Article) GetArticlesKey() string {
	keys := []string{
		e.CACHE_ARTICLE,
		"LIST",
	}

	if a.ID> 0{
		keys = append(keys,strconv.Itoa(a.ID))
	}
	if a.tagID > 0{
		keys = append(keys, strconv.Itoa(a.tagID))
	}
	if a.State >= 0{
		keys = append(keys, strconv.Itoa(a.State))
	}
	if a.PageNum > 0 {
		keys = append(keys, strconv.Itoa(a.PageNum))
	}
	if a.PageSize > 0{
		keys = append(keys, strconv.Itoa(a.PageNum))
	}
	return strings.Join(keys,"_")
}
