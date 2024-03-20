package repoutil

func GetOffset(pageSize, pageId int) int {
	return pageSize * (pageId - 1)
}
