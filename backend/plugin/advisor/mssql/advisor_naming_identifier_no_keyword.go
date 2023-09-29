// Package mssql is the advisor for MSSQL database.
package mssql

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/bytebase/tsql-parser"
	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/plugin/advisor"
	"github.com/bytebase/bytebase/backend/plugin/advisor/db"
	tsqlparser "github.com/bytebase/bytebase/backend/plugin/parser/tsql"
)

var (
	_ advisor.Advisor = (*NamingIdentifierNoKeywordAdvisor)(nil)
)

func init() {
	advisor.Register(db.MSSQL, advisor.MSSQLIdentifierNamingNoKeyword, &NamingIdentifierNoKeywordAdvisor{})
}

// NamingIdentifierNoKeywordAdvisor is the advisor checking for identifier naming convention without keyword..
type NamingIdentifierNoKeywordAdvisor struct {
}

// Check checks for identifier naming convention without keyword..
func (*NamingIdentifierNoKeywordAdvisor) Check(ctx advisor.Context, _ string) ([]advisor.Advice, error) {
	tree, ok := ctx.AST.(antlr.Tree)
	if !ok {
		return nil, errors.Errorf("failed to convert to Tree")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(ctx.Rule.Level)
	if err != nil {
		return nil, err
	}

	listener := &namingIdentifierNoKeywordChecker{
		level: level,
		title: string(ctx.Rule.Type),
	}

	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	return listener.generateAdvice()
}

// namingIdentifierNoKeywordChecker is the listener for identifier naming convention without keyword.
type namingIdentifierNoKeywordChecker struct {
	*parser.BaseTSqlParserListener

	level advisor.Status
	title string

	adviceList []advisor.Advice
}

// generateAdvice returns the advices generated by the listener, the advices must not be empty.
func (l *namingIdentifierNoKeywordChecker) generateAdvice() ([]advisor.Advice, error) {
	if len(l.adviceList) == 0 {
		l.adviceList = append(l.adviceList, advisor.Advice{
			Status:  advisor.Success,
			Code:    advisor.Ok,
			Title:   "OK",
			Content: "",
		})
	}
	return l.adviceList, nil
}

// EnterId_ is called when production id_ is entered.
func (l *namingIdentifierNoKeywordChecker) EnterId_(ctx *parser.Id_Context) {
	if ctx == nil {
		return
	}

	parent := ctx.GetParent()
	switch parent.(type) {
	case *parser.Column_definitionContext:
	case *parser.Table_constraintContext:
	case *parser.Create_schemaContext:
	case *parser.Create_databaseContext:
	case *parser.Create_indexContext:
	case *parser.Table_nameContext:
	default:
		return
	}
	if ctx.GetText() == "" {
		return
	}

	normalizedID := tsqlparser.NormalizeTSQLIdentifier(ctx)
	if tsqlparser.IsTSQLKeyword(normalizedID, false) {
		l.adviceList = append(l.adviceList, advisor.Advice{
			Status:  l.level,
			Code:    advisor.NameIsKeywordIdentifier,
			Title:   l.title,
			Content: fmt.Sprintf("Identifier [%s] is a keyword identifier and should be avoided.", normalizedID),
			Line:    ctx.GetStart().GetLine(),
		})
	}
}
