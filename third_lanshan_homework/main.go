package main

import "fmt"

//定义结构体(父类)
type Goods struct {
	class              string
	price              float64
	inventory_quantity int
}

//定义结构体(子类)
type E_goods struct {
	Goods
	electronic_product_type string
	brand                   string
	model                   string
}

type F_goods struct {
	Goods
	brand          string
	furnitune_type string
	size_w         float64
	size_l         float64
}

type C_goods struct {
	Goods
	brand        string
	clothes_type string
	clothes_size string
}

//设计接口(父类)
type Goods_manager interface {
	Check_quantity()
	New_quantity(q int)
	Print_info()
}

//设计接口(子类)
type Goods_manager_e interface {
	Print_info_e()
}

type Goods_manager_f interface {
	Print_info_f()
}

type Goods_manager_c interface {
	Print_info_c()
}

//定义方法(父类)
func (g *Goods) Check_quantity() {
	fmt.Printf("该产品还剩余%v个\n", g.inventory_quantity)
}

func (g *Goods) New_quantity(q int) {
	g.inventory_quantity = q
}

func (g *Goods) Print_info() {
	fmt.Printf("商品种类：%v  ，  库存：%v  ，  单价：%v\n", g.class, g.inventory_quantity, g.price)
}

//定义方法(子类)
func (eg *E_goods) Print_info_e() {
	fmt.Printf("电子产品种类：%v  ，  品牌：%v  ，  型号：%v\n", eg.electronic_product_type, eg.brand, eg.model)
}

func (fg *F_goods) Print_info_f() {
	fmt.Printf("家具类型：%v  ， 品牌：%v  ，    长度：%v  ，  宽度：%v\n", fg.furnitune_type, fg.brand, fg.size_l, fg.size_w)
}

func (cg *C_goods) Print_info_c() {
	fmt.Printf(" 衣服类型：%v  ， 品牌：%v  ，   尺码：%v\n", cg.clothes_type, cg.brand, cg.clothes_size)
}
