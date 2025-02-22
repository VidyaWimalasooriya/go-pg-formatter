package cmd

type configurationType struct {
	Anonymize     bool
	CommaBreak    bool
	CommaEnd      bool
	CommaStart    bool
	FunctionCase  int
	KeywordCase   int
	NoRcFile      bool
	Placeholder   string
	Spaces        int
	StripComments bool
	Tabs          bool
}
