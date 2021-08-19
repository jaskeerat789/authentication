package product

import (
	"encoding/json"
	"fmt"
	"io"
)

type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
}

type Products []Product

var products = Products{
	{Id: 1, Name: "World of Authcraft", Slug: "world-of-authcraft", Description: "Battle bugs and protect yourself from invaders while you explore a scary world with no security"},
	{Id: 2, Name: "Ocean Explorer", Slug: "ocean-explorer", Description: "Explore the depths of the sea in this one of a kind underwater experience"},
	{Id: 3, Name: "Dinosaur Park", Slug: "dinosaur-park", Description: "Go back 65 million years in the past and ride a T-Rex"},
	{Id: 4, Name: "Cars VR", Slug: "cars-vr", Description: "Get behind the wheel of the fastest cars in the world."},
	{Id: 5, Name: "Robin Hood", Slug: "robin-hood", Description: "Pick up the bow and arrow and master the art of archery"},
	{Id: 6, Name: "Real World VR", Slug: "real-world-vr", Description: "Explore the seven wonders of the world in VR"},
}

func GetProducts() *Products {
	return &products
}

func FindBySlug(slug string) (*Product, error) {
	for _, p := range products {
		if p.Slug == slug {
			return &p, nil
		}
	}
	return &Product{}, fmt.Errorf("cannot find product with slug=%s", slug)
}

func (ps *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(ps)
}

func (p *Product) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
