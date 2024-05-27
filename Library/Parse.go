package Library

import 	(
    "fmt"
    "path/filepath"
    "strings"
    "time"

	"strconv"
)

func ParseInt(param string) int {
	parsedData, err := strconv.Atoi(param);
	if err != nil {
		panic(err)
	}

	return parsedData
}

// Fungsi untuk menghasilkan nama file unik
func GenerateUniqueFileName(originalName string) string {
    // Anda dapat menggunakan waktu, GUID, atau metode lain untuk membuat nama file unik
    // Contoh sederhana: menambahkan timestamp ke nama file asli
    timestamp := time.Now().UnixNano()
    extension := filepath.Ext(originalName)
    filename := strings.TrimSuffix(originalName, extension)
    return fmt.Sprintf("%s_%d%s", filename, timestamp, extension)
}