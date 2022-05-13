package model

type LoginTab struct {
	Id           uint64 `gorm:"column:id;primary_key"`
	Name         string `gorm:"column:name"`
	Key          string `gorm:"column:key"`
	Token        string `gorm:"column:token"`
	ReadTimes    uint32 `gorm:"column:readtimes"`
	WrittenTimes uint32 `gorm:"column:writetimes"`
}

func (*LoginTab) Table() string {
	return "test_login_tab"
}

func (*LoginTab) DB() string {
	return "db"
}
