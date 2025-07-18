package mysql

import (
	"context"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/pkg/errors"

	mysql "github.com/bytebase/mysql-parser"

	"github.com/bytebase/bytebase/backend/common"
	storepb "github.com/bytebase/bytebase/backend/generated-go/store"
	"github.com/bytebase/bytebase/backend/plugin/advisor"
	mysqlparser "github.com/bytebase/bytebase/backend/plugin/parser/mysql"
)

var (
	_ advisor.Advisor = (*StatementWhereDisallowUsingFunctionAdvisor)(nil)
)

func init() {
	advisor.Register(storepb.Engine_MYSQL, advisor.MySQLStatementWhereDisallowUsingFunction, &StatementWhereDisallowUsingFunctionAdvisor{})
}

type StatementWhereDisallowUsingFunctionAdvisor struct {
}

func (*StatementWhereDisallowUsingFunctionAdvisor) Check(_ context.Context, checkCtx advisor.Context) ([]*storepb.Advice, error) {
	stmtList, ok := checkCtx.AST.([]*mysqlparser.ParseResult)
	if !ok {
		return nil, errors.Errorf("failed to convert to mysql parse result")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(checkCtx.Rule.Level)
	if err != nil {
		return nil, err
	}
	checker := &statementWhereDisallowUsingFunctionChecker{
		level: level,
		title: string(checkCtx.Rule.Type),
	}

	for _, stmt := range stmtList {
		checker.baseLine = stmt.BaseLine
		antlr.ParseTreeWalkerDefault.Walk(checker, stmt.Tree)
	}

	return checker.adviceList, nil
}

type statementWhereDisallowUsingFunctionChecker struct {
	*mysql.BaseMySQLParserListener

	baseLine      int
	adviceList    []*storepb.Advice
	level         storepb.Advice_Status
	title         string
	text          string
	isSelect      bool
	inWhereClause bool
}

func (checker *statementWhereDisallowUsingFunctionChecker) EnterQuery(ctx *mysql.QueryContext) {
	checker.text = ctx.GetParser().GetTokenStream().GetTextFromRuleContext(ctx)
}

func (checker *statementWhereDisallowUsingFunctionChecker) EnterSelectStatement(*mysql.SelectStatementContext) {
	checker.isSelect = true
}

func (checker *statementWhereDisallowUsingFunctionChecker) ExitSelectStatement(*mysql.SelectStatementContext) {
	checker.isSelect = false
}

func (checker *statementWhereDisallowUsingFunctionChecker) EnterWhereClause(*mysql.WhereClauseContext) {
	checker.inWhereClause = true
}

func (checker *statementWhereDisallowUsingFunctionChecker) ExitWhereClause(*mysql.WhereClauseContext) {
	checker.inWhereClause = false
}

func (checker *statementWhereDisallowUsingFunctionChecker) EnterFunctionCall(ctx *mysql.FunctionCallContext) {
	if !checker.isSelect || !checker.inWhereClause {
		return
	}

	pi := ctx.PureIdentifier()
	if pi != nil {
		checker.adviceList = append(checker.adviceList, &storepb.Advice{
			Status:        checker.level,
			Code:          advisor.DisabledFunction.Int32(),
			Title:         checker.title,
			Content:       fmt.Sprintf("Function is disallowed in where clause, but \"%s\" uses", checker.text),
			StartPosition: common.ConvertANTLRLineToPosition(checker.baseLine + ctx.GetStart().GetLine()),
		})
	}
}
