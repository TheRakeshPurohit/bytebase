// Package snowflake is the advisor for snowflake database.
package snowflake

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	parser "github.com/bytebase/snowsql-parser"
	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/plugin/advisor"
	"github.com/bytebase/bytebase/backend/plugin/advisor/db"
	snowsqlparser "github.com/bytebase/bytebase/backend/plugin/parser/snowflake"
)

var (
	_ advisor.Advisor = (*NamingTableNoKeywordAdvisor)(nil)
)

func init() {
	advisor.Register(db.Snowflake, advisor.SnowflakeTableNamingNoKeyword, &NamingTableNoKeywordAdvisor{})
}

// NamingTableNoKeywordAdvisor is the advisor checking for table naming convention without keyword.
type NamingTableNoKeywordAdvisor struct {
}

// Check checks for table naming convention without keyword.
func (*NamingTableNoKeywordAdvisor) Check(ctx advisor.Context, _ string) ([]advisor.Advice, error) {
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
	*parser.BaseSnowflakeParserListener

	level advisor.Status
	title string

	adviceList []advisor.Advice
}

// generateAdvice returns the advices generated by the listener, the advices must not be empty.
func (l *namingTableNoKeywordChecker) generateAdvice() ([]advisor.Advice, error) {
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

// EnterCreate_table is called when production create_table is entered.
func (l *namingTableNoKeywordChecker) EnterCreate_table(ctx *parser.Create_tableContext) {
	tableName := snowsqlparser.NormalizeSnowSQLObjectNamePart(ctx.Object_name().GetO())
	if snowsqlparser.IsSnowflakeKeyword(tableName, false) {
		l.adviceList = append(l.adviceList, advisor.Advice{
			Status:  l.level,
			Code:    advisor.NameIsKeywordIdentifier,
			Title:   l.title,
			Content: fmt.Sprintf("Table name %q is a keyword identifier and should be avoided.", tableName),
		})
	}
}

// EnterAlter_table is called when production alter_table is entered.
func (l *namingTableNoKeywordChecker) EnterAlter_table(ctx *parser.Alter_tableContext) {
	if ctx.RENAME() == nil {
		return
	}

	tableName := snowsqlparser.NormalizeSnowSQLObjectNamePart(ctx.Object_name(1).GetO())
	if snowsqlparser.IsSnowflakeKeyword(tableName, false) {
		l.adviceList = append(l.adviceList, advisor.Advice{
			Status:  l.level,
			Code:    advisor.NameIsKeywordIdentifier,
			Title:   l.title,
			Content: fmt.Sprintf("Table name %q is a keyword identifier and should be avoided.", tableName),
		})
	}
}
