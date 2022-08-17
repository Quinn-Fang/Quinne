package handler

import (
	"quinn007.com/parser"
)

func ArgumentsContextHandler(contextParser *parser.ArgumentsContext) error {
	children := contextParser.GetChildren()

	for _, child := range children {
		switch parserContext := child.(type) {
		case *parser.ExpressionListContext:
			{
				ExpressionListContextHandler(parserContext)
			}
		}
	}

	return nil
}