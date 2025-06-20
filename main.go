package main

func main() {
	csv, err := NewCSV("./file.csv", ',')
	if err != nil {
		panic(err)
	}
	lines, err := csv.yield()
	if err != nil {
		panic(err)
	}

	for line := range lines {
		//fmt.Println(" -- ", line.Get("id"), line.Get("city"))
		_ = line.Get("id")
		_ = line.Get("city")
	}

	// file, err := os.Open("file.csv")
	// if err != nil {
	// 	panic(err)
	// }
	// parts := make([]string, 0)
	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	_ = Split(scanner.Text(), ",", &parts)
	// }
}
