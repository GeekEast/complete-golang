package G10_collection_datatypes

func TestMap() {
	p := map[string]int{
		"alice":30,
	}
	age,ok := p["aliced"]
	println(age,ok)
}