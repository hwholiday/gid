package entity

type Segments struct {
	BizTag     string `json:"biz_tag" xorm:"not null pk VARCHAR(128) 'biz_tag'"`
	MaxId      int64  `json:"max_id" xorm:"BIGINT(20) 'max_id'"`
	Step       int    `json:"step" xorm:"INT(11) 'step'"`
	Remark     string `json:"remark" xorm:"VARCHAR(200) 'remark'"`
	CreateTime int64  `json:"create_time" xorm:"BIGINT(20) 'create_time'"`
	UpdateTime int64  `json:"update_time" xorm:"BIGINT(20) 'update_time'"`
}

func GetSegmentsTableName() string {
	return "segments"
}
