package models

type QuestionAnswer struct {
	Question string
	Answer   string
}

type OkcService struct {
	Id               string
	Title            string
	ShortDescription string
	Description      string
	Link             string
	IconUrl          string
	Photo390x260     string
	Webp390x260      string
	OnMain           bool
	DetailUrl        string
	QuestionAnswers  []QuestionAnswer
}

var Services = []OkcService{
	{
		Id:               "1",
		Title:            "Межевание участка",
		ShortDescription: "Проведем для Вас межевание земельного участка",
		Description:      "Подготовка межевого плана<br/>Постановка участка на кадастровый учет<br/> Уточнение границ земельного участка",
		IconUrl:          "/static/images/samples/390x260/image_01.jpg",
		Photo390x260:     "/static/images/okc46/services/390x260/mega.jpg",
		Webp390x260:      "/static/images/okc46/services/390x260/mega.webp",
		OnMain:           true,
		DetailUrl:        "/servicedetails/%s",
		Link:             "/service/1",
	},
	{
		Id:               "2",
		Title:            "Выдел пая",
		ShortDescription: "Проведем для Вас межевание земельного участка",
		Description:      "Межевание земельного участка<br/>Подготовка кадастрового плана<br/>Информационное сопровождение",
		IconUrl:          "images/icon/3.png",
		Photo390x260:     "/static/images/okc46/services/390x260/pie.jpg",
		Webp390x260:      "/static/images/okc46/services/390x260/mega.webp",
		OnMain:           false,
		DetailUrl:        "",
		Link:             "/service/2",
	},
	{
		Id:               "3",
		Title:            "Топографическая съемка",
		ShortDescription: "Наши специалисты выполнят топографическую съемку и профессиональную обработу результатов",
		Description:      "Топографическая съемка земельного участка<br/>Топографическая съемка линейных сооружений<br/>Камеральные исследования<br/>",
		IconUrl:          "",
		Photo390x260:     "/static/images/okc46/services/390x260/topo.jpg",
		Webp390x260:      "/static/images/okc46/services/390x260/topo.webp",
		OnMain:           true,
		DetailUrl:        "",
		Link:             "/service/3",
	},
	{
		Id:               "4",
		Title:            "Оценка недвижимости",
		ShortDescription: "Сертифицированные специалисты проведут оценку квартиры, дома, строения и земельного участка",
		Description:      "Оценка квартиры, дома для ипотеки<br/>Оценка объекта недвижимости для суда<br/>Оценка недвижимости для наследства<br/>",
		IconUrl:          "",
		Photo390x260:     "/static/images/okc46/services/390x260/estimation.jpg",
		Webp390x260:      "/static/images/okc46/services/390x260/estimation.webp",
		OnMain:           true,
		DetailUrl:        "",
		Link:             "/service/4",
	},
	{
		Id:               "5",
		Title:            "Подготовка плана объекта",
		ShortDescription: "Подготовим для Вас карту объекта землеустройства",
		Description:      "Получене сведений государственного кадастра<br/>Определение границ на местности<br/>Проведение согласований",
		IconUrl:          "images/icon/4.png",
		Photo390x260:     "/static/images/okc46/services/390x260/objectplan.jpg",
		Webp390x260:      "/static/images/okc46/services/390x260/objectplan.webp",
		OnMain:           false,
		DetailUrl:        "",
		Link:             "/service/5",
	},
	{
		Id:               "6",
		Title:            "Технический план",
		ShortDescription: "Подготовим для Вас технический план здания, помещения",
		Description:      "Подготовка технического плана здания, помещения<br/>Регистрация права на объект недвижимости<br/>Внесение сведений об изменении в ЕГРН",
		IconUrl:          "images/icon/2.png",
		Photo390x260:     "/static/images/okc46/services/390x260/techplan.jpg",
		Webp390x260:      "/static/images/okc46/services/390x260/objectplan.webp",
		OnMain:           false,
		DetailUrl:        "",
		Link:             "/service/6",
	},
	{
		Id:               "7",
		Title:            "Вынос точек",
		ShortDescription: "Вынос границ земельного участка в натуру",
		Description:      "Вынос координат и точек с соблюдением требований СНиП<br/>Установка временных или постоянных межевых знаков<br/>Составление акта",
		IconUrl:          "images/icon/1.png",
		Photo390x260:     "/static/images/okc46/services/390x260/vinos.jpg",
		Webp390x260:      "/static/images/okc46/services/390x260/vinos.webp",
		OnMain:           false,
		DetailUrl:        "",
		Link:             "/service/7",
	},
	{
		Id:               "8",
		Title:            "Энергоаудит",
		ShortDescription: "Энергоаудит",
		Description:      "Испытание зданий, жилых помещений на воздхо и тепло проницаемость<br/>Поиск утечек тепла, замеры, консльтации, отчеты</br>Сертифицированное обордование",
		IconUrl:          "images/icon/1.png",
		Photo390x260:     "/static/images/okc46/services/390x260/teplovisor.jpg",
		Webp390x260:      "/static/images/okc46/services/390x260/teplovisor.webp",
		OnMain:           false,
		DetailUrl:        "",
		Link:             "/service/8",
	},
}
