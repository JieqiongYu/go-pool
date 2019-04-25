package functions

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/bigquery"
)

/*Data BigQueryにデータを追加するための構造体、タグで変数とキーを紐づける。*/
type Data struct {
	Name     string    `bigquery:"NAME"`
	Datetime time.Time `bigquery:"DATETIME"`
}

/*ToBigQuery HTTP triggerで実行される*/
func ToBigQuery(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet: // GETの場合

		// GET parameterの"name"から値を取り出す
		names, err := r.URL.Query()["name"]

		// 取り出せない場合はエラーとして処理を終了する
		if !err || len(names[0]) < 1 {
			fmt.Fprint(w, "パラメータに:\"name\"がありません。\r\n")
			return
		}

		// BigQueryへ出力する関数を呼び出す
		WriteToBigQuery(names[0])

	default: // GET以外の場合はエラー
		http.Error(w, "405 - Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

/*WriteToBigQuery is the function for writing to bigquery.*/
func WriteToBigQuery(name string) {

	// 流し込むデータを構造へ格納する
	data := Data{}
	data.Name = name
	data.Datetime = time.Now()

	// コンテクストを取得する
	ctx := context.Background()

	os.Setenv("GCP_PROJECT", "kouzoh-p-jieqiong-yu")
	// プロジェクトIDを取得する
	projectID := os.Getenv("GCP_PROJECT")

	// BigQueryを操作するクライアントを作成する、エラーの場合はLoggingへ出力する
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		log.Printf("BigQuery接続エラー Error:%T message: %v", err, err)
		return
	}

	// deferでクライアントを閉じる
	defer client.Close()

	// クライアントからテーブルを操作するためのアップローダーを取得する
	u := client.Dataset("go_dataset").Table("names").Uploader()

	items := []Data{data}

	// テーブルへデータの追加を行う、エラーの場合はLoggingへ出力する
	err = u.Put(ctx, items)
	if err != nil {
		log.Printf("データ書き込みエラー　Error: %T message: %v", err, err)
		return
	}
}
