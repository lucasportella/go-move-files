package types

type InnerPaths struct {
	Src_dir string
	Dst_dir string
}

type Paths map[string]map[string]InnerPaths