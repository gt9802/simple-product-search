package handlers

type Searchhandler struct {
	Products []string
}

func NewSearchHandler(products []string) *Searchhandler {
	return &Searchhandler{Products: products}
}
