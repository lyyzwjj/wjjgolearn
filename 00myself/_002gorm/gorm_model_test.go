package _002gorm

//type User struct {
//	ID           uint
//	Name         string
//	Email        *string
//	Age          uint8
//	Birthday     *time.Time
//	MemberNumber sql.NullString
//	ActivatedAt  sql.NullTime
//	CreatedAt    time.Time
//	UpdatedAt    time.Time
//}

// gorm.Model 的定义
//type Model struct {
//	ID        uint           `gorm:"primaryKey"`
//	CreatedAt time.Time
//	UpdatedAt time.Time
//	DeletedAt gorm.DeletedAt `gorm:"index"`
//}

// 字段级权限控制
//type User struct {
//	Name string `gorm:"<-:create"`          // 允许读和创建
//	Name string `gorm:"<-:update"`          // 允许读和更新
//	Name string `gorm:"<-"`                 // 允许读和写（创建和更新）
//	Name string `gorm:"<-:false"`           // 允许读，禁止写
//	Name string `gorm:"->"`                 // 只读（除非有自定义配置，否则禁止写）
//	Name string `gorm:"->;<-:create"`       // 允许读和写
//	Name string `gorm:"->:false;<-:create"` // 仅创建（禁止从 db 读）
//	Name string `gorm:"-"`                  // 通过 struct 读写会忽略该字段
//}

// 创建/更新时间追踪（纳秒、毫秒、秒、Time
// 如果您想要保存 UNIX（毫/纳）秒时间戳，而不是 time，您只需简单地将 time.Time 修改为 int 即可
//type User struct {
//	CreatedAt time.Time // Set to current time if it is zero on creating
//	UpdatedAt int       // Set to current unix seconds on updating or if it is zero on creating
//	Updated   int64     `gorm:"autoUpdateTime:nano"`  // Use unix nano seconds as updating time
//	Updated   int64     `gorm:"autoUpdateTime:milli"` // Use unix milli seconds as updating time
//	Created   int64     `gorm:"autoCreateTime"`       // Use unix seconds as creating time
//}

// 嵌入结构体
// 对于匿名字段，GORM 会将其字段包含在父结构体中，例如：
//type User struct {
//	gorm.Model
//	Name string
//}
//// 等效于
//type User struct {
//	ID        uint           `gorm:"primaryKey"`
//	CreatedAt time.Time
//	UpdatedAt time.Time
//	DeletedAt gorm.DeletedAt `gorm:"index"`
//	Name string
//}

// 对于正常的结构体字段，你也可以通过标签 embedded 将其嵌入，例如：
//type Author struct {
//	Name  string
//	Email string
//}
//
//type Blog struct {
//	ID      int
//	Author  Author `gorm:"embedded"`
//	Upvotes int32
//}
//// 等效于
//type Blog struct {
//	ID    int64
//	Name  string
//	Email string
//	Upvotes  int32
//}
// 并且，您可以使用标签 embeddedPrefix 来为 db 中的字段名添加前缀，例如：
//type Blog struct {
//	ID      int
//	Author  Author `gorm:"embedded;embeddedPrefix:author_"`
//	Upvotes int32
//}
//// 等效于
//type Blog struct {
//	ID          int64
//	AuthorName  string
//	AuthorEmail string
//	Upvotes     int32
//}
