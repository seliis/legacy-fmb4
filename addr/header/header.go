package header

type Nav struct {
	Menu string
}

func GetNavs() [6]Nav {
	var arr [6]Nav
	arr[0].Menu = "aviators"
	arr[1].Menu = "training"
	arr[2].Menu = "document"
	arr[3].Menu = "campaign"
	arr[4].Menu = "database"
	arr[5].Menu = "signs-in"

	return arr
}
