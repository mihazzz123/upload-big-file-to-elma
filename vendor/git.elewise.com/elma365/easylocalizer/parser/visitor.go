package parser

import (
	"go/ast"
	"strings"
)

type visitor struct {
	*Parser

	name string
}

// Visit реализует интерфейс ast.Visitor.
func (v *visitor) Visit(node ast.Node) (w ast.Visitor) {
	switch node := node.(type) {
	case *ast.Package:
		return v
	case *ast.File:
		v.PkgName = node.Name.String()
		return v
	case *ast.GenDecl:
		skip, explicit, _ := v.needType(node.Doc)
		if skip || explicit {
			for _, nodeSpec := range node.Specs {
				switch nodeSpecType := nodeSpec.(type) {
				case *ast.TypeSpec:
					nodeSpecType.Doc = node.Doc
				}
			}
		}

		return v
	case *ast.TypeSpec:
		skip, explicit, alias := v.needType(node.Doc)
		if skip {
			return nil
		}
		if !explicit && !v.AllStructs {
			return nil
		}

		v.name = node.Name.String()

		if explicit {
			v.StructNames = append(v.StructNames, StructInfo{
				Name:              v.name,
				LocalizationAlias: alias,
			})
			return nil
		}

		return v
	case *ast.StructType:
		v.StructNames = append(v.StructNames, StructInfo{
			Name:              v.name,
			LocalizationAlias: "",
		})
		return nil
	}

	return nil
}

func (p *Parser) needType(comments *ast.CommentGroup) (skip, explicit bool, alias string) {
	if comments == nil {
		return
	}

	for _, c := range comments.List {
		commentText := c.Text

		if len(commentText) > 2 {
			switch commentText[1] {
			case '/':
				commentText = commentText[2:]
			case '*':
				commentText = commentText[2 : len(commentText)-2]
			}
		}

		for _, comment := range strings.Split(commentText, "\n") {
			comment = strings.TrimSpace(comment)

			if strings.HasPrefix(comment, structSkipComment) {
				return true, false, ""
			}
			if strings.HasPrefix(comment, structCommentTemplate) {
				alias := strings.Replace(comment, structCommentTemplate, "", -1)
				alias = strings.Split(alias, " ")[0]
				alias = strings.TrimSpace(alias)
				return false, true, alias
			}
		}
	}

	return
}
