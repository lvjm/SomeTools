package model

type scm_arrival_bill struct {
	Id          int    `gorm:"id" json:"id"`                     // 主键id
	Cid         int    `gorm:"cid" json:"cid"`                   // 所属企业
	BillNum     string `gorm:"bill_num" json:"bill_num"`         // 到货单号
	VendorId    int    `gorm:"vendor_id" json:"vendor_id"`       // 供应商id,可以为空
	Uid         int    `gorm:"uid" json:"uid"`                   // 操作人员
	ContractId  int    `gorm:"contract_id" json:"contract_id"`   // 合同id，可以不关联 为空
	BillStatus  int    `gorm:"bill_status" json:"bill_status"`   // 单据状态,0:已创建，1:已提交，2:已审核，3:审核拒绝，4:到货 ,5:入库，6完成
	QcStatus    int    `gorm:"qc_status" json:"qc_status"`       // 质检标志：0 未质检  1 已质检
	QcInspector int    `gorm:"qc_inspector" json:"qc_inspector"` // 质检操作人
	Remark      string `gorm:"remark" json:"remark"`             // 备注
	CreateAt    string `gorm:"create_at" json:"create_at"`       // 创建时间
	UpdateAt    string `gorm:"update_at" json:"update_at"`       // 修改时间
	Status      int    `gorm:"status" json:"status"`             // 状态：0 正常  1 停用 9 删除
}

func (scm_arrival_bill) TableName() string {
	return "scm_arrival_bill"
}

type scm_contract_policy_performance struct {
	Id               int    `gorm:"id" json:"id"`                                 // 主键id
	Cid              int    `gorm:"cid" json:"cid"`                               // 所属采购企业
	ContractPolicyId int    `gorm:"contract_policy_id" json:"contract_policy_id"` // 采购合同政策id
	Uid              int    `gorm:"uid" json:"uid"`                               // 操作人
	Comment          string `gorm:"comment" json:"comment"`                       // 内容
	Remark           string `gorm:"remark" json:"remark"`                         // 备注
	CreateAt         string `gorm:"create_at" json:"create_at"`                   // 创建时间
	UpdateAt         string `gorm:"update_at" json:"update_at"`                   // 修改时间
	Status           int    `gorm:"status" json:"status"`                         // 状态：0 正常  1 停用 9 删除
}

func (scm_contract_policy_performance) TableName() string {
	return "scm_contract_policy_performance"
}

type scm_contract_sku struct {
	Id             int    `gorm:"id" json:"id"`                             // 主键id
	ContractId     int    `gorm:"contract_id" json:"contract_id"`           // 合同id
	SalesCorpId    int    `gorm:"sales_corp_id" json:"sales_corp_id"`       // 销售方企业id
	SalesSkuId     int    `gorm:"sales_sku_id" json:"sales_sku_id"`         // 销售方sku
	PurchaseCorpId int    `gorm:"purchase_corp_id" json:"purchase_corp_id"` // 采购方企业id
	PurchaseSkuId  int    `gorm:"purchase_sku_id" json:"purchase_sku_id"`   // 采购方sku
	Price          int    `gorm:"price" json:"price"`                       // 单价，单位分
	Num            int    `gorm:"num" json:"num"`                           // 数量
	Remark         string `gorm:"remark" json:"remark"`                     // 备注
	CreateAt       string `gorm:"create_at" json:"create_at"`               // 创建时间
	UpdateAt       string `gorm:"update_at" json:"update_at"`               // 修改时间
	Status         int    `gorm:"status" json:"status"`                     // 状态：0 正常  1 停用 9 删除
}

func (scm_contract_sku) TableName() string {
	return "scm_contract_sku"
}

type scm_outbound_order struct {
	Id          int    `gorm:"id" json:"id"`                     // 主键id
	Cid         int    `gorm:"cid" json:"cid"`                   // 所属企业id
	WarehouseId int    `gorm:"warehouse_id" json:"warehouse_id"` // 所属仓库
	BillNum     string `gorm:"bill_num" json:"bill_num"`         // 出库单号
	Uid         int    `gorm:"uid" json:"uid"`                   // 操作人员
	BillType    int    `gorm:"bill_type" json:"bill_type"`       // 出库类型，1 调拨出库，2 销售出库，3 直接出库，4 盘点出库，5 报损出库，6 其他出库
	BillStatus  int    `gorm:"bill_status" json:"bill_status"`   // 单据状态，0:已创建 ，1:已提交，2:已审核，3:审核拒绝，4:出库，5完成
	Remark      string `gorm:"remark" json:"remark"`             // 备注
	CreateAt    string `gorm:"create_at" json:"create_at"`       // 创建时间
	UpdateAt    string `gorm:"update_at" json:"update_at"`       // 修改时间
	Status      int    `gorm:"status" json:"status"`             // 状态：0 正常  1 停用 9 删除
}

func (scm_outbound_order) TableName() string {
	return "scm_outbound_order"
}

type scm_packaging struct {
	Id            int    `gorm:"id" json:"id"`                         // 主键id
	Cid           int    `gorm:"cid" json:"cid"`                       // 所属企业id
	ModelId       int    `gorm:"model_id" json:"model_id"`             // 所属型号id
	PackagingUnit string `gorm:"packaging_unit" json:"packaging_unit"` // 包装单位
	PackagingNum  int    `gorm:"packaging_num" json:"packaging_num"`   // 包装数量
	Weight        int    `gorm:"weight" json:"weight"`                 // 重量
	WeightUnit    int    `gorm:"weight_unit" json:"weight_unit"`       // 重量单位 0:克，1:千克，2:吨
	Length        int    `gorm:"length" json:"length"`                 // 长
	Width         int    `gorm:"width" json:"width"`                   // 宽
	Height        int    `gorm:"height" json:"height"`                 // 高
	LengthUnit    int    `gorm:"length_unit" json:"length_unit"`       // 长度单位 0:厘米，1:米
	Remark        string `gorm:"remark" json:"remark"`                 // 备注
	CreateAt      string `gorm:"create_at" json:"create_at"`           // 创建时间
	UpdateAt      string `gorm:"update_at" json:"update_at"`           // 修改时间
	Status        int    `gorm:"status" json:"status"`                 // 状态：0 正常  1 停用 9 删除
}

func (scm_packaging) TableName() string {
	return "scm_packaging"
}

type scm_sales_order_detail struct {
	Id           int    `gorm:"id" json:"id"`                         // 主键id
	SalesOrderId int    `gorm:"sales_order_id" json:"sales_order_id"` // 订单主表id
	PackagingId  int    `gorm:"packaging_id" json:"packaging_id"`     // 产品sku
	SkuAlias     string `gorm:"sku_alias" json:"sku_alias"`           // 产品sku别名，默认为空
	Price        int    `gorm:"price" json:"price"`                   // 产品价格
	Num          int    `gorm:"num" json:"num"`                       // 产品数量
	Discount     string `gorm:"discount" json:"discount"`             // 折扣信息
	Subamount    int    `gorm:"subamount" json:"subamount"`           // 小计，单位分
	Remark       string `gorm:"remark" json:"remark"`                 // 备注
	CreateAt     string `gorm:"create_at" json:"create_at"`           // 创建时间
	UpdateAt     string `gorm:"update_at" json:"update_at"`           // 修改时间
	Status       int    `gorm:"status" json:"status"`                 // 状态：0 正常  1 停用 9 删除
}

func (scm_sales_order_detail) TableName() string {
	return "scm_sales_order_detail"
}

type scm_specs struct {
	Id         int    `gorm:"id" json:"id"`                   // 主键id
	Cid        int    `gorm:"cid" json:"cid"`                 // 所属企业id
	ModelId    int    `gorm:"model_id" json:"model_id"`       // 所属型号id
	Sequence   int    `gorm:"sequence" json:"sequence"`       // 排列顺序，数值从小到大排列
	Name       string `gorm:"name" json:"name"`               // 规格类型名称
	ValueUnits string `gorm:"value_units" json:"value_units"` // 规格值单位
	ValueType  int    `gorm:"value_type" json:"value_type"`   // 规格值类型,文本: 0,数值:1
	Remark     string `gorm:"remark" json:"remark"`           // 备注
	CreateAt   string `gorm:"create_at" json:"create_at"`     // 创建时间
	UpdateAt   string `gorm:"update_at" json:"update_at"`     // 修改时间
	Status     int    `gorm:"status" json:"status"`           // 状态：0 正常 1 停用 9 删除
}

func (scm_specs) TableName() string {
	return "scm_specs"
}

type scm_corp_info struct {
	Id          int    `gorm:"id" json:"id"`                     // 主键id
	Cid         int    `gorm:"cid" json:"cid"`                   // 企业唯一id
	Name        string `gorm:"name" json:"name"`                 // 企业名称
	Trade       int    `gorm:"trade" json:"trade"`               // 行业类别
	Type        int    `gorm:"type" json:"type"`                 // 企业性质；国营：1，民营：2 合资：3 外资：4
	Scale       int    `gorm:"scale" json:"scale"`               // 企业规模： 20人以下：1，100人以下：2，500人以下：3，2000人以下：4，5000人以下：5，10000人以下：6，10000人以上：7
	Trn         string `gorm:"trn" json:"trn"`                   // 纳税人识别号
	Address     string `gorm:"address" json:"address"`           // 地址
	BankName    string `gorm:"bank_name" json:"bank_name"`       // 开户行
	BankAccount string `gorm:"bank_account" json:"bank_account"` // 银行账号
	Contacts    string `gorm:"contacts" json:"contacts"`         // 联系人
	Mobile      string `gorm:"mobile" json:"mobile"`             // 联系人手机
	Email       string `gorm:"email" json:"email"`               // 联系人电子邮箱
	Telephone   string `gorm:"telephone" json:"telephone"`       // 座机
	Remark      string `gorm:"remark" json:"remark"`             // 备注
	CreateAt    string `gorm:"create_at" json:"create_at"`       // 创建时间
	UpdateAt    string `gorm:"update_at" json:"update_at"`       // 修改时间
	Status      int    `gorm:"status" json:"status"`             // 状态：0 正常  1 停用 9 删除
}

func (scm_corp_info) TableName() string {
	return "scm_corp_info"
}

type scm_purchase_requestion_detail struct {
	Id                   int    `gorm:"id" json:"id"`                                         // 主键id
	PurchaseRequestionId int    `gorm:"purchase_requestion_id" json:"purchase_requestion_id"` // 订单主表id
	PackagingId          int    `gorm:"packaging_id" json:"packaging_id"`                     // 产品sku
	Name                 string `gorm:"name" json:"name"`                                     // 请购商品名称
	Num                  int    `gorm:"num" json:"num"`                                       // 产品数量
	Remark               string `gorm:"remark" json:"remark"`                                 // 备注
	CreateAt             string `gorm:"create_at" json:"create_at"`                           // 创建时间
	UpdateAt             string `gorm:"update_at" json:"update_at"`                           // 修改时间
	Status               int    `gorm:"status" json:"status"`                                 // 状态：0 正常  1 停用 9 删除
}

func (scm_purchase_requestion_detail) TableName() string {
	return "scm_purchase_requestion_detail"
}

type scm_vendor struct {
	Id       int    `gorm:"id" json:"id"`               // 主键id
	Cid      int    `gorm:"cid" json:"cid"`             // 所属企业id
	VendorId int    `gorm:"vendor_id" json:"vendor_id"` // 供应商企业id
	Remark   string `gorm:"remark" json:"remark"`       // 备注
	CreateAt string `gorm:"create_at" json:"create_at"` // 创建时间
	UpdateAt string `gorm:"update_at" json:"update_at"` // 修改时间
	Status   int    `gorm:"status" json:"status"`       // 状态：0 正常  1 停用 9 删除
}

func (scm_vendor) TableName() string {
	return "scm_vendor"
}

type scm_warehouse struct {
	Id                  int    `gorm:"id" json:"id"`                                       // 主键id
	Cid                 int    `gorm:"cid" json:"cid"`                                     // 所属企业id
	WarehouseCategoryId int    `gorm:"warehouse_category_id" json:"warehouse_category_id"` // 所属企业id类别id
	Name                string `gorm:"name" json:"name"`                                   // 仓库名称
	Address             string `gorm:"address" json:"address"`                             // 地址
	Longitude           string `gorm:"longitude" json:"longitude"`                         // 经度,转换
	Latitude            string `gorm:"latitude" json:"latitude"`                           // 纬度，转换
	ContactsId          int    `gorm:"contacts_id" json:"contacts_id"`                     // 联系人id
	ContactsName        string `gorm:"contacts_name" json:"contacts_name"`                 // 联系人名称
	PhoneNumber         string `gorm:"phone_number" json:"phone_number"`                   // 联系电话
	Remark              string `gorm:"remark" json:"remark"`                               // 备注
	CreateAt            string `gorm:"create_at" json:"create_at"`                         // 创建时间
	UpdateAt            string `gorm:"update_at" json:"update_at"`                         // 修改时间
	UsedStatus          int    `gorm:"used_status" json:"used_status"`                     // 状态：0 未使用  1 已使用,当该仓库在关联库存等数据后就默认为被使用，之后删除操作无法执行，只能停用
	Status              int    `gorm:"status" json:"status"`                               // 状态：0 正常 1 停用 9 删除
}

func (scm_warehouse) TableName() string {
	return "scm_warehouse"
}

type scm_batch_number_format struct {
	Id       int    `gorm:"id" json:"id"`               // 主键id
	Value    string `gorm:"value" json:"value"`         // 日期格式: 如 yyyy年mm月dd日、yyyymmdd
	Remark   string `gorm:"remark" json:"remark"`       // 备注
	CreateAt string `gorm:"create_at" json:"create_at"` // 创建时间
	UpdateAt string `gorm:"update_at" json:"update_at"` // 修改时间
	Status   int    `gorm:"status" json:"status"`       // 状态：0 正常  1 停用 9 删除
}

func (scm_batch_number_format) TableName() string {
	return "scm_batch_number_format"
}

type scm_model struct {
	Id       int    `gorm:"id" json:"id"`               // 主键id
	Cid      int    `gorm:"cid" json:"cid"`             // 所属企业id
	GoodsId  int    `gorm:"goods_id" json:"goods_id"`   // 所属产品id
	Sequence int    `gorm:"sequence" json:"sequence"`   // 排列顺序，数值从小到大排列
	Name     string `gorm:"name" json:"name"`           // 名称
	Comment  string `gorm:"comment" json:"comment"`     // 说明
	Notice   string `gorm:"notice" json:"notice"`       // 特别提醒
	Intro    string `gorm:"intro" json:"intro"`         // 简述
	ImgUrls  string `gorm:"img_urls" json:"img_urls"`   // 描述图片，可多个
	Detail   string `gorm:"detail" json:"detail"`       // 详述
	ExtInfo  string `gorm:"ext_info" json:"ext_info"`   // 型号参数
	Remark   string `gorm:"remark" json:"remark"`       // 备注
	CreateAt string `gorm:"create_at" json:"create_at"` // 创建时间
	UpdateAt string `gorm:"update_at" json:"update_at"` // 修改时间
	Status   int    `gorm:"status" json:"status"`       // 状态：0 正常  1 停用 9 删除
}

func (scm_model) TableName() string {
	return "scm_model"
}

type scm_contract_goods struct {
	Id         int    `gorm:"id" json:"id"`                     // 主键id
	ContractId int    `gorm:"contract_id" json:"contract_id"`   // 合同id
	SysGoodsId int    `gorm:"sys_goods_id" json:"sys_goods_id"` // 系统商品id
	Price      int    `gorm:"price" json:"price"`               // 单价，单位分
	Remark     string `gorm:"remark" json:"remark"`             // 备注
	CreateAt   string `gorm:"create_at" json:"create_at"`       // 创建时间
	UpdateAt   string `gorm:"update_at" json:"update_at"`       // 修改时间
	Status     int    `gorm:"status" json:"status"`             // 状态：0 正常  1 停用 9 删除
}

func (scm_contract_goods) TableName() string {
	return "scm_contract_goods"
}

type scm_dispatch_bill struct {
	Id                  int    `gorm:"id" json:"id"`                                       // 主键id
	Cid                 int    `gorm:"cid" json:"cid"`                                     // 所属企业
	BillNum             string `gorm:"bill_num" json:"bill_num"`                           // 发货单号
	BcId                int    `gorm:"bc_id" json:"bc_id"`                                 // 商业客户id
	Uid                 int    `gorm:"uid" json:"uid"`                                     // 操作人员
	WarehouseId         int    `gorm:"warehouse_id" json:"warehouse_id"`                   // 发货仓库
	ContractId          int    `gorm:"contract_id" json:"contract_id"`                     // 合同id，可以不关联 为空
	LogisticsTrackingId int    `gorm:"logistics_tracking_id" json:"logistics_tracking_id"` // 物流跟踪信息表
	DispatchWay         int    `gorm:"dispatch_way" json:"dispatch_way"`                   // 发货方式: 1 自提，2 送货
	DispatchTime        string `gorm:"dispatch_time" json:"dispatch_time"`                 // 发货时间
	DispatchAddress     string `gorm:"dispatch_address" json:"dispatch_address"`           // 发货地址
	Receiver            string `gorm:"receiver" json:"receiver"`                           // 客户联系人
	Phone               string `gorm:"phone" json:"phone"`                                 // 联系电话
	ShippingAddress     string `gorm:"shipping_address" json:"shipping_address"`           // 收货地址
	BillStatus          int    `gorm:"bill_status" json:"bill_status"`                     // 单据状态,0:已创建，1:已提交，2:已审核，3:审核拒绝，4:发货 ,5:出库，6完成
	QcStatus            int    `gorm:"qc_status" json:"qc_status"`                         // 质检标志：0 未质检  1 已质检
	QcInspector         int    `gorm:"qc_inspector" json:"qc_inspector"`                   // 质检操作人
	Remark              string `gorm:"remark" json:"remark"`                               // 备注
	CreateAt            string `gorm:"create_at" json:"create_at"`                         // 创建时间
	UpdateAt            string `gorm:"update_at" json:"update_at"`                         // 修改时间
	Status              int    `gorm:"status" json:"status"`                               // 状态：0 正常  1 停用 9 删除
}

func (scm_dispatch_bill) TableName() string {
	return "scm_dispatch_bill"
}

type scm_order_reference struct {
	Id           int    `gorm:"id" json:"id"`                         // 主键id
	Cid          int    `gorm:"cid" json:"cid"`                       // 所属企业
	OrderId      int    `gorm:"order_id" json:"order_id"`             // 单据id
	RefOrderId   int    `gorm:"ref_order_id" json:"ref_order_id"`     // 引用单据id
	RefOrderType int    `gorm:"ref_order_type" json:"ref_order_type"` // 引用单据类型: 10 请购单，21 采购单，22 销售单， 31 付款单，32 收款单，41 到货单，42 发货单，51 入库单，52 出库单
	Remark       string `gorm:"remark" json:"remark"`                 // 备注
	CreateAt     string `gorm:"create_at" json:"create_at"`           // 创建时间
	UpdateAt     string `gorm:"update_at" json:"update_at"`           // 修改时间
	Status       int    `gorm:"status" json:"status"`                 // 状态：0 正常  1 停用 9 删除
}

func (scm_order_reference) TableName() string {
	return "scm_order_reference"
}

type scm_specs_detail struct {
	Id       int    `gorm:"id" json:"id"`               // 主键id
	Cid      int    `gorm:"cid" json:"cid"`             // 所属企业id
	ModelId  int    `gorm:"model_id" json:"model_id"`   // 所属型号id
	SpecsId  int    `gorm:"specs_id" json:"specs_id"`   // 所属规格id
	Value    string `gorm:"value" json:"value"`         // 规格值
	ImgUrls  string `gorm:"img_urls" json:"img_urls"`   // 规格图片，可多个
	Remark   string `gorm:"remark" json:"remark"`       // 备注
	CreateAt string `gorm:"create_at" json:"create_at"` // 创建时间
	UpdateAt string `gorm:"update_at" json:"update_at"` // 修改时间
	Status   int    `gorm:"status" json:"status"`       // 状态：0 正常 1 停用 9 删除
}

func (scm_specs_detail) TableName() string {
	return "scm_specs_detail"
}

type scm_business_customer struct {
	Id             int    `gorm:"id" json:"id"`                           // 主键id
	Cid            int    `gorm:"cid" json:"cid"`                         // 所属企业id
	BcId           int    `gorm:"bc_id" json:"bc_id"`                     // 商业客户
	FollowupStatus int    `gorm:"followup_status" json:"followup_status"` // 跟进状态：0 否  1 是
	Remark         string `gorm:"remark" json:"remark"`                   // 备注
	CreateAt       string `gorm:"create_at" json:"create_at"`             // 创建时间
	UpdateAt       string `gorm:"update_at" json:"update_at"`             // 修改时间
	Status         int    `gorm:"status" json:"status"`                   // 状态：0 正常  1 停用 9 删除
}

func (scm_business_customer) TableName() string {
	return "scm_business_customer"
}

type scm_logistics_tracking_detail struct {
	Id                  int    `gorm:"id" json:"id"`                                       // 主键id
	Cid                 int    `gorm:"cid" json:"cid"`                                     // 所属企业
	LogisticsTrackingId int    `gorm:"logistics_tracking_id" json:"logistics_tracking_id"` // 物流跟踪信息表id
	OccurrenceTime      string `gorm:"occurrence_time" json:"occurrence_time"`             // 发生`时间
	OccurrenceInfo      string `gorm:"occurrence_info" json:"occurrence_info"`             // 发生事件
	Remark              string `gorm:"remark" json:"remark"`                               // 备注
	CreateAt            string `gorm:"create_at" json:"create_at"`                         // 创建时间
	UpdateAt            string `gorm:"update_at" json:"update_at"`                         // 修改时间
	Status              int    `gorm:"status" json:"status"`                               // 状态：0 正常  1 停用 9 删除
}

func (scm_logistics_tracking_detail) TableName() string {
	return "scm_logistics_tracking_detail"
}

type scm_pay_type struct {
	Id       int    `gorm:"id" json:"id"`               // 主键id
	Name     string `gorm:"name" json:"name"`           // 支付方式名称： 1 现金，2 银行转账，3 微信，4 支付宝，5 网银，6 现金支票，7 银行支票，8 电汇
	Value    int    `gorm:"value" json:"value"`         // 编码值
	Remark   string `gorm:"remark" json:"remark"`       // 备注
	CreateAt string `gorm:"create_at" json:"create_at"` // 创建时间
	UpdateAt string `gorm:"update_at" json:"update_at"` // 修改时间
	Status   int    `gorm:"status" json:"status"`       // 状态：0 正常  1 停用 9 删除
}

func (scm_pay_type) TableName() string {
	return "scm_pay_type"
}

type scm_purchase_order struct {
	Id           int    `gorm:"id" json:"id"`                       // 主键id
	Cid          int    `gorm:"cid" json:"cid"`                     // 所属企业
	OrderNum     string `gorm:"order_num" json:"order_num"`         // 订单流水号
	VendorId     int    `gorm:"vendor_id" json:"vendor_id"`         // 供应商id
	Uid          int    `gorm:"uid" json:"uid"`                     // 采购人员
	Amount       int    `gorm:"amount" json:"amount"`               // 总金额，单位分
	PayType      int    `gorm:"pay_type" json:"pay_type"`           // 支付方式： 1 现金，2 银行转账，3 微信，4 支付宝，5 网银，6 现金支票，7 银行支票，8 电汇
	OrderStatus  int    `gorm:"order_status" json:"order_status"`   // 订单状态,0:已创建 ，1:已提交，2:已审核，3:审核拒绝，4:发货，5:入库，6:付款，7完成
	PurchaseTime string `gorm:"purchase_time" json:"purchase_time"` // 采购时间
	ContractId   int    `gorm:"contract_id" json:"contract_id"`     // 采购合同id，可以不关联 为空
	Remark       string `gorm:"remark" json:"remark"`               // 备注
	CreateAt     string `gorm:"create_at" json:"create_at"`         // 创建时间
	UpdateAt     string `gorm:"update_at" json:"update_at"`         // 修改时间
	Status       int    `gorm:"status" json:"status"`               // 状态：0 正常  1 停用 9 删除
}

func (scm_purchase_order) TableName() string {
	return "scm_purchase_order"
}

type scm_sku struct {
	Id             int    `gorm:"id" json:"id"`                             // 主键id
	Cid            int    `gorm:"cid" json:"cid"`                           // 所属企业id
	GoodsId        int    `gorm:"goods_id" json:"goods_id"`                 // 所属产品id
	ModelId        int    `gorm:"model_id" json:"model_id"`                 // 所属型号id
	SpecsDetailIds string `gorm:"specs_detail_ids" json:"specs_detail_ids"` // 所属规格详情ids,存的顺序按照规格的顺序排列
	PackagingId    int    `gorm:"packaging_id" json:"packaging_id"`         // 所属包装规格id
	Barcode        string `gorm:"barcode" json:"barcode"`                   // 统一条码
	SkuBarcode     string `gorm:"sku_barcode" json:"sku_barcode"`           // 内部sku条码,格式: productId["颜色":"红色","内存":"128G"]
	Intro          string `gorm:"intro" json:"intro"`                       // sku描述,格式: 商品名、型号名规格包装规格描述
	MaxAlert       int    `gorm:"max_alert" json:"max_alert"`               // 库存上限预警
	MinAlert       int    `gorm:"min_alert" json:"min_alert"`               // 库存下限预警
	Remark         string `gorm:"remark" json:"remark"`                     // 备注
	CreateAt       string `gorm:"create_at" json:"create_at"`               // 创建时间
	UpdateAt       string `gorm:"update_at" json:"update_at"`               // 修改时间
	Status         int    `gorm:"status" json:"status"`                     // 状态：0 正常  1 停用 9 删除
}

func (scm_sku) TableName() string {
	return "scm_sku"
}

type scm_stock struct {
	Id             int    `gorm:"id" json:"id"`                           // 主键id
	Cid            int    `gorm:"cid" json:"cid"`                         // 所属企业id
	WarehouseId    int    `gorm:"warehouse_id" json:"warehouse_id"`       // 所属仓库
	SkuId          int    `gorm:"sku_id" json:"sku_id"`                   // skuid
	Barcode        string `gorm:"barcode" json:"barcode"`                 // sku统一条码
	SkuBarcode     string `gorm:"sku_barcode" json:"sku_barcode"`         // sku条码,格式: productId:["颜色":"红色","内存":"128G"]
	Unit           string `gorm:"unit" json:"unit"`                       // 单位
	BatchNumber    string `gorm:"batch_number" json:"batch_number"`       // 生产批号
	ProductionDate string `gorm:"production_date" json:"production_date"` // 生产日期
	ExpiriesDate   string `gorm:"expiries_date" json:"expiries_date"`     // 失效日期
	ExtInfo        string `gorm:"ext_info" json:"ext_info"`               // 额外参数
	CurrentStock   int    `gorm:"current_stock" json:"current_stock"`     // 当前库存
	AvailableStock int    `gorm:"available_stock" json:"available_stock"` // 可用库存
	OutgoingStock  int    `gorm:"outgoing_stock" json:"outgoing_stock"`   // 待出库数
	PendingStock   int    `gorm:"pending_stock" json:"pending_stock"`     // 待入库数
	MaxAlert       int    `gorm:"max_alert" json:"max_alert"`             // 库存上限预警
	MinAlert       int    `gorm:"min_alert" json:"min_alert"`             // 库存下限预警
	Remark         string `gorm:"remark" json:"remark"`                   // 备注
	CreateAt       string `gorm:"create_at" json:"create_at"`             // 创建时间
	UpdateAt       string `gorm:"update_at" json:"update_at"`             // 修改时间
	Status         int    `gorm:"status" json:"status"`                   // 状态：0 正常  1 停用 9 删除
}

func (scm_stock) TableName() string {
	return "scm_stock"
}

type scm_warehouse_category struct {
	Id       int    `gorm:"id" json:"id"`               // 主键id
	Cid      int    `gorm:"cid" json:"cid"`             // 所属企业id
	Name     string `gorm:"name" json:"name"`           // 类别名称
	ParentId int    `gorm:"parent_id" json:"parent_id"` // 父类别
	Level    int    `gorm:"level" json:"level"`         // 层级，越大层级越深，顶级为一级：1
	Creator  int    `gorm:"creator" json:"creator"`     // 创建者的uid
	Remark   string `gorm:"remark" json:"remark"`       // 备注
	CreateAt string `gorm:"create_at" json:"create_at"` // 创建时间
	UpdateAt string `gorm:"update_at" json:"update_at"` // 修改时间
	Status   int    `gorm:"status" json:"status"`       // 状态：0 正常  1 停用 9 删除
}

func (scm_warehouse_category) TableName() string {
	return "scm_warehouse_category"
}

type scm_changelog struct {
	Id              int    `gorm:"id" json:"id"`                             // 主键id
	Cid             int    `gorm:"cid" json:"cid"`                           // 所属企业id
	TableName       string `gorm:"table_name" json:"table_name"`             // 修改记录的表名
	RecordId        int    `gorm:"record_id" json:"record_id"`               // 记录id
	ChangeBefore    string `gorm:"change_before" json:"change_before"`       // 修改之前的记录，格式为[{"column_name":"name",value":"zhangsan"},{"column_name":"give_name",value":"zhangsana"}]
	ChangeAfter     string `gorm:"change_after" json:"change_after"`         // 修改之后的记录,格式为[{"column_name":"name",value":"zhangsan"},{"column_name":"give_name",value":"zhangsana"}]
	EffectiveStatus int    `gorm:"effective_status" json:"effective_status"` // 生效状态：0 未生效  1 已生效
	Remark          string `gorm:"remark" json:"remark"`                     // 备注
	CreateAt        string `gorm:"create_at" json:"create_at"`               // 创建时间
	UpdateAt        string `gorm:"update_at" json:"update_at"`               // 修改时间
	Status          int    `gorm:"status" json:"status"`                     // 状态：0 正常  1 停用 9 删除
}

func (scm_changelog) Tablename() string {
	return "scm_changelog"
}

type scm_contract struct {
	Id               int    `gorm:"id" json:"id"`                                 // 主键id
	TaskId           int    `gorm:"task_id" json:"task_id"`                       // 对应的子任务id
	Cid              int    `gorm:"cid" json:"cid"`                               // 所属公司id
	Uid              int    `gorm:"uid" json:"uid"`                               // 处理人id
	Title            string `gorm:"title" json:"title"`                           // 表单的名称
	SalesCorpId      int    `gorm:"sales_corp_id" json:"sales_corp_id"`           // 销售方企业id
	SalesCorpName    string `gorm:"sales_corp_name" json:"sales_corp_name"`       // 销售方企业名称
	PurchaseCorpId   int    `gorm:"purchase_corp_id" json:"purchase_corp_id"`     // 采购方企业id
	PurchaseCorpName string `gorm:"purchase_corp_name" json:"purchase_corp_name"` // 采购方企业名称
	Code             string `gorm:"code" json:"code"`                             // 合同编码
	StartDate        string `gorm:"start_date" json:"start_date"`                 // 开始日期
	EndDate          string `gorm:"end_date" json:"end_date"`                     // 结束日期
	RemindDate       string `gorm:"remind_date" json:"remind_date"`               // 提醒日期,默认提前三个月
	Amount           int    `gorm:"amount" json:"amount"`                         // 总金额，单位分
	SettlementType   int    `gorm:"settlement_type" json:"settlement_type"`       // 结算方式 1:按订单结算 2:按入库结算 3:按预付款结算 4:按实销月结 5:按发货月结 6:按发货几个月后结算 7:按铺货多少金额后几个月后结算
	SettlementValue  string `gorm:"settlement_value" json:"settlement_value"`     // 结算方式值，月数，或金额(单位分) 形如[{"code":6,"month":3},{"code":7,"month":3,"amount":"23423423"}]
	PayType          int    `gorm:"pay_type" json:"pay_type"`                     // 支付方式： 1 现金，2 银行转账，3 微信，4 支付宝，5 网银，6 现金支票，7 银行支票，8 电汇
	Voucher          string `gorm:"voucher" json:"voucher"`                       // 凭证，文件链接，string数组
	Remark           string `gorm:"remark" json:"remark"`                         // 备注
	IsSplit          int    `gorm:"is_split" json:"is_split"`                     // 是否表单拆分而来：0 否，1 是
	MergedTables     string `gorm:"merged_tables" json:"merged_tables"`           // 数据来源表单，多对一就合并[{"task_id":1,"form_id":99},{"task_id":1,"form_id":99}]，一对一时填单元素数组,没有为[]
	WriteOffTables   string `gorm:"write_off_tables" json:"write_off_tables"`     // 被核销的表单的 formid [{"task_id":1,"form_id":99},{"task_id":1,"form_id":99}]，没有就为[]
	RelatedTables    string `gorm:"related_tables" json:"related_tables"`         // 被关联的表单的 formid [{"task_id":1,"form_id":99},{"task_id":1,"form_id":99}]，没有就为[]
	BusiStatus       int    `gorm:"busi_status" json:"busi_status"`               // 业务状态：-1 审核不过 0 提交待审核 1 审核中 2 已通过 3/4/5(业务决定)...
	DeriveStatus     int    `gorm:"derive_status" json:"derive_status"`           // 数据流转状态：0 已生成，1 已流转，90 已核销，100 已完成（包括已核销）
	OrderNo          int    `gorm:"order_no" json:"order_no"`                     // 排序号，越小越靠前
	CreateAt         string `gorm:"create_at" json:"create_at"`                   // 创建时间
	UpdateAt         string `gorm:"update_at" json:"update_at"`                   // 修改时间
	Status           int    `gorm:"status" json:"status"`                         // 状态：0 正常 1 删除
}

func (scm_contract) TableName() string {
	return "scm_contract"
}

type scm_contract_policy struct {
	Id              int    `gorm:"id" json:"id"`                             // 主键id
	ContractId      int    `gorm:"contract_id" json:"contract_id"`           // 合同id
	Title           string `gorm:"title" json:"title"`                       // 标题
	Content         string `gorm:"content" json:"content"`                   // 条款
	IsAllGoods      int    `gorm:"is_all_goods" json:"is_all_goods"`         // 是否针对所有商品：0 否  1 是
	SysGoods        string `gorm:"sys_goods" json:"sys_goods"`               // 针对系统商品列表,如果针对所有商品，这里的商品列表无效 系统商品id [1,2]
	SettlemenTiming int    `gorm:"settlemen_timing" json:"settlemen_timing"` // 结算时间点: 每月计算、每半年计算、每年计算
	PayType         int    `gorm:"pay_type" json:"pay_type"`                 // 支付方式： 1 通过货款抵扣 2 其他费用
	Remark          string `gorm:"remark" json:"remark"`                     // 备注
	CreateAt        string `gorm:"create_at" json:"create_at"`               // 创建时间
	UpdateAt        string `gorm:"update_at" json:"update_at"`               // 修改时间
	Status          int    `gorm:"status" json:"status"`                     // 状态：0 正常  1 停用 9 删除
}

func (scm_contract_policy) TableName() string {
	return "scm_contract_policy"
}

type scm_goods_category struct {
	Id       int    `gorm:"id" json:"id"`               // 主键id
	Cid      int    `gorm:"cid" json:"cid"`             // 所属企业id
	Name     string `gorm:"name" json:"name"`           // 类别名称
	ParentId int    `gorm:"parent_id" json:"parent_id"` // 父类别
	Level    int    `gorm:"level" json:"level"`         // 层级，越大层级越深，顶级为一级：1
	Creator  int    `gorm:"creator" json:"creator"`     // 创建者的uid
	Remark   string `gorm:"remark" json:"remark"`       // 备注
	CreateAt string `gorm:"create_at" json:"create_at"` // 创建时间
	UpdateAt string `gorm:"update_at" json:"update_at"` // 修改时间
	Status   int    `gorm:"status" json:"status"`       // 状态：0 正常  1 停用 9 删除
}

func (scm_goods_category) TableName() string {
	return "scm_goods_category"
}

type scm_purchase_order_detail struct {
	Id              int    `gorm:"id" json:"id"`                               // 主键id
	PurchaseOrderId int    `gorm:"purchase_order_id" json:"purchase_order_id"` // 订单主表id
	PackagingId     int    `gorm:"packaging_id" json:"packaging_id"`           // 产品sku
	Price           int    `gorm:"price" json:"price"`                         // 产品价格
	Num             int    `gorm:"num" json:"num"`                             // 产品数量
	Subamount       int    `gorm:"subamount" json:"subamount"`                 // 小计，单位分
	Remark          string `gorm:"remark" json:"remark"`                       // 备注
	CreateAt        string `gorm:"create_at" json:"create_at"`                 // 创建时间
	UpdateAt        string `gorm:"update_at" json:"update_at"`                 // 修改时间
	Status          int    `gorm:"status" json:"status"`                       // 状态：0 正常  1 停用 9 删除
}

func (scm_purchase_order_detail) TableName() string {
	return "scm_purchase_order_detail"
}

type scm_purchase_requestion struct {
	Id          int    `gorm:"id" json:"id"`                     // 主键id
	Cid         int    `gorm:"cid" json:"cid"`                   // 所属企业
	OrderNum    string `gorm:"order_num" json:"order_num"`       // 订单流水号
	Uid         int    `gorm:"uid" json:"uid"`                   // 请购人员
	Comment     string `gorm:"comment" json:"comment"`           // 请购原因
	OrderStatus int    `gorm:"order_status" json:"order_status"` // 订单状态,0:已创建 ，1:已提交，2:已审核，3:审核拒绝，4:采购，4:到货，5:入库，6:付款，7完成
	Datetime    string `gorm:"datetime" json:"datetime"`         // 销售时间
	Remark      string `gorm:"remark" json:"remark"`             // 备注
	CreateAt    string `gorm:"create_at" json:"create_at"`       // 创建时间
	UpdateAt    string `gorm:"update_at" json:"update_at"`       // 修改时间
	Status      int    `gorm:"status" json:"status"`             // 状态：0 正常  1 停用 9 删除
}

func (scm_purchase_requestion) TableName() string {
	return "scm_purchase_requestion"
}

type scm_sales_order struct {
	Id          int    `gorm:"id" json:"id"`                     // 主键id
	Cid         int    `gorm:"cid" json:"cid"`                   // 所属企业
	OrderNum    string `gorm:"order_num" json:"order_num"`       // 订单流水号
	BcId        int    `gorm:"bc_id" json:"bc_id"`               // 商业客户id
	Uid         int    `gorm:"uid" json:"uid"`                   // 销售人员
	Discount    string `gorm:"discount" json:"discount"`         // 折扣信息
	Amount      int    `gorm:"amount" json:"amount"`             // 总金额，单位分
	PayType     int    `gorm:"pay_type" json:"pay_type"`         // 支付方式： 1 现金，2 银行转账，3 微信，4 支付宝，5 网银，6 现金支票，7 银行支票，8 电汇
	OrderStatus int    `gorm:"order_status" json:"order_status"` // 订单状态,0:已创建 ，1:已提交，2:已审核，3:审核拒绝，4:发货，5:出库，6:收款，7完成
	SalesTime   string `gorm:"sales_time" json:"sales_time"`     // 销售时间
	ContractId  int    `gorm:"contract_id" json:"contract_id"`   // 合同id，可以不关联 为空
	Remark      string `gorm:"remark" json:"remark"`             // 备注
	CreateAt    string `gorm:"create_at" json:"create_at"`       // 创建时间
	UpdateAt    string `gorm:"update_at" json:"update_at"`       // 修改时间
	Status      int    `gorm:"status" json:"status"`             // 状态：0 正常  1 停用 9 删除
}

func (scm_sales_order) TableName() string {
	return "scm_sales_order"
}

type scm_goods struct {
	Id             int    `gorm:"id" json:"id"`                           // 主键id
	Cid            int    `gorm:"cid" json:"cid"`                         // 所属企业id
	CategoryIds    string `gorm:"category_ids" json:"category_ids"`       // 类别ids
	Name           string `gorm:"name" json:"name"`                       // 通用名称
	GivenName      string `gorm:"given_name" json:"given_name"`           // 俗称
	Groups         string `gorm:"groups" json:"groups"`                   // 产品展示分类
	Brand          string `gorm:"brand" json:"brand"`                     // 品牌名称
	IsManufacturer int    `gorm:"is_manufacturer" json:"is_manufacturer"` // 是否为生产商：0 不是  1 是
	Manufacturer   string `gorm:"manufacturer" json:"manufacturer"`       // 如果不是自产 生产厂商
	Madein         string `gorm:"madein" json:"madein"`                   // 产地
	MinUnit        string `gorm:"min_unit" json:"min_unit"`               // 最小单位
	BatchConfig    int    `gorm:"batch_config" json:"batch_config"`       // 是否启用批号管理：0 否 ,1 是
	BatchNumberId  int    `gorm:"batch_number_id" json:"batch_number_id"` // 批号管理id
	Notice         string `gorm:"notice" json:"notice"`                   // 特别提醒
	Intro          string `gorm:"intro" json:"intro"`                     // 简述
	Remark         string `gorm:"remark" json:"remark"`                   // 备注
	CreateAt       string `gorm:"create_at" json:"create_at"`             // 创建时间
	UpdateAt       string `gorm:"update_at" json:"update_at"`             // 修改时间
	UsedStatus     int    `gorm:"used_status" json:"used_status"`         // 状态：0 未使用  1 已使用,当该商品在关联库存等数据后就默认为被使用，之后删除操作无法执行，只能停用
	Status         int    `gorm:"status" json:"status"`                   // 状态：0 正常  1 停用 9 删除
}

func (scm_goods) TableName() string {
	return "scm_goods"
}

type scm_outbound_order_detail struct {
	Id              int    `gorm:"id" json:"id"`                               // 主键id
	OutboundOrderId int    `gorm:"outbound_order_id" json:"outbound_order_id"` // 出库单主表id
	SkuId           int    `gorm:"sku_id" json:"sku_id"`                       // 产品sku
	Num             int    `gorm:"num" json:"num"`                             // 产品数量
	BatchNumber     string `gorm:"batch_number" json:"batch_number"`           // 生产批号
	ProductionDate  string `gorm:"production_date" json:"production_date"`     // 生产日期
	ExpiringDate    string `gorm:"expiring_date" json:"expiring_date"`         // 失效日期
	ExtInfo         string `gorm:"ext_info" json:"ext_info"`                   // 额外参数
	Remark          string `gorm:"remark" json:"remark"`                       // 备注
	CreateAt        string `gorm:"create_at" json:"create_at"`                 // 创建时间
	UpdateAt        string `gorm:"update_at" json:"update_at"`                 // 修改时间
	Status          int    `gorm:"status" json:"status"`                       // 状态：0 正常 1 停用 9 删除
}

func (scm_outbound_order_detail) TableName() string {
	return "scm_outbound_order_detail"
}

type scm_settlemen_term struct {
	Id             int    `gorm:"id" json:"id"`                           // 主键id
	ContractId     int    `gorm:"contract_id" json:"contract_id"`         // 合同id
	PolicyId       int    `gorm:"policy_id" json:"policy_id"`             // 合同商业政策id
	Content        string `gorm:"content" json:"content"`                 // 条款内容：进货超过多少金额，结算支付比例，优惠是多少点；根据购进数量；根据时间点支付金额
	IsAllGoods     int    `gorm:"is_all_goods" json:"is_all_goods"`       // 是否针对所有商品：0 否  1 是
	SysGoods       string `gorm:"sys_goods" json:"sys_goods"`             // 针对系统商品列表 商品id [1,2]
	StartDate      string `gorm:"start_date" json:"start_date"`           // 计算期间的开始日期
	EndData        string `gorm:"end_data" json:"end_data"`               // 计算期间的结束时间
	StatisticsForm int    `gorm:"statistics_form" json:"statistics_form"` // 统计形式：0 按照发货金额  1 按回款金额
	Remark         string `gorm:"remark" json:"remark"`                   // 备注
	CreateAt       string `gorm:"create_at" json:"create_at"`             // 创建时间
	UpdateAt       string `gorm:"update_at" json:"update_at"`             // 修改时间
	Status         int    `gorm:"status" json:"status"`                   // 状态：0 正常  1 停用 9 删除
}

func (scm_settlemen_term) TableName() string {
	return "scm_settlemen_term"
}

type scm_logistics_tracking struct {
	Id               int    `gorm:"id" json:"id"`                               // 主键id
	Cid              int    `gorm:"cid" json:"cid"`                             // 所属企业
	LogisticsCompany string `gorm:"logistics_company" json:"logistics_company"` // 物流公司名称
	LogisticsNumber  string `gorm:"logistics_number" json:"logistics_number"`   // 物流单号
	BcId             int    `gorm:"bc_id" json:"bc_id"`                         // 商业客户id
	Uid              int    `gorm:"uid" json:"uid"`                             // 操作人员
	DispatchAddress  string `gorm:"dispatch_address" json:"dispatch_address"`   // 发货地址
	Receiver         string `gorm:"receiver" json:"receiver"`                   // 客户联系人
	Phone            string `gorm:"phone" json:"phone"`                         // 联系电话
	ShippingAddress  string `gorm:"shipping_address" json:"shipping_address"`   // 收货地址
	Remark           string `gorm:"remark" json:"remark"`                       // 备注
	CreateAt         string `gorm:"create_at" json:"create_at"`                 // 创建时间
	UpdateAt         string `gorm:"update_at" json:"update_at"`                 // 修改时间
	Status           int    `gorm:"status" json:"status"`                       // 状态：0 正常  1 停用 9 删除
}

func (scm_logistics_tracking) TableName() string {
	return "scm_logistics_tracking"
}

type scm_arrival_bill_detail struct {
	Id           int    `gorm:"id" json:"id"`                       // 主键id
	ArrivalBill  int    `gorm:"arrival_bill" json:"arrival_bill"`   // 到货单主表id
	PackagingId  int    `gorm:"packaging_id" json:"packaging_id"`   // 产品sku
	Num          int    `gorm:"num" json:"num"`                     // 产品数量
	QualifiedNum int    `gorm:"qualified_num" json:"qualified_num"` // 合格数量
	Remark       string `gorm:"remark" json:"remark"`               // 备注
	CreateAt     string `gorm:"create_at" json:"create_at"`         // 创建时间
	UpdateAt     string `gorm:"update_at" json:"update_at"`         // 修改时间
	Status       int    `gorm:"status" json:"status"`               // 状态：0 正常 1 停用 9 删除
}

func (scm_arrival_bill_detail) TableName() string {
	return "scm_arrival_bill_detail"
}

type scm_batch_number struct {
	Id                   int    `gorm:"id" json:"id"`                                         // 主键id
	Cid                  int    `gorm:"cid" json:"cid"`                                       // 所属企业id
	ProductionDateFormat string `gorm:"production_date_format" json:"production_date_format"` // 生产日期格式: 如 yyyy年mm月dd日、yyyymmdd
	ProductionDateAlias  string `gorm:"production_date_alias" json:"production_date_alias"`   // 生产日期别名
	ExpiresDateFormat    string `gorm:"expires_date_format" json:"expires_date_format"`       // 失效日期格式: 如 yyyy年mm月dd日、yyyymmdd
	ExpiresDateAlias     string `gorm:"expires_date_alias" json:"expires_date_alias"`         // 失效日期别名
	ValidityYear         int    `gorm:"validity_year" json:"validity_year"`                   // 有效期，年数
	ValidityMonth        int    `gorm:"validity_month" json:"validity_month"`                 // 有效期，月份
	ValidityDay          int    `gorm:"validity_day" json:"validity_day"`                     // 有效期，天数
	ValidityAlias        string `gorm:"validity_alias" json:"validity_alias"`                 // 有效期别名
	DisplayOrder         int    `gorm:"display_order" json:"display_order"`                   // 显示顺序：1:批号、生产日期、失效日期，2:批号、生产日期、失效日期，3:生产日期、批号、失效日期，4:生产日期、失效日期、批号，5:失效日期、批号、失效日期，6:失效日期、失效日期、批号
	Remark               string `gorm:"remark" json:"remark"`                                 // 备注
	CreateAt             string `gorm:"create_at" json:"create_at"`                           // 创建时间
	UpdateAt             string `gorm:"update_at" json:"update_at"`                           // 修改时间
	Status               int    `gorm:"status" json:"status"`                                 // 状态：0 正常  1 停用 9 删除
}

func (scm_batch_number) TableName() string {
	return "scm_batch_number"
}

type scm_bc_contact struct {
	Id        int    `gorm:"id" json:"id"`               // 主键id
	BcId      int    `gorm:"bc_id" json:"bc_id"`         // 商业客户id
	Contacts  string `gorm:"contacts" json:"contacts"`   // 联系人
	Mobile    string `gorm:"mobile" json:"mobile"`       // 联系人手机
	Email     string `gorm:"email" json:"email"`         // 联系人电子邮箱
	Telephone string `gorm:"telephone" json:"telephone"` // 座机
	Remark    string `gorm:"remark" json:"remark"`       // 备注
	CreateAt  string `gorm:"create_at" json:"create_at"` // 创建时间
	UpdateAt  string `gorm:"update_at" json:"update_at"` // 修改时间
	Status    int    `gorm:"status" json:"status"`       // 状态：0 正常  1 停用 9 删除
}

func (scm_bc_contact) TableName() string {
	return "scm_bc_contact"
}

type scm_bc_followup struct {
	Id       int    `gorm:"id" json:"id"`               // 主键id
	Cid      int    `gorm:"cid" json:"cid"`             // 所属企业id
	BcId     int    `gorm:"bc_id" json:"bc_id"`         // 商业客户id
	Uid      int    `gorm:"uid" json:"uid"`             // 跟进人员
	Date     string `gorm:"date" json:"date"`           // 跟进时间
	Content  string `gorm:"content" json:"content"`     // 跟进内容
	NextDate string `gorm:"next_date" json:"next_date"` // 下次跟进时间
	Remark   string `gorm:"remark" json:"remark"`       // 备注
	CreateAt string `gorm:"create_at" json:"create_at"` // 创建时间
	UpdateAt string `gorm:"update_at" json:"update_at"` // 修改时间
	Status   int    `gorm:"status" json:"status"`       // 状态：0 正常  1 停用 9 删除
}

func (scm_bc_followup) TableName() string {
	return "scm_bc_followup"
}

type scm_dispatch_bill_detail struct {
	Id             int    `gorm:"id" json:"id"`                             // 主键id
	DispatchBillId int    `gorm:"dispatch_bill_id" json:"dispatch_bill_id"` // 发货单主表id
	PackagingId    int    `gorm:"packaging_id" json:"packaging_id"`         // 产品sku
	SkuAlias       string `gorm:"sku_alias" json:"sku_alias"`               // 产品sku别名，默认为空
	Num            int    `gorm:"num" json:"num"`                           // 产品数量
	Remark         string `gorm:"remark" json:"remark"`                     // 备注
	CreateAt       string `gorm:"create_at" json:"create_at"`               // 创建时间
	UpdateAt       string `gorm:"update_at" json:"update_at"`               // 修改时间
	Status         int    `gorm:"status" json:"status"`                     // 状态：0 正常 1 停用 9 删除
}

func (scm_dispatch_bill_detail) TableName() string {
	return "scm_dispatch_bill_detail"
}

type scm_inbound_order struct {
	Id          int    `gorm:"id" json:"id"`                     // 主键id
	Cid         int    `gorm:"cid" json:"cid"`                   // 所属企业id
	WarehouseId int    `gorm:"warehouse_id" json:"warehouse_id"` // 所属仓库
	BillNum     string `gorm:"bill_num" json:"bill_num"`         // 入库单号
	Uid         int    `gorm:"uid" json:"uid"`                   // 操作人员
	BillType    int    `gorm:"bill_type" json:"bill_type"`       // 入库类型，1 调拨出库，2 销售出库，3 直接出库，4 盘点出库，5 报损出库，6 其他出库
	BillStatus  int    `gorm:"bill_status" json:"bill_status"`   // 单据状态，0:已创建 ，1:已提交，2:已审核，3:审核拒绝，4:入库，5完成
	Remark      string `gorm:"remark" json:"remark"`             // 备注
	CreateAt    string `gorm:"create_at" json:"create_at"`       // 创建时间
	UpdateAt    string `gorm:"update_at" json:"update_at"`       // 修改时间
	Status      int    `gorm:"status" json:"status"`             // 状态：0 正常  1 停用 9 删除
}

func (scm_inbound_order) TableName() string {
	return "scm_inbound_order"
}

type scm_inbound_order_detail struct {
	Id             int    `gorm:"id" json:"id"`                             // 主键id
	InboundOrderId int    `gorm:"inbound_order_id" json:"inbound_order_id"` // 入库单主表id
	SkuId          int    `gorm:"sku_id" json:"sku_id"`                     // 产品sku
	Num            int    `gorm:"num" json:"num"`                           // 产品数量
	BatchNumber    string `gorm:"batch_number" json:"batch_number"`         // 生产批号
	ProductionDate string `gorm:"production_date" json:"production_date"`   // 生产日期
	ExpiringDate   string `gorm:"expiring_date" json:"expiring_date"`       // 失效日期
	ExtInfo        string `gorm:"ext_info" json:"ext_info"`                 // 额外参数
	Remark         string `gorm:"remark" json:"remark"`                     // 备注
	CreateAt       string `gorm:"create_at" json:"create_at"`               // 创建时间
	UpdateAt       string `gorm:"update_at" json:"update_at"`               // 修改时间
	Status         int    `gorm:"status" json:"status"`                     // 状态：0 正常 1 停用 9 删除
}

func (scm_inbound_order_detail) TableName() string {
	return "scm_inbound_order_detail"
}

type scm_settlement_type struct {
	Id       int    `gorm:"id" json:"id"`               // 主键id
	Name     string `gorm:"name" json:"name"`           // 结算方式名称 1:按订单结算 2:按入库结算 3:按预付款结算 4:按实销月结 5:按发货月结 6:按发货几个月后结算 7:按铺货多少金额后几个月后结算
	Value    int    `gorm:"value" json:"value"`         // 编码值
	Remark   string `gorm:"remark" json:"remark"`       // 备注
	CreateAt string `gorm:"create_at" json:"create_at"` // 创建时间
	UpdateAt string `gorm:"update_at" json:"update_at"` // 修改时间
	Status   int    `gorm:"status" json:"status"`       // 状态：0 正常  1 停用 9 删除
}

func (scm_settlement_type) TableName() string {
	return "scm_settlement_type"
}

type scm_trade struct {
	Id       int    `gorm:"id" json:"id"`               // 主键id
	Name     string `gorm:"name" json:"name"`           // 行业名称
	Remark   string `gorm:"remark" json:"remark"`       // 备注
	CreateAt string `gorm:"create_at" json:"create_at"` // 创建时间
	UpdateAt string `gorm:"update_at" json:"update_at"` // 修改时间
	Status   int    `gorm:"status" json:"status"`       // 状态：0 正常  1 停用 9 删除
}

func (scm_trade) TableName() string {
	return "scm_trade"
}
