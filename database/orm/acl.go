package orm

type Acl struct {
	AclId    int64  `gorm:"type:bigint NOT NULL auto_increment;primary_key;" json:"AclId,omitempty"`
	ParentId string `gorm:"type:varchar(100) NOT NULL;" json:"ParentId,omitempty"`
	Key      string `gorm:"type:varchar(100) NOT NULL; index:idx_acl_key"  json:"Key,omitempty"`
	Name     string `gorm:"type:varchar(100) NOT NULL;" json:"name,omitempty"`
}
