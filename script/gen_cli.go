//go:generate go run gen_cli.go

package main

import (
	"blue/commands"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

const code_ = `package main

import (
	g "blue/api/go"
)

type CmdFunc func(*g.Client, []string) (string, error)

func Exec(conn *g.Client, s []string) (string, error) {
	if len(s) == 0 {
		return "", ErrCommand(s[0])
	}

	f, ok := funcMap[s[0]]
	if !ok {
		return "", ErrCommand(s[0])
	}
	return f(conn, s)
}

`

const select_ = `func Select() CmdFunc {
	return func(conn *g.Client, s []string) (string, error) {
		if len(s) > 2 {
			return "", ErrArgu(s[0])
		}
		if len(s) == 2 {
			return conn.Select(s[1])
		}
		return conn.Select()
	}
}

`

func main() {
	// 读取 JSON 文件
	files := make([]string, 0) // 添加更多文件名

	err := filepath.Walk("../commands", func(path string, info fs.FileInfo, err error) error {
		if filepath.Ext(path) != ".json" {
			return nil
		}

		abs, err := filepath.Abs(path)
		if err != nil {
			return nil
		}

		files = append(files, abs)
		return nil
	})
	if err != nil {
		fmt.Println("Error walking directory:", err)
		return
	}

	var cmds []commands.Cmd
	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		// 解析 JSON 数据
		var fileCommands commands.Cmd
		err = json.Unmarshal(data, &fileCommands)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return
		}

		cmds = append(cmds, fileCommands)
	}

	var code strings.Builder
	code.WriteString("// Code generated by go generate; DO NOT EDIT.\n")
	code.WriteString("// Code generated by go generate; DO NOT EDIT.\n")
	code.WriteString("// Code generated by go generate; DO NOT EDIT.\n\n")
	code.WriteString(code_)

	for _, cmd := range cmds {
		if cmd.Name == "select" {
			selectFunc(&code, cmd)
			continue
		}

		writeFunc(&code, cmd)
	}

	code.WriteString(fmt.Sprintf("var funcMap = map[string]CmdFunc{\n"))

	for _, cmd := range cmds {
		upName := strings.ToUpper(cmd.Name[:1]) + cmd.Name[1:]
		code.WriteString(fmt.Sprintf("	\"%s\": %s(),\n", cmd.Name, upName))
	}

	code.WriteString("}\n")

	// 将代码写入文件
	err = os.WriteFile("../blue-client/exec.go", []byte(code.String()), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("client generated successfully!")
}

func writeFunc(code *strings.Builder, cmd commands.Cmd) {
	para := ""
	for i := 0; i < cmd.Arity; i++ {
		para += fmt.Sprintf("s[%d],", i+1)
	}
	if para != "" {
		para = para[:len(para)-1]
	}
	upName := strings.ToUpper(cmd.Name[:1]) + cmd.Name[1:]
	tmp := fmt.Sprintf("%s(%s)", upName, para)

	code.WriteString(fmt.Sprintf(`func %s() CmdFunc {
	return func(conn *g.Client, s []string) (string, error) {
		if len(s) != %d {
			return "", ErrArgu(s[0])
		}
		return conn.%s
	}
}
`, upName, cmd.Arity+1, tmp))
	code.WriteString("\n")
}

func selectFunc(code *strings.Builder, cmd commands.Cmd) {
	code.WriteString(select_)
}
