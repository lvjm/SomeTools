package model

type scm_contract struct {
	Id               int64  `gorm:"id" json:"id"`                                 // 主键id
	TaskId           int64  `gorm:"task_id" json:"task_id"`                       // 对应的子任务id
	Cid              int64  `gorm:"cid" json:"cid"`                               // 所属公司id
	Uid              int64  `gorm:"uid" json:"uid"`                               // 处理人id
	Title            string `gorm:"title" json:"title"`                           // 表单的名称
	SalesCorpId      int64  `gorm:"sales_corp_id" json:"sales_corp_id"`           // 销售方企业id
	SalesCorpName    string `gorm:"sales_corp_name" json:"sales_corp_name"`       // 销售方企业名称
	PurchaseCorpId   int64  `gorm:"purchase_corp_id" json:"purchase_corp_id"`     // 采购方企业id
	PurchaseCorpName string `gorm:"purchase_corp_name" json:"purchase_corp_name"` // 采购方企业名称
	Code             string `gorm:"code" json:"code"`                             // 合同编码
	StartDate        string `gorm:"start_date" json:"start_date"`                 // 开始日期
	EndDate          string `gorm:"end_date" json:"end_date"`                     // 结束日期
	RemindDate       string `gorm:"remind_date" json:"remind_date"`               // 提醒日期,默认提前三个月
	Amount           int64  `gorm:"amount" json:"amount"`                         // 总金额，单位分
	SettlementType   int64  `gorm:"settlement_type" json:"settlement_type"`       // 结算方式 1:按订单结算 2:按入库结算 3:按预付款结算 4:按实销月结 5:按发货月结 6:按发货几个月后结算 7:按铺货多少金额后几个月后结算
	SettlementValue  string `gorm:"settlement_value" json:"settlement_value"`     // 结算方式值，月数，或金额(单位分) 形如[{"code":6,"month":3},{"code":7,"month":3,"amount":"23423423"}]
	PayType          int64  `gorm:"pay_type" json:"pay_type"`                     // 支付方式： 1 现金，2 银行转账，3 微信，4 支付宝，5 网银，6 现金支票，7 银行支票，8 电汇
	Voucher          string `gorm:"voucher" json:"voucher"`                       // 凭证，文件链接，string数组
	Remark           string `gorm:"remark" json:"remark"`                         // 备注
	IsSplit          int64  `gorm:"is_split" json:"is_split"`                     // 是否表单拆分而来：0 否，1 是
	MergedTables     string `gorm:"merged_tables" json:"merged_tables"`           // 数据来源表单，多对一就合并[{"task_id":1,"form_id":99},{"task_id":1,"form_id":99}]，一对一时填单元素数组,没有为[]
	WriteOffTables   string `gorm:"write_off_tables" json:"write_off_tables"`     // 被核销的表单的 formid [{"task_id":1,"form_id":99},{"task_id":1,"form_id":99}]，没有就为[]
	RelatedTables    string `gorm:"related_tables" json:"related_tables"`         // 被关联的表单的 formid [{"task_id":1,"form_id":99},{"task_id":1,"form_id":99}]，没有就为[]
	BusiStatus       int64  `gorm:"busi_status" json:"busi_status"`               // 业务状态：-1 审核不过 0 提交待审核 1 审核中 2 已通过 3/4/5(业务决定)...
	DeriveStatus     int64  `gorm:"derive_status" json:"derive_status"`           // 数据流转状态：0 已生成，1 已流转，90 已核销，100 已完成（包括已核销）
	OrderNo          int64  `gorm:"order_no" json:"order_no"`                     // 排序号，越小越靠前
	CreateAt         string `gorm:"create_at" json:"create_at"`                   // 创建时间
	UpdateAt         string `gorm:"update_at" json:"update_at"`                   // 修改时间
	Status           int64  `gorm:"status" json:"status"`                         // 状态：0 正常 1 删除
}

func (scm_contract) TableName() string {
	return "scm_contract"
}
