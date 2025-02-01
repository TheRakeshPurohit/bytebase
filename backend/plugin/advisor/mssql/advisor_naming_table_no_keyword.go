// Package mssql is the advisor for MSSQL database.
package mssql

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/bytebase/tsql-parser"
	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/plugin/advisor"
	tsqlparser "github.com/bytebase/bytebase/backend/plugin/parser/tsql"
	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
)

var (
	_ advisor.Advisor = (*NamingTableNoKeywordAdvisor)(nil)
)

func init() {
	advisor.Register(storepb.Engine_MSSQL, advisor.MSSQLTableNamingNoKeyword, &NamingTableNoKeywordAdvisor{})
}

// NamingTableNoKeywordAdvisor is the advisor checking for table naming convention without keyword..
type NamingTableNoKeywordAdvisor struct {
}

// Check checks for table naming convention without keyword..
func (*NamingTableNoKeywordAdvisor) Check(ctx advisor.Context, _ string) ([]*storepb.Advice, error) {
	tree, ok := ctx.AST.(antlr.Tree)
	if !ok {
		return nil, errors.Errorf("failed to convert to Tree")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(ctx.Rule.Level)
	if err != nil {
		return nil, err
	}

	listener := &namingTableNoKeywordChecker{
		level: level,
		title: string(ctx.Rule.Type),
	}

	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	return listener.generateAdvice()
}

// namingTableNoKeywordChecker is the listener for table naming convention without keyword.
type namingTableNoKeywordChecker struct {
	*parser.BaseTSqlParserListener

	level storepb.Advice_Status
	title string

	adviceList []*storepb.Advice
}

// generateAdvice returns the advices generated by the listener, the advices must not be empty.
func (l *namingTableNoKeywordChecker) generateAdvice() ([]*storepb.Advice, error) {
	return l.adviceList, nil
}

// EnterCreate_table is called when production create_table is entered.
func (l *namingTableNoKeywordChecker) EnterCreate_table(ctx *parser.Create_tableContext) {
	tableName := ctx.Table_name().GetTable()
	_, normalizedTableName := tsqlparser.NormalizeTSQLIdentifier(tableName)
	if tsqlparser.IsTSQLReservedKeyword(normalizedTableName, false) {
		l.adviceList = append(l.adviceList, &storepb.Advice{
			Status:  l.level,
			Code:    advisor.NameIsKeywordIdentifier.Int32(),
			Title:   l.title,
			Content: fmt.Sprintf("Table name [%s] is a reserved keyword and should be avoided.", normalizedTableName),
			StartPosition: &storepb.Position{
				Line: int32(tableName.GetStart().GetLine()),
			},
		})
	}
}
