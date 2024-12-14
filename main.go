package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func splitSentences(text string) []string {
	// 使用strings.FieldsFunc函数来分割字符串
	// 这里使用了一个匿名函数来定义分割的规则
	// 分割的条件是遇到句号、问号或分号
	splitFunc := func(r rune) bool {
		return r == '.' || r == '?' || r == ';'
	}
	// 使用FieldsFunc函数进行分割，并将结果存储在一个切片中
	fields := strings.FieldsFunc(text, splitFunc)
	// 遍历切片，去除每个元素前后的空格，并将结果存储在一个新的切片中
	var sentences []string
	for _, field := range fields {
		sentences = append(sentences, strings.TrimSpace(field))
	}
	// 返回分割后的句子切片
	return sentences
}

// 处理请求的函数
func handleRequest(w http.ResponseWriter, r *http.Request) {
	// 从请求中获取 ID 参数
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	// 将 ID 转换为整数
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	// 根据 ID 获取文章
	fileName := getFileNameByID(id)
	// 生成题库
	// 读取文件
	article, err := os.ReadFile(fileName)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}
	// 将文章转换为字符串
	articleStr := string(article)
	// 分割文章为句子
	sentences := splitSentences(articleStr)

	// 将题库作为响应返回
	fmt.Fprint(w, sentences)
}
func generateOptions(word string) []string {
	// 生成选项
	options := make([]string, 4)
	options[0] = word
	for i := 1; i < 4; i++ {
		// 随机生成一个单词作为选项
		options = generateRandomWord()
	}
	// 将选项打乱
	rand.Shuffle(len(options), func(i, j int) {
		options[i], options[j] = options[j], options[i]
	})
	return options
}

func generateRandomWord() []string {
	return []string{}
}

func chooseWords(sentences []string) []Question {
	var questions = make([]Question, 0)
	for i, sentence := range sentences {
		// 生成题目
		prefix, subfix, word := chooseWord(sentence)
		// 生成选项
		options := generateOptions(word)
		// 将题目和选项添加到题库中
		question := Question{
			ID:         i,
			UserChoose: -1,
			Prefix:     prefix,
			Subfix:     subfix,
			Options:    options,
		}
		questions = append(questions, question)
		// 将题目添加到题库中
	}
	return questions
}

func chooseWord(sentence string) (prefix string, subfix string, word string) {
	// 分割句子为单词
	words := strings.Fields(sentence)
	// 随机选择一个单词作为答案
	wordIndex := rand.Intn(len(words))
	word = words[wordIndex]
	// 生成前缀和后缀
	prefix = strings.Join(words[:wordIndex], " ")
	subfix = strings.Join(words[wordIndex+1:], " ")
	return
}

// 语料文章
type Article struct {
	ID       int    `json:"id"`
	FileName string `json:"file_name"`
}

type Question struct {
	ID         int      `json:"id"`
	UserChoose int      `json:"user_choose"`
	Prefix     string   `json:"prefix"`
	Subfix     string   `json:"subfix"`
	Options    []string `json:"options"`
}

// 生成题库的函数

// Easy_IELTS_Band_4-5.txt
// Hard_IELTS_Band_8-9.txt
// Midium_IELTS_Band_6-7.txt
var articles = []Article{
	{ID: 1, FileName: "Easy_IELTS_Band_4-5.txt"},
	{ID: 2, FileName: "Hard_IELTS_Band_8-9.txt"},
	{ID: 3, FileName: "Midium_IELTS_Band_6-7.txt"},
}

func getFileNameByID(id int) string {
	for _, article := range articles {
		if article.ID == id {
			return article.FileName
		}
	}
	return "" // 如果没有找到对应的文章ID，返回空字符串
}

func main() {
	// 设置路由和处理函数
	http.HandleFunc("/choose", handleRequest)

	// 启动服务器
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
