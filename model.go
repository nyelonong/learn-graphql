package main

// ws_user
var (
	UserData map[int]User

	Zaki   User
	Yoshua User
)

// ws_shop
var (
	ShopData map[int]Shop

	Komopedia Shop
	ElShadai  Shop
)

// ws_product
var (
	ProductData map[int]Product

	LP1 Product
	LP2 Product
	LP3 Product
)

type User struct {
	ID   int
	Name string
	Shop int
}

type Shop struct {
	ID       int
	Name     string
	Products []int
}

type Product struct {
	ID    int
	Name  string
	Price float64
}

func init() {
	Zaki = User{
		ID:   2391,
		Name: "Zaki",
		Shop: 1,
	}

	Yoshua = User{
		ID:   1234,
		Name: "Yoshua",
		Shop: 2,
	}

	Komopedia = Shop{
		ID:       1,
		Name:     "Komopedia",
		Products: []int{1},
	}

	ElShadai = Shop{
		ID:       2,
		Name:     "El Shadai",
		Products: []int{2, 3},
	}

	LP1 = Product{
		ID:    1,
		Name:  "LP 1",
		Price: 1000,
	}

	LP2 = Product{
		ID:    3,
		Name:  "LP 2",
		Price: 2000,
	}

	LP3 = Product{
		ID:    3,
		Name:  "LP 3",
		Price: 3000,
	}

	UserData = map[int]User{
		1: Zaki,
		2: Yoshua,
	}

	ShopData = map[int]Shop{
		1: Komopedia,
		2: ElShadai,
	}

	ProductData = map[int]Product{
		1: LP1,
		2: LP2,
		3: LP3,
	}
}
