package main

import (
"bytes"
"crypto/md5"
"encoding/hex"
"encoding/json"
"fmt"
"io"
"log"
"net/http"
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
LinksFile           = "src/data/links.json"
GeminiModel         = "gemini-1.5-flash"
)

var rssFeeds = []string{
"https://pubmed.ncbi.nlm.nih.gov/rss/search/1-y_L-7-8-9-10-11-12-13-14-15-16-17-18-19-20/",
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

func loadLinks() map[string]string {
links := make(map[string]string)
content, err := os.ReadFile(LinksFile)
if err != nil {
return links
}
json.Unmarshal(content, &links)
return links
}

func fetchFullContent(url string) string {
resp, err := http.Get(url)
if err != nil {
return ""
}
defer resp.Body.Close()
body, _ := io.ReadAll(resp.Body)
// Simple cleanup: just return the body, Gemini will handle the HTML
return string(body)
}

func summarizeWithAI(title, content string) (string, float64) {
apiKey := os.Getenv("GEMINI_API_KEY")
if apiKey == "" {
return "Error: No API Key", 0
}

prompt := fmt.Sprintf("You are a medical AI analyst. Summarize the following healthcare AI news article for a professional audience. Title: %s. Content: %s. Include: 1. Executive Summary, 2. Key Technical/Clinical Findings, 3. Strategic Impact on Healthcare, 4. Potential Limitations. Format in Markdown with clear headers.", title, content)
if len(prompt) > 30000 {
prompt = prompt[:30000]
}

url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent?key=%s", GeminiModel, apiKey)
payload := map[string]interface{}{
"contents": []interface{}{
map[string]interface{}{
"parts": []interface{}{
map[string]interface{}{"text": prompt},
},
},
},
}
jsonData, _ := json.Marshal(payload)
resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
if err != nil {
return "Error calling AI", 0
}
defer resp.Body.Close()

var result struct {
Candidates []struct {
Content struct {
Parts []struct {
Text string `json:"text"`
} `json:"parts"`
} `json:"content"`
} `json:"candidates"`
}
json.NewDecoder(resp.Body).Decode(&result)

if len(result.Candidates) > 0 && len(result.Candidates[0].Content.Parts) > 0 {
return result.Candidates[0].Content.Parts[0].Text, 0.0005 // Estimated cost for Flash
}
return "Summary unavailable", 0
}

func generateMarkdown(article Article, links map[string]string) string {
fullContent := fetchFullContent(article.URL)
aiSummary, cost := summarizeWithAI(article.Title, fullContent)
logPnL(fmt.Sprintf("AI Deep Summary: %s", article.Title), cost, 0)

recommendedTool := ""
for keyword, link := range links {
if strings.Contains(strings.ToLower(article.Title), strings.ToLower(keyword)) ||
strings.Contains(strings.ToLower(aiSummary), strings.ToLower(keyword)) {
recommendedTool = fmt.Sprintf("\n\n> **Recommended Resource:** Check out this [%s Tool/Course](%s) related to this breakthrough.", keyword, link)
break
}
}

var sb strings.Builder
sb.WriteString("---\n")
sb.WriteString(fmt.Sprintf("title: \"%s\"\n", strings.ReplaceAll(article.Title, "\"", "'")))
sb.WriteString(fmt.Sprintf("date: %s\n", time.Now().Format(time.RFC3339)))
sb.WriteString(fmt.Sprintf("source_url: \"%s\"\n", article.URL))
sb.WriteString("---\n\n")
sb.WriteString(aiSummary)
sb.WriteString(recommendedTool)
sb.WriteString(fmt.Sprintf("\n\n*Read more at [Source](%s)*\n", article.URL))

return sb.String()
}

func main() {
os.MkdirAll(DataDir, 0755)
os.MkdirAll("src/data", 0755)

links := loadLinks()
fp := gofeed.NewParser()
processedHashes := getProcessedHashes()

count := 0
for _, feedURL := range rssFeeds {
feed, err := fp.ParseURL(feedURL)
if err != nil {
continue
}

for _, item := range feed.Items {
if count >= 3 { // Limit to 3 high-quality posts per run to manage costs/time
break
}

h := md5.New()
io.WriteString(h, item.Link)
urlHash := hex.EncodeToString(h.Sum(nil))

if !processedHashes[urlHash] {
article := Article{
Title: item.Title,
URL:   item.Link,
Hash:  urlHash,
}
content := generateMarkdown(article, links)
filename := filepath.Join(DataDir, article.Hash+".md")
os.WriteFile(filename, []byte(content), 0644)
saveHash(article.Hash)
fmt.Printf("📝 Published High-Quality: %s\n", article.Title)
count++
}
}
if count >= 3 {
break
}
}
}
