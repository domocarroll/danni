package styles

import (
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/glamour/ansi"
)

// GetGlamourRenderer returns a Glamour renderer with Danni theming
// Hot Pink + Gold + Black aesthetic for markdown rendering
func GetGlamourRenderer(width int) (*glamour.TermRenderer, error) {
	styleConfig := ansi.StyleConfig{
		Document: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Color: stringPtr("#FFFFFF"), // White - clear text
			},
		},
		Heading: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Color:       stringPtr("#FF1493"), // HotPink - bold headers
				Bold:        boolPtr(true),
				BlockPrefix: "\n",
				BlockSuffix: "\n",
			},
		},
		H1: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Color: stringPtr("#FF1493"), // HotPink - primary headers
				Bold:  boolPtr(true),
			},
		},
		H2: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Color: stringPtr("#FFD700"), // Gold - secondary headers
				Bold:  boolPtr(true),
			},
		},
		H3: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Color: stringPtr("#FF69B4"), // HotPinkAlt - tertiary headers
				Bold:  boolPtr(true),
			},
		},
		H4: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Color: stringPtr("#DAA520"), // GoldMuted - subtle headers
				Bold:  boolPtr(true),
			},
		},
		Code: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Color:           stringPtr("#FFD700"),  // Gold - inline code
				BackgroundColor: stringPtr("#1A1A1A"),  // BlackRich
			},
		},
		CodeBlock: ansi.StyleCodeBlock{
			StyleBlock: ansi.StyleBlock{
				StylePrimitive: ansi.StylePrimitive{
					Color:           stringPtr("#FFFFFF"),  // White
					BackgroundColor: stringPtr("#0A0A0A"),  // BlackSoft
				},
			},
			Chroma: &ansi.Chroma{
				Text:                ansi.StylePrimitive{Color: stringPtr("#FFFFFF")},
				Error:               ansi.StylePrimitive{Color: stringPtr("#FF1493")},
				Comment:             ansi.StylePrimitive{Color: stringPtr("#808080")},
				CommentPreproc:      ansi.StylePrimitive{Color: stringPtr("#FF69B4")},
				Keyword:             ansi.StylePrimitive{Color: stringPtr("#FF1493"), Bold: boolPtr(true)},
				KeywordType:         ansi.StylePrimitive{Color: stringPtr("#FFD700")},
				Operator:            ansi.StylePrimitive{Color: stringPtr("#FF1493")},
				Punctuation:         ansi.StylePrimitive{Color: stringPtr("#CCCCCC")},
				Name:                ansi.StylePrimitive{Color: stringPtr("#FFFFFF")},
				NameBuiltin:         ansi.StylePrimitive{Color: stringPtr("#FF69B4")},
				NameTag:             ansi.StylePrimitive{Color: stringPtr("#FF1493")},
				NameAttribute:       ansi.StylePrimitive{Color: stringPtr("#FFD700")},
				NameClass:           ansi.StylePrimitive{Color: stringPtr("#FFD700"), Bold: boolPtr(true)},
				NameConstant:        ansi.StylePrimitive{Color: stringPtr("#FFA500")},
				NameDecorator:       ansi.StylePrimitive{Color: stringPtr("#FF69B4")},
				NameFunction:        ansi.StylePrimitive{Color: stringPtr("#FFD700")},
				LiteralNumber:       ansi.StylePrimitive{Color: stringPtr("#FFA500")},
				LiteralString:       ansi.StylePrimitive{Color: stringPtr("#DAA520")},
				LiteralStringEscape: ansi.StylePrimitive{Color: stringPtr("#FF69B4")},
				GenericDeleted:      ansi.StylePrimitive{Color: stringPtr("#FF1493")},
				GenericEmph:         ansi.StylePrimitive{Italic: boolPtr(true)},
				GenericInserted:     ansi.StylePrimitive{Color: stringPtr("#FFD700")},
				GenericStrong:       ansi.StylePrimitive{Bold: boolPtr(true)},
				GenericSubheading:   ansi.StylePrimitive{Color: stringPtr("#FF69B4")},
				Background:          ansi.StylePrimitive{BackgroundColor: stringPtr("#0A0A0A")},
			},
		},
		Link: ansi.StylePrimitive{
			Color:     stringPtr("#FF1493"), // HotPink - links
			Underline: boolPtr(true),
		},
		LinkText: ansi.StylePrimitive{
			Color: stringPtr("#FF1493"), // HotPink
			Bold:  boolPtr(true),
		},
		List: ansi.StyleList{
			StyleBlock: ansi.StyleBlock{
				StylePrimitive: ansi.StylePrimitive{
					Color: stringPtr("#FFFFFF"), // White
				},
			},
		},
		Emph: ansi.StylePrimitive{
			Color:  stringPtr("#FFD700"), // Gold - emphasis
			Italic: boolPtr(true),
		},
		Strong: ansi.StylePrimitive{
			Color: stringPtr("#FF1493"), // HotPink - strong emphasis
			Bold:  boolPtr(true),
		},
		Strikethrough: ansi.StylePrimitive{
			Color:          stringPtr("#808080"), // GrayMed
			CrossedOut:     boolPtr(true),
		},
		BlockQuote: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Color:  stringPtr("#CCCCCC"), // GrayLight
				Italic: boolPtr(true),
			},
			Indent: uintPtr(2),
		},
		Table: ansi.StyleTable{
			StyleBlock: ansi.StyleBlock{
				StylePrimitive: ansi.StylePrimitive{
					Color: stringPtr("#FFFFFF"),
				},
			},
		},
		DefinitionDescription: ansi.StylePrimitive{
			BlockPrefix: "\n🔹 ",
		},
	}

	return glamour.NewTermRenderer(
		glamour.WithStyles(styleConfig),
		glamour.WithWordWrap(width),
	)
}

func stringPtr(s string) *string {
	return &s
}

func boolPtr(b bool) *bool {
	return &b
}

func uintPtr(u uint) *uint {
	return &u
}
