package cli

import (
	"bufio"
	"fmt"
	"gbm/util"
	"os"
	"regexp"
	"strings"
)

const (
	__script_var       = `(?<var_name>[a-zA-Z0-9]+)\s?=\s?(?<var_val>.*)`
	__script_func_call = `^(?<func_name>\S+)\s?\(\s?(?<func_param>.+)\)`
)

type ScriptCtx struct {
	regexFuncCall *regexp.Regexp
	regexVar      *regexp.Regexp
	variables     map[string]string
}

func Script() {
	flags := util.NewFlagSet("script")
	var scriptFile string
	flags.StringVar(&scriptFile, "f", "", "script file location")
	flags.ParseCmdFlags()

	flags.ValidateStringNotEmpty(scriptFile, "script file is required")

	file, err := os.Open(scriptFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	script := &ScriptCtx{
		regexFuncCall: regexp.MustCompile(__script_func_call),
		regexVar:      regexp.MustCompile(__script_var),
		variables:     make(map[string]string),
	}

	for scanner.Scan() {
		line := scanner.Text()
		script.ExecStatement(line)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func (ctx *ScriptCtx) getFuncParams(param string) ([]string, error) {
	parts := strings.Split(param, ",")
	var params []string
	i := 0
	for i < len(parts) {
		if strings.HasPrefix(parts[i], "\"") {
			str := ""
			for {
				str += parts[i]
				if strings.HasSuffix(parts[i], "\"") {
					break
				}
				i += 1
			}
			params = append(params, strings.TrimSpace(str[1:len(str)-1]))
		} else {
			p := strings.TrimSpace(parts[i])
			if v, ok := ctx.variables[p]; ok {
				p = v
			}
			params = append(params, p)
		}
		i += 1
	}
	return params, nil
}

func (ctx *ScriptCtx) ExecStatement(line string) {
	if match := ctx.regexFuncCall.FindStringSubmatch(line); match != nil {
		params, err := ctx.getFuncParams(match[2])
		if err != nil {
			panic("error exec statement: " + line)
		}
		ctx.handleFuncCall(match[1], params)
	} else if match := ctx.regexVar.FindStringSubmatch(line); match != nil {
		err := ctx.handleVarialbeAssignment(match[1], match[2])
		if err != nil {
			panic("error exec statement: " + line)
		}
	}
}

func (ctx *ScriptCtx) handleVarialbeAssignment(name, val string) error {
	if strings.HasPrefix(val, "\"") && strings.HasSuffix(val, "\"") {
		ctx.variables[name] = strings.TrimSpace(val[1 : len(val)-1])
	} else if match := ctx.regexFuncCall.FindStringSubmatch(val); match != nil {
		params, err := ctx.getFuncParams(match[2])
		if err != nil {
			panic("error exec statement: " + val)
		}
		ctx.variables[name] = ctx.handleFuncCall(match[1], params)
	} else {
		ctx.variables[name] = val
	}
	return nil
}

func (ctx *ScriptCtx) handleFuncCall(name string, param []string) string {
	if name == "print" {
		fmt.Println(strings.Join(param, " "))
		return "<nil>"
	}

	return "<nil>"
}
