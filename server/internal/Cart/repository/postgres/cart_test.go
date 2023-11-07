package repository

import (
	"errors"
	"reflect"
	"server/internal/domain/entity"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestCreateCartSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &CartRepo{
		DB: db,
	}

	var userID uint
	userID = 1

	rows := sqlmock.
		NewRows([]string{"id"}).AddRow(1)

	mock.
		ExpectExec(`INSERT INTO cart`).
		WithArgs(userID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.
		ExpectQuery("SELECT ID FROM cart WHERE").
		WithArgs(userID).
		WillReturnRows(rows)

	id, err := repo.CreateCart(userID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if id != 1 {
		t.Errorf("bad id: want %v, have %v", id, 1)
		return
	}

}

func TestCreateCartFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &CartRepo{
		DB: db,
	}

	var userID uint
	userID = 1

	testErr := errors.New("test")

	mock.
		ExpectExec(`INSERT INTO cart`).
		WithArgs(userID).
		WillReturnError(testErr)

	id, err := repo.CreateCart(userID)
	if err != entity.ErrInternalServerError {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if id != 0 {
		t.Errorf("bad id: want %v, have %v", id, 1)
		return
	}

	mock.
		ExpectExec(`INSERT INTO cart`).
		WithArgs(userID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.
		ExpectQuery("SELECT ID FROM cart WHERE").
		WithArgs(userID).
		WillReturnError(testErr)

	id, err = repo.CreateCart(userID)
	if err != entity.ErrInternalServerError {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if id != 0 {
		t.Errorf("bad id: want %v, have %v", id, 1)
		return
	}

}

func TestGetCartByUserIDSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &CartRepo{
		DB: db,
	}

	row := sqlmock.
		NewRows([]string{"id", "user_id"})
	expect := &entity.Cart{
		ID:     1,
		UserID: 1,
	}

	row = row.AddRow(expect.ID, expect.UserID)

	var elemID = 1

	mock.
		ExpectQuery("SELECT id, user_id FROM cart WHERE").WithArgs(elemID).
		WillReturnRows(row)

	cart, err := repo.GetCartByUserID(uint(elemID))
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if !reflect.DeepEqual(cart, expect) {
		t.Errorf("results not match, want %v, have %v", expect, cart)
		return
	}

}

func TestGetCartByUserIDFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &CartRepo{
		DB: db,
	}

	testErr := errors.New("test")

	var elemID = 1

	mock.
		ExpectQuery("SELECT id, user_id FROM cart WHERE").WithArgs(elemID).
		WillReturnError(testErr)

	cart, err := repo.GetCartByUserID(uint(elemID))

	if err != entity.ErrInternalServerError {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if cart != nil {
		t.Errorf("not nil cart while error")
		return
	}
}

func TestGetCartProductsByCartIDSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &CartRepo{
		DB: db,
	}

	rows := sqlmock.
		NewRows([]string{"id", "product_id", "cart_id", "item_count"})
	expect := []*entity.CartProduct{
		{
			ID:        1,
			ProductID: 1,
			CartID:    1,
			ItemCount: 6,
		},
		{
			ID:        2,
			ProductID: 3,
			CartID:    1,
			ItemCount: 6,
		},
	}

	for _, product := range expect {
		rows = rows.AddRow(product.ID, product.ProductID, product.CartID, product.ItemCount)
	}

	var elemID = 1

	mock.
		ExpectQuery("SELECT id, product_id, cart_id, item_count FROM cart_product WHERE").WithArgs(elemID).
		WillReturnRows(rows)

	products, err := repo.GetCartProductsByCartID(uint(elemID))
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if !reflect.DeepEqual(products[0], expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], products[0])
		return
	}

}

func TestGetCartProductsByCartIDFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &CartRepo{
		DB: db,
	}

	testErr := errors.New("test")

	var elemID = 1

	mock.
		ExpectQuery("SELECT id, product_id, cart_id, item_count FROM cart_product WHERE").WithArgs(elemID).
		WillReturnError(testErr)

	products, err := repo.GetCartProductsByCartID(uint(elemID))

	if err != testErr {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if products != nil {
		t.Errorf("carts not nil while error")
	}
}

func TestAddProductToCartSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &CartRepo{
		DB: db,
	}

	var cartID, productID uint
	cartID = 1
	productID = 1
	mock.
		ExpectExec(`INSERT INTO cart_product`).
		WithArgs(cartID, productID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.AddProductToCart(cartID, productID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

}

func TestAddProductToCartFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &CartRepo{
		DB: db,
	}

	var cartID, productID uint
	cartID = 1
	productID = 1

	testErr := errors.New("test")

	mock.
		ExpectExec(`INSERT INTO cart_product`).
		WithArgs(cartID, productID).
		WillReturnError(testErr)

	err = repo.AddProductToCart(cartID, productID)
	if err != entity.ErrInternalServerError {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestDeleteProductFromCartSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &CartRepo{
		DB: db,
	}

	var cartID, productID uint
	cartID = 1
	productID = 1
	mock.
		ExpectExec(`DELETE FROM cart_product WHERE`).
		WithArgs(cartID, productID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.DeleteProductFromCart(cartID, productID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestDeleteProductFromCartFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &CartRepo{
		DB: db,
	}

	var cartID, productID uint
	cartID = 1
	productID = 1

	testErr := errors.New("test")

	mock.
		ExpectExec(`DELETE FROM cart_product WHERE`).
		WithArgs(cartID, productID).
		WillReturnError(testErr)

	err = repo.DeleteProductFromCart(cartID, productID)
	if err != entity.ErrInternalServerError {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestUpdateItemCountUpSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &CartRepo{
		DB: db,
	}

	var cartID, productID uint
	cartID = 1
	productID = 1

	mock.
		ExpectExec(`UPDATE cart_product SET`).
		WithArgs(cartID, productID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UpdateItemCountUp(cartID, productID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

}
