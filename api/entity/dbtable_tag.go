package entity

type DBTableTag struct {
	ID   int64  `gorm:"column:id;primaryKey" json:"id"`
	Name string `gorm:"column:name;uniqueIndex" json:"name"`
}

func (DBTableTag) TableName() string {
	return "tags"
}
