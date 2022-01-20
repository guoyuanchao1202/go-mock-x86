package consts

type PagingMode string

const (
	NormalPagingMode PagingMode = "32-bit paging"
	PAEPaging        PagingMode = "PAE paging"
	FourLevelPaging  PagingMode = "4-level paging"
	FiveLevelPaging  PagingMode = "5-level paging"
)
