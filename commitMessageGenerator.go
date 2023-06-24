package main

import (
	"log"
	"os/exec"
	"strings"
)

// ステージングされたファイルの変更内容を取得してコミットメッセージを生成する
func commitMessageGenerator() string {
	//ステージングされたファイルの変更内容を取得
	stagedFiles := getStagedFiles()
	//変更内容からコミットメッセージを生成
	commitMessage := createCommitMessage(stagedFiles)
	return commitMessage
}

// ステージングされたファイルの変更内容を取得
func getStagedFiles() []string {
	//git diff --cached --name-onlyでステージングされたファイルの変更内容を取得
	cmd := exec.Command("git", "diff", "--cached", "--name-only")
	//コマンド実行
	out, err := cmd.Output()
	if err != nil {
		log.Fatalln(err)
	}
	//取得した変更内容をスライスに格納
	stagedFiles := strings.Split(string(out), "\n")
	return stagedFiles
}

// 変更内容からコミットメッセージを生成
func createCommitMessage(stagedFiles []string) string {
	//コミットメッセージのテンプレートを作成
	commitMessage := "commit: "
	//変更内容の数だけループ
	for _, stagedFile := range stagedFiles {
		//変更内容がある場合
		if stagedFile != "" {
			//変更内容をコミットメッセージに追加
			commitMessage = commitMessage + stagedFile + ", "
		}
	}
	//コミットメッセージの最後のカンマを削除
	commitMessage = strings.TrimRight(commitMessage, ", ")
	return commitMessage
}

func main() {
	commitMessage := commitMessageGenerator()
	log.Println(commitMessage)
}
