package handler

import (
	"quinn007.com/navigator"
	"quinn007.com/parser"
	"quinn007.com/sym_tables"
)

func SimpleStmtContextHandler(contextParser *parser.SimpleStmtContext) error {
	curCursor, _ := navigator.GetCursor()
	curCursor.NewStatement()
	curStatement := curCursor.GetStatement()
	children := contextParser.GetChildren()

	for _, child := range children {
		switch parserContext := child.(type) {
		case *parser.ShortVarDeclContext:
			{
				ShortVarDeclContextHandler(parserContext)
			}
		case *parser.VarSpecContext:
			{
				VarSpecContextHandler(parserContext)
			}
		case *parser.ExpressionStmtContext:
			{
				ExpressionStmtContextHandler(parserContext)
			}
		}
	}
	rightValues := curStatement.GetRightValues()
	curSymTable := sym_tables.GetCurSymTable()
	curStatement.Assign()

	// add assigned variable to symbol table
	for _, variable := range rightValues {
		curSymTable.AddVariable(variable)
	}

	return nil
}