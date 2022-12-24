package sql2gorm

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	sql := `create table user
(
    uid                 int unsigned auto_increment comment '用户ID'
        primary key,
    telephone           char(15)                     not null comment '电话号码',
    username          varchar(40)  default '游客'  not  null comment '用户名',

    status           tinyint(1)       default 1   not null comment '帐号状态：1：正常  2:锁定 9：禁用 10:游客',
    ban_reason varchar(50)      default ''  not null comment '禁用原因'    
)
    comment '用户表';`

	data, err := Parse(sql, WithJsonTag())
	assert.Nil(t, err)
	t.Log(data)
	fmt.Println(" ")
	fmt.Println(data.StructCode)

}
