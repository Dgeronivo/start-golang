package questlisting

type questInListing struct {
	Id          string
	Title       string
	Description string
	Author      string
}

func provideQuests() []questInListing {
	return getQuestsFromDB()
}

// todo add query to DB
func getQuestsFromDB() []questInListing {
	return []questInListing{
		{Id: "uuid", Title: "title1", Description: "Lorem ipsum dolor sit asd", Author: "Admin"},
		{Id: "uuid", Title: "title1", Description: "Lorem ipsum dolor sit amet...", Author: "Admin"},
		{Id: "uuid", Title: "title1", Description: "Lorem ipsum dolor sit amet...", Author: "Admin"},
		{Id: "uuid", Title: "title1", Description: "Lorem ipsum dolor sit dsa .", Author: "Admin"},
		{Id: "uuid", Title: "title1", Description: "Lorem ipsum dolor sit amet...", Author: "Admin"},
	}
}
