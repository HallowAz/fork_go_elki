package entity

type Cart struct {
	ID     uint
	UserID uint
}

type CartProduct struct {
	ID        uint
	ProductID uint
	CartID    uint
	ItemCount int
}

type CartWithRestaurant struct {
	RestaurantId uint
	Products     []*CartProduct
	PromoId      uint
}
