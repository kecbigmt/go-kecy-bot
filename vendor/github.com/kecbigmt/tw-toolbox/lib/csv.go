package lib
import(
  "encoding/csv"
  "os"
  "path/filepath"
  "log"
)

func ExportCsv(header []string, rows [][]string, exportPath string) string {
  // O_WRONLY:書き込みモード開く, O_CREATE:無かったらファイルを作成
  p, err := os.Getwd()
  if err != nil{
    log.Fatal("Error:", err)
  }
  path := filepath.Join(p, exportPath)
  file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0600)
  if err != nil{
    log.Fatal("Error:", err)
  }

  defer file.Close()

  err = file.Truncate(0) // ファイルを空っぽにする(実行2回目以降用)
  if err != nil{
    log.Fatal("Error:", err)
  }

  writer := csv.NewWriter(file)
  writer.Write(header)
  for _, r := range rows {
    writer.Write(r)
  }
  writer.Flush()
  return path
}
