package models

import (
	"gopkg.in/mgo.v2"
	"net/url"
)

//grabInt is a simple handler for grabbing an integer out of a url.Values
//with a default value if there are any errors
//func grabInt(v url.Values, key string, def int) int {
//	val := v.Get(key)
//	n, err := strconv.ParseInt(val, 10, 0)
//	if err != nil || n < 0 {
//		return def
//	}
//	return int(n)
//}

//listParse takes a collection and some query values and generates an iterator
//for the objects that should be returned on that page
func ListParse(c string, page int) (*mgo.Iter, int, int) {
	//parse out sorting
	//sort := map[string]int{}
	//for key, _ := range v {
	//	val := v.Get(key)
	//	if len(key) > 5 && strings.HasPrefix(key, "sort_") {
	//		field := key[5:]
	//		switch strings.ToLower(val) {
	//		case "asc":
	//			sort[field] = 1
	//		case "desc":
	//			sort[field] = -1
	//		}
	//	}
	//}
	//set up the query with the correct sort order
	query := db.C(c).Find(nil)
	//if sort {
	//	query.Sort(sort)
	//}
	numpage := 20
	//pagination
	if page < 1 {
		page = 1
	}
	if numpage < 1 {
		numpage = 1
	}
	//pages are 1 indexed.
	query = query.Skip(numpage * (page - 1)).Limit(numpage)

	return query.Iter(), page, numpage
}

//Pagination helps generate lists of pages for the List view.
type Pagination struct {
	Pages       int
	CurrentPage int
	Query       url.Values
}

//PageList returns a list of integers of size n around the current page. For example
//
//	p := Pagination{8, 4}
//	p.PageList(5) => []int{1,2,3,4,5,6,7,8}
func (p Pagination) PageList(n int) []int {
	bottom, top := p.CurrentPage-n, p.CurrentPage+n
	if bottom < 1 {
		bottom = 1
	}
	if top > p.Pages {
		top = p.Pages
	}

	ints := make([]int, 0, top-bottom+1)
	for i := bottom; i <= top; i++ {
		ints = append(ints, i)
	}

	return ints
}

//IsCurrent helps you determine if a given page number is the current page in
//templates.
func (p Pagination) IsCurrent(page int) bool {
	return page == p.CurrentPage
}

//Next returns the numerical value of the next page. It returns 0 if it would be
//past the last page.
func (p Pagination) Next() int {
	if p.CurrentPage >= p.Pages {
		return 0
	}
	return p.CurrentPage + 1
}

//Prev returns the numerical value of the previous page. It returns 0 if it would
//be before the fist page.
func (p Pagination) Prev() int {
	if p.CurrentPage <= 1 {
		return 0
	}
	return p.CurrentPage - 1
}

//Page adds the requested page to the passed in url.Values and returns the
//encoded result.
//func (p Pagination) Page(n int) string {
//	p.query.Set("page", fmt.Sprint(n))
//	return "?" + p.query.Encode()
//}