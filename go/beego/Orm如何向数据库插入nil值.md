问题描述:
beege Orm Insert方法,默认会将字段的零值,插入数据库,当我们想插入null,的时候就没有办法.特别是这个字段为唯一字段的时候,特别麻烦,后发现如果字段定义为指针,即可插入null的数据

orm无法插入null
```

type User struct {
	Id         int       `orm:"size(11);column(id)"`
	Email      string    `field:"邮箱" orm:"null" valid:"Email"`                              //邮箱

}
```
orm可插入nil
```
type User struct {
	Id         int       `orm:"size(11);column(id)"`
	Email      *string    `field:"邮箱" orm:"null" valid:"Email"`                              //邮箱

}
```
将Email 字段定义为指针,不赋值的情况下,Orm会使用nil去构建sql,从而达到插入null字段的目的