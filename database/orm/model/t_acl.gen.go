package model

const TableNameTAcl = "t_acl"

// TAcl 權限表
type TAcl struct {
	AclID    int64  `gorm:"column:AclId;primaryKey;autoIncrement:true;comment:ACL ID" json:"AclId"` // ACL ID
	ParentID int64  `gorm:"column:ParentId;comment:父ACL ID" json:"ParentId"`                        // 父ACL ID
	Key      string `gorm:"column:Key;comment:ACL碼" json:"Key"`                                     // ACL碼
	Name     string `gorm:"column:Name;comment:ACL名稱" json:"Name"`                                  // ACL名稱
	API      string `gorm:"column:Api;comment:API Uri" json:"Api"`                                  // API Uri
	IsLog    bool   `gorm:"column:IsLog;not null;comment:是否寫log,0-不寫入,1-寫入" json:"IsLog"`           // 是否寫log,0-不寫入,1-寫入
}

// TableName TAcl's table name
func (*TAcl) TableName() string {
	return TableNameTAcl
}
