package mysql

// Framework code is generated by the generator.

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/pkg/errors"

	mysql "github.com/bytebase/mysql-parser"

	"github.com/bytebase/bytebase/backend/plugin/advisor"
	mysqlparser "github.com/bytebase/bytebase/backend/plugin/parser/mysql"
	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
)

var (
	_ advisor.Advisor = (*ColumnDisallowChangingOrderAdvisor)(nil)
)

func init() {
	advisor.Register(storepb.Engine_MYSQL, advisor.MySQLColumnDisallowChangingOrder, &ColumnDisallowChangingOrderAdvisor{})
	advisor.Register(storepb.Engine_MARIADB, advisor.MySQLColumnDisallowChangingOrder, &ColumnDisallowChangingOrderAdvisor{})
	advisor.Register(storepb.Engine_OCEANBASE, advisor.MySQLColumnDisallowChangingOrder, &ColumnDisallowChangingOrderAdvisor{})
}

// ColumnDisallowChangingOrderAdvisor is the advisor checking for disallow changing column order.
type ColumnDisallowChangingOrderAdvisor struct {
}

// Check checks for disallow changing column order.
func (*ColumnDisallowChangingOrderAdvisor) Check(ctx advisor.Context, _ string) ([]*storepb.Advice, error) {
	stmtList, ok := ctx.AST.([]*mysqlparser.ParseResult)
	if !ok {
		return nil, errors.Errorf("failed to convert to mysql parser result")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(ctx.Rule.Level)
	if err != nil {
		return nil, err
	}
	checker := &columnDisallowChangingOrderChecker{
		level: level,
		title: string(ctx.Rule.Type),
	}

	for _, stmtNode := range stmtList {
		checker.baseLine = stmtNode.BaseLine
		antlr.ParseTreeWalkerDefault.Walk(checker, stmtNode.Tree)
	}

	return checker.adviceList, nil
}

type columnDisallowChangingOrderChecker struct {
	*mysql.BaseMySQLParserListener

	baseLine   int
	adviceList []*storepb.Advice
	level      storepb.Advice_Status
	title      string
	text       string
}

func (checker *columnDisallowChangingOrderChecker) EnterQuery(ctx *mysql.QueryContext) {
	checker.text = ctx.GetParser().GetTokenStream().GetTextFromRuleContext(ctx)
}

func (checker *columnDisallowChangingOrderChecker) EnterAlterTable(ctx *mysql.AlterTableContext) {
	if !mysqlparser.IsTopMySQLRule(&ctx.BaseParserRuleContext) {
		return
	}
	if ctx.AlterTableActions() == nil {
		return
	}
	if ctx.AlterTableActions().AlterCommandList() == nil {
		return
	}
	if ctx.AlterTableActions().AlterCommandList().AlterList() == nil {
		return
	}

	for _, item := range ctx.AlterTableActions().AlterCommandList().AlterList().AllAlterListItem() {
		if item == nil {
			continue
		}

		switch {
		// modify column.
		case item.MODIFY_SYMBOL() != nil && item.ColumnInternalRef() != nil:
			// do nothing.
		// change column
		case item.CHANGE_SYMBOL() != nil && item.ColumnInternalRef() != nil && item.Identifier() != nil:
			// do nothing.
		default:
			continue
		}

		if item.Place() != nil {
			checker.adviceList = append(checker.adviceList, &storepb.Advice{
				Status:  checker.level,
				Code:    advisor.ChangeColumnOrder.Int32(),
				Title:   checker.title,
				Content: fmt.Sprintf("\"%s\" changes column order", checker.text),
				StartPosition: &storepb.Position{
					Line: int32(checker.baseLine + ctx.GetStart().GetLine()),
				},
			})
		}
	}
}
