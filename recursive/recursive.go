// 递归查询

package main

import (
	"fmt"
	"yunplus/erp_scm/database"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Result struct {
	ID       int
	Name     string
	ParentId int
}

type GoodsCategoryNode struct {
	ID       int                 `json:"id"`                 // 分类ID
	Name     string              `json:"name"`               // 类别名称
	Children []GoodsCategoryNode `json:"children,omitempty"` // 子节点，可为空
}

func main() {
	var results []*Result
	results, _ = GetCategoryChildren(1)
	var ret = new([]GoodsCategoryNode)
	// 将查询结果组织成树状结构
	for _, v := range results {
		/*选择需要开始遍历的节点*/
		find(v, ret)
	}
	fmt.Print(ret)
}

// 将查询结果组织成树状结构
func find(currentCate *Result, ret *[]GoodsCategoryNode) {

	if len(*ret) == 0 || currentCate.ParentId == 0 {
		bb := GoodsCategoryNode{
			ID:   currentCate.ID,
			Name: currentCate.Name,
		}
		*ret = append(*ret, bb)
		return
	}

	for i, _ := range *ret {
		//当前节点中有该父级
		if currentCate.ParentId == (*ret)[i].ID {
			aa := GoodsCategoryNode{
				ID:   currentCate.ID,
				Name: currentCate.Name,
			}

			if (*ret)[i].Children == nil {
				(*ret)[i].Children = make([]GoodsCategoryNode, 0)
			}

			(*ret)[i].Children = append((*ret)[i].Children, aa)
			return
		}

		//否则遍历该节点的子节点
		if (*ret)[i].Children == nil || len((*ret)[i].Children) == 0 {
			continue
		}

		//有子节点就遍历出来
		find(currentCate, &(*ret)[i].Children)
	}
	return
}

// GetCategoryChildren  获取数据
func GetCategoryChildren(id int) ([]*Result, error) {
	var results []*Result
	db := database.MysqlDB.GetDB()
	if err := db.Raw(`select  id,name,parent_id from (select * from scm_goods_category  where status = ? order by parent_id, id) cate_sorted,(
                           select @pv := ?) initialisation
                           where find_in_set(parent_id, @pv)
                           and length(@pv := concat(@pv, ',', id))`, 0, id).Scan(&results).Error; err != nil {

		return nil, err
	}
	return results, nil
}
