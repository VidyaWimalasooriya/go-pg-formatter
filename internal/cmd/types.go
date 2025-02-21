package cmd

type configurationType struct {
	Anonymize     bool   `json:"anonymize,omitempty"`
	CommaBreak    bool   `json:"commaBreak,omitempty"`
	FunctionCase  int    `json:"functionCase,omitempty"`
	KeywordCase   int    `json:"keywordCase,omitempty"`
	NoRcFile      bool   `json:"noRcFile,omitempty"`
	Placeholder   string `json:"placeholder,omitempty"`
	Spaces        int    `json:"spaces,omitempty"`
	StripComments bool   `json:"stripComments,omitempty"`
	Tabs          bool   `json:"tabs,omitempty"`
}
