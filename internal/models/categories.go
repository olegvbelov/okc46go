package models

type Category struct {
	Name string
	Url  string
}

var Categories = []Category{
	{"кадастр", "/services"},
	{"проект", "/services"},
	{"кадастровый учет", "/services"},
	{"межевой план", "/details/1"},
	{"оценка", "/services"},
	{"технический план", "/services"},
	{"межевание", "/details/1"},
	{"геодезия", "/services"},
	{"перепланировка", "/services"},
	{"энергоаудит", "/services"},
}
