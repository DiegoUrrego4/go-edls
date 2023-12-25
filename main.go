package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"regexp"
	"runtime"
	"strings"
	"time"
)

func main() {
	// flag pattern
	flagPattern := flag.String("p", "", "filter by pattern")
	//flagAll := flag.Bool("a", false, "all files including hidden files")
	flagNumberRecords := flag.Int("n", 0, "numbers of records")

	// order flags
	//hasOrderByTime := flag.Bool("t", false, "sort by time, oldest first")
	//hasOrderBySize := flag.Bool("s", false, "sort by size, smallest first")
	//hasOrderReverse := flag.Bool("r", false, "reverse order while sorting")

	flag.Parse() // Mapea cada uno de los flags y almacena en variables

	path := flag.Arg(0)
	if path == "" {
		path = "."
	}
	dirs, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	var files []file

	for _, dir := range dirs {
		f, err := getFile(dir, false)
		if err != nil {
			panic(err)
		}

		// Determinar si se incluye o excluye
		// "(?i)" --> Se agrega para indicarle al programa que sea insensitive (aceptar mayúsculas y minúsculas)
		isMatched, err := regexp.MatchString("(?i)"+*flagPattern, f.name)
		if err != nil {
			panic(err)
		}

		if !isMatched {
			continue
		}

		files = append(files, f)
	}

	if *flagNumberRecords == 0 || *flagNumberRecords > len(files) {
		*flagNumberRecords = len(files)
	}

	printList(files, *flagNumberRecords)

}

func printList(fs []file, nRecords int) {
	for _, file := range fs[:nRecords] {
		style := mapStyleByFileType[file.fileType]
		fmt.Printf("%s %s %s %10d %s %s %s%s\n", file.mode, file.userName, file.groupName, file.size,
			file.modificationTime.Format(time.DateTime), style.icon, file.name, style.symbol)
	}
}

func getFile(dir fs.DirEntry, isHidden bool) (file, error) {
	info, err := dir.Info()
	if err != nil {
		return file{}, fmt.Errorf("dir.Info(): %v", err)
	}

	f := file{
		name:             dir.Name(),
		isDir:            dir.IsDir(),
		isHidden:         isHidden,
		userName:         "durrego",
		groupName:        "MELI",
		size:             info.Size(),
		modificationTime: info.ModTime(),
		mode:             info.Mode().String(),
	}

	setFile(&f)

	return f, nil
}

func setFile(f *file) {
	switch {
	case isLink(*f):
		f.fileType = fileLink
	case f.isDir:
		f.fileType = fileDirectory
	case isExec(*f):
		f.fileType = fileExecutable
	case isCompress(*f):
		f.fileType = fileCompress
	case isImage(*f):
		f.fileType = fileImage
	default:
		f.fileType = fileRegular
	}
}

// Funciones auxiliares

func isLink(f file) bool {
	return strings.HasPrefix(strings.ToUpper(f.mode), "L")
}

func isExec(f file) bool {
	if runtime.GOOS == Windows {
		return strings.HasSuffix(f.name, exe)
	}

	return strings.Contains(f.mode, "x")
}

func isCompress(f file) bool {
	return strings.HasSuffix(f.name, zip) ||
		strings.HasSuffix(f.name, gz) ||
		strings.HasSuffix(f.name, tar) ||
		strings.HasSuffix(f.name, rar) ||
		strings.HasSuffix(f.name, deb)
}

func isImage(f file) bool {
	return strings.HasSuffix(f.name, png) ||
		strings.HasSuffix(f.name, jpg) ||
		strings.HasSuffix(f.name, gif)
}
