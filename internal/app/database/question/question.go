package question

type Question struct {
	id          uint `gorm:"primaryKey" json:"id`
	title       string
	description string
	options     option
	_type       int
	auto_judge  bool
	result      string
	create_time Datatime
	bank_id     uint
}

type option struct {
	value string
}
