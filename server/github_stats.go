package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type GitHubStats struct {
	TotalCommits       int    `json:"total_commits"`
	TotalStars         int    `json:"total_stars"`
	TotalPRs           int    `json:"total_prs"`
	TotalIssues        int    `json:"total_issues"`
	ContributedTo      int    `json:"contributed_to"`
	TotalContributions int    `json:"total_contributions"`
	CurrentStreak      int    `json:"current_streak"`
	LongestStreak      int    `json:"longest_streak"`
	Name               string `json:"name"`
	AvatarURL          string `json:"avatar_url"`
}

func main() {
	fmt.Println("Starting server...")
	err := godotenv.Load(".env.local")
	if err != nil {
		fmt.Println("Error loading .env.local file:", err)
	} else {
		fmt.Println(".env.local file loaded successfully")
	}

	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		fmt.Println("Warning: GITHUB_TOKEN is not set")
	} else {
		fmt.Println("GITHUB_TOKEN is set")
	}

	http.HandleFunc("/api/github-stats", enableCors(getGitHubStats))
	http.HandleFunc("/", enableCors(healthCheck))
	fmt.Println("Routes registered")
	fmt.Println("Server is running on http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func enableCors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func getGitHubStats(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getGitHubStats function called")
	token := os.Getenv("GITHUB_TOKEN")
	fmt.Printf("GitHub token: %s\n", token) // This will print the token (be careful with this in production)
	if token == "" {
		fmt.Println("Error: GitHub token not set")
		http.Error(w, "GitHub token not set", http.StatusInternalServerError)
		return
	}
	fmt.Println("GitHub token found")

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.github.com/user", nil)
	req.Header.Set("Authorization", "token "+token)
	fmt.Println("Sending request to GitHub API")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request to GitHub API: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	fmt.Printf("GitHub API response status: %s\n", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		http.Error(w, "Error reading response body", http.StatusInternalServerError)
		return
	}
	fmt.Printf("GitHub API response body: %s\n", string(body))

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		http.Error(w, "Error parsing JSON", http.StatusInternalServerError)
		return
	}

	fmt.Printf("GitHub API response: %+v\n", data)

	stats := GitHubStats{
		TotalStars:         0,
		TotalCommits:       0,
		TotalPRs:           0,
		TotalIssues:        0,
		ContributedTo:      0,
		TotalContributions: 0,
		CurrentStreak:      0,
		LongestStreak:      0,
		Name:               "",
		AvatarURL:          "",
	}

	if name, ok := data["name"].(string); ok {
		stats.Name = name
	}
	if avatarURL, ok := data["avatar_url"].(string); ok {
		stats.AvatarURL = avatarURL
	}

	stats.TotalStars = getTotalStars(token)
	stats.TotalCommits = getTotalCommits(token)
	stats.TotalPRs = getTotalPRs(token)
	stats.TotalIssues = getTotalIssues(token)
	stats.ContributedTo = getContributedTo(token)
	stats.TotalContributions, stats.CurrentStreak, stats.LongestStreak = getContributionStats(token)

	fmt.Printf("Final stats: %+v\n", stats)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

func getTotalStars(token string) int {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.github.com/user/repos", nil)
	req.Header.Set("Authorization", "token "+token)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error fetching repos: %v\n", err)
		return 0
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading repos response body: %v\n", err)
		return 0
	}

	var repos []map[string]interface{}
	if err := json.Unmarshal(body, &repos); err != nil {
		fmt.Printf("Error parsing repos JSON: %v\n", err)
		return 0
	}

	totalStars := 0
	for _, repo := range repos {
		if stargazersCount, ok := repo["stargazers_count"].(float64); ok {
			totalStars += int(stargazersCount)
		}
	}
	return totalStars
}

func getTotalCommits(token string) int {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.github.com/search/commits?q=author:YOUR_GITHUB_USERNAME", nil)
	req.Header.Set("Authorization", "token "+token)
	req.Header.Set("Accept", "application/vnd.github.cloak-preview")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error fetching commits: %v\n", err)
		return 0
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading commits response body: %v\n", err)
		return 0
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Printf("Error parsing commits JSON: %v\n", err)
		return 0
	}

	if totalCount, ok := result["total_count"].(float64); ok {
		return int(totalCount)
	}
	return 0
}

func getTotalPRs(token string) int {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.github.com/search/issues?q=author:YOUR_GITHUB_USERNAME+type:pr", nil)
	req.Header.Set("Authorization", "token "+token)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error fetching PRs: %v\n", err)
		return 0
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading PRs response body: %v\n", err)
		return 0
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Printf("Error parsing PRs JSON: %v\n", err)
		return 0
	}

	if totalCount, ok := result["total_count"].(float64); ok {
		return int(totalCount)
	}
	return 0
}

func getTotalIssues(token string) int {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.github.com/search/issues?q=author:YOUR_GITHUB_USERNAME+type:issue", nil)
	req.Header.Set("Authorization", "token "+token)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error fetching issues: %v\n", err)
		return 0
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading issues response body: %v\n", err)
		return 0
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Printf("Error parsing issues JSON: %v\n", err)
		return 0
	}

	if totalCount, ok := result["total_count"].(float64); ok {
		return int(totalCount)
	}
	return 0
}

func getContributedTo(token string) int {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.github.com/user/repos?type=all", nil)
	req.Header.Set("Authorization", "token "+token)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error fetching contributed repos: %v\n", err)
		return 0
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading contributed repos response body: %v\n", err)
		return 0
	}

	var repos []map[string]interface{}
	if err := json.Unmarshal(body, &repos); err != nil {
		fmt.Printf("Error parsing contributed repos JSON: %v\n", err)
		return 0
	}

	contributedTo := 0
	for _, repo := range repos {
		if fork, ok := repo["fork"].(bool); ok && !fork {
			contributedTo++
		}
	}
	return contributedTo
}

func getContributionStats(token string) (int, int, int) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.github.com/users/YOUR_GITHUB_USERNAME/events", nil)
	req.Header.Set("Authorization", "token "+token)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error fetching contribution stats: %v\n", err)
		return 0, 0, 0
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading contribution stats response body: %v\n", err)
		return 0, 0, 0
	}

	var events []map[string]interface{}
	if err := json.Unmarshal(body, &events); err != nil {
		fmt.Printf("Error parsing contribution stats JSON: %v\n", err)
		return 0, 0, 0
	}

	totalContributions := 0
	currentStreak := 0
	longestStreak := 0
	lastContributionDate := time.Time{}

	for _, event := range events {
		if eventType, ok := event["type"].(string); ok {
			if eventType == "PushEvent" || eventType == "PullRequestEvent" || eventType == "IssuesEvent" {
				if createdAt, ok := event["created_at"].(string); ok {
					contributionDate, err := time.Parse(time.RFC3339, createdAt)
					if err == nil {
						totalContributions++
						if lastContributionDate.IsZero() || contributionDate.YearDay() != lastContributionDate.YearDay() {
							if contributionDate.Sub(lastContributionDate) <= 24*time.Hour {
								currentStreak++
								if currentStreak > longestStreak {
									longestStreak = currentStreak
								}
							} else {
								currentStreak = 1
							}
							lastContributionDate = contributionDate
						}
					}
				}
			}
		}
	}

	return totalContributions, currentStreak, longestStreak
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is running"))
}
