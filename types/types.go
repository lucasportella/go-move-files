package types

type Configuration struct {
	Default  map[string]Paths       `json:"default"`
	WithDate DatePaths              `json:"withDate"`
	Delete   map[string]DeletePaths `json:"delete"`
}

type Paths struct {
	SrcDir string `json:"src_dir"`
	DstDir string `json:"dst_dir"`
}

type DatePaths struct {
	WithDateDaily   map[string]Paths `json:"withDateDaily"`
	WithDateMonthly map[string]Paths `json:"withDateMonthly"`
	WithDateYearly  map[string]Paths `json:"withDateYearly"`
}

type DeletePaths struct {
	SrcDir string `json:"src_dir"`
}
