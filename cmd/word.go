package cmd

import (
	"log"
	"strings"

	"github.com/axzed/go-tools/internal/word"
	"github.com/spf13/cobra"
)

var str string
var mode int8

const (
	ModeUpper                      = iota + 1 // 全部单词转为大写
	ModeLower                                 // 全部单词转为小写
	ModeUnderscoreToUpperCamelCase            // 下划线单词转为大写驼峰单词
	ModeUnderscoreToLowerCamelCase            // 下划线单词转为小写驼峰单词
	ModeCamelCaseToUnderScore                 // 驼峰单词转为下划线单词
)

// 设置集成子命令
var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换,模式如下:",
	"1: 全部单词转为大写",
	"2: 全部单词转为小写",
	"3: 下划线单词转为大写驼峰单词",
	"4: 下划线单词转为小写驼峰单词",
	"5: 驼峰单词转为下划线单词",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeCamelCaseToUnderScore:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暂不支持该转换模式,请执行 help word 查看帮助文档")
		}
		log.Printf("输出结果: %s", content)
	},
}

func init() {
	// VarP系列方法 1.要绑定的变量 2.接收该参数的完整的命令标志 3.对应短标识 4.参数默认值 5.使用说明
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换模式")
}
