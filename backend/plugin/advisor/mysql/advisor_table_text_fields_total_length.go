package mysql

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/pkg/errors"

	mysql "github.com/bytebase/mysql-parser"

	"github.com/bytebase/bytebase/backend/common"
	storepb "github.com/bytebase/bytebase/backend/generated-go/store"
	"github.com/bytebase/bytebase/backend/plugin/advisor"
	"github.com/bytebase/bytebase/backend/plugin/advisor/catalog"
	mysqlparser "github.com/bytebase/bytebase/backend/plugin/parser/mysql"
)

var (
	_ advisor.Advisor = (*TableMaximumVarcharLengthAdvisor)(nil)
)

func init() {
	advisor.Register(storepb.Engine_MYSQL, advisor.MySQLTableTextFieldsTotalLength, &TableMaximumVarcharLengthAdvisor{})
}

type TableMaximumVarcharLengthAdvisor struct {
}

func (*TableMaximumVarcharLengthAdvisor) Check(_ context.Context, checkCtx advisor.Context) ([]*storepb.Advice, error) {
	stmtList, ok := checkCtx.AST.([]*mysqlparser.ParseResult)
	if !ok {
		return nil, errors.Errorf("failed to convert to mysql parse result")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(checkCtx.Rule.Level)
	if err != nil {
		return nil, err
	}
	payload, err := advisor.UnmarshalNumberTypeRulePayload(checkCtx.Rule.Payload)
	if err != nil {
		return nil, err
	}
	checker := &tableFieldsMaximumVarcharLengthChecker{
		level:   level,
		title:   string(checkCtx.Rule.Type),
		catalog: checkCtx.Catalog,
		maximum: payload.Number,
	}

	for _, stmt := range stmtList {
		checker.baseLine = stmt.BaseLine
		antlr.ParseTreeWalkerDefault.Walk(checker, stmt.Tree)
	}

	return checker.adviceList, nil
}

type tableFieldsMaximumVarcharLengthChecker struct {
	*mysql.BaseMySQLParserListener

	catalog    *catalog.Finder
	adviceList []*storepb.Advice
	level      storepb.Advice_Status
	title      string
	baseLine   int
	maximum    int
}

func (checker *tableFieldsMaximumVarcharLengthChecker) EnterCreateTable(ctx *mysql.CreateTableContext) {
	if !mysqlparser.IsTopMySQLRule(&ctx.BaseParserRuleContext) {
		return
	}
	if ctx.TableElementList() == nil || ctx.TableName() == nil {
		return
	}

	_, tableName := mysqlparser.NormalizeMySQLTableName(ctx.TableName())
	if tableName == "" {
		return
	}
	tableInfo := checker.catalog.Final.FindTable(&catalog.TableFind{TableName: tableName})
	if tableInfo == nil {
		return
	}
	total := getTotalTextLength(tableInfo)
	if total > int64(checker.maximum) {
		checker.adviceList = append(checker.adviceList, &storepb.Advice{
			Status:        checker.level,
			Code:          advisor.IndexCountExceedsLimit.Int32(),
			Title:         checker.title,
			Content:       fmt.Sprintf("Table %q total text column length (%d) exceeds the limit (%d).", tableName, total, checker.maximum),
			StartPosition: common.ConvertANTLRLineToPosition(checker.baseLine + ctx.GetStart().GetLine()),
		})
	}
}

func (checker *tableFieldsMaximumVarcharLengthChecker) EnterAlterTable(ctx *mysql.AlterTableContext) {
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

	_, tableName := mysqlparser.NormalizeMySQLTableRef(ctx.TableRef())
	if tableName == "" {
		return
	}
	tableInfo := checker.catalog.Final.FindTable(&catalog.TableFind{TableName: tableName})
	if tableInfo == nil {
		return
	}
	total := getTotalTextLength(tableInfo)
	if total > int64(checker.maximum) {
		checker.adviceList = append(checker.adviceList, &storepb.Advice{
			Status:        checker.level,
			Code:          advisor.TotalTextLengthExceedsLimit.Int32(),
			Title:         checker.title,
			Content:       fmt.Sprintf("Table %q total text column length (%d) exceeds the limit (%d).", tableName, total, checker.maximum),
			StartPosition: common.ConvertANTLRLineToPosition(checker.baseLine + ctx.GetStart().GetLine()),
		})
	}
}

func getTotalTextLength(tableInfo *catalog.TableState) int64 {
	var total int64
	columns := tableInfo.ListColumns()
	for _, column := range columns {
		total += getTextLength(column.Type())
	}
	return total
}

func getTextLength(s string) int64 {
	s = strings.ToLower(s)
	switch s {
	case "char", "binary":
		return 1
	case "tinyblob", "tinytext":
		return 255
	case "blob", "text":
		return 65_535
	case "mediumblob", "mediumtext":
		return 16_777_215
	case "longblob", "longtext":
		return 4_294_967_295
	default:
		re := regexp.MustCompile(`[a-z]+\((\d+)\)`)
		match := re.FindStringSubmatch(s)
		if len(match) >= 2 {
			n, err := strconv.ParseInt(match[1], 10, 64)
			if err == nil {
				return int64(n)
			}
		}
	}
	return 0
}
