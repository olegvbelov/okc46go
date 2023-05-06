package models

import "html/template"

type QuestionAnswer struct {
	Id       string
	Question string
	Answer   string
}

type PlanDetail struct {
	Description template.HTML
	Cost        string
}

type OkcService struct {
	Id                string
	Title             string
	ShortDescription  template.HTML
	Description       template.HTML
	Link              string
	IconUrl           string
	Photo390x260      string
	Webp390x260       string
	OnMain            bool
	DetailUrl         string
	QuestionAnswers   []QuestionAnswer
	DetailImageUrl    []string
	DetailDescription []string
	DetailPlan        []PlanDetail
}

var Services = []OkcService{
	{
		Id:               "1",
		Title:            "Межевание участка",
		ShortDescription: "Проведем для Вас межевание земельного участкаи и подготовим пакет документов для постановки на учет",
		Description:      "Подготовка межевого плана<br/>Постановка участка на кадастровый учет<br/> Уточнение границ земельного участка",
		Photo390x260:     "/static/images/okc46/services/390x260/mega.jpg",
		Webp390x260:      "/static/images/okc46/services/390x260/mega.webp",
		OnMain:           true,
		DetailUrl:        "/details/1",
		QuestionAnswers: []QuestionAnswer{
			{"accordion-owner", "Возможно ли оформление меживание не собственником", "Да, при наличии нотариально заверенной доверенности."},
			{"accordion-agreement", "Кто проводит согласование кадастрового плана с соседями?", "Согласование проводит собственник участка."},
			{"accordion-passport", "Где можно получить кадастровый паспорт", "Кадастровый паспорт получают в БТИ."},
		},
		Link:           "/services",
		DetailImageUrl: []string{"mega.jpg", "mega1.jpg"},
		DetailDescription: []string{"Межевание земельного участка - установка на топографической карте местности координат земельного участка, определение и согласование его границ с целью постановки земельного участка на кадастровый учет.",
			"Межевание позволит решить возможные конфликты с соседями, получить право собственности и совершать юридические операции с земельным участком: дарить, менять, продавать."},
		DetailPlan: []PlanDetail{{"Подготовить документы на замельный участок: <br/><ul><li class=\"template-arrow-right-2\">заявление, составленное в свободной форме;</li><li class=\"template-arrow-right-2\">паспорт владельца;</li><li class=\"template-arrow-right-2\">правоустанавливающие документы на земельный участок;</li><li class=\"template-arrow-right-2\">правоудостоверяющие документы на земельный участок;</li><li class=\"template-arrow-right-2\">кадастровая выписка;</li><li class=\"template-arrow-right-2\">кадастровый паспорт;</li></ul>",
			""},
			{"Заключить договор с кадастровым инженером/землеустроительной организацией", "6000 рублей"},
			{"Проведение кадастровым инженером замеров и подготовка кадастрового плана", ""},
			{"Ознакомление с составленным кадастровым планом", ""},
			{"Согласование кадастрового плана с владельцами смежных участков", ""},
			{"Отправка заявления и кадастрового плана в орган регистрации прав", ""},
			{"Получение выписки из ЕГРН", ""},
		},
	},
	{
		Id:               "2",
		Title:            "Выдел пая",
		ShortDescription: "Выдел пая из земель сельхозназначения",
		Description:      "Межевание земельного участка<br/>Подготовка кадастрового плана<br/>Информационное сопровождение",
		Photo390x260:     "/static/images/okc46/services/390x260/pie.jpg",
		Webp390x260:      "/static/images/okc46/services/390x260/mega.webp",
		OnMain:           false,
		DetailUrl:        "/services",
		Link:             "/services",
	},
	{
		Id:               "3",
		Title:            "Топографическая съемка",
		ShortDescription: "Наши специалисты выполнят топографическую съемку и профессиональную обработу результатов",
		Description:      "Топографическая съемка земельного участка<br/>Топографическая съемка линейных сооружений<br/>Камеральные исследования<br/>",
		Photo390x260:     "/static/images/okc46/services/390x260/topo.jpg",
		Webp390x260:      "/static/images/okc46/services/390x260/topo.webp",
		OnMain:           true,
		DetailUrl:        "/services",
		Link:             "/services",
	},
	{
		Id:               "4",
		Title:            "Оценка недвижимости",
		ShortDescription: "Сертифицированные специалисты проведут оценку квартиры, дома, строения и земельного участка",
		Description:      "Оценка квартиры, дома для ипотеки<br/>Оценка объекта недвижимости для суда<br/>Оценка недвижимости для наследства<br/>",
		Photo390x260:     "/static/images/okc46/services/390x260/estimation.jpg",
		Webp390x260:      "/static/images/okc46/services/390x260/estimation.webp",
		OnMain:           true,
		DetailUrl:        "/services",
		Link:             "/services",
	},
	{
		Id:               "5",
		Title:            "Подготовка плана объекта",
		ShortDescription: "Подготовим для Вас карту объекта землеустройства",
		Description:      "Получене сведений государственного кадастра<br/>Определение границ на местности<br/>Проведение согласований",
		Photo390x260:     "/static/images/okc46/services/390x260/objectplan.jpg",
		Webp390x260:      "/static/images/okc46/services/390x260/objectplan.webp",
		OnMain:           false,
		DetailUrl:        "/services",
		Link:             "/services",
	},
	{
		Id:               "6",
		Title:            "Технический план",
		ShortDescription: "Подготовим для Вас технический план здания, помещения",
		Description:      "Подготовка технического плана здания, помещения<br/>Регистрация права на объект недвижимости<br/>Внесение сведений об изменении в ЕГРН",
		Photo390x260:     "/static/images/okc46/services/390x260/techplan.jpg",
		Webp390x260:      "/static/images/okc46/services/390x260/objectplan.webp",
		OnMain:           false,
		DetailUrl:        "/services",
		Link:             "/services",
	},
	{
		Id:               "7",
		Title:            "Вынос точек",
		ShortDescription: "Вынос границ земельного участка в натуру",
		Description:      "Вынос координат и точек с соблюдением требований СНиП<br/>Установка временных или постоянных межевых знаков<br/>Составление акта",
		Photo390x260:     "/static/images/okc46/services/390x260/vinos.jpg",
		Webp390x260:      "/static/images/okc46/services/390x260/vinos.webp",
		OnMain:           false,
		DetailUrl:        "/services",
		Link:             "/services",
	},
	{
		Id:               "8",
		Title:            "Энергоаудит",
		ShortDescription: "Энергоаудит",
		Description:      "Испытание зданий, жилых помещений на воздхо и тепло проницаемость<br/>Поиск утечек тепла, замеры, консльтации, отчеты</br>Сертифицированное обордование",
		Photo390x260:     "/static/images/okc46/services/390x260/teplovisor.jpg",
		Webp390x260:      "/static/images/okc46/services/390x260/teplovisor.webp",
		OnMain:           false,
		DetailUrl:        "/services",
		Link:             "/services",
	},
}

type Contact struct {
	Id       int
	City     string
	Address  template.HTML
	Phone    string
	Email    string
	WorkTime template.HTML
	Link     string
	MapUrl   string
	IsMain   bool
}

var Offices = []Contact{
	{
		Id:       1,
		City:     "Курск",
		Address:  "305006, Курск<br/>ул. Садовая, 12, офис 1",
		Phone:    "+7 (471-2)36 03 08",
		Email:    "sales@okc46.ru",
		WorkTime: "Пон - Пят: 09.00 - 18.00<br/>Суб: 10.00 - 16.00",
		Link:     "kursk",
		MapUrl:   "",
		IsMain:   true,
	},
	{
		Id:       2,
		City:     "Фатеж",
		Address:  "307100, Фатеж<br>ул. Тихая, 35",
		Phone:    "+7 (471-2)36 03 08",
		Email:    "sales@okc46.ru",
		WorkTime: "Пон - Пят: 09.00 - 18.00",
		Link:     "phatezh",
		MapUrl:   "",
		IsMain:   false,
	},
	{
		Id:       4,
		City:     "Конышевка",
		Address:  "307620, п. Конышевка<br>ул. Маяковского, 1а",
		Phone:    "+7 (471-2)36 03 08",
		Email:    "sales@okc46.ru",
		WorkTime: "Пон - Пят: 09.00 - 15.00",
		Link:     "konishevka",
		MapUrl:   "",
		IsMain:   false,
	},
	{
		Id:       5,
		City:     "Курчатов",
		Address:  "307251, Курчатов<br>Коммунистический пр-т, 30",
		Phone:    "+7 (471-2)36 03 08",
		Email:    "sales@okc46.ru",
		WorkTime: "Пон - Пят: 09.00 - 15.00",
		Link:     "kurchatov",
		MapUrl:   "",
		IsMain:   false,
	},
	{
		Id:       6,
		City:     "Прямицино",
		Address:  "307200, пгт. Прямицино,<br/>ул. Октябрьская, 138",
		Phone:    "+7 (471-2)36 03 08",
		Email:    "sales@okc46.ru",
		WorkTime: "Пон - Пят: 09.00 - 15.00<br/>Суб: 9.00 - 12.00",
		Link:     "pryamitsino",
		MapUrl:   "",
		IsMain:   false,
	},
}

type SiteMessage struct {
	Name    string
	Email   string
	Phone   string
	Message string
}
