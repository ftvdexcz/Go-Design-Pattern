package main

// Nike is factory for nike product
type Nike struct {
}

// NikeShoes implements IShoes
type NikeShoes struct {
	Shoes
}

// NikeShirt implements IShirt
type NikeShirt struct {
	Shirt
}

func (n *Nike) createShoes() IShoes {
	return &NikeShoes{
		Shoes{
			logo: "nike",
			size: 14,
		},
	}
}

func (n *Nike) createShirt() IShirt {
	return &NikeShirt{
		Shirt{
			logo: "nike",
			size: 14,
		},
	}
}
