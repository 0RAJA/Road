package settions

import "time"

type AllSettings struct {
	Upload    Upload    `yaml:"Upload"`
	Server    Server    `yaml:"Server"`
	App       App       `yaml:"App"`
	Log       Log       `yaml:"Log"`
	Redis     Redis     `yaml:"Redis"`
	Email     Email     `yaml:"Email"`
	Mysql     Mysql     `yaml:"Mysql"`
	Pagelines Pagelines `yaml:"Pagelines"`
	Token     Token     `yaml:"Token"`
	Github    Github    `yaml:"Github"`
	Rule      Rule      `yaml:"Rule"`
}

type Rule struct {
	UsernameLen           int           `yaml:"UsernameLen"`
	PasswordLen           int           `yaml:"PasswordLen"`
	DefaultAvatarUrl      string        `yaml:"DefaultAvatarUrl"`
	CommentLen            int           `yaml:"CommentLen"`
	TitleLen              int           `yaml:"TitleLen"`
	AbstractLen           int           `yaml:"AbstractLen"`
	TagLen                int           `yaml:"TagLen"`
	EndurancePostStarTime time.Duration `yaml:"EndurancePostStarTime"`
	EndurancePostViews    time.Duration `yaml:"EndurancePostViews"`
	EnduranceViewTime     time.Duration `yaml:"EnduranceViewTime"`
}

type Github struct {
	ClientID     string `yaml:"ClientID"`
	ClientSecret string `yaml:"ClientSecret"`
	RedirectUri  string `yaml:"RedirectUri"`
}

type Pagelines struct {
	DefaultPageSize int32  `yaml:"DefaultPageSize"`
	PageKey         string `yaml:"PageKey"`
	PageSizeKey     string `yaml:"PageSizeKey"`
	MaxPageSize     int32  `yaml:"MaxPageSize"`
}

type File struct {
	Type      string   `yaml:"Type"`
	MaxSize   int      `yaml:"MaxSize"`
	UrlPrefix string   `yaml:"UrlPrefix"`
	LocalPath string   `yaml:"LocalPath"`
	Suffix    []string `yaml:"Suffix"`
}

type Server struct {
	Address               string        `yaml:"Address"`
	ReadTimeout           time.Duration `yaml:"ReadTimeout"`
	WriteTimeout          time.Duration `yaml:"WriteTimeout"`
	DefaultContextTimeout time.Duration `yaml:"DefaultContextTimeout"`
	RunMode               string        `yaml:"RunMode"`
}

type Redis struct {
	Address         string        `yaml:"Address"`
	DB              int           `yaml:"DB"`
	Password        string        `yaml:"Password"`
	PoolSize        int           `yaml:"PoolSize"`
	PostTimeout     time.Duration `yaml:"PostTimeout"`
	PostInfoTimeout time.Duration `yaml:"PostInfoTimeout"`
}

type Email struct {
	Port     int      `yaml:"Port"`
	UserName string   `yaml:"UserName"`
	Password string   `yaml:"Password"`
	IsSSL    bool     `yaml:"IsSSL"`
	From     string   `yaml:"From"`
	To       []string `yaml:"To"`
	Host     string   `yaml:"Host"`
}

type Mysql struct {
	DriverName   string `yaml:"DriverName"`
	SourceName   string `yaml:"SourceName"`
	MaxOpenConns int    `yaml:"MaxOpenConns"`
	MaxIdleConns int    `yaml:"MaxIdleConns"`
}

type Token struct {
	AuthorizationType    string        `yaml:"AuthorizationType"`
	Key                  string        `yaml:"Key"`
	AssessTokenDuration  time.Duration `yaml:"AssessTokenDuration"`
	RefreshTokenDuration time.Duration `yaml:"RefreshTokenDuration"`
}

type Upload struct {
	StaticPath string `yaml:"StaticPath"`
	Image      Image  `yaml:"Image"`
	File       File   `yaml:"File"`
}

type Image struct {
	MaxSize   int      `yaml:"MaxSize"`
	UrlPrefix string   `yaml:"UrlPrefix"`
	LocalPath string   `yaml:"LocalPath"`
	Suffix    []string `yaml:"Suffix"`
	Type      string   `yaml:"Type"`
}

type App struct {
	Name      string `yaml:"Name"`
	Version   string `yaml:"Version"`
	StartTime string `yaml:"StartTime"`
	Format    string `yaml:"Format"`
	Logo      string `yaml:"Logo"`
}

type Log struct {
	Level         string `yaml:"Level"`
	MaxSize       int    `yaml:"MaxSize"`
	MaxAge        int    `yaml:"MaxAge"`
	MaxBackups    int    `yaml:"MaxBackups"`
	Compress      bool   `yaml:"Compress"`
	LogSavePath   string `yaml:"LogSavePath"`
	HighLevelFile string `yaml:"HighLevelFile"`
	LowLevelFile  string `yaml:"LowLevelFile"`
	LogFileExt    string `yaml:"LogFileExt"`
}
