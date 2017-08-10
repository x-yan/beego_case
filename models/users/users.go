package users

import (
	"github.com/astaxie/beego/orm"
)

type Users struct {
	Id int64         `orm:"pk;column(id);"`//多个设置间使用 ; 分隔，设置的值如果是多个，使用 , 分隔。
	Title string
	Name string
	AnyField string `orm:"-"`//设置 - 即可忽略 struct 中的字段
	//Name string `orm:"column(user_name)"` 为字段设置 db 字段的名称
}

/**
 设定默认名字
 */
func (u *Users) TableName() string {
	return "t_impression_picture_data"
}


func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Users))

	//orm.RegisterDriver("mysql", orm.DRMySQL)

	//orm.RegisterModelWithPrefix("prefix_", new(User)) //注册表前缀

	//db, err := orm.GetDB()返回已注册的数据库名

	//orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8")

	//o1 := orm.NewOrm()
	//o1.Using("default") 可选择数据库
	//dr.Name() 返回名字
}

//func  read()  {
//
//	o := orm.NewOrm()
//	user := Users{Id: 1}
//
//	err := o.Read(&user)
//
//	if err == orm.ErrNoRows {
//		fmt.Println("查询不到")
//	} else if err == orm.ErrMissPK {
//		fmt.Println("找不到主键")
//	} else {
//		fmt.Println(user.Id, user.Name)
//	}
//
//	//user := User{Name: "slene"}
//	//err = o.Read(&user, "Name")
//
//   //设置返回字段
//	type Post struct {
//		Id      int
//		Title   string
//		Content string
//		Status  int
//	}
//
//	// 只返回 Id 和 Title
//	var posts []Post
//	o.QueryTable("post").Filter("Status", 1).All(&posts, "Id", "Title")
//
//
//	//以键值对的形式返回
//
//	var maps []orm.Params
//	num, err := o.QueryTable("user").Values(&maps)
//	if err == nil {
//		fmt.Printf("Result Nums: %d\n", num)
//		for _, m := range maps {
//			fmt.Println(m["Id"], m["Name"])
//		}
//	}
//}
//
//func  insert()  {
//	o := orm.NewOrm()
//	var user Users
//	//user.Name = "slene"
//	//user.IsActive = true
//
//	id, err := o.Insert(&user)
//	if err == nil {
//		fmt.Println(id)
//	}
//
//	//方法二
//
//	var users []*Users
//	qs := o.QueryTable("user")
//	i, _ := qs.PrepareInsert()
//	for _, user := range users {
//		id, err := i.Insert(user)
//		if err == nil {
//
//		}
//	}
//	// PREPARE INSERT INTO user (`name`, ...) VALUES (?, ...)
//	// EXECUTE INSERT INTO user (`name`, ...) VALUES ("slene", ...)
//	// EXECUTE ...
//	// ...
//	i.Close() // 别忘记关闭 statement
//
//
//}
//
//func InsertMulti()  {
//	o := orm.NewOrm()
//	users := []Users{
//		{Name: "slene"},
//		{Name: "astaxie"},
//		{Name: "unknown"},
//}
//	successNums, err := o.InsertMulti(100, users)// bulk 为并列插入的数量
//}
//
//func update()  {
//	//默认更新所有字段
//	o := orm.NewOrm()
//	user := Users{Id: 1}
//	if o.Read(&user) == nil {
//		user.Name = "MyName"
//		if num, err := o.Update(&user); err == nil {
//			fmt.Println(num)
//		}
//	}
//	// 只更新 Name
//	o.Update(&user, "Name")
//	// 指定多个字段
//	// o.Update(&user, "Field1", "Field2", ...)
//}
//
//func delete()  {
//	o := orm.NewOrm()
//	if num, err := o.Delete(&Users{Id: 1}); err == nil {
//		fmt.Println(num)
//	}
//}
//
//func Other()  {
//	o := orm.NewOrm()
//
//	// 获取 QuerySeter 对象，user 为表名
//	qs := o.QueryTable("user")
//
//	// 也可以直接使用对象作为表名
//	user := new(Users)
//	qs = o.QueryTable(user) // 返回 QuerySeter
//
//	qs.Filter("name", "slene") // WHERE name = 'slene'
//	qs.Filter("name__iexact", "slene")// WHERE name LIKE 'slene'
//	qs.Filter("name__contains", "slene")// WHERE name LIKE BINARY '%slene%'
//	qs.Filter("name__icontains", "slene")// WHERE name LIKE '%slene%'
//	qs.Filter("profile__age__in", 17, 18, 19, 20)// WHERE profile.age IN (17, 18, 19, 20)
//	ids:=[]int{17,18,19,20}
//	qs.Filter("profile__age__in", ids)// WHERE profile.age IN (17, 18, 19, 20)
//	qs.Filter("profile__age__gt", 17)	// WHERE profile.age > 17
//	qs.Filter("profile__age__gte", 18)	// WHERE profile.age >= 18
//	qs.Filter("profile__age__lt", 17) // WHERE profile.age < 17
//	qs.Filter("profile__age__lte", 18) // WHERE profile.age <= 18
//	qs.Filter("name__startswith", "slene") // WHERE name LIKE BINARY 'slene%'
//	qs.Filter("name__endswith", "slene") // WHERE name LIKE BINARY '%slene'
//	qs.Filter("name__iendswithi", "slene") // WHERE name LIKE '%slene'
//	qs.Filter("profile__isnull", true)
//	qs.Filter("profile_id__isnull", true) // WHERE profile_id IS NULL
//	qs.Filter("profile__isnull", false) // WHERE profile_id IS NOT NULL
//	qs.Filter("profile__isnull", true).Filter("name", "slene") // WHERE profile_id IS NULL AND name = 'slene'
//	qs.Exclude("profile__isnull", true).Filter("name", "slene") // WHERE NOT profile_id IS NULL AND name = 'slene'
//
//	cond := orm.NewCondition()
//	cond1 := cond.And("profile__isnull", false).AndNot("status__in", 1).Or("profile__age__gt", 2000)
//	qs := orm.QueryTable("user")
//	qs = qs.SetCond(cond1)
//	// WHERE ... AND ... AND NOT ... OR ...
//	cond2 := cond.AndCond(cond1).OrCond(cond.And("name", "slene"))
//	qs = qs.SetCond(cond2).Count()
//	// WHERE (... AND ... AND NOT ... OR ...) OR ( ... )
//	qs.Limit(10) // LIMIT 10
//	qs.Limit(10, 20) // LIMIT 10 OFFSET 20 注意跟 SQL 反过来的
//	qs.GroupBy("id", "age") // GROUP BY id,age
//	qs.Offset(20) // LIMIT 1000 OFFSET 20
//	qs.OrderBy("id", "-profile__age") // ORDER BY id ASC, profile.age DESC
//	qs.OrderBy("-profile__age", "profile") // ORDER BY profile.age DESC, profile_id ASC
//	qs.Distinct() // SELECT DISTINCT
//	exist := o.QueryTable("user").Filter("UserName", "Name").Exist()
//	cnt, err := o.QueryTable("user").Count() // SELECT COUNT(*) FROM USER
//	// 假设 user struct 里有一个 nums int 字段
//	num, err := o.QueryTable("user").Update(orm.Params{
//		"nums": orm.ColValue(orm.ColAdd, 100),
//	})
//	// SET nums = nums + 100
//	//ColAdd      // 加
//	//ColMinus    // 减
//	//ColMultiply // 乘
//	//ColExcept   // 除
//
//	num, err := o.QueryTable("user").Filter("name", "slene").Delete()
//	fmt.Printf("Affected Num: %s, %s", num, err)
//	// DELETE FROM user WHERE name = "slene"
//
//
//}
//
// func  RelatedSel(){
//	 o := orm.NewOrm()
//	 qs := o.QueryTable("post")
//	 qs.RelatedSel("user") // INNER JOIN user ...
//
// }
///**
//构造 查询
// */
//func bulder()  {
//	// User 包装了下面的查询结果
//	type User struct {
//		Name string
//		Age  int
//	}
//	var users []User
//
//	// 获取 QueryBuilder 对象. 需要指定数据库驱动参数。
//	// 第二个返回值是错误对象，在这里略过
//	qb, _ := orm.NewQueryBuilder("mysql")
//
//	// 构建查询对象
//	qb.Select("user.name",
//		"profile.age").
//		From("user").
//		InnerJoin("profile").On("user.id_user = profile.fk_user").
//		Where("age > ?").
//		OrderBy("name").Desc().
//		Limit(10).Offset(0)
//
//	// 导出 SQL 语句
//	sql := qb.String()
//
//	// 执行 SQL 语句
//	o := orm.NewOrm()
//	o.Raw(sql, 20).QueryRows(&users)
//
//
//	//type QueryBuilder interface {
//	//	Select(fields ...string) QueryBuilder
//	//	From(tables ...string) QueryBuilder
//	//	InnerJoin(table string) QueryBuilder
//	//	LeftJoin(table string) QueryBuilder
//	//	RightJoin(table string) QueryBuilder
//	//	On(cond string) QueryBuilder
//	//	Where(cond string) QueryBuilder
//	//	And(cond string) QueryBuilder
//	//	Or(cond string) QueryBuilder
//	//	In(vals ...string) QueryBuilder
//	//	OrderBy(fields ...string) QueryBuilder
//	//	Asc() QueryBuilder
//	//	Desc() QueryBuilder
//	//	Limit(limit int) QueryBuilder
//	//	Offset(offset int) QueryBuilder
//	//	GroupBy(fields ...string) QueryBuilder
//	//	Having(cond string) QueryBuilder
//	//	Subquery(sub string, alias string) string
//	//	String() string
//	//}
//}