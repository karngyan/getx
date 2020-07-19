package utils

func GetSourceUri(id string) string {
	return "/files/" + id + ".html"
}

func GetFilePath(id string) string {
	return "app/files/" + id + ".html"
}
