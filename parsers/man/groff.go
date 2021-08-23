package man

/**
  Source: https://man7.org/linux/man-pages/man7/groff_man.7.html

  .B      Bold                         Font style macros
  .BI     Bold, italic alternating     Font style macros
  .BR     Bold, roman alternating      Font style macros
  .I      Italic                       Font style macros
  .IB     Italic, bold alternating     Font style macros
  .IR     Italic, roman alternating    Font style macros
  .RB     Roman, bold alternating      Font style macros
  .RI     Roman, italic alternating    Font style macros
  .SB     Small bold                   Font style macros
  .SM     Small                        Font style macros

  .EE     Example end                  Document structure macros
  .EX     Example begin                Document structure macros
  .SH     Section heading              Document structure macros
  .SS     Subsection heading           Document structure macros
  .RE     Relative inset end           Document structure macros
  .RS     Relative inset start         Document structure macros
  .TH     Title heading                Document structure macros

  .IP     Indented paragraph           Paragraph macros
  .LP     (Left) paragraph             Paragraph macros
  .TP     Tagged paragraph             Paragraph macros
  .TQ     Supplemental paragraph tag   Paragraph macros
  .P      Paragraph                    Paragraph macros
  .PP     Paragraph                    Paragraph macros

  .ME     Mail-to end                  Hyperlink and email macros
  .MT     Mail-to start                Hyperlink and email macros
  .UE     URL end                      Hyperlink and email macros
  .UR     URL start                    Hyperlink and email macros

  .OP     (Command-line) option        Command synopsis macros
  .SY     Synopsis start               Command synopsis macros
  .YS     Synopsis end                 Command synopsis macros
*/

const MacroFlag = "Fl"
const MacroArgListStart = "Xo"
const MacroArgListEnd = "Xc"
const MacroUnix = "Ux"
const MacroNoSpacesOn =".Sm off"
const MacroNoSpacesOff =".Sm on"
const MacroArgument = "Ar"
const MacroSquareBracketStart = ".Oo"
const MacroSquareBracketEnd = "Oc"
const MacroBrackets = "Pq"
const MacroName = "Nm"
const MacroBeginList = "Bl"
const MacroEndList = "El"
const MacroArgWithoutDash = "Cm"
const MacroExtendedArgList = "It"
const MacroNoSpace = "Ns"
const MacroManPageName = "Xr"
const MacroManPageReference = "Xr"
const MacroDoubleQuote = "Dq"
const MacroQuoteLiteral = "Ql"
const MacroSingleQuote = "Sq"
const MacroDoubleQuoteOpen = "Do"
const MacroDoubleQuoteClose = "Dc"

// .Dq	.Do	.Dc
const MacroFontBold = "B"
const MacroFontBoldItalic = "BI"
const MacroFontBoldRoman = "BR"
const MacroFontItalic = "I"
const MacroFontUnderline = "Pa"
const MacroFontItalicBold = "IB"
const MacroFontItalicRoman = "IR"
const MacroFontRomanBold = "RB"
const MacroFontRomanItalic = "RI"
const MacroFontSmallBold = "SB"
const MacroFontSmall = "SM"

const MacroStructureExampleEnd = "EE"
const MacroStructureExampleStart = "EX"
const MacroStructureSectionHeading = "SH"
const MacroStructureTitleHeading = "TH"
const MacroStructureSubSectionHeading = "SS"
const MacroStructureRelativeInsetEnd = "RE"
const MacroStructureRelativeInsetStart = "RS"

const MacroParagraphIndented = "IP"
const MacroParagraphLeft = "LP"
const MacroParagraphTag = "TP"
const MacroParagraphSupplementalTag = "TP"
const MacroParagraph = "P"
const MacroParagraph2 = "Pp"

const MacroLinkMailEnd = "ME"
const MacroLinkMailStart = "MT"
const MacroLinkUrlStart = "UR"
const MacroLinkUrlEnd = "UE"

const MacroSynopsisOption = "OP"
const MacroSynopsisStart = "SY"
const MacroSynopsisEnd = "YS"

const SectionDescription = "DESCRIPTION"
const SectionOptions = "OPTIONS"
