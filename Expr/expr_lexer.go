// Code generated from Expr.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 11, 50, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4,
	6, 4, 27, 10, 4, 13, 4, 14, 4, 28, 3, 5, 6, 5, 32, 10, 5, 13, 5, 14, 5,
	33, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 6, 10, 45, 10,
	10, 13, 10, 14, 10, 46, 3, 10, 3, 10, 2, 2, 11, 3, 3, 5, 4, 7, 5, 9, 6,
	11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 3, 2, 5, 4, 2, 12, 12, 15, 15, 3,
	2, 50, 59, 4, 2, 11, 11, 34, 34, 2, 52, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2,
	2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2,
	2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 3, 21, 3,
	2, 2, 2, 5, 23, 3, 2, 2, 2, 7, 26, 3, 2, 2, 2, 9, 31, 3, 2, 2, 2, 11, 35,
	3, 2, 2, 2, 13, 37, 3, 2, 2, 2, 15, 39, 3, 2, 2, 2, 17, 41, 3, 2, 2, 2,
	19, 44, 3, 2, 2, 2, 21, 22, 7, 42, 2, 2, 22, 4, 3, 2, 2, 2, 23, 24, 7,
	43, 2, 2, 24, 6, 3, 2, 2, 2, 25, 27, 9, 2, 2, 2, 26, 25, 3, 2, 2, 2, 27,
	28, 3, 2, 2, 2, 28, 26, 3, 2, 2, 2, 28, 29, 3, 2, 2, 2, 29, 8, 3, 2, 2,
	2, 30, 32, 9, 3, 2, 2, 31, 30, 3, 2, 2, 2, 32, 33, 3, 2, 2, 2, 33, 31,
	3, 2, 2, 2, 33, 34, 3, 2, 2, 2, 34, 10, 3, 2, 2, 2, 35, 36, 7, 45, 2, 2,
	36, 12, 3, 2, 2, 2, 37, 38, 7, 47, 2, 2, 38, 14, 3, 2, 2, 2, 39, 40, 7,
	44, 2, 2, 40, 16, 3, 2, 2, 2, 41, 42, 7, 49, 2, 2, 42, 18, 3, 2, 2, 2,
	43, 45, 9, 4, 2, 2, 44, 43, 3, 2, 2, 2, 45, 46, 3, 2, 2, 2, 46, 44, 3,
	2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 49, 8, 10, 2, 2, 49,
	20, 3, 2, 2, 2, 6, 2, 28, 33, 46, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'('", "')'", "", "", "'+'", "'-'", "'*'", "'/'",
}

var lexerSymbolicNames = []string{
	"", "", "", "NEWLINE", "INT", "SUM", "SUB", "MULTI", "DIV", "WHITESPACE",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "NEWLINE", "INT", "SUM", "SUB", "MULTI", "DIV", "WHITESPACE",
}

type ExprLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewExprLexer(input antlr.CharStream) *ExprLexer {

	l := new(ExprLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Expr.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ExprLexer tokens.
const (
	ExprLexerT__0       = 1
	ExprLexerT__1       = 2
	ExprLexerNEWLINE    = 3
	ExprLexerINT        = 4
	ExprLexerSUM        = 5
	ExprLexerSUB        = 6
	ExprLexerMULTI      = 7
	ExprLexerDIV        = 8
	ExprLexerWHITESPACE = 9
)
