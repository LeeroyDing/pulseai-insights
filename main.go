package main

import (
"crypto/md5"
"encoding/hex"
"fmt"
"io"
"log"
"os"
"path/filepath"
"strings"
"time"

"github.com/mmcdole/gofeed"
)

const (
DataDir             = "hugo/content/posts"
ProcessedHashesFile = "src/data/processed_hashes.txt"
PnLFile             = "pnl_tracker.csv"
)

var rssFeeds = []string{
"https://pubmed.ncbi.nlm.nih.gov/rss/search/1-y_L-7-8-9-10-11-12-13-14-15-16-17-18-19-20/", // Retrying with standard search RSS
"https://bmjdigitalhealth.bmj.com/rss/recent.xml",
"https://www.sciencedaily.com/rss/health_medicine/artificial_intelligence.xml",
"https://www.medgadget.com/category/artificial-intelligence/feed",
"https://www.technologyreview.com/topic/artificial-intelligence/feed/",
}

type Article struct {
Title     string
URL       string
Summary   string
Published string
Hash      string
}

func logPnL(event string, cost, revenue float64) {
date := time.Now().Format("2006-01-02 15:04:05")
f, err := os.OpenFile(PnLFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
log.Printf("Error opening PnL file: %v", err)
return
}
defer f.Close()
fmt.Fprintf(f, "%s,%s,%.5f,%.5f,%.5f,USD\n", date, event, cost, revenue, revenue-cost)
}

func getProcessedHashes() map[string]bool {
hashes := make(map[string]bool)
content, err := os.ReadFile(ProcessedHashesFile)
if err != nil {
return hashes
}
lines := strings.Split(string(content), "\n")
for _, line := range lines {
if h := strings.TrimSpace(line); h != "" {
hashes[h] = true
}
}
return hashes
}

func saveHash(h string) {
f, err := os.OpenFile(ProcessedHashesFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
log.Printf("Error saving hash: %v", err)
return
}
defer f.Close()
fmt.Fprintln(f, h)
}

func fetchNews() []Article {
fmt.Println("🚀 Scouting for news...")
processedHashes := getProcessedHashes()
var newArticles []Article
fp := gofeed.NewParser()

for _, url := range rssFeeds {
feed, err := fp.ParseURL(url)
if err != nil {
log.Printf("⚠️ Feed Error [%s]: %v", url, err)
continue
}

for _, item := range feed.Items {
h := md5.New()
io.WriteString(h, item.Link)
urlHash := hex.EncodeToString(h.Sum(nil))

if !processedHashes[urlHash] {
newArticles = append(newArticles, Article{
Title:     item.Title,
URL:       item.Link,
Summary:   item.Description,
Published: item.Published,
Hash:      urlHash,
})
}
}
}
return newArticles
}

func generateMarkdown(article Article) string {
// In production, this calls the LLM API
apiCost := 0.00012
logPnL(fmt.Sprintf("AI Summary: %s", article.Title), apiCost, 0)

var sb strings.Builder
sb.WriteString("---\n")
sb.WriteString(fmt.Sprintf("title: \"%s\"\n", strings.ReplaceAll(article.Title, "\"", "'")))
sb.WriteString(fmt.Sprintf("date: %s\n", time.Now().Format(time.RFC3339)))
sb.WriteString(fmt.Sprintf("source_url: \"%s\"\n", article.URL))
sb.WriteString("---\n\n")
sb.WriteString("### Summary\n")
sb.WriteString("- **What happened:** [AI Summary Placeholder]\n")
sb.WriteString("- **Why it matters:** [AI Analysis Placeholder]\n")
sb.WriteString("- **Clinical Implication:** [AI Insight Placeholder]\n\n")
sb.WriteString(fmt.Sprintf("*Read more at [Source](%s)*\n", article.URL))

return sb.String()
}

func main() {
os.MkdirAll(DataDir, 0755)
os.MkdirAll("src/data", 0755)

articles := fetchNews()
fmt.Printf("✅ Found %d new articles.\n", len(articles))

count := 0
for _, article := range articles {
if count >= 5 { break }
content := generateMarkdown(article)
filename := filepath.Join(DataDir, article.Hash+".md")
os.WriteFile(filename, []byte(content), 0644)
saveHash(article.Hash)
fmt.Printf("📝 Published: %s\n", article.Title)
count++
}
}
