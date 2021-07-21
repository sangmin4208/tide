package main

func main() {
	tides := []*Tide{}
	fileNames := getFileNames(INPUT_PATH)
	for _, fn := range fileNames {
		tide := convertToTide(fn)
		tides = append(tides, tide)
	}

	for _, tide := range tides {
		writeFileWithTide(tide)
	}
}
