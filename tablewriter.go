package tablewriter

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

ASCII Table Writer
即时生成 ASCII 表...安装很简单
	go get github.com/olekukonko/tablewriter

Features
Automatic Padding											// 自动填充
Support Multiple Lines										// 支持多行
Supports Alignment											// 支持对齐
Support Custom Separators									// 支持自定义分隔符
Automatic Alignment of numbers & percentage					// 数字和百分比的自动对齐
Write directly to http , file etc via io.Writer				// 通过 io.Writer 直接写入 http ，文件等
Read directly from CSV file									// 直接从 CSV 文件读取
Optional row line via SetRowLine							// 通过 SetRowLine 的可选行线
Normalise table header										// 规范化表头
Make CSV Headers optional									// 将 CSV 标头设为可选
Enable or disable table border								// 启用或禁用表格边框
Set custom footer support									// 设置自定义页脚支持
Optional identical cells merging							// 可选的相同单元格合并
Set custom caption											// 设置自定义标题
Optional reflowing of paragraphs in multi-line cells.		// 多行单元格中段落的可选重排。

示例 1 - 基本
data := [][]string{
	[]string{"A", "The Good", "500"},
	[]string{"B", "The Very very Bad Man", "288"},
	[]string{"C", "The Ugly", "120"},
	[]string{"D", "The Gopher", "800"},
}
table := tablewriter.NewWriter(os.Stdout)					// Start New Table 直接拿io.Writer
table.SetHeader([]string{"Name", "Sign", "Rating"})			// 设置表头
for _, v := range data {
	table.Append(v)											// 将行追加到表格
}
table.Render() // Send output

Output 1
+------+-----------------------+--------+
| NAME |         SIGN          | RATING |
+------+-----------------------+--------+
|  A   |       The Good        |    500 |
|  B   | The Very very Bad Man |    288 |
|  C   |       The Ugly        |    120 |
|  D   |      The Gopher       |    800 |
+------+-----------------------+--------+

========================================================================================================================

示例 2 - 无边框/页脚/批量追加
data := [][]string{
	[]string{"1/1/2014", "Domain name", "2233", "$10.98"},
	[]string{"1/1/2014", "January Hosting", "2233", "$54.95"},
	[]string{"1/4/2014", "February Hosting", "2233", "$51.00"},
	[]string{"1/4/2014", "February Extra Bandwidth", "2233", "$30.00"},
}
table := tablewriter.NewWriter(os.Stdout)										// Start New Table 直接拿io.Writer
table.SetHeader([]string{"Date", "Description", "CV2", "Amount"})				// 设置表头
table.SetFooter([]string{"", "", "Total", "$146.93"}) // Add Footer				// 设置表格页脚
table.SetBorder(false)                                // Set Border to false	// 无边框 设置表格边框这将启用/禁用表格周围的线条
table.AppendBulk(data)                                // Add Bulk Data			// 批量追加
table.Render()

Output 2

   DATE    |       DESCRIPTION        |  CV2  | AMOUNT
-----------+--------------------------+-------+----------
  1/1/2014 | Domain name              |  2233 | $10.98
  1/1/2014 | January Hosting          |  2233 | $54.95
  1/4/2014 | February Hosting         |  2233 | $51.00
  1/4/2014 | February Extra Bandwidth |  2233 | $30.00
-----------+--------------------------+-------+----------
									    TOTAL | $146 93
									  --------+----------

========================================================================================================================

示例 3 - CSV
table, _ := tablewriter.NewCSV(os.Stdout, "testdata/test_info.csv", true)	// 通过从 CSV 文件导入开始一个新表采用 io.Writer 和 csv 文件名
table.SetAlignment(tablewriter.ALIGN_LEFT)   // Set Alignment				// 设置表格对齐方式
table.Render()

Output 3
+----------+--------------+------+-----+---------+----------------+
|  FIELD   |     TYPE     | NULL | KEY | DEFAULT |     EXTRA      |
+----------+--------------+------+-----+---------+----------------+
| user_id  | smallint(5)  | NO   | PRI | NULL    | auto_increment |
| username | varchar(10)  | NO   |     | NULL    |                |
| password | varchar(100) | NO   |     | NULL    |                |
+----------+--------------+------+-----+---------+----------------+

========================================================================================================================

示例 4 - 自定义分隔符
table, _ := tablewriter.NewCSV(os.Stdout, "testdata/test.csv", true)	// 通过从 CSV 文件导入开始一个新表采用 io.Writer 和 csv 文件名
table.SetRowLine(true)         	// Enable row line			// 启用行线	// 设置行线这将启用/禁用表的每一行上的一行
table.SetCenterSeparator("*")	// Change table lines		// 改变表格行 设置中心分隔符
table.SetColumnSeparator("╪")								// 设置列分隔符
table.SetRowSeparator("-")									// 设置行分隔符
table.SetAlignment(tablewriter.ALIGN_LEFT)					// 设置表格对齐方式
table.Render()

Output 4
*------------*-----------*---------*
╪ FIRST NAME ╪ LAST NAME ╪   SSN   ╪
*------------*-----------*---------*
╪ John       ╪ Barry     ╪ 123456  ╪
*------------*-----------*---------*
╪ Kathy      ╪ Smith     ╪ 687987  ╪
*------------*-----------*---------*
╪ Bob        ╪ McCornick ╪ 3979870 ╪
*------------*-----------*---------*

========================================================================================================================

示例 5 - Markdown 格式
data := [][]string{
	[]string{"1/1/2014", "Domain name", "2233", "$10.98"},
	[]string{"1/1/2014", "January Hosting", "2233", "$54.95"},
	[]string{"1/4/2014", "February Hosting", "2233", "$51.00"},
	[]string{"1/4/2014", "February Extra Bandwidth", "2233", "$30.00"},
}
table := tablewriter.NewWriter(os.Stdout)													// Start New Table 直接拿io.Writer
table.SetHeader([]string{"Date", "Description", "CV2", "Amount"})							// 设置表头
table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
table.SetCenterSeparator("|")																// 设置中心分隔符
table.AppendBulk(data) // Add Bulk Data														// 允许支持批量追加消除重复的 for 循环
table.Render()

Output 5
|   DATE   |       DESCRIPTION        | CV2  | AMOUNT |
|----------|--------------------------|------|--------|
| 1/1/2014 | Domain name              | 2233 | $10.98 |
| 1/1/2014 | January Hosting          | 2233 | $54.95 |
| 1/4/2014 | February Hosting         | 2233 | $51.00 |
| 1/4/2014 | February Extra Bandwidth | 2233 | $30.00 |

========================================================================================================================

示例 6 - 相同的单元格合并
data := [][]string{
	[]string{"1/1/2014", "Domain name", "1234", "$10.98"},
	[]string{"1/1/2014", "January Hosting", "2345", "$54.95"},
	[]string{"1/4/2014", "February Hosting", "3456", "$51.00"},
	[]string{"1/4/2014", "February Extra Bandwidth", "4567", "$30.00"},
}
table := tablewriter.NewWriter(os.Stdout)								// Start New Table 直接拿io.Writer
table.SetHeader([]string{"Date", "Description", "CV2", "Amount"})		// 设置表头
table.SetFooter([]string{"", "", "Total", "$146.93"})					// 设置表格页脚
table.SetAutoMergeCells(true)											// 设置自动合并单元格这将启用/禁用具有相同值的单元格的合并
table.SetRowLine(true)													// 设置行线这将启用/禁用表的每一行上的一行
table.AppendBulk(data)													// 允许支持批量追加消除重复的 for 循环
table.Render()

Output 6
+----------+--------------------------+-------+---------+
|   DATE   |       DESCRIPTION        |  CV2  | AMOUNT  |
+----------+--------------------------+-------+---------+
| 1/1/2014 | Domain name              |  1234 | $10.98  |
+          +--------------------------+-------+---------+
|          | January Hosting          |  2345 | $54.95  |
+----------+--------------------------+-------+---------+
| 1/4/2014 | February Hosting         |  3456 | $51.00  |
+          +--------------------------+-------+---------+
|          | February Extra Bandwidth |  4567 | $30.00  |
+----------+--------------------------+-------+---------+
|                                       TOTAL | $146 93 |
+----------+--------------------------+-------+---------+

========================================================================================================================

示例 7 - 相同的单元格合并（指定要合并的列索引）
data := [][]string{
	[]string{"1/1/2014", "Domain name", "1234", "$10.98"},
	[]string{"1/1/2014", "January Hosting", "1234", "$10.98"},
	[]string{"1/4/2014", "February Hosting", "3456", "$51.00"},
	[]string{"1/4/2014", "February Extra Bandwidth", "4567", "$30.00"},
}
table := tablewriter.NewWriter(os.Stdout)								// Start New Table 直接拿io.Writer
table.SetHeader([]string{"Date", "Description", "CV2", "Amount"})		// 设置表头
table.SetFooter([]string{"", "", "Total", "$146.93"})					// 设置表格页脚
table.SetAutoMergeCellsByColumnIndex([]int{2, 3})						// 按列索引设置自动合并单元格这将为特定列启用/禁用具有相同值的单元格合并如果 cols 为空，则与 SetAutoMergeCells(true) 相同。
table.SetRowLine(true)													// 设置行线这将启用/禁用表的每一行上的一行
table.AppendBulk(data)													// 允许支持批量追加消除重复的 for 循环
table.Render()

Output 7
+----------+--------------------------+-------+---------+
|   DATE   |       DESCRIPTION        |  CV2  | AMOUNT  |
+----------+--------------------------+-------+---------+
| 1/1/2014 | Domain name              |  1234 | $10.98  |
+----------+--------------------------+       +         +
| 1/1/2014 | January Hosting          |       |         |
+----------+--------------------------+-------+---------+
| 1/4/2014 | February Hosting         |  3456 | $51.00  |
+----------+--------------------------+-------+---------+
| 1/4/2014 | February Extra Bandwidth |  4567 | $30.00  |
+----------+--------------------------+-------+---------+
|                                       TOTAL | $146.93 |
+----------+--------------------------+-------+---------+

========================================================================================================================

Table with color	带颜色的表
data := [][]string{
	[]string{"1/1/2014", "Domain name", "2233", "$10.98"},
	[]string{"1/1/2014", "January Hosting", "2233", "$54.95"},
	[]string{"1/4/2014", "February Hosting", "2233", "$51.00"},
	[]string{"1/4/2014", "February Extra Bandwidth", "2233", "$30.00"},
}
table := tablewriter.NewWriter(os.Stdout)										// Start New Table 直接拿io.Writer
table.SetHeader([]string{"Date", "Description", "CV2", "Amount"})				// 设置表头
table.SetFooter([]string{"", "", "Total", "$146.93"}) // Add Footer				// 设置表格页脚
table.SetBorder(false)                                // Set Border to false	// 设置表格边框这将启用/禁用表格周围的线条

table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor},	// 添加标题颜色（ANSI 代码）
	tablewriter.Colors{tablewriter.FgHiRedColor, tablewriter.Bold, tablewriter.BgBlackColor},
	tablewriter.Colors{tablewriter.BgRedColor, tablewriter.FgWhiteColor},
	tablewriter.Colors{tablewriter.BgCyanColor, tablewriter.FgWhiteColor})

table.SetColumnColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},	// 添加列颜色（ANSI 代码）
	tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
	tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
	tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor})

table.SetFooterColor(tablewriter.Colors{}, tablewriter.Colors{},						// 添加列颜色（ANSI 代码）
	tablewriter.Colors{tablewriter.Bold},
	tablewriter.Colors{tablewriter.FgHiRedColor})

table.AppendBulk(data)															// 允许支持批量追加消除重复的 for 循环
table.Render()

========================================================================================================================

示例 - 8 带颜色的表格单元格
来自 func Rich 的单个单元格颜色优先于列颜色
data := [][]string{
	[]string{"Test1Merge", "HelloCol2 - 1", "HelloCol3 - 1", "HelloCol4 - 1"},
	[]string{"Test1Merge", "HelloCol2 - 2", "HelloCol3 - 2", "HelloCol4 - 2"},
	[]string{"Test1Merge", "HelloCol2 - 3", "HelloCol3 - 3", "HelloCol4 - 3"},
	[]string{"Test2Merge", "HelloCol2 - 4", "HelloCol3 - 4", "HelloCol4 - 4"},
	[]string{"Test2Merge", "HelloCol2 - 5", "HelloCol3 - 5", "HelloCol4 - 5"},
	[]string{"Test2Merge", "HelloCol2 - 6", "HelloCol3 - 6", "HelloCol4 - 6"},
	[]string{"Test2Merge", "HelloCol2 - 7", "HelloCol3 - 7", "HelloCol4 - 7"},
	[]string{"Test3Merge", "HelloCol2 - 8", "HelloCol3 - 8", "HelloCol4 - 8"},
	[]string{"Test3Merge", "HelloCol2 - 9", "HelloCol3 - 9", "HelloCol4 - 9"},
	[]string{"Test3Merge", "HelloCol2 - 10", "HelloCol3 -10", "HelloCol4 - 10"},
}

table := tablewriter.NewWriter(os.Stdout)						// Start New Table 直接拿io.Writer
table.SetHeader([]string{"Col1", "Col2", "Col3", "Col4"})		// 设置表头
table.SetFooter([]string{"", "", "Footer3", "Footer4"})			// 设置表格页脚
table.SetBorder(false)											// 设置表格边框这将启用/禁用表格周围的线条

table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor},	// 添加标题颜色（ANSI 代码）
	tablewriter.Colors{tablewriter.FgHiRedColor, tablewriter.Bold, tablewriter.BgBlackColor},
	tablewriter.Colors{tablewriter.BgRedColor, tablewriter.FgWhiteColor},
	tablewriter.Colors{tablewriter.BgCyanColor, tablewriter.FgWhiteColor})

table.SetColumnColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},	// 添加列颜色（ANSI 代码）
	tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
	tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
	tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor})

table.SetFooterColor(tablewriter.Colors{}, tablewriter.Colors{},						// 添加列颜色（ANSI 代码）
	tablewriter.Colors{tablewriter.Bold},
	tablewriter.Colors{tablewriter.FgHiRedColor})

colorData1 := []string{"TestCOLOR1Merge", "HelloCol2 - COLOR1", "HelloCol3 - COLOR1", "HelloCol4 - COLOR1"}
colorData2 := []string{"TestCOLOR2Merge", "HelloCol2 - COLOR2", "HelloCol3 - COLOR2", "HelloCol4 - COLOR2"}

for i, row := range data {
	if i == 4 {		// 将行附加到具有颜色属性的表
		table.Rich(colorData1, []tablewriter.Colors{tablewriter.Colors{}, tablewriter.Colors{tablewriter.Normal, tablewriter.FgCyanColor}, tablewriter.Colors{tablewriter.Bold, tablewriter.FgWhiteColor}, tablewriter.Colors{}})
		table.Rich(colorData2, []tablewriter.Colors{tablewriter.Colors{tablewriter.Normal, tablewriter.FgMagentaColor}, tablewriter.Colors{}, tablewriter.Colors{tablewriter.Bold, tablewriter.BgRedColor}, tablewriter.Colors{tablewriter.FgHiGreenColor, tablewriter.Italic, tablewriter.BgHiCyanColor}})
	}
	table.Append(row)				// 将行追加到表格
}

table.SetAutoMergeCells(true)		// 设置自动合并单元格这将启用/禁用具有相同值的单元格的合并
table.Render()

========================================================================================================================

示例 9 - 设置表格标题
data := [][]string{
	[]string{"A", "The Good", "500"},
	[]string{"B", "The Very very Bad Man", "288"},
	[]string{"C", "The Ugly", "120"},
	[]string{"D", "The Gopher", "800"},
}

table := tablewriter.NewWriter(os.Stdout)						// Start New Table 直接拿io.Writer
table.SetHeader([]string{"Name", "Sign", "Rating"})				// 设置表头
table.SetCaption(true, "Movie ratings.")						// 设置表格标题

for _, v := range data {
	table.Append(v)												// 将行追加到表格
}
table.Render() // Send output
注意：标题文本将以渲染表格的总宽度换行。
Output 9
+------+-----------------------+--------+
| NAME |         SIGN          | RATING |
+------+-----------------------+--------+
|  A   |       The Good        |    500 |
|  B   | The Very very Bad Man |    288 |
|  C   |       The Ugly        |    120 |
|  D   |      The Gopher       |    800 |
+------+-----------------------+--------+
Movie ratings.

========================================================================================================================

示例 10 - 设置 NoWhiteSpace 和 TablePadding 选项
data := [][]string{
	{"node1.example.com", "Ready", "compute", "1.11"},
	{"node2.example.com", "Ready", "compute", "1.11"},
	{"node3.example.com", "Ready", "compute", "1.11"},
	{"node4.example.com", "NotReady", "compute", "1.11"},
}
table := tablewriter.NewWriter(os.Stdout)									// Start New Table 直接拿io.Writer
table.SetHeader([]string{"Name", "Status", "Role", "Version"})				// 设置表头
table.SetAutoWrapText(false)												// 打开/关闭自动多行文本调整。默认为开（真）。
table.SetAutoFormatHeaders(true)											// 打开/关闭标题自动格式设置。默认为开（真）。
table.SetHeaderAlignment(ALIGN_LEFT)										// 设置页眉对齐
table.SetAlignment(ALIGN_LEFT)												// 设置表格对齐方式
table.SetCenterSeparator("")												// 设置中心分隔符
table.SetColumnSeparator("")												// 设置列分隔符
table.SetRowSeparator("")													// 设置行分隔符
table.SetHeaderLine(false)													// 设置标题行这将启用/禁用标题后的一行
table.SetBorder(false)														// 设置表格边框这将启用/禁用表格周围的线条
table.SetTablePadding("\t") // pad with tabs								// 设置表格填充
table.SetNoWhiteSpace(true)													// 设置无空白
table.AppendBulk(data) // Add Bulk Data										// 允许支持批量追加消除重复的 for 循环
table.Render()

Output 10
NAME             	STATUS  	ROLE   	VERSION
node1.example.com	Ready   	compute	1.11
node2.example.com	Ready   	compute	1.11
node3.example.com	Ready   	compute	1.11
node4.example.com	NotReady	compute	1.11

========================================================================================================================

Render table into a string	将表格渲染成字符串
除了将表格渲染到 io.Stdout 之外，您还可以将其渲染为字符串。 Go 1.10 引入了 strings.Builder 类型，它实现了 io.Writer 接口，因此可以用于此任务。例子：
package main
import (
	"strings"
	"fmt"
	"github.com/olekukonko/tablewriter"
)
func main() {
	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)								// Start New Table 直接拿io.Writer
	/*
	 * Code to fill the table	填写表格的代码
	 */
	table.Render()
	fmt.Println(tableString.String())
}

========================================================================================================================

Constants
const (
	CENTER  = "+"
	ROW     = "-"
	COLUMN  = "|"
	SPACE   = " "
	NEWLINE = "\n"
)

const (
	ALIGN_DEFAULT = iota
	ALIGN_CENTER
	ALIGN_RIGHT
	ALIGN_LEFT
)

const (
	BgBlackColor int = iota + 40
	BgRedColor
	BgGreenColor
	BgYellowColor
	BgBlueColor
	BgMagentaColor
	BgCyanColor
	BgWhiteColor
)

const (
	FgBlackColor int = iota + 30
	FgRedColor
	FgGreenColor
	FgYellowColor
	FgBlueColor
	FgMagentaColor
	FgCyanColor
	FgWhiteColor
)

const (
	BgHiBlackColor int = iota + 100
	BgHiRedColor
	BgHiGreenColor
	BgHiYellowColor
	BgHiBlueColor
	BgHiMagentaColor
	BgHiCyanColor
	BgHiWhiteColor
)

const (
	FgHiBlackColor int = iota + 90
	FgHiRedColor
	FgHiGreenColor
	FgHiYellowColor
	FgHiBlueColor
	FgHiMagentaColor
	FgHiCyanColor
	FgHiWhiteColor
)

const (
	Normal          = 0
	Bold            = 1
	UnderlineSingle = 4
	Italic
)

const ESC = "\033"

const (
	MAX_ROW_WIDTH = 30
)

const SEP = ";"

Functions
func Color(colors ...int) []int
func ConditionString(cond bool, valid, inValid string) string		// 字符串的简单条件 根据条件返回值
func DisplayWidth(str string) int
func Pad(s, pad string, width int) string							// Pad String 尝试将字符串放在中心
func PadLeft(s, pad string, width int) string						// Pad String Left position 这会将字符串放在屏幕的右侧
func PadRight(s, pad string, width int) string						// Pad String Right position 这会将字符串放在屏幕的左侧
func Title(name string) string										// 格式表头替换 _ , .和空间
func WrapString(s string, lim int) ([]string, int)					// Wrap 将 s 包裹成一段长度为 lim 的行，具有最小的不规则性。

func WrapWords(words []string, spc, lim, pen int) [][]string
// WrapWords 是低级别的换行算法，如果您需要更多地控制文本换行过程的细节，它会很有用。对于大多数用途，WrapString 就足够了，而且更方便。
// WrapWords 将单词列表拆分为具有最小“参差不齐”的行，将每个符文视为一个单元，计算每行上相邻单词之间的 spc 单元，并尝试将行限制为 lim 单元。
// 不规则度是所有线条的总误差，其中误差是线条长度与 lim 之差的平方。太长的行（仅当单个单词长于 lim 单位时才会发生）将笔惩罚单位添加到错误中。

Types
1.type Border struct {
	Left   bool
	Right  bool
	Top    bool
	Bottom bool
}
2.type Colors []int

3.type Table struct {
	// contains filtered or unexported fields
}
func NewCSV(writer io.Writer, fileName string, hasHeader bool) (*Table, error)	// 通过从 CSV 文件导入开始一个新表采用 io.Writer 和 csv 文件名
func NewCSVReader(writer io.Writer, csvReader *csv.Reader, hasHeader bool) (*Table, error)	// 使用 csv.Reader 启动一个新的 Table Writer // 这可以启用自定义，例如 reader.Comma = ';
func NewWriter(writer io.Writer) *Table											// Start New Table 直接拿io.Writer
func (t *Table) Append(row []string)											// 将行追加到表格
func (t *Table) AppendBulk(rows [][]string)										// 允许支持批量追加消除重复的 for 循环
func (t *Table) ClearFooter()													// 清除页脚
func (t *Table) ClearRows()														// 清除行
func (t *Table) NumLines() int													// NumLines 获取行数
func (t *Table) Render()														// 渲染表输出
func (t *Table) Rich(row []string, colors []Colors)								// 将行附加到具有颜色属性的表
func (t *Table) SetAlignment(align int)											// 设置表格对齐方式
func (t *Table) SetAutoFormatHeaders(auto bool)									// 打开/关闭标题自动格式设置。默认为开（真）。
func (t *Table) SetAutoMergeCells(auto bool)									// 设置自动合并单元格这将启用/禁用具有相同值的单元格的合并
func (t *Table) SetAutoMergeCellsByColumnIndex(cols []int)						// 按列索引设置自动合并单元格这将为特定列启用/禁用具有相同值的单元格合并如果 cols 为空，则与 SetAutoMergeCells(true) 相同。
func (t *Table) SetAutoWrapText(auto bool)										// 打开/关闭自动多行文本调整。默认为开（真）。
func (t *Table) SetBorder(border bool)											// 设置表格边框这将启用/禁用表格周围的线条
func (t *Table) SetBorders(border Border)
func (t *Table) SetCaption(caption bool, captionText ...string)					// 设置表格标题
func (t *Table) SetCenterSeparator(sep string)									// 设置中心分隔符
func (t *Table) SetColMinWidth(column int, width int)							// 设置列的最小宽度
func (t *Table) SetColWidth(width int)											// 设置默认列宽
func (t *Table) SetColumnAlignment(keys []int)
func (t *Table) SetColumnColor(colors ...Colors)								// 添加列颜色（ANSI 代码）
func (t *Table) SetColumnSeparator(sep string)									// 设置列分隔符
func (t *Table) SetFooter(keys []string)										// 设置表格页脚
func (t *Table) SetFooterAlignment(fAlign int)									// 设置页脚对齐
func (t *Table) SetFooterColor(colors ...Colors)								// 添加列颜色（ANSI 代码）
func (t *Table) SetHeader(keys []string)										// 设置表头
func (t *Table) SetHeaderAlignment(hAlign int)									// 设置页眉对齐
func (t *Table) SetHeaderColor(colors ...Colors)								// 添加标题颜色（ANSI 代码）
func (t *Table) SetHeaderLine(line bool)										// 设置标题行这将启用/禁用标题后的一行
func (t *Table) SetNewLine(nl string)											// 设置新行
func (t *Table) SetNoWhiteSpace(allow bool)										// 设置无空白
func (t *Table) SetReflowDuringAutoWrap(auto bool)								// 重新换行时自动重排多行文本。默认为开（真）。
func (t *Table) SetRowLine(line bool)											// 设置行线这将启用/禁用表的每一行上的一行
func (t *Table) SetRowSeparator(sep string)										// 设置行分隔符
func (t *Table) SetTablePadding(padding string)									// 设置表格填充
