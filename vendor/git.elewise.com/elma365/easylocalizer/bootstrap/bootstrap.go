package bootstrap

import (
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"git.elewise.com/elma365/easylocalizer/parser"
)

const (
	tempFilePostfix = ".tmp"
	pkgGenerator    = "git.elewise.com/elma365/easylocalizer/generator"
	PkgGoText       = "github.com/leonelquinteros/gotext"
	PkgCommonI18n   = "git.elewise.com/elma365/common/pkg/i18n"
	PkgErrors       = "github.com/pkg/errors"
)

// Generator генератор Go файлов локализации.
type Generator struct {
	PkgPath, PkgName string
	Types            []parser.StructInfo
	OutName          string
}

// Run выполняет генерацию кода.
func (g *Generator) Run() error {
	if err := g.writeStub(); err != nil {
		return err
	}

	path, err := g.writeMain()
	if err != nil {
		return err
	}
	defer func() {
		dir := filepath.Dir(path)
		os.RemoveAll(dir)
	}()

	file, err := os.Create(g.OutName + tempFilePostfix)
	if err != nil {
		return err
	}
	defer os.Remove(file.Name())

	execArgs := []string{"run", filepath.Base(path)}
	cmd := exec.Command("go", execArgs...)
	cmd.Stdout = file
	cmd.Stderr = os.Stderr
	cmd.Dir = filepath.Dir(path)
	if err = cmd.Run(); err != nil {
		return err
	}
	file.Close()

	in, err := ioutil.ReadFile(file.Name())
	if err != nil {
		return err
	}
	out, err := format.Source(in)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(g.OutName, out, 0644)
}

func (g *Generator) writeStub() error {
	file, err := os.Create(g.OutName)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Fprintln(file, "// TEMPORARY AUTOGENERATED FILE: localizer stub code to make the package")
	fmt.Fprintln(file, "// compilable during generation.")
	fmt.Fprintln(file)
	fmt.Fprintln(file, "package ", g.PkgName)

	if len(g.Types) > 0 {
		fmt.Fprintln(file)
		fmt.Fprintln(file, "import (")
		fmt.Fprintln(file, `  "`+PkgGoText+`"`)
		fmt.Fprintln(file, `  "`+PkgCommonI18n+`"`)
		fmt.Fprintln(file, ")")
	}

	for _, t := range g.Types {
		fmt.Fprintln(file)
		fmt.Fprintln(file, "func (entity ", t.Name, ") ExtractPO(poCtxt *i18n.EntityPOContext) ([]byte, error) { return nil, nil }")
		fmt.Fprintln(file, "func (entity *", t.Name, ") ApplyTranslation(poCtxt *i18n.EntityPOContext, translator gotext.Translator) { return }")
		fmt.Fprintln(file, "type Localizer_exporter_"+t.Name+" *"+t.Name)
		fmt.Fprintln(file)
	}
	return nil
}

func (g *Generator) writeMain() (path string, err error) {
	fileDir := filepath.Join(filepath.Dir(g.OutName), "main")
	if _, dirErr := os.Stat(fileDir); !os.IsNotExist(dirErr) {
		err = os.RemoveAll(fileDir)
		if err != nil {
			return "", err
		}
	}
	err = os.Mkdir(fileDir, 0755)
	if err != nil {
		return "", err
	}
	file, err := ioutil.TempFile(fileDir, "localizer-bootstrap")
	if err != nil {
		return "", err
	}

	fmt.Fprintln(file, "// TEMPORARY AUTOGENERATED FILE: localizer bootstapping code to launch")
	fmt.Fprintln(file, "// the actual generator.")
	fmt.Fprintln(file)
	fmt.Fprintln(file, "package main")
	fmt.Fprintln(file)
	fmt.Fprintln(file, "import (")
	fmt.Fprintln(file, `  "fmt"`)
	fmt.Fprintln(file, `  "os"`)
	fmt.Fprintln(file)
	fmt.Fprintf(file, "  %q\n", pkgGenerator)
	if len(g.Types) > 0 {
		fmt.Fprintln(file)
		fmt.Fprintf(file, "  pkg %q\n", g.PkgPath)
	}
	fmt.Fprintln(file, ")")
	fmt.Fprintln(file)

	fmt.Fprintln(file, "func main() {")
	fmt.Fprintf(file, "  g := generator.New(%q)\n", filepath.Base(g.OutName))
	fmt.Fprintf(file, "  g.SetPkg(%q, %q)\n", g.PkgName, g.PkgPath)
	if len(g.Types) > 0 {
		for _, t := range g.Types {
			fmt.Fprintf(file, "  g.AddStruct(%q, %q, %s)\n", t.Name, t.LocalizationAlias, "pkg.Localizer_exporter_"+t.Name+"(nil)")
		}
	}
	fmt.Fprintln(file, "  if err := g.Run(os.Stdout); err != nil {")
	fmt.Fprintln(file, "    fmt.Fprintln(os.Stderr, err)")
	fmt.Fprintln(file, "    os.Exit(1)")
	fmt.Fprintln(file, "  }")
	fmt.Fprintln(file, "}")

	src := file.Name()
	if err = file.Close(); err != nil {
		return src, err
	}

	dest := src + ".go"
	return dest, os.Rename(src, dest)
}
