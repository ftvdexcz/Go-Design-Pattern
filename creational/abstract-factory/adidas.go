package main

// Adidas is factory for adidas product
type Adidas struct {
}

// AdidasShoes implements IShoes
type AdidasShoes struct {
	Shoes
}

// AdidasShirt implements IShirt
type AdidasShirt struct {
	Shirt
}

func (a *Adidas) createShoes() IShoes {
	return &AdidasShoes{
		Shoes{
			logo: "adidas",
			size: 14,
		},
	}
}

func (a *Adidas) createShirt() IShirt {
	return &AdidasShirt{
		Shirt{
			logo: "adidas",
			size: 14,
		},
	}
}
