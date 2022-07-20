package parser

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

const (
	structCommentTemplate = "localizer:"
	structSkipComment     = "localizer:skip"
)

// Parser анализатор абстрактного дерева кода.
type Parser struct {
	PkgPath     string
	PkgName     string
	StructNames []StructInfo
	AllStructs  bool
}

// StructInfo содержит информацию о структуре.
type StructInfo struct {
	Name              string
	LocalizationAlias string
}

// Parse разбирает содержимое указанного файла.
func (p *Parser) Parse(fileName string, isDirectory bool) error {
	if packagePath, err := getPkgPath(fileName, isDirectory); err != nil {
		return err
	} else {
		p.PkgPath = packagePath
	}

	fileSet := token.NewFileSet()
	if isDirectory {
		packages, err := parser.ParseDir(fileSet, fileName, excludeTestFiles, parser.ParseComments)
		if err != nil {
			return err
		}

		for _, pkg := range packages {
			ast.Walk(&visitor{Parser: p}, pkg)
		}
	} else {
		file, err := parser.ParseFile(fileSet, fileName, nil, parser.ParseComments)
		if err != nil {
			return err
		}

		ast.Walk(&visitor{Parser: p}, file)
	}

	return nil
}

func excludeTestFiles(fi os.FileInfo) bool {
	return !strings.HasSuffix(fi.Name(), "_test.go")
}

func getPkgPath(fileName string, isDirectory bool) (string, error) {
	if !filepath.IsAbs(fileName) {
		pwd, err := os.Getwd()
		if err != nil {
			return "", err
		}
		fileName = filepath.Join(pwd, fileName)
	}

	goModPath, _ := goModPath(fileName, isDirectory)
	if strings.Contains(goModPath, "go.mod") {
		pkgPath, err := getPkgPathFromGoMod(fileName, isDirectory, goModPath)
		if err != nil {
			return "", err
		}

		return pkgPath, nil
	}

	return getPkgPathFromGOPATH(fileName, isDirectory)
}

var goModPathCache = struct {
	paths map[string]string
	sync.RWMutex
}{
	paths: make(map[string]string),
}

func goModPath(fileName string, isDirectory bool) (string, error) {
	root := fileName
	if !isDirectory {
		root = filepath.Dir(fileName)
	}

	goModPathCache.RLock()
	goModPath, ok := goModPathCache.paths[root]
	goModPathCache.RUnlock()
	if ok {
		return goModPath, nil
	}

	defer func() {
		goModPathCache.Lock()
		goModPathCache.paths[root] = goModPath
		goModPathCache.Unlock()
	}()

	cmd := exec.Command("go", "env", "GOMOD")
	cmd.Dir = root

	stdout, err := cmd.Output()
	if err != nil {
		return "", err
	}

	goModPath = string(bytes.TrimSpace(stdout))

	return goModPath, nil
}

func getPkgPathFromGoMod(fileName string, isDirectory bool, goModPath string) (string, error) {
	modulePath := getModulePath(goModPath)
	if modulePath == "" {
		return "", fmt.Errorf("cannot determine module path from %s", goModPath)
	}

	rel := path.Join(modulePath, filePathToPackagePath(strings.TrimPrefix(fileName, filepath.Dir(goModPath))))

	if !isDirectory {
		return path.Dir(rel), nil
	}

	return path.Clean(rel), nil
}

var pkgPathFromGoModCache = struct {
	paths map[string]string
	sync.RWMutex
}{
	paths: make(map[string]string),
}

func getModulePath(goModPath string) string {
	pkgPathFromGoModCache.RLock()
	pkgPath, ok := pkgPathFromGoModCache.paths[goModPath]
	pkgPathFromGoModCache.RUnlock()
	if ok {
		return pkgPath
	}

	defer func() {
		pkgPathFromGoModCache.Lock()
		pkgPathFromGoModCache.paths[goModPath] = pkgPath
		pkgPathFromGoModCache.Unlock()
	}()

	data, err := ioutil.ReadFile(goModPath)
	if err != nil {
		return ""
	}
	pkgPath = modulePath(data)
	return pkgPath
}

func getPkgPathFromGOPATH(fileName string, isDirectory bool) (string, error) {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}

	for _, p := range strings.Split(gopath, string(filepath.ListSeparator)) {
		prefix := filepath.Join(p, "src") + string(filepath.Separator)
		rel, err := filepath.Rel(prefix, fileName)
		if err == nil && !strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
			if !isDirectory {
				return path.Dir(filePathToPackagePath(rel)), nil
			} else {
				return path.Clean(filePathToPackagePath(rel)), nil
			}
		}
	}

	return "", fmt.Errorf("file '%v' is not in GOPATH '%v'", fileName, gopath)
}

func filePathToPackagePath(path string) string {
	return filepath.ToSlash(path)
}
