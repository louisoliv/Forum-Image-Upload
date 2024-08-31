package module

var (
	Error     []Data
	Error2    Data
	database  Database
	user_data User_info
	// DBPosts     []Posts
	DBComments      []Comments
	Loaded          = false
	Loadcomment     = false
	Loadlike        = false
	LoadedUser      = false
	previousMessage string
	ErrorMessage3   string
	ErrorMessage4   string
	ErrorMessage5   string

	// previousUsername string
)

type Data struct {
	ErrorMessage string
}
type Database struct {
	User          User_info
	User_public   []User_public
	Posts         []Posts
	FilteredPosts []FilteredPosts
	Comments      []Comments
	LikE          []LikE
	DislikE       []DislikE
}
type User_info struct {
	Id            int
	Username      string
	Firstname     string
	Lastname      string
	Email         string
	Role          string
	Password      string
	Connexion     bool
	Token         string
	ErrorMessage  string
	ErrorMessage2 string
	ErrorMessage3 string
	ErrorMessage4 string
	ErrorMessage5 string
	TokenError    string
	Filtered      bool
	Category      string
	Userpublic    string
}
type User_public struct {
	Id             int
	Username       string
	Email          string
	Firstname      string
	Lastname       string
	Role           string
	Token          string
	Date           string
	Messages       int
	Localisation   string
	Statut         string
	Loisirs        string
	Date_Naissance string
	Sexe           string
}
type Posts struct {
	Id         int
	Username   string
	Date       string
	Token      string
	Message    string
	Golang     bool
	JavaScript bool
	Python     bool
	Rust       bool
	HTML_CSS   bool
	Angular    bool
	Autre      bool
	Like       int
	Dislike    int
	Image      string
}
type FilteredPosts struct {
	Id            int
	Username      string
	Date          string
	Token         string
	Message       string
	Golang        bool
	JavaScript    bool
	Python        bool
	Rust          bool
	HTML_CSS      bool
	Angular       bool
	Autre         bool
	Like          int
	Dislike       int
	Image         string
	ErrorMessage5 string
}

type Comments struct {
	Id              int
	Username        string
	Date            string
	TokenComment    string
	Message_comment string
	TokenID         string
	Like            int
	Dislike         int
	Image_comment   string
}
type LikE struct {
	Username string
	Token    string
}
type DislikE struct {
	Username string
	Tokens   string
}
