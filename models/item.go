package models

type Item struct {
	ImageUrl    string
	Title       string
	Category    string
	Subcategory string
	Price       string
}

var Data = []Item{
	{
		ImageUrl:    "https://images.unsplash.com/photo-1589782431746-713d84eef3c4?q=80&w=1974&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Title:       "Nero",
		Category:    "Accessories",
		Subcategory: "Chanel",
		Price:       "$120",
	},
	{
		ImageUrl:    "https://images.unsplash.com/photo-1520256862855-398228c41684?q=80&w=2069&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Title:       "Addidas",
		Category:    "Sport Shoes",
		Subcategory: "Addidas",
		Price:       "$100",
	},
	{
		ImageUrl:    "https://images.unsplash.com/photo-1519027356316-9f99e11d8bac?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Title:       "Costume di lana",
		Category:    "Clothes Mix",
		Subcategory: "Man Mix",
		Price:       "$100",
	},
	{
		ImageUrl:    "https://images.unsplash.com/photo-1516461240763-822a87484851?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MTZ8fGNsYXNzaWMlMjB3YXRjaGVzfGVufDB8fDB8fHww",
		Title:       "Fossil",
		Category:    "Accessories",
		Subcategory: "Classic watch",
		Price:       "$300",
	},
	{
		ImageUrl:    "https://images.unsplash.com/photo-1502716119720-b23a93e5fe1b?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MTl8fHdvbWFuJTIwY2xvdGhlc3xlbnwwfHwwfHx8MA%3D%3D",
		Title:       "Red Dress",
		Category:    "Clothes Mix",
		Subcategory: "Woman Mix",
		Price:       "$110",
	},
	{
		ImageUrl:    "https://images.unsplash.com/photo-1515955656352-a1fa3ffcd111?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Title:       "Nike Bleu",
		Category:    "Sport Shoes",
		Subcategory: "Nike",
		Price:       "$100",
	},
	{
		ImageUrl:    "https://images.unsplash.com/photo-1497382706140-605ee76b3b55?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Title:       "Daniel Wellington",
		Category:    "Accessories",
		Subcategory: "Classic watch",
		Price:       "$300",
	},
	{
		ImageUrl:    "https://images.unsplash.com/photo-1492633397843-92adffad3d1c?q=80&w=1973&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Title:       "T-shirt Marrone",
		Category:    "Clothes Mix",
		Subcategory: "Man Mix",
		Price:       "$100",
	},
	{
		ImageUrl:    "https://images.unsplash.com/photo-1720516494265-ce32a91ada94?q=80&w=1932&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Title:       "Fossil Q",
		Category:    "Accessories",
		Subcategory: "Classic watch",
		Price:       "$300",
	},
	{
		ImageUrl:    "https://images.unsplash.com/photo-1521774971864-62e842046145?q=80&w=2069&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Title:       "Azzure",
		Category:    "Sport Shoes",
		Subcategory: "Rebook",
		Price:       "$100",
	},
	{
		ImageUrl:    "https://images.unsplash.com/photo-1485230895905-ec40ba36b9bc?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Title:       "Red Sweatshirt",
		Category:    "Clothes Mix",
		Subcategory: "Woman Mix",
		Price:       "$110",
	},
	{
		ImageUrl:    "https://images.unsplash.com/photo-1523293182086-7651a899d37f?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Title:       "Bleu",
		Category:    "Accessories",
		Subcategory: "Chanel",
		Price:       "$120",
	},
	{
		ImageUrl:    "https://images.unsplash.com/photo-1539185441755-769473a23570?q=80&w=2071&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Title:       "New Balance",
		Category:    "Sport Shoes",
		Subcategory: "New Balance",
		Price:       "$100",
	},
	{
		ImageUrl:    "https://images.unsplash.com/photo-1503342217505-b0a15ec3261c?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Title:       "Black T-shirt",
		Category:    "Clothes Mix",
		Subcategory: "Woman Mix",
		Price:       "$110",
	},
	{
		ImageUrl:    "https://images.unsplash.com/photo-1514348871858-1d3c20902571?q=80&w=1974&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Title:       "Nr.5",
		Category:    "Accessories",
		Subcategory: "Chanel",
		Price:       "$120",
	},
	{
		ImageUrl:    "https://images.unsplash.com/photo-1542291026-7eec264c27ff?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Title:       "Nike Red",
		Category:    "Sport Shoes",
		Subcategory: "Nike",
		Price:       "$100",
	},
}
